package dbManager

import (
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbPlotOpr"
)

type PlotDiskMgr struct {
	CommonDeleteDataMgr
	loadDataTime int64 //加载数据的时间单位
}

var PlotDiskMgrObj *PlotDiskMgr

func init() {
	PlotDiskMgrObj = NewPlotDiskMgr()
}

func NewPlotDiskMgr() *PlotDiskMgr {
	obj := &PlotDiskMgr{}
	obj.Init()
	return obj
}

func (s *PlotDiskMgr) loadDataProc(tableUpdatetime int64) {
	if s.loadDataTime != 0 && tableUpdatetime >= 0 && tableUpdatetime <= s.loadDataTime {
		zap.L().Info("loadDataProc skip ", zap.Int64("tableUpdatetime", tableUpdatetime),
			zap.Int64("s.loadDataTime", s.loadDataTime),
		)
		return
	}
	if utime, b := s.loadDataFromDb(s.loadDataTime); b {
		s.loadDataTime = utime
		zap.L().Info(fmt.Sprintf("UpdatedAt loadDataTime:%d", s.loadDataTime))
	}
}

func (s *PlotDiskMgr) loadDataFromDb(unixUpdatedAt int64) (int64, bool) {
	var rUpdatetime int64 = unixUpdatedAt
	var bRs bool = false
	items := dbPlotOpr.GetPlotDiskListByUpdateTime(unixUpdatedAt)
	if len(items) > 0 {
		LogInfo(fmt.Sprintf("PlotDiskMgr.loadDataFromDb Count:%d", len(items)))
	}
	for _, vn := range items {
		updatetime := vn.UpdatedAt.Unix()
		if updatetime > rUpdatetime {
			rUpdatetime = updatetime
		}
		bRs = true
		v := &PlotDisk{Id: vn.Id}
		s.AddItem(v, vn)
		//设置当前节点分组需更新标志的值
		zap.L().Info(fmt.Sprintf("NodeId:%s Mount:%s  updatetime:%d (load PlotDisks)unixupdatetime:%d",
			vn.NodeId, vn.Mount, vn.UpdatedAt.Unix(), unixUpdatedAt))
	}
	if s.Length() > 0 {
		s.loadDeleteDataFromDb(s, new(dbPlotOpr.PlotDisk))
	}
	return rUpdatetime, bRs
}

func (s *PlotDiskMgr) LoadDeleteDataFromDb() {
	s.loadDeleteDataFromDb(s, &dbPlotOpr.PlotDisk{})
}

func (s *PlotDiskMgr) AddItem(item *PlotDisk, chgItem interface{}) error {
	s.CommonItemMgr.AddItem(item, chgItem)
	return nil
}

func (s *PlotDiskMgr) GetItems() []*PlotDisk {
	dataSlice := s.CommonItemMgr.GetItems()
	List := make([]*PlotDisk, 0, len(dataSlice))
	for i := 0; i < len(dataSlice); i++ {
		List = append(List, (dataSlice[i]).(*PlotDisk))
	}
	return List
}

func (s *PlotDiskMgr) GetItemByNodeId(nodeId string, mount string) *PlotDisk {
	items := s.CommonItemMgr.GetItems()
	for _, v := range items {
		item := v.(*PlotDisk)
		if item.NodeId == nodeId && item.Mount == mount {
			return item
		}
	}
	return nil
}

func (s *PlotDiskMgr) GetItemsByNodeId(nodeId string) []*PlotDisk {
	dataSlice := s.CommonItemMgr.GetItems()
	List := make([]*PlotDisk, 0, len(dataSlice))
	for i := 0; i < len(dataSlice); i++ {
		item := (dataSlice[i]).(*PlotDisk)
		if item.NodeId == nodeId {
			List = append(List, item)
		}
	}
	return List
}
