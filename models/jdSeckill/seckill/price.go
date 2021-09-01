package seckill

import (
	"encoding/json"
	"fmt"
	"github.com/taiwer/miner/common/myutils"
	"log"
	"net/http"
	"strings"
)

type PriceInfo struct {
	Id  string `json:"id"`
	P   string `json:"p"`
	OP  string `json:"op"`
	Cbf string `json:"cbf"`
	M   string `json:"m"`
}

func (s *Seckill) GetPrices(skuId ...string) []*PriceInfo {
	//https://api.m.jd.com/api?appid=o2_channels&functionId=pcSeckillBrandGoods&client=pc&clientVersion=1.0.0&callback=pcSeckillBrandGoods&jsonp=pcSeckillBrandGoods&body=%7B%22brandId%22%3A%2251559%22%2C%22skuIds%22%3A%22%22%7D&_=1630260803270
	//https://api.m.jd.com/api?appid=o2_channels&functionId=pcSeckillBrandGoods&client=pc&clientVersion=1.0.0&callback=pcSeckillBrandGoods&jsonp=pcSeckillBrandGoods&body={"brandId":"51559","skuIds":""}&_=1630260803270
	//https://item-soa.jd.com/getWareBusiness?callback=jQuery4276045&skuId=100005537245&cat=670%2C671%2C2694&area=22_1930_50948_57091&shopId=1000000326&venderId=1000000326&paramJson=%7B%22platform2%22%3A%22100000000001%22%2C%22specialAttrStr%22%3A%22p0ppp1pppppp2p1p1pppppppppp%22%2C%22skuMarkStr%22%3A%2200%22%7D&num=1
	req := s.NewRequest()
	req.SetHeader("Referer", "https://order.jd.com/center/list.action")
	params := []string{""}
	params = append(params, "callback=jQuery8074630&type=1&area=22_1930_50948_57091.4805419862&pdtk=&pduid=&pduid=16302221852292129865716")
	params = append(params, "pdpin=18980683091_p&pin=18980683091_p&pdbp=0")
	params = append(params, "skuIds=J_"+strings.Join(skuId, "%2cJ_"))
	params = append(params, "&ext=11100000&source=item-pc")
	req.SetUrl("https://p.3.cn/prices/mgets?" + strings.Join(params, "&"))
	req.SetMethod("get")
	req.Send()
	resp, body, err := req.End()
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("访问商品详情失败")
		return nil
	}
	list := make([]*PriceInfo, 0)
	ret := myutils.GetBetweenStr(body, "(", ")")
	json.Unmarshal([]byte(ret), &list)
	fmt.Printf("%v", list)
	return list
}
