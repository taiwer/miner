package dbJd

import (
	"fmt"
	"github.com/taiwer/miner/common/dbOrm"
	"go.uber.org/zap"
)

//节点表
type JdUser struct {
	Id              int64
	Username        string `gorm:"unique;size(32)" json:"username"`
	Password        string `gorm:"size(32)" json:"password"`
	Nickname        string `gorm:"unique;size(32)" json:"nickname"`
	Token           string `gorm:"unique;size(2000)" json:"token"`
	dbOrm.TimeModel        //时间模型
}

func (n *JdUser) TableName() string {
	return TABLE_HEADER + "jd_users"
}

func (s *JdUser) InserTodb(args map[string]interface{}) (int64, error) {
	db := getDb()
	db.Create(s)
	return db.RowsAffected, db.Error
}

func (s *JdUser) InserOrUpdate(args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("Username='%s'", s.Username)
	if count, _ := dbOrm.GetCount(s, where); count == 0 {
		return s.InserTodb(args)
	} else {
		newargs := make(map[string]interface{}, 10)
		newargs["password"] = s.Password
		newargs["nickname"] = s.Nickname
		newargs["token"] = s.Token
		RowsAffected, err := dbOrm.UpdateByWhere(s, "", "", where, newargs)
		return RowsAffected, err
	}
}

//===============
//add
func AddJdUser(u *JdUser, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelJdUserById(Id int64) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.DelByWhere(&JdUser{}, where)
}

//del list
func DelJdUserByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&JdUser{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetJdUserByIp(ip string) (item *JdUser) {
	item = &JdUser{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("ip='%s'", ip)); err != nil {
		zap.L().Error("GetJdUserById", zap.String("err", err.Error()))
	}
	return item
}

func GetJdUserByName(name string) (item *JdUser) {
	item = &JdUser{}
	if n, err := dbOrm.GetByWhere(item, fmt.Sprintf("name='%s'", name)); err != nil {
		zap.L().Error("GetJdUserByName", zap.String("err", err.Error()))
		return nil
	} else {
		if n == 0 {
			return nil
		}
	}
	return item
}

//get list
func GetJdUserList(field ...string) (list []*JdUser) {
	if _, err := dbOrm.GetList(&JdUser{}, &list, field...); err != nil {
		zap.L().Error("GetJdUserlist", zap.String("err", err.Error()))
	}
	return list
}

func GetJdUserListByWhere(where string, field ...string) (list []*JdUser) {
	if _, err := dbOrm.GetListByWhere(&JdUser{}, &list, where, field...); err != nil {
		zap.L().Error("GetJdUserlist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetJdUserListByUpdateTime(updateTime int64) (list []*JdUser) {
	if _, err := dbOrm.GetListByUpdateTime(&JdUser{}, &list, updateTime); err != nil {
		zap.L().Error("GetJdUserlist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

func GetFreeJdUserList() (list []*JdUser) {
	where := fmt.Sprintf("role=%d", 0)
	return GetJdUserListByWhere(where)
}

//get  Page
func GetJdUserListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []JdUser
	return dbOrm.GetListWithPage(&JdUser{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdateJdUserByName(uname string, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("username='%s'", uname)
	return dbOrm.UpdateByWhere(&JdUser{}, "", "", where, args)
}

func UpdateJdUserByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&JdUser{}, "", "", where, args)
}
