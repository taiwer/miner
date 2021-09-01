package seckill

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ItemInfo struct {
	Price struct {
		Discount        string `json:"discount"`
		Epp             string `json:"epp"`
		HagglePromotion bool   `json:"hagglePromotion"`
		Id              string `json:"id"`
		L               string `json:"l"`
		M               string `json:"m"`
		Nup             string `json:"nup"`
		Op              string `json:"op"`
		P               string `json:"p"`
		PlusTag         struct {
			Limit     bool `json:"limit"`
			Min       int  `json:"min"`
			Max       int  `json:"max"`
			Overlying bool `json:"overlying"`
		} `json:"plusTag"`
		Pp  string `json:"pp"`
		Sdp string `json:"sdp"`
		Sp  string `json:"sp"`
		Tkp string `json:"tkp"`
		Tpp string `json:"tpp"`
	} `json:"price"`
	YuyueInfo struct {
		BtnText          string   `json:"btnText"`
		BuyTime          string   `json:"buyTime"`
		CdPrefix         string   `json:"cdPrefix"`
		Countdown        int      `json:"countdown"`
		HidePrice        int      `json:"hidePrice"`
		HideText         string   `json:"hideText"`
		Num              int      `json:"num"`
		PlusText         string   `json:"plusText"`
		PlusType         int      `json:"plusType"`
		SellWhilePresell string   `json:"sellWhilePresell"`
		ShowDraw         bool     `json:"showDraw"`
		ShowPlusTime     bool     `json:"showPlusTime"`
		ShowPromoPrice   int      `json:"showPromoPrice"`
		State            int      `json:"state"`
		SupportOther     int      `json:"supportOther"`
		Type             string   `json:"type"`
		Url              string   `json:"url"`
		UserType         string   `json:"userType"`
		Yuyue            bool     `json:"yuyue"`
		YuyueRuleText    []string `json:"yuyueRuleText"`
		YuyueTime        string   `json:"yuyueTime"`
	} `json:"yuyueInfo"`
	MiaoshaInfo struct {
		EndTime           int64  `json:"endTime"`
		Miaosha           bool   `json:"miaosha"`
		MiaoshaRemainTime int    `json:"miaoshaRemainTime"`
		MsTrailer         string `json:"msTrailer"`
		OriginPrice       string `json:"originPrice"`
		SeckillType       string `json:"seckillType"`
		StartTime         int64  `json:"startTime"`
		State             int    `json:"state"`
		Title             string `json:"title"`
	} `json:"miaoshaInfo"`
}

//获取秒杀列表
func (s *Seckill) GetItem(skuId string) (*ItemInfo, error) {
	//s.GetHttpHtmlContent("https://yushou.jd.com/member/qualificationList.action", "body", `document.querySelector("body")`)
	req := s.NewRequest()
	req.SetUrl("https://item-soa.jd.com/getWareBusiness")
	req.SetParam("skuId", skuId)
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return nil, errors.New("访问商品详情失败")
	}
	itemInfo := &ItemInfo{}
	json.Unmarshal([]byte(body), itemInfo)
	return itemInfo, nil
}

//添加到购物车
func (s *Seckill) AddItemToCart(pid string, count int) error {
	//s.GetHttpHtmlContent("https://yushou.jd.com/member/qualificationList.action", "body", `document.querySelector("body")`)
	req := s.NewRequest()
	req.SetUrl("https://cart.jd.com/gate.action")
	req.SetParam("pid", pid)
	req.SetParam("ptype", "1")
	req.SetParam("pcount", strconv.Itoa(count))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 302 {
			Location := resp.Header.Get("Location")
			if strings.Index(Location, "cart.jd.com/addToCart.html") >= 0 {
				return nil
			}
			if strings.Index(Location, "passport.jd.com/new/login.aspx") >= 0 {
				return errors.New("账号未登录")
			}
		}
		log.Println("添加到购物车是失败")
		return errors.New("添加到购物车是失败")
	}
	itemInfo := &ItemInfo{}
	json.Unmarshal([]byte(body), itemInfo)
	return nil
}
