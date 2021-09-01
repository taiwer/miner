package seckill

import (
	"context"
	"errors"
	"fmt"
	"github.com/Albert-Zhan/httpc"
	"github.com/PuerkitoBio/goquery"
	"github.com/taiwer/miner/models/jdSeckill/email"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Seckill struct {
	Request
	ctx         context.Context
	cancel      context.CancelFunc
	reserveList map[string]*ProductInfo
}

func NewSeckill() *Seckill {
	seckill := &Seckill{reserveList: make(map[string]*ProductInfo, 0)}
	seckill.ctx, seckill.cancel = context.WithCancel(context.Background())
	seckill.Request.Init()
	return seckill
}

func (this *Seckill) startSeckill(skuid string, taskNum int) {
	for i := 1; i <= taskNum; i++ {
		go func() {
			ctx, _ := context.WithCancel(this.ctx)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if this.RequestSeckillUrl(skuid) {
						this.SeckillPage(skuid)
						this.SubmitSeckillOrder(skuid, "")
					}
				}
			}
		}()
	}
}

func (this *Seckill) SkuTitle(skuId string) (title, action string, err error) {
	//skuId := this.conf.Read("config", "sku_id")
	//https://item.m.jd.com/product/100011483893.html?from=qrcode
	req := httpc.NewRequest(this.client)
	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Referer", "https://order.jd.com/center/list.action")
	req.SetUrl(fmt.Sprintf("https://item.jd.com/%s.html", skuId)).SetMethod("get")
	resp, body, err := req.Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", "", errors.New("访问商品详情失败")
	}
	html := strings.NewReader(body)
	doc, _ := goquery.NewDocumentFromReader(html)
	title = strings.TrimSpace(doc.Find(".sku-name").Text())
	action = strings.TrimSpace(doc.Find("#btn-reservation").Text())
	return title, action, nil
}

func (this *Seckill) MakeReserve(skuId string) {
	if skuId == "" {
		skuId = this.conf.Read("config", "sku_id")
	}
	shopTitle, action, err := this.SkuTitle(skuId)
	if err != nil {
		log.Println("获取商品信息失败")
	} else {
		log.Println("商品名称:"+shopTitle, action)
	}
	//skuId := this.conf.Read("config", "sku_id")
	req := httpc.NewRequest(this.client)
	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Referer", fmt.Sprintf("https://item.jd.com/%s.html", skuId))
	resp, body, err := req.SetUrl("https://yushou.jd.com/youshouinfo.action?callback=fetchJSON&sku=" + skuId + "&_=" + strconv.Itoa(int(time.Now().Unix()*1000))).SetMethod("get").Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("预约商品失败")
	} else {
		reserveUrl := gjson.Get(body, "url").String()
		req = httpc.NewRequest(this.client)
		_, _, _ = req.SetUrl("https:" + reserveUrl).SetMethod("get").Send().End()
		log.Println("预约成功，已获得抢购资格 / 您已成功预约过了，无需重复预约")
	}
}

func (this *Seckill) getSeckillUrl() (string, error) {
	log.Println("获取秒杀Url")
	skuId := this.conf.Read("config", "sku_id")
	req := httpc.NewRequest(this.client)
	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Host", "itemko.jd.com")
	req.SetHeader("Referer", fmt.Sprintf("https://item.jd.com/%s.html", skuId))
	req.SetUrl("https://itemko.jd.com/itemShowBtn?callback=jQuery{}" + strconv.Itoa(Rand(1000000, 9999999)) + "&skuId=" + skuId + "&from=pc&_=" + strconv.Itoa(int(time.Now().Unix()*1000))).SetMethod("get")
	resp, body, err := req.Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("抢购链接获取失败，稍后自动重试")
		return "", errors.New("抢购链接获取失败，稍后自动重试")
	}
	url := gjson.Get(body, "url").String()
	if url == "" {
		log.Println("抢购链接获取失败，稍后自动重试")
		return "", errors.New("抢购链接获取失败，稍后自动重试")
	}
	//https://divide.jd.com/user_routing?skuId=8654289&sn=c3f4ececd8461f0e4d7267e96a91e0e0&from=pc
	url = strings.ReplaceAll(url, "divide", "marathon")
	//https://marathon.jd.com/captcha.html?skuId=8654289&sn=c3f4ececd8461f0e4d7267e96a91e0e0&from=pc
	url = strings.ReplaceAll(url, "user_routing", "captcha.html")
	return url, nil
}

func (this *Seckill) RequestSeckillUrl(skuId string) bool {
	userInfo, err := this.GetUserInfo()
	if err != nil {
		log.Println("获取用户信息失败")
	} else {
		log.Println("用户:" + userInfo)
	}
	shopTitle, action, err := this.SkuTitle(skuId)
	if err != nil {
		log.Println("获取商品信息失败")
	} else {
		log.Println("商品名称:"+shopTitle, action)
	}
	if action == "抢购" {
		return true
	} else {
		return false
	}
	//url, _ := this.getSeckillUrl()
	//if url != "" {
	//	if skuId == "" {
	//		skuId = this.conf.Read("config", "sku_id")
	//	}
	//	req := httpc.NewRequest(this.client)
	//	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	//	req.SetHeader("Host", "marathon.jd.com")
	//	req.SetHeader("Referer", fmt.Sprintf("https://item.jd.com/%s.html", skuId))
	//	_, _, _ = req.SetUrl(url).SetMethod("get").Send().End()
	//	return true
	//}
	//return false
}

func (this *Seckill) SeckillPage(skuId string) {
	log.Println("访问抢购订单结算页面...")
	seckillNum := this.conf.Read("config", "seckill_num")
	req := httpc.NewRequest(this.client)
	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Host", "marathon.jd.com")
	req.SetHeader("Referer", fmt.Sprintf("https://item.jd.com/%s.html", skuId))
	req.SetUrl("https://marathon.jd.com/seckill/seckill.action?skuId=" + skuId + "&num=" + seckillNum + "&rid=" + strconv.Itoa(int(time.Now().Unix()))).SetMethod("get")
	_, _, _ = req.Send().End()
}

func (this *Seckill) SeckillInitInfo(skuId, seckillNum string) (string, error) {
	log.Println("获取秒杀初始化信息...")
	if skuId == "" {
		skuId = this.conf.Read("config", "sku_id")
	}
	if seckillNum == "" {
		seckillNum = this.conf.Read("config", "seckill_num")
	}
	req := httpc.NewRequest(this.client)
	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Host", "marathon.jd.com")
	req.SetParam("sku", skuId)
	req.SetParam("num", seckillNum)
	req.SetParam("isModifyAddress", "false")
	req.SetUrl("https://marathon.jd.com/seckillnew/orderService/pc/init.action").SetMethod("post")
	resp, body, err := req.Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("初始化秒杀信息失败")
		return "", errors.New("初始化秒杀信息失败")
	}
	return body, nil
}

func (this *Seckill) SubmitSeckillOrder(skuId, seckillNum string) bool {
	if skuId == "" {
		skuId = this.conf.Read("config", "sku_id")
	}
	if seckillNum == "" {
		seckillNum = this.conf.Read("config", "seckill_num")
	}
	eid := this.conf.Read("config", "eid")
	fp := this.conf.Read("config", "fp")
	paymentPwd := this.conf.Read("account", "payment_pwd")
	initInfo, _ := this.SeckillInitInfo(skuId, seckillNum)
	address := gjson.Get(initInfo, "addressList").Array()
	defaultAddress := address[0]
	isinvoiceInfo := gjson.Get(initInfo, "invoiceInfo").Exists()
	invoiceTitle := "-1"
	invoiceContentType := "-1"
	invoicePhone := ""
	invoicePhoneKey := ""
	if isinvoiceInfo {
		invoiceTitle = gjson.Get(initInfo, "invoiceInfo.invoiceTitle").String()
		invoiceContentType = gjson.Get(initInfo, "invoiceInfo.invoiceContentType").String()
		invoicePhone = gjson.Get(initInfo, "invoiceInfo.invoicePhone").String()
		invoicePhoneKey = gjson.Get(initInfo, "invoiceInfo.invoicePhoneKey").String()
	}
	invoiceInfo := "false"
	if isinvoiceInfo {
		invoiceInfo = "true"
	}
	token := gjson.Get(initInfo, "token").String()
	log.Println("提交抢购订单...")
	req := httpc.NewRequest(this.client)
	req.SetHeader("User-Agent", this.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Host", "marathon.jd.com")
	req.SetHeader("Referer", fmt.Sprintf("https://marathon.jd.com/seckill/seckill.action?skuId=%s&num=%s&rid=%d", skuId, seckillNum, int(time.Now().Unix())))
	req.SetParam("skuId", skuId)
	req.SetParam("num", seckillNum)
	req.SetParam("addressId", defaultAddress.Get("id").String())
	req.SetParam("yuShou", "true")
	req.SetParam("isModifyAddress", "false")
	req.SetParam("name", defaultAddress.Get("name").String())
	req.SetParam("provinceId", defaultAddress.Get("provinceId").String())
	req.SetParam("cityId", defaultAddress.Get("cityId").String())
	req.SetParam("countyId", defaultAddress.Get("countyId").String())
	req.SetParam("townId", defaultAddress.Get("townId").String())
	req.SetParam("addressDetail", defaultAddress.Get("addressDetail").String())
	req.SetParam("mobile", defaultAddress.Get("mobile").String())
	req.SetParam("mobileKey", defaultAddress.Get("mobileKey").String())
	req.SetParam("email", defaultAddress.Get("email").String())
	req.SetParam("postCode", "")
	req.SetParam("invoiceTitle", invoiceTitle)
	req.SetParam("invoiceCompanyName", "")
	req.SetParam("invoiceContent", invoiceContentType)
	req.SetParam("invoiceTaxpayerNO", "")
	req.SetParam("invoiceEmail", "")
	req.SetParam("invoicePhone", invoicePhone)
	req.SetParam("invoicePhoneKey", invoicePhoneKey)
	req.SetParam("invoice", invoiceInfo)
	req.SetParam("password", paymentPwd)
	req.SetParam("codTimeType", "3")
	req.SetParam("paymentType", "4")
	req.SetParam("areaCode", "")
	req.SetParam("overseas", "0")
	req.SetParam("phone", "")
	req.SetParam("eid", eid)
	req.SetParam("fp", fp)
	req.SetParam("token", token)
	req.SetParam("pru", "")
	resp, body, err := req.SetUrl("https://marathon.jd.com/seckillnew/orderService/pc/submitOrder.action?skuId=" + skuId).SetMethod("post").Send().End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("抢购失败，网络错误")
		if this.conf.Read("messenger", "enable") == "true" && this.conf.Read("messenger", "type") == "smtp" {
			email := email.NerEmail(this.conf)
			_ = email.SendMail([]string{this.conf.Read("messenger", "email")}, "茅台抢购通知", "抢购失败，网络错误")
		}
		return false
	}
	if !gjson.Valid(body) {
		log.Println("抢购失败，返回信息:" + Substr(body, 0, 128))
		if this.conf.Read("messenger", "enable") == "true" && this.conf.Read("messenger", "type") == "smtp" {
			email := email.NerEmail(this.conf)
			_ = email.SendMail([]string{this.conf.Read("messenger", "email")}, "茅台抢购通知", "抢购失败，返回信息:"+Substr(body, 0, 128))
		}
		return false
	}
	if gjson.Get(body, "success").Bool() {
		orderId := gjson.Get(body, "orderId").String()
		totalMoney := gjson.Get(body, "totalMoney").String()
		payUrl := "https:" + gjson.Get(body, "pcUrl").String()
		log.Println(fmt.Sprintf("抢购成功，订单号:%s, 总价:%s, 电脑端付款链接:%s", orderId, totalMoney, payUrl))
		if this.conf.Read("messenger", "enable") == "true" && this.conf.Read("messenger", "type") == "smtp" {
			email := email.NerEmail(this.conf)
			_ = email.SendMail([]string{this.conf.Read("messenger", "email")}, "茅台抢购通知", fmt.Sprintf("抢购成功，订单号:%s, 总价:%s, 电脑端付款链接:%s", orderId, totalMoney, payUrl))
		}
		return true
	} else {
		log.Println("抢购失败，返回信息:" + body)
		if this.conf.Read("messenger", "enable") == "true" && this.conf.Read("messenger", "type") == "smtp" {
			email := email.NerEmail(this.conf)
			_ = email.SendMail([]string{this.conf.Read("messenger", "email")}, "茅台抢购通知", "抢购失败，返回信息:"+body)
		}
		return false
	}
}
