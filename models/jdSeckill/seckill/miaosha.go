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

type MiaoshaItemInfo struct {
	//"wareId": "10035163404173",
	//"wname": "御峥汽车载U盘带歌曲2021抖音网络新歌网红流行音乐MP3听歌无损高音质优盘 超值64G【至尊版DTS无损音乐+视频】",
	//"shortWname": "32G 抖音热门歌曲 车载音乐U盘",
	//"imageurl": "//m.360buyimg.com/seckillcms/s250x250_jfs/t1/205041/33/3348/93850/6129ee46E16bc1e85/dfe1384154673733.jpg",
	//"good": "",
	//"jdPrice": "199",
	//"book": "false",
	//"promotion": "false",
	//"spuId": "10020375384048",
	//"adword": "",
	//"message": "",
	//"canBuy": "true",
	//"miaoSha": "true",
	//"rate": "1.5折",
	//"startRemainTime": -4074,
	//"endRemainTime": 82324,
	//"miaoShaPrice": "29.9",
	//"discount": "169.10",
	//"activeId": "7199931",
	//"canFreeRead": "false",
	//"moreFunId": "searchCatelogy",
	//"cid": "",
	//"cName": "",
	//"sourceValue": "42_1_10035163404173_0_3_0_0_0",
	//"sourceValueMapStr": "{\"preSale\":\"0\",\"isQK\":\"0\",\"fpgLabel\":\"0\",\"sku\":\"10035163404173\",\"type\":\"0\",\"isPromo\":\"0\",\"BiInfo\":\"3\"}",
	//"operateWord": "限时秒杀 抢先提醒",
	//"specificationLabel": "限量2500件",
	//"promotionId": "110383690190",
	//"soldRate": 86,
	//"almostSoldoutViewUser": 0,
	//"tagType": 15,
	//"tagText": "超级秒杀",
	//"clockNum": 5872,
	//"lowestPriceDaysInfo": "21天历史最低价",
	//"startTimeShow": "18:00",
	//"resultSort": 11,
	//"startTimeMills": 1630317600000,
	//"isNewGoods": 0,
	//"clk": "10035163404173#0#0",
	//"promoType": 1,
	//"seckillProdType": "0#3",
	//"bubbleTips": "",
	//"isHistoryProduct": false,
	//"sourceType": 1,
	//"priceDiscountText": "省169.1元",
	//"soldRateText": "86%",
	//"displayPrice": "29.9",
	//"displayPriceType": 0,
	//"hashKey": 10035163404173
	WareId          string `json:"wareId"`
	Wname           string `json:"wname"`           //商品名称
	ShortWname      string `json:"shortWname"`      //商品短名称
	Imageurl        string `json:"imageurl"`        //图片地址
	JdPrice         string `json:"jdPrice"`         //京东价格
	CanBuy          string `json:"canBuy"`          //是否可以购买
	IsMiaoSha       string `json:"miaoSha"`         //是否秒杀
	Rate            string `json:"rate"`            //折扣
	StartRemainTime int64  `json:"startRemainTime"` //开始了时间 开始后可能会为负值
	EndRemainTime   int64  `json:"endRemainTime"`   //还有多久结束 单位秒
	MiaoShaPrice    string `json:"miaoShaPrice"`    //秒杀价格
	TagText         string `json:"tagText"`         //超级秒杀
	SoldRate        int32  `json:"soldRate"`        //已售百分比
}

//获取秒杀列表
func (s *Seckill) GetMiaoShaList() ([]*MiaoshaItemInfo, error) {
	//s.GetHttpHtmlContent("https://yushou.jd.com/member/qualificationList.action", "body", `document.querySelector("body")`)
	req := s.NewRequest()
	req.SetHeader("Referer", "https://miaosha.jd.com/")
	req.SetUrl("https://api.m.jd.com/api")
	req.SetParam("appid", "o2_channels")
	req.SetParam("functionId", "pcMiaoShaAreaList")
	req.SetParam("client", "pc")
	req.SetParam("clientVersion", "1.0.0")
	//req.SetParam("callback", "pcMiaoShaAreaList")
	//req.SetParam("jsonp", "pcMiaoShaAreaList")
	req.SetParam("body", "{}")
	req.SetParam("_", strconv.Itoa(int(time.Now().Unix())))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return nil, errors.New("访问商品详情失败")
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
	return list, nil
}
