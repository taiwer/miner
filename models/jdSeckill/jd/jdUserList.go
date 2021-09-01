package jd

import (
	"github.com/taiwer/miner/models/jdSeckill/seckill"
	"sync"
)

type JdUserList struct {
	list map[string]*seckill.Seckill
	l    sync.Mutex
}

var JdUserListObj *JdUserList

func init() {
	JdUserListObj = &JdUserList{
		list: make(map[string]*seckill.Seckill),
	}
}

func GetSeckill(jdUserName string) *seckill.Seckill {
	return JdUserListObj.Get(jdUserName)
}

func (s *JdUserList) Add(jdUserName string, token string) {
	sk := s.get(jdUserName)
	if token != "" {
		sk.SetToken(token)
	}
}

func (s *JdUserList) Get(jdUserName string) *seckill.Seckill {
	s.l.Lock()
	defer s.l.Unlock()
	return s.get(jdUserName)
}

func (s *JdUserList) get(jdUserName string) *seckill.Seckill {
	if sk, ok := s.list[jdUserName]; ok {
		return sk
	} else {
		sk := seckill.NewSeckill()
		s.list[jdUserName] = sk
		return sk
	}
}
