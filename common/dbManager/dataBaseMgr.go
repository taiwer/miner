package dbManager

import (
	"goplot/common/dbPlotOpr"
	"time"
)

type DataBaseMgr struct {
}

var DataBaseMgrObj DataBaseMgr

func LoadDbConfig(mgr InCommonMgr, TableName string) {
	item := UpdateLogMgrObj.GetItemById(TableName)
	if item != nil {
		mgr.loadDataProc(item.UpdatedAt.Unix())
	} else {
		mgr.loadDataProc(0)
	}
}

func (s *DataBaseMgr) LoadDbConfig() {
	UpdateLogMgrObj.loadDataProc(-1)
	LoadDbConfig(NodeMgrObj, (&dbPlotOpr.Node{}).TableName())
	LoadDbConfig(DeleteLogMgrObj, (&dbPlotOpr.DeleteLog{}).TableName())
	LoadDbConfig(NodeMgrObj, (&dbPlotOpr.Node{}).TableName())
	LoadDbConfig(PlotDiskMgrObj, (&dbPlotOpr.PlotDisk{}).TableName())
}

func (s *DataBaseMgr) LoadDbConfigLoop() {
	/*动态加载数据中的内容，这里不科学，需要优化*/
	var clearTime int64 = 0
	for {
		if time.Now().UnixNano() > clearTime {
			clearTime = time.Now().UnixNano() + 1*int64(time.Hour)
			s.ClearTableDatas()
		}
		s.LoadDbConfig()
		time.Sleep(5 * time.Second)
	}
}

func (s *DataBaseMgr) ClearTableDatas() {
	_, _ = dbPlotOpr.ClearDeleteLogByUpdateLast(time.Hour * 2)
}

func (s *DataBaseMgr) Run() {
	/*
		这里需要优化，添加memecached
	*/

	DeleteLogMgrObj.Addhandle("NodeMgrObj", NodeMgrObj)
	//s.LoadDbConfig() //先加载一次
	go s.LoadDbConfigLoop()
}
