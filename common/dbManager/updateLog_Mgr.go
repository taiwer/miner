package dbManager

import (
	"fmt"
	"github.com/taiwer/miner/common/dbPlotOpr"
	"go.uber.org/zap"
)

type UpdateLogMgr struct {
	CommonDeleteDataMgr
	loadDataTime int64 //加载数据的时间单位
}

var UpdateLogMgrObj *UpdateLogMgr

func init() {
	UpdateLogMgrObj = NewUpdateLogMgr()
}

func NewUpdateLogMgr() *UpdateLogMgr {
	obj := &UpdateLogMgr{}
	obj.Init()
	return obj
}

func (s *UpdateLogMgr) loadDataProc(tableUpdatetime int64) {
	if s.loadDataTime != 0 && tableUpdatetime >= 0 && tableUpdatetime <= s.loadDataTime {
		zap.L().Info("loadDataProc skip ", zap.Int64("tableUpdatetime", tableUpdatetime),
			zap.Int64("s.loadDataTime", s.loadDataTime),
		)
		return
	}
	if utime, b := s.loadDataFromDb(s.loadDataTime); b {
		s.loadDataTime = utime
		//zap.L().Info(fmt.Sprintf("loadDataProc \t loadDataTime:%d", s.loadDataTime))
	}
}

func (s *UpdateLogMgr) loadDataFromDb(unixUpdatedAt int64) (int64, bool) {
	var rUpdatedAt int64 = unixUpdatedAt
	var bRs bool = false
	items := dbPlotOpr.GetUpdateLogListByUpdateTime(unixUpdatedAt)
	if len(items) > 0 {
		zap.L().Info(fmt.Sprintf("UpdateLogMgr.loadDataFromDb Count:%d", len(items)))
	}
	for _, vn := range items {
		UpdatedAt := vn.UpdatedAt.Unix()
		if UpdatedAt > rUpdatedAt {
			rUpdatedAt = UpdatedAt
		}
		bRs = true
		v := &UpdateLog{UpdateTable: vn.UpdateTable}
		_ = s.AddItem(v, vn)
		zap.L().Info(fmt.Sprintf("Info dataBaseUpdate name:%s", "UpdateLog"),
			zap.String("UpdateTable", vn.UpdateTable),
			zap.Int64("UpdatedAt", vn.UpdatedAt.Unix()),
			zap.Int64("rUpdatedAt", rUpdatedAt),
		)
	}
	return rUpdatedAt, bRs
}

func (s *UpdateLogMgr) AddItem(item *UpdateLog, chgItem interface{}) error {
	_ = s.CommonItemMgr.AddItem(item, chgItem)
	return nil
}

func (s *UpdateLogMgr) GetItems() []*UpdateLog {
	dataSlice := s.CommonItemMgr.GetItems()
	List := make([]*UpdateLog, 0, len(dataSlice))
	for i := 0; i < len(dataSlice); i++ {
		List = append(List, (dataSlice[i]).(*UpdateLog))
	}
	return List
}

func (s *UpdateLogMgr) GetItemById(Id string) *UpdateLog {
	item := s.CommonItemMgr.GetItemById(Id)
	if item != nil {
		return item.(*UpdateLog)
	} else {
		return nil
	}
}
