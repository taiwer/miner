package dbManager

import (
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbOrm"
	"goplot/common/dbPlotOpr"
	"time"
)

type CommonDeleteDataMgr struct {
	CommonItemMgr
	deleteUpdatedAt int64
}
type InCommonMgr interface {
	loadDataProc(tableUpdatetime int64)
}

func (s *CommonDeleteDataMgr) OnDeleteEnvet(id interface{}, obj interface{}, deletetime int64) {
	for _, v := range s.onChangeEvent {
		v.RemoveItemById(id, obj, deletetime)
	}
}

func (s *CommonDeleteDataMgr) loadDeleteDataFromDb(Obj InCommonMgr, tableObj dbOrm.InTable) {
	var rUpdatedAt int64 = 0
	if s.deleteUpdatedAt == 0 {
		s.deleteUpdatedAt = int64(time.Now().Unix()) - 3600*24
	}
	list := dbPlotOpr.GetDeleteLogByUpdateTime(tableObj, s.deleteUpdatedAt)
	if len(list) > 0 {
		zap.L().Info(fmt.Sprintf("GetDeleteLogByUpdateTime count:%d deleteUpdatedAt:%d", len(list), time.Unix(s.deleteUpdatedAt, 0)))
		for _, vn := range list {
			UpdatedAt := vn.UpdatedAt.Unix()
			if UpdatedAt > rUpdatedAt {
				rUpdatedAt = UpdatedAt
			}
			if vn.DeleteId != 0 {

				delItem, _ := s.RemoveItemById(vn.DeleteId, vn.UpdatedAt.Unix())
				s.OnDeleteEnvet(vn.DeleteId, delItem, vn.UpdatedAt.Unix())
			} else {
				delItem, _ := s.RemoveItemById(vn.DeleteName, vn.UpdatedAt.Unix())
				s.OnDeleteEnvet(vn.DeleteName, delItem, vn.UpdatedAt.Unix())
			}
		}
		if rUpdatedAt > 0 {
			s.deleteUpdatedAt = rUpdatedAt
		}
	}
}

func (s *CommonDeleteDataMgr) LoadDbConfigLoop(Obj InCommonMgr) {
	/*动态加载数据中的内容，这里不科学，需要优化*/
	for {
		Obj.loadDataProc(0)
		time.Sleep(5 * time.Second)
	}
}

func (s *CommonDeleteDataMgr) Run(Obj InCommonMgr) {
	/*
		这里需要优化，添加memecached
	*/
	Obj.loadDataProc(0) //先加载一次
	go s.LoadDbConfigLoop(Obj)
}
