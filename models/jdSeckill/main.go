package jdSeckill

import (
	"errors"
	"fmt"
	"github.com/Albert-Zhan/httpc"
	"github.com/taiwer/miner/models/jdSeckill/conf"
	"github.com/taiwer/miner/models/jdSeckill/seckill"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

var client *httpc.HttpClient

var cookieJar *httpc.CookieJar
var config *conf.Config
var wg *sync.WaitGroup

var miaosha *seckill.Seckill

func init() {
	//客户端设置初始化
	miaosha = seckill.NewSeckill()
	client = httpc.NewHttpClient()
	cookieJar = httpc.NewCookieJar()
	client.SetCookieJar(cookieJar)
	client.SetRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
	//配置文件初始化
	confFile := "./conf/jd_seckill_conf.ini"
	if !seckill.Exists(confFile) {
		log.Println("配置文件不存在，程序退出")
		os.Exit(0)
	}
	config = &conf.Config{}
	config.InitConfig(confFile)
	wg = new(sync.WaitGroup)
	wg.Add(1)
}

func RunSeckill() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//用户登录
	user := seckill.NewSeckill()

	user.LoadCookie("") //jd189806cks
	unick, _ := user.GetUserInfo()
	if unick == "" {
		wlfstkSmdl, err := user.QrLogin()
		if err != nil {
			os.Exit(0)
		}
		dir, _ := os.Getwd()
		seckill.OpenImage(dir + "/" + seckill.QrImgPath + "/" + wlfstkSmdl + ".png")
		ticket := ""
		//等待登录
		for {
			_, _, ticket, err = user.QrcodeTicket(wlfstkSmdl)
			if err == nil && ticket != "" {
				break
			}
			time.Sleep(2 * time.Second)
		}
		_, _, err = user.TicketInfo(ticket)
		if err != nil {
			log.Println("登录失败")
			return
		}
	}

	log.Println("登录成功")
	go user.GoTickPreserve()
	//刷新用户状态和获取用户信息
	if status := user.RefreshStatus(); status == nil {
		userInfo, _ := user.GetUserInfo()
		log.Println("用户:" + userInfo)
		//开始预约,预约过的就重复预约
		seckill := user
		seckill.MakeReserve("")
		//获取预约商品列表
		text, err := seckill.GetReserveList()
		log.Println("预约商品列表")
		fmt.Println(text)
		log.Println(err)
		log.Println("预约商品列表===============")
		//等待抢购/开始抢购
		nowLocalTime := time.Now().UnixNano() / 1e6
		jdTime, _ := getJdTime()
		buyDate := config.Read("config", "buy_time")
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", buyDate, loc)
		buyTime := t.UnixNano() / 1e6
		diffTime := nowLocalTime - jdTime
		log.Println(fmt.Sprintf("正在等待到达设定时间:%s，检测本地时间与京东服务器时间误差为【%d】毫秒", buyDate, diffTime))
		timerTime := (buyTime + diffTime) - jdTime
		if timerTime <= 0 {
			log.Println("请设置抢购时间")
			os.Exit(0)
		}
		//time.Sleep(time.Duration(timerTime) * time.Millisecond)
		//开启任务
		log.Println("时间到达，开始执行……")
		//start(seckill, 1)
		wg.Wait()
	}
}

func getJdTime() (int64, error) {
	req := httpc.NewRequest(client)
	resp, body, err := req.SetUrl("https://a.jd.com//ajax/queryServerData.html").SetMethod("get").Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("获取京东服务器时间失败")
		return 0, errors.New("获取京东服务器时间失败")
	}
	return gjson.Get(body, "serverTime").Int(), nil
}

func start(sk *seckill.Seckill, taskNum int) {
	for i := 1; i <= taskNum; i++ {
		go func(seckill *seckill.Seckill) {
			for {
				if seckill.RequestSeckillUrl("") {
					seckill.SeckillPage("")
					seckill.SubmitSeckillOrder("", "")
				}
				time.Sleep(time.Second * 2)
			}

		}(sk)
	}
}
