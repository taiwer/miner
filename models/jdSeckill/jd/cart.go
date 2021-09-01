package jd

//添加到购物车
func AddItemToCart(jdusername string, pid string, count int) error {
	miaosha := GetSeckill(jdusername)
	return miaosha.AddItemToCart(pid, count)
}

//添加到购物车
func GetCartList(jdusername string) (string, error) {
	miaosha := GetSeckill(jdusername)
	return miaosha.GetCartList()
}

func CartCheckSingle(jdusername string, id string, num int) (string, error) {
	miaosha := GetSeckill(jdusername)
	return miaosha.CartCheckSingle(id, num)
}

func CartUnCheckSingle(jdusername string, id string, num int) (string, error) {
	miaosha := GetSeckill(jdusername)
	return miaosha.CartUnCheckSingle(id, num)
}

func CartUnCheckAll(jdusername string) (string, error) {
	miaosha := GetSeckill(jdusername)
	return miaosha.CartUnCheckAll()
}
