package jd

import "github.com/taiwer/miner/models/jdSeckill/seckill"

var user *seckill.Seckill

func init() {
	user = seckill.NewSeckill()
}
func QrLogin() (string, error) {
	user := GetSeckill("")
	id, err := user.QrLogin()
	return id, err
}

func QrcodeTicket(wlfstkSmdl string) (code int64, msg string, ticket string, err error) {
	user := GetSeckill("")
	return user.QrcodeTicket(wlfstkSmdl)
}

func TicketInfo(ticket string) (thor, unick string, err error) {
	user := GetSeckill("")
	return user.TicketInfo(ticket)
}
