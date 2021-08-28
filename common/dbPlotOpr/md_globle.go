package dbPlotOpr

import (
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbOrm"
)

//节点表
type Global struct {
	Name            string `gorm:"size:255;primary_key" json:"name"`
	Val             int64  `gorm:"default:0" json:"val"`  //玩家IP
	Text            string `gorm:"size:255" json:"text"`  //字符串内容
	Descs           string `gorm:"size:255" json:"descs"` //描述
	dbOrm.TimeModel        //时间模型
}

func (n *Global) TableName() string {
	return TABLE_HEADER + "globles"
}

func (s *Global) InserTodb(args map[string]interface{}) (int64, error) {
	db := getDb()
	db.Create(s)
	return db.RowsAffected, db.Error
}

//===============
//add
func AddGlobal(u *Global, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelGlobalById(Id int64) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.DelByWhere(&Global{}, where)
}

//del list
func DelGlobalByIds(Ids []string) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("name='%s'", v)
		} else {
			where = fmt.Sprintf("%s or name='%s'", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&Global{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetGlobalByName(name string) (item *Global) {
	item = &Global{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("name='%s'", name)); err != nil {
		zap.L().Error("GetGlobalById", zap.String("err", err.Error()))
	}
	return item
}

//get list
func GetGlobalList() (list []*Global) {
	if _, err := dbOrm.GetList(&Global{}, &list); err != nil {
		zap.L().Error("GetGloballist", zap.String("err", err.Error()))
	}
	return list
}

func GetGlobalListByWhere(where string) (list []*Global) {
	if _, err := dbOrm.GetListByWhere(&Global{}, &list, where); err != nil {
		zap.L().Error("GetGloballist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetGlobalListByUpdateTime(updateTime int64) (list []*Global) {
	if _, err := dbOrm.GetListByUpdateTime(&Global{}, &list, updateTime); err != nil {
		zap.L().Error("GetGloballist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//get  Page
func GetGlobalListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []Global
	return dbOrm.GetListWithPage(&Global{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdateGlobalByName(name string, args map[string]interface{}) (int64, error) {
	obj := &Global{Name: name}
	dbOrm.SetObjValue(obj, args)
	return dbOrm.InsertOrUpdate(obj, args)
}

func UpdateGlobalByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&Global{}, "", "", where, args)
}

func GetGlobalText(globleName string) string {
	globle := GetGlobalByName(globleName)
	fmt.Println("globle", globle)
	if globle != nil {
		return globle.Text
	}
	return ""
}
