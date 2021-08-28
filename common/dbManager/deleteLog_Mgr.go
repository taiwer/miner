package dbManager

import ()

type Inxxx interface {
	LoadDeleteDataFromDb()
}

type DeleteLogMgr struct {
	list         map[string]Inxxx
	loadDataTime int64 //加载数据的时间单位
}

var DeleteLogMgrObj *DeleteLogMgr

func init() {
	DeleteLogMgrObj = NewDeleteLogMgr()
}

func NewDeleteLogMgr() *DeleteLogMgr {
	obj := &DeleteLogMgr{}
	obj.list = make(map[string]Inxxx)
	return obj
}

func (s *DeleteLogMgr) loadDataProc(tableUpdatetime int64) {
	if s.loadDataTime == tableUpdatetime {
		return
	}
	for _, v := range s.list {
		v.LoadDeleteDataFromDb()
	}
	s.loadDataTime = tableUpdatetime
}

func (s *DeleteLogMgr) Addhandle(name string, xxx Inxxx) {
	s.list[name] = xxx
}
