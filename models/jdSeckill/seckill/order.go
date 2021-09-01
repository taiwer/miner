package seckill

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//获取订单信息
func (s *Seckill) GetOrderInfo() (string, error) {
	req := s.NewRequest()
	//req.SetHeader("Referer", "https://cart.jd.com/cart_index")
	req.SetUrl("https://trade.jd.com/shopping/order/getOrderInfo.action")
	req.SetParam("_", strconv.Itoa(int(time.Now().Unix())))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", errors.New("访问商品详情失败")
	}
	html := strings.NewReader(body)
	doc, _ := goquery.NewDocumentFromReader(html)
	elemneedPay := doc.Find("#needPay")
	if elemneedPay != nil {
		needPay, _ := elemneedPay.Attr("value")
		return needPay, nil
	}
	//<input type="hidden" id="needPay" value="89.90"/>
	//<input type="hidden" id="lastneedPay" value="89.90"/>
	//<input type="hidden" id="btNeedPay" value="89.90"/>
	return "", nil
}

//提交订单
func (s *Seckill) SubmitOrder() (string, error) {
	req := s.NewRequest()
	req.SetHeader("Referer", "https://trade.jd.com/shopping/order/getOrderInfo.action")
	req.SetUrl("https://trade.jd.com/shopping/order/submitOrder.action")
	req.SetParam("presaleStockSign", "1")
	req.SetMethod("post")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("提交订单")
		return "", errors.New("提交订单")
	}
	return body, nil
}
