package seckill

import (
	"github.com/Albert-Zhan/httpc"
	"github.com/taiwer/miner/models/jdSeckill/conf"
	"log"
	"net/http"
)

type Request struct {
	client    *httpc.HttpClient
	cookieJar *httpc.CookieJar
	conf      *conf.Config
}

func (s *Request) Init() {
	client := httpc.NewHttpClient()
	s.cookieJar = httpc.NewCookieJar()
	client.SetCookieJar(s.cookieJar)
	client.SetRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
	s.client = client
	//配置文件初始化
	confFile := "./conf/jd_seckill_conf.ini"
	if !Exists(confFile) {
		log.Println("配置文件不存在，程序退出")
	}
	s.conf = &conf.Config{}
	s.conf.InitConfig(confFile)
}

func (s *Request) NewRequest() *httpc.Request {
	req := httpc.NewRequest(s.client)
	req.SetHeader("User-Agent", s.conf.Read("config", "DEFAULT_USER_AGENT"))
	req.SetHeader("Connection", "keep-alive")
	return req
}

//返回
func (s *Request) OnResponse() {
	log.Println("OnResponse")
}
