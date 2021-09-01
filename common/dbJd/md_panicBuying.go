package dbJd

import (
	"fmt"
	"github.com/taiwer/miner/common/dbOrm"
	"go.uber.org/zap"
	"time"
)

//节点表
type PanicBuying struct {
	Id              int64     `json:"id"`
	UserName        string    `gorm:"unique;size(32)" json:"user_name"`          //用户名称
	ItemId          string    `gorm:"unique;size(32)" json:"item_id"`            //物品ID
	ItemName        string    `gorm:"unique;size(512)" json:"item_name"`         //物品名称
	Num             uint32    `gorm:"default:0" json:"num"`                      //购买数量
	LimitPrice      uint32    `gorm:"default:0" json:"limit_price"`              //限制价格
	Enable          bool      `gorm:"default:0" json:"enable"`                   //开关
	StartAt         time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP" json:"start_at"` //开始时间
	StopAt          time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP " json:"stop_at"` //结束时间
	dbOrm.TimeModel           //时间模型
}

func (n *PanicBuying) TableName() string {
	return TABLE_HEADER + "panic_buyings"
}

func (s *PanicBuying) InserTodb(args map[string]interface{}) (int64, error) {
	db := getDb()
	db = db.Create(s)
	return db.RowsAffected, db.Error
}

func (s *PanicBuying) InserOrUpdate(args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("item_id='%s'", s.ItemId)
	if count, _ := dbOrm.GetCount(s, where); count == 0 {
		return s.InserTodb(args)
	} else {
		newargs := make(map[string]interface{}, 10)
		newargs["item_name"] = s.ItemName
		newargs["num"] = s.Num
		newargs["limit_price"] = s.LimitPrice
		newargs["enable"] = s.Enable
		newargs["start_at"] = s.StartAt
		newargs["stop_at"] = s.StopAt
		RowsAffected, err := dbOrm.UpdateByWhere(s, "", "", where, newargs)
		return RowsAffected, err
	}
}

//===============
//add
func AddPanicBuying(u *PanicBuying, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelPanicBuyingById(Id int64) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.DelByWhere(&PanicBuying{}, where)
}

//del list
func DelPanicBuyingByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&PanicBuying{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetPanicBuyingByIp(ip string) (item *PanicBuying) {
	item = &PanicBuying{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("ip='%s'", ip)); err != nil {
		zap.L().Error("GetPanicBuyingById", zap.String("err", err.Error()))
	}
	return item
}

func GetPanicBuyingByName(name string) (item *PanicBuying) {
	item = &PanicBuying{}
	if n, err := dbOrm.GetByWhere(item, fmt.Sprintf("name='%s'", name)); err != nil {
		zap.L().Error("GetPanicBuyingByName", zap.String("err", err.Error()))
		return nil
	} else {
		if n == 0 {
			return nil
		}
	}
	return item
}

//get list
func GetPanicBuyingList(field ...string) (list []*PanicBuying) {
	if _, err := dbOrm.GetList(&PanicBuying{}, &list, field...); err != nil {
		zap.L().Error("GetPanicBuyinglist", zap.String("err", err.Error()))
	}
	return list
}

func GetPanicBuyingListByWhere(where string, field ...string) (list []*PanicBuying) {
	if _, err := dbOrm.GetListByWhere(&PanicBuying{}, &list, where, field...); err != nil {
		zap.L().Error("GetPanicBuyinglist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetPanicBuyingListByUpdateTime(updateTime int64) (list []*PanicBuying) {
	if _, err := dbOrm.GetListByUpdateTime(&PanicBuying{}, &list, updateTime); err != nil {
		zap.L().Error("GetPanicBuyinglist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

func GetFreePanicBuyingList() (list []*PanicBuying) {
	where := fmt.Sprintf("role=%d", 0)
	return GetPanicBuyingListByWhere(where)
}

//get  Page
func GetPanicBuyingListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []PanicBuying
	return dbOrm.GetListWithPage(&PanicBuying{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdatePanicBuyingById(id int64, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("id=%d", id)
	return dbOrm.UpdateByWhere(&PanicBuying{}, "", "", where, args)
}

func UpdatePanicBuyingByName(userName string, itemId string, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("user_name='%s' AND item_id='%s'", userName, itemId)
	return dbOrm.UpdateByWhere(&PanicBuying{}, "", "", where, args)
}

func UpdatePanicBuyingByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&PanicBuying{}, "", "", where, args)
}
