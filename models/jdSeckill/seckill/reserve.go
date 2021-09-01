package seckill

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/taiwer/miner/common/myutils"
	"log"
	"net/http"
	"strings"
	"time"
)

//预约商品

type ProductInfo struct {
	Id            string
	Title         string
	Price         string
	DownStartTime string //抢购时间
	DownEndTime   string //抢购距离
	IsRuning      bool   //是否已经在抢购
}

//监听收藏商品 时间到了自动抢购
func (s *Seckill) GoTickPreserve() {
	ctx, _ := context.WithCancel(s.ctx)
	tick := time.NewTicker(time.Second * 10)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			s.GetReserveList()
			s.checkDoSeckill()
		}
	}
}

//检查产品是否需要秒杀
func (s *Seckill) checkDoSeckill() {
	nowTime := time.Now().UTC().Add(time.Hour * 8)
	for _, v := range s.reserveList {
		pTime, _ := time.Parse("2006-01-02 15:04:05", v.DownStartTime)
		subTime := pTime.Sub(nowTime) / time.Second
		if subTime <= 10 { //提前10秒开枪
			if !v.IsRuning {
				v.IsRuning = true
				s.startSeckill(v.Id, 1)
			}
		} else {
			if subTime < 3600*24 {
				log.Println(v.DownStartTime, "等待时间", myutils.FormatTimeDuration(int64(subTime)), v.Title)
			}
		}
	}
}

func (s *Seckill) GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("User-Agent", s.conf.Read("config", "DEFAULT_USER_AGENT")),
		chromedp.Flag("Referer", "https://order.jd.com/center/list.action"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()
	log.Printf("Chrome visit page %s\n", url)
	//headers := make(map[string]interface{})
	//headers["User-Agent"] = s.conf.Read("config", "DEFAULT_USER_AGENT")
	//headers["Referer"] = "https://order.jd.com/center/list.action"
	//
	//network.SetExtraHTTPHeaders(headers)
	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		//chromedp.Navigate(url),
		//chromedp.WaitVisible(selector),
		//chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
		chromedp.Navigate(url),
		chromedp.WaitVisible("body"),
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
	)

	if err != nil {

		return "", err
	}
	log.Println(htmlContent)
	return htmlContent, nil
}

//获取预约商品列表
func (s *Seckill) GetReserveList() (string, error) {
	//s.GetHttpHtmlContent("https://yushou.jd.com/member/qualificationList.action", "body", `document.querySelector("body")`)
	req := s.NewRequest()
	req.SetHeader("Referer", "https://order.jd.com/center/list.action")
	req.SetUrl("https://yushou.jd.com/member/qualificationList.action")
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", errors.New("访问商品详情失败")
	}
	html := strings.NewReader(body)
	doc, _ := goquery.NewDocumentFromReader(html)
	elem := doc.Find(".cont-box")
	productList := []*ProductInfo{}
	for _, node := range elem.Nodes {
		var productInfo ProductInfo
		product := goquery.NewDocumentFromNode(node)
		productInfo.Title = product.Find(".prod-title").Text()
		//<input type="hidden" id="100011553443_buystime" value="2021-08-31 15:00:00">
		//<input type="hidden" id="100011553443_buyetime" value="2021-08-31 17:00:00">
		id := myutils.GetBetweenStr(product.Text(), "loadSkuPrice('", "'")
		productInfo.Id = id
		priceList := s.GetPrices(id)
		if len(priceList) > 0 {
			productInfo.Price = priceList[0].P
		}
		inputs := product.Find("#" + id + "_buystime")
		productInfo.DownStartTime, _ = inputs.Attr("value")
		inputs = product.Find("#" + id + "_buyetime")
		productInfo.DownEndTime, _ = inputs.Attr("value")
		productList = append(productList, &productInfo)
	}
	for _, v := range productList {
		if _, ok := s.reserveList[v.Id]; !ok {
			s.reserveList[v.Id] = v
		}
	}
	itemList := make([]string, 0)
	for k, productInfo := range productList {
		text := fmt.Sprintf("%d Id:%s 商品名称:[%s] price:[%s] sTime:[%s] eTime:[%s]", k,
			productInfo.Id, productInfo.Title, productInfo.Price, productInfo.DownStartTime, productInfo.DownEndTime)
		itemList = append(itemList, text)
	}
	text := "\n" + strings.Join(itemList, "\n")
	return strings.TrimSpace(text), nil
}
