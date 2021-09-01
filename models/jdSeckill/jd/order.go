package jd

func GetOrderInfo(jdusername string) (string, error) {
	miaosha := GetSeckill(jdusername)
	return miaosha.GetOrderInfo()
}

func SubmitOrder(jdusername string) (string, error) {
	miaosha := GetSeckill(jdusername)
	return miaosha.SubmitOrder()
}
