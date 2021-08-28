package dbPlotOpr

import (
	"fmt"
	"github.com/taiwer/miner/common/dbOrm"
	"log"
	"time"
	//"github.com/astaxie/beego/validation"
)

//节点表
type DeleteLog struct {
	Id          int64
	DeleteId    int64  `gorm:"default:0"` //删除的行Id
	DeleteName  string `gorm:"size:255"`  //被删除行的主键字符串Id
	DeleteTable string `gorm:"size:255"`  //被删除项目的表名称
	dbOrm.TimeModel
}

func (n *DeleteLog) TableName() string {
	return TABLE_HEADER + "delete_logs"
}

func NewDeleteLog() *DeleteLog {
	return &DeleteLog{}
}

//获取实例 通过更新时间
func GetDeleteLogByUpdateTime(deleteTableIn dbOrm.InTable, updateTime int64) (list []*DeleteLog) {
	tableName := deleteTableIn.TableName()
	where := fmt.Sprintf("unix_timestamp('%s') < unix_timestamp(updated_at) and delete_table = '%s'", time.Unix(updateTime, 0), tableName)
	if _, err := dbOrm.GetListByWhere(&DeleteLog{}, &list, where); err == nil {
	} else {
		log.Println("Err:", err)
	}
	return list
}

func ClearDeleteLogByUpdateLast(d time.Duration) (int64, error) {
	return dbOrm.ClearByUpdateLast(&DeleteLog{}, d)
}
