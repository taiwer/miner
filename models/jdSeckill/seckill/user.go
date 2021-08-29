package seckill

import (
	"errors"
	"fmt"
	"github.com/Albert-Zhan/httpc"
	"github.com/Unknwon/goconfig"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const cookieFile = "./conf/cookie.ini"

func (this *Seckill) loginPage() {
	req := this.NewRequest()
	req.SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	_, _, _ = req.SetUrl("https://passport.jd.com/new/login.aspx").SetMethod("get").Send().End()
}

func (this *Seckill) LoadCookie(unick string) (bool, error) {
	if unick == "" {
		unick = this.conf.Read("account", "unick")
	}
	config, err := goconfig.LoadConfigFile(cookieFile)
	if err != nil {
		log.Println(err)
		return false, err
	}
	thor, err := config.GetValue("thor", unick)
	if thor != "" {
		this.cookieJar.SetCookies(&url.URL{Host: "jd.com", Scheme: "https"}, []*http.Cookie{&http.Cookie{Name: "thor", Value: thor, Path: "/", Domain: "jd.com"}})
		return true, nil
	}
	return false, err
}

func (this *Seckill) QrLogin() (string, error) {
	//登录页面
	this.GetReserveList()
	this.loginPage()
	//二维码登录
	req := this.NewRequest()
	req.SetHeader("Referer", "https://passport.jd.com/new/login.aspx")
	resp, err := req.SetUrl("https://qr.m.jd.com/show?appid=133&size=147&t="+strconv.Itoa(int(time.Now().Unix()*1000))).SetMethod("get").Send().EndFile("./", "qr_code.png")
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("获取二维码失败")
		return "", errors.New("获取二维码失败")
	}
	cookies := resp.Cookies()
	wlfstkSmdl := ""
	for _, cookie := range cookies {
		if cookie.Name == "wlfstk_smdl" {
			wlfstkSmdl = cookie.Value
			break
		}
	}
	log.Println("二维码获取成功，请打开京东APP扫描")
	dir, _ := os.Getwd()
	OpenImage(dir + "/qr_code.png")
	return wlfstkSmdl, nil
}

func (this *Seckill) QrcodeTicket(wlfstkSmdl string) (string, error) {
	req := this.NewRequest()
	req.SetHeader("Referer", "https://passport.jd.com/new/login.aspx")
	params := []string{}
	//params = append(params, "callback=jQuery"+strconv.Itoa(int(Rand(1000000, 9999999))))
	params = append(params, "callback=jQuery111")
	params = append(params, "token="+wlfstkSmdl)
	params = append(params, "_="+strconv.Itoa(int(time.Now().Unix())))
	resp, body, err := req.SetUrl("https://qr.m.jd.com/check?appid=133&" + strings.Join(params, "&")).SetMethod("get").Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("获取二维码扫描结果异常")
		return "", errors.New("获取二维码扫描结果异常")
	}
	if gjson.Get(body, "code").Int() != 200 {
		log.Printf("Code: %s, Message: %s", gjson.Get(body, "code").String(), gjson.Get(body, "msg").String())
		return "", errors.New(fmt.Sprintf("Code: %s, Message: %s", gjson.Get(body, "code").String(), gjson.Get(body, "msg").String()))
	}
	log.Println("已完成手机客户端确认")
	return gjson.Get(body, "ticket").String(), nil
}

func (this *Seckill) TicketInfo(ticket string) (string, error) {
	req := this.NewRequest()
	req.SetHeader("Referer", "https://passport.jd.com/uc/login?ltype=logout")
	resp, body, err := req.SetUrl("https://passport.jd.com/uc/qrCodeTicketValidation?t=" + ticket).SetMethod("get").Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("二维码信息校验失败")
		return "", errors.New("二维码信息校验失败")
	}
	if gjson.Get(body, "returnCode").Int() == 0 {
		log.Println("二维码信息校验成功")
		//保存cookie
		unick := ""
		thor := ""
		cookies := resp.Cookies()
		for _, cookie := range cookies {
			switch cookie.Name {
			case "thor":
				thor = cookie.Value
			case "unick":
				unick = cookie.Value
			}
		}
		if unick != "" && thor != "" {
			if !Exists(cookieFile) {
				f, _ := os.Create(cookieFile)
				f.Close()
			}
			config, err := goconfig.LoadConfigFile(cookieFile)
			if err != nil {
				log.Println(err)
				if config == nil {
					config = &goconfig.ConfigFile{}
				}
			}
			config.SetValue("thor", unick, thor)
			goconfig.SaveConfigFile(config, cookieFile)
		}
		return "", nil
	} else {
		log.Println("二维码信息校验失败")
		return "", errors.New("二维码信息校验失败")
	}
}

func (this *Seckill) RefreshStatus() error {
	req := this.NewRequest()
	resp, _, err := req.SetUrl("https://order.jd.com/center/list.action?rid=" + strconv.Itoa(int(time.Now().Unix()*1000))).SetMethod("get").Send().End()
	if err == nil && resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New("登录失效")
	}
}

func (this *Seckill) GetUserInfo() (string, error) {
	req := httpc.NewRequest(this.client)
	req.SetHeader("Referer", "https://order.jd.com/center/list.action")
	resp, body, err := req.SetUrl("https://passport.jd.com/user/petName/getUserInfoForMiniJd.action?callback=" + strconv.Itoa(Rand(1000000, 9999999)) + "&_=" + strconv.Itoa(int(time.Now().Unix()*1000))).SetMethod("get").Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("获取用户信息失败")
		return "", errors.New("获取用户信息失败")
	} else {
		b, _ := GbkToUtf8([]byte(gjson.Get(body, "nickName").String()))
		log.Println("获取用户信息成功", string(b))
		return string(b), nil
	}
}
