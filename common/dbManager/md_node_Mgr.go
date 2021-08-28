package dbManager

import (
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbPlotOpr"
)

type NodeMgr struct {
	CommonDeleteDataMgr
	loadDataTime int64 //加载数据的时间单位
}

var NodeMgrObj *NodeMgr

func init() {
	NodeMgrObj = NewNodeMgr()
}

func NewNodeMgr() *NodeMgr {
	obj := &NodeMgr{}
	obj.Init()
	return obj
}

func (s *NodeMgr) loadDataProc(tableUpdatetime int64) {
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

func (s *NodeMgr) loadDataFromDb(unixUpdatedAt int64) (int64, bool) {
	var rUpdatetime int64 = unixUpdatedAt
	var bRs bool = false
	items := dbPlotOpr.GetNodeListByUpdateTime(unixUpdatedAt)
	if len(items) > 0 {
		LogInfo(fmt.Sprintf("NodeMgr.loadDataFromDb Count:%d", len(items)))
	}
	for _, vn := range items {
		updatetime := vn.UpdatedAt.Unix()
		if updatetime > rUpdatetime {
			rUpdatetime = updatetime
		}
		bRs = true
		v := &Node{Name: vn.Name}
		s.AddItem(v, vn)
		//设置当前节点分组需更新标志的值
		zap.L().Info(fmt.Sprintf("id:%x  updatetime:%d (load nodes)unixupdatetime:%d",
			vn.Ip, vn.UpdatedAt.Unix(), unixUpdatedAt))
	}
	if s.Length() > 0 {
		s.loadDeleteDataFromDb(s, new(dbPlotOpr.Node))
	}
	return rUpdatetime, bRs
}

func (s *NodeMgr) LoadDeleteDataFromDb() {
	s.loadDeleteDataFromDb(s, &dbPlotOpr.Node{})
}

func (s *NodeMgr) AddItem(item *Node, chgItem interface{}) error {
	s.CommonItemMgr.AddItem(item, chgItem)
	return nil
}

func (s *NodeMgr) GetItems() []*Node {
	dataSlice := s.CommonItemMgr.GetItems()
	List := make([]*Node, 0, len(dataSlice))
	for i := 0; i < len(dataSlice); i++ {
		List = append(List, (dataSlice[i]).(*Node))
	}
	return List
}

func (s *NodeMgr) GetItemById(Id string) *Node {
	item := s.CommonItemMgr.GetItemById(Id)
	if item != nil {
		return item.(*Node)
	} else {
		return nil
	}
}
