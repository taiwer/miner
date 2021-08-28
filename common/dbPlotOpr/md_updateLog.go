package dbPlotOpr

import (
	"goplot/common/dbOrm"
	"log"
	//"github.com/astaxie/beego/validation"
)

//节点表
type UpdateLog struct {
	UpdateTable     string `gorm:"size:100;primary_key"` //被删除项目的表名称
	Tag             int
	dbOrm.TimeModel //时间模型
}

func (n *UpdateLog) TableName() string {
	return TABLE_HEADER + "update_logs"
}

func NewUpdateLog() *DeleteLog {
	return &DeleteLog{}
}

//获取实例 通过更新时间
func GetUpdateLogListByUpdateTime(updateTime int64) (list []*UpdateLog) {
	if _, err := dbOrm.GetListByUpdateTime(&UpdateLog{}, &list, updateTime); err == nil {
	} else {
		log.Println("Err:", err)
	}
	return list
}
