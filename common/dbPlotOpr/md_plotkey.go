package dbPlotOpr

import (
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbOrm"
)

//节点表
type PlotKey struct {
	Id                  int64  `json:"id"`
	UserName            string `gorm:"size:255;" json:"user_name"`              //节点名称第一次为空 登录后由服务器发送一个新名称
	KeyName             string `gorm:"size:255;" json:"key_name"`               //节点名称第一次为空 登录后由服务器发送一个新名称
	PlotFarmerPublicKey string `gorm:"size:255;" json:"plot_farmer_public_key"` //农民key
	PlotPoolPublicKey   string `gorm:"size:255;" json:"plot_pool_public_key"`   //矿池key
	dbOrm.TimeModel            //时间模型
}

func (n *PlotKey) TableName() string {
	return TABLE_HEADER + "plot_keys"
}

func (s *PlotKey) InserTodb(args map[string]interface{}) (int64, error) {
	db := getDb()
	db = db.Create(s)
	return db.RowsAffected, db.Error
}

//===============
//add
func AddPlotKey(u *PlotKey, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}



func DelPlotKeyById(id int64) (int64, error) {
	where := fmt.Sprintf("id='%d'", id)
	return dbOrm.DelByWhere(&PlotKey{}, where)
}

func DelPlotKeyByName(name string) (int64, error) {
	where := fmt.Sprintf("name='%s'", name)
	return dbOrm.DelByWhere(&PlotKey{}, where)
}

//del list
func DelPlotKeyByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&PlotKey{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetPlotKeyByIp(ip string) (item *PlotKey) {
	item = &PlotKey{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("ip='%s'", ip)); err != nil {
		zap.L().Error("GetPlotKeyById", zap.String("err", err.Error()))
	}
	return item
}

func GetPlotKeyByName(name string) (item *PlotKey) {
	item = &PlotKey{}
	if n, err := dbOrm.GetByWhere(item, fmt.Sprintf("name='%s'", name)); err != nil {
		zap.L().Error("GetPlotKeyByName", zap.String("err", err.Error()))
		return nil
	} else {
		if n == 0 {
			return nil
		}
	}
	return item
}

//get list
func GetPlotKeyList(field ...string) (list []*PlotKey) {
	if _, err := dbOrm.GetList(&PlotKey{}, &list, field...); err != nil {
		zap.L().Error("GetPlotKeylist", zap.String("err", err.Error()))
	}
	return list
}

func GetPlotKeyListByWhere(where string, field ...string) (list []*PlotKey) {
	if _, err := dbOrm.GetListByWhere(&PlotKey{}, &list, where, field...); err != nil {
		zap.L().Error("GetPlotKeylist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetPlotKeyListByUpdateTime(updateTime int64) (list []*PlotKey) {
	if _, err := dbOrm.GetListByUpdateTime(&PlotKey{}, &list, updateTime); err != nil {
		zap.L().Error("GetPlotKeylist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

func GetFreePlotKeyList() (list []*PlotKey) {
	where := fmt.Sprintf("role=%d", 0)
	return GetPlotKeyListByWhere(where)
}

//get  Page
func GetPlotKeyListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []PlotKey
	return dbOrm.GetListWithPage(&PlotKey{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdatePlotKeyById(Id int64, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.UpdateByWhere(&PlotKey{}, "", "", where, args)
}

func UpdatePlotKeyByName(name string, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("name='%s'", name)
	return dbOrm.UpdateByWhere(&PlotKey{}, "", "", where, args)
}

func UpdatePlotKeyByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&PlotKey{}, "", "", where, args)
}
