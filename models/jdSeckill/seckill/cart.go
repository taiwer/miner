package seckill

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TheSku struct {
	Id      string `json:"Id"`
	Num     int    `json:"num"`
	SkuUuid string `json:"skuUuid"`
	UseUuid bool   `json:"useUuid"`
}

type Operation struct {
	TheSkus []TheSku `json:"TheSkus"`
}
type SerInfo struct {
	Area    string `json:"area"`
	UserKey string `json:"user-key"`
}
type JDC_mall_cart struct {
	Operations []Operation `json:"operations"`
	SerInfo    SerInfo     `json:"serInfo"`
}

//获取秒杀列表
func (s *Seckill) GetCartList() (string, error) {
	//https://api.m.jd.com/api?functionId=pcCart_jc_getCurrentCart&appid=JDC_mall_cart&loginType=3&body={"serInfo":{"area":"22_1930_50949_52153","user-key":"ee6655ac-b117-485c-ac81-2d7a031cfacb"},"cartExt":{"specialId":1}}
	req := s.NewRequest()
	req.SetHeader("Referer", "https://cart.jd.com/cart_index")
	req.SetUrl("https://api.m.jd.com/api")
	req.SetParam("appid", "JDC_mall_cart")
	req.SetParam("functionId", "pcCart_jc_getCurrentCart")
	req.SetParam("loginType", "3")
	req.SetParam("body", "{\"serInfo\":{\"area\":\"22_1930_50949_52153\",\"user-key\":\"ee6655ac-b117-485c-ac81-2d7a031cfacb\"},\"cartExt\":{\"specialId\":1}}")
	req.SetParam("_", strconv.Itoa(int(time.Now().Unix())))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", errors.New("访问商品详情失败")
	}
	jo := gjson.Parse(body)
	miaoShaList := jo.Get("miaoShaList").Array()
	list := make([]*MiaoshaItemInfo, 0)
	for k, v := range miaoShaList {
		//	fmt.Println(v.String())
		fmt.Println(k)
		miaoshaItemInfo := &MiaoshaItemInfo{}
		json.Unmarshal([]byte(v.String()), miaoshaItemInfo)
		list = append(list, miaoshaItemInfo)
		fmt.Printf("%+v\n", miaoshaItemInfo)
	}
	return body, nil
}

//获取秒杀列表
func (s *Seckill) CartCheckSingle(id string, num int) (string, error) {
	req := s.NewRequest()
	req.SetHeader("Referer", "https://cart.jd.com/cart_index")
	req.SetUrl("https://api.m.jd.com/api")
	req.SetParam("appid", "JDC_mall_cart")
	req.SetParam("functionId", "pcCart_jc_cartCheckSingle")
	req.SetParam("loginType", "3")
	//req.SetParam("body", "{\"serInfo\":{\"area\":\"22_1930_50949_52153\",\"user-key\":\"ee6655ac-b117-485c-ac81-2d7a031cfacb\"},\"cartExt\":{\"specialId\":1}}")
	var mall_cart = JDC_mall_cart{}
	mall_cart.SerInfo.Area = "22_1930_50949_52153"
	mall_cart.SerInfo.UserKey = "ee6655ac-b117-485c-ac81-2d7a031cfacb"
	mall_cart.Operations = append(mall_cart.Operations, Operation{
		TheSkus: []TheSku{
			TheSku{
				Id:      id,
				Num:     num,
				SkuUuid: "",
				UseUuid: false,
			},
		},
	})
	sbody, _ := json.Marshal(mall_cart)
	req.SetParam("body", string(sbody))
	req.SetParam("_", strconv.Itoa(int(time.Now().Unix())))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", errors.New("访问商品详情失败")
	}
	jo := gjson.Parse(body)
	miaoShaList := jo.Get("miaoShaList").Array()
	list := make([]*MiaoshaItemInfo, 0)
	for k, v := range miaoShaList {
		//	fmt.Println(v.String())
		fmt.Println(k)
		miaoshaItemInfo := &MiaoshaItemInfo{}
		json.Unmarshal([]byte(v.String()), miaoshaItemInfo)
		list = append(list, miaoshaItemInfo)
		fmt.Printf("%+v\n", miaoshaItemInfo)
	}
	return body, nil
}

//获取秒杀列表
func (s *Seckill) CartUnCheckSingle(id string, num int) (string, error) {
	req := s.NewRequest()
	req.SetHeader("Referer", "https://cart.jd.com/cart_index")
	req.SetUrl("https://api.m.jd.com/api")
	req.SetParam("appid", "JDC_mall_cart")
	req.SetParam("functionId", "pcCart_jc_cartUnCheckSingle")
	req.SetParam("loginType", "3")
	//req.SetParam("body", "{\"serInfo\":{\"area\":\"22_1930_50949_52153\",\"user-key\":\"ee6655ac-b117-485c-ac81-2d7a031cfacb\"},\"cartExt\":{\"specialId\":1}}")
	var mall_cart = JDC_mall_cart{}
	mall_cart.SerInfo.Area = "22_1930_50949_52153"
	mall_cart.SerInfo.UserKey = "ee6655ac-b117-485c-ac81-2d7a031cfacb"
	mall_cart.Operations = append(mall_cart.Operations, Operation{
		TheSkus: []TheSku{
			TheSku{
				Id:      id,
				Num:     num,
				SkuUuid: "",
				UseUuid: false,
			},
		},
	})
	sbody, _ := json.Marshal(mall_cart)
	req.SetParam("body", string(sbody))
	req.SetParam("_", strconv.Itoa(int(time.Now().Unix())))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", errors.New("访问商品详情失败")
	}
	jo := gjson.Parse(body)
	miaoShaList := jo.Get("miaoShaList").Array()
	list := make([]*MiaoshaItemInfo, 0)
	for k, v := range miaoShaList {
		//	fmt.Println(v.String())
		fmt.Println(k)
		miaoshaItemInfo := &MiaoshaItemInfo{}
		json.Unmarshal([]byte(v.String()), miaoshaItemInfo)
		list = append(list, miaoshaItemInfo)
		fmt.Printf("%+v\n", miaoshaItemInfo)
	}
	return body, nil
}

//取消所有选中
func (s *Seckill) CartUnCheckAll() (string, error) {
	//https://api.m.jd.com/api?functionId=pcCart_jc_cartUnCheckAll&appid=JDC_mall_cart&loginType=3&body={"serInfo":{"area":"22_1930_50949_52153","user-key":"ee6655ac-b117-485c-ac81-2d7a031cfacb"}}
	req := s.NewRequest()
	req.SetHeader("Referer", "https://cart.jd.com/cart_index")
	req.SetUrl("https://api.m.jd.com/api")
	req.SetParam("appid", "JDC_mall_cart")
	req.SetParam("functionId", "pcCart_jc_cartUnCheckAll")
	req.SetParam("loginType", "3")
	//req.SetParam("body", "{\"serInfo\":{\"area\":\"22_1930_50949_52153\",\"user-key\":\"ee6655ac-b117-485c-ac81-2d7a031cfacb\"},\"cartExt\":{\"specialId\":1}}")
	var mall_cart = JDC_mall_cart{}
	mall_cart.SerInfo.Area = "22_1930_50949_52153"
	mall_cart.SerInfo.UserKey = "ee6655ac-b117-485c-ac81-2d7a031cfacb"
	sbody, _ := json.Marshal(mall_cart)
	req.SetParam("body", string(sbody))
	req.SetParam("_", strconv.Itoa(int(time.Now().Unix())))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return "", errors.New("访问商品详情失败")
	}
	jo := gjson.Parse(body)
	miaoShaList := jo.Get("miaoShaList").Array()
	list := make([]*MiaoshaItemInfo, 0)
	for k, v := range miaoShaList {
		//	fmt.Println(v.String())
		fmt.Println(k)
		miaoshaItemInfo := &MiaoshaItemInfo{}
		json.Unmarshal([]byte(v.String()), miaoshaItemInfo)
		list = append(list, miaoshaItemInfo)
		fmt.Printf("%+v\n", miaoshaItemInfo)
	}
	return body, nil
}
