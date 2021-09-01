package jd

import (
	"github.com/taiwer/miner/models/jdSeckill/seckill"
	"log"
	"os"
	"time"
)

var miaosha *seckill.Seckill

func init() {
	miaosha = seckill.NewSeckill()
	miaosha.LoadCookie("") //jd189806cks
	unick, _ := miaosha.GetUserInfo()
	if unick == "" {
		wlfstkSmdl, err := miaosha.QrLogin()
		if err != nil {
			os.Exit(0)
		}
		ticket := ""
		//等待登录
		for {
			_, _, ticket, err = miaosha.QrcodeTicket(wlfstkSmdl)
			if err == nil && ticket != "" {
				break
			}
			time.Sleep(2 * time.Second)
		}
		_, _, err = miaosha.TicketInfo(ticket)
		if err != nil {
			log.Println("登录失败")
			return
		}
	}
}

func GetMiaoshaList(jdusername string) []*seckill.MiaoshaItemInfo {
	miaosha := GetSeckill(jdusername)
	list, _ := miaosha.GetMiaoShaList()
	return list
}

func GetItemInfo(jdusername string, skuId string) *seckill.ItemInfo {
	miaosha := GetSeckill(jdusername)
	list, _ := miaosha.GetItem(skuId)
	return list
}
