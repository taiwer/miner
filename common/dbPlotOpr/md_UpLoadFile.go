package dbPlotOpr

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/taiwer/miner/common/dbOrm"
	"go.uber.org/zap"
)

type UpLoadFile struct {
	Id        int64  `json:"id"`
	FileName  string `gorm:"size:255" json:"file_name"`  //客户创建订单时，生成的实例订单，这里作为主键
	FileClass string `gorm:"size:255" json:"file_class"` //上传文件的类型
	Version   string `gorm:"size:255" json:"version"`    //版本
	Descs     string `gorm:"size:255" json:"descs"`      //描述
	dbOrm.TimeModel
}

func (n *UpLoadFile) TableName() string {
	return TABLE_HEADER + "up_load_files"
}

func (s *UpLoadFile) InserTodb(args map[string]interface{}) (int64, error) {
	if GetUpLoadFileByFileName(s.FileName) == nil { //判断是否存在
		db := getDb()
		db.Create(s)
		return db.RowsAffected, db.Error
	}
	return 0, nil
}

func (s *UpLoadFile) PareseData() {

}

//===============
//add
func AddUpLoadFile(u *UpLoadFile, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelUpLoadFileById(Id int64) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.DelByWhere(&UpLoadFile{}, where)
}

//del list
func DelUpLoadFileByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&UpLoadFile{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetUpLoadFileById(id int64) (item *UpLoadFile) {
	item = &UpLoadFile{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("id=%d", id)); err != nil {
		zap.L().Error("GetUpLoadFileById", zap.String("err", err.Error()))
	}
	return item
}

func GetUpLoadFileByFileName(fileName string) (item *UpLoadFile) {
	item = &UpLoadFile{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("file_name='%s'", fileName)); err != nil {
		zap.L().Error("GetUpLoadFileByFileName", zap.String("err", err.Error()))
		return nil
	}
	return item
}

func GetUpLoadFileCount() (int64, error) {
	return dbOrm.GetCount(&UpLoadFile{}, "")
}

//get list
func GetUpLoadFileList() (list []*UpLoadFile) {
	if _, err := dbOrm.GetList(&UpLoadFile{}, &list); err != nil {
		zap.L().Error("GetUpLoadFilelist", zap.String("err", err.Error()))
	}
	return list
}

func GetUpLoadFileListByWhere(where string) (list []*UpLoadFile) {
	if _, err := dbOrm.GetListByWhere(&UpLoadFile{}, &list, where); err != nil {
		zap.L().Error("GetUpLoadFilelist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetUpLoadFileListByUpdateTime(updateTime int64) (list []*UpLoadFile) {
	if _, err := dbOrm.GetListByUpdateTime(&UpLoadFile{}, &list, updateTime); err != nil {
		zap.L().Error("GetUpLoadFilelist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//get  Page
func GetUpLoadFileListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []UpLoadFile
	return dbOrm.GetListWithPage(&UpLoadFile{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdateUpLoadFileById(id int64, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("id=%d", id)
	return dbOrm.UpdateByWhere(&UpLoadFile{}, "", "", where, args)
}

//更新用户
func UpdateUpLoadFileByName(name string, args orm.Params) (int64, error) {
	zap.L().Info("UpdateUpLoadFileByName")
	where := fmt.Sprintf("account='%s'", name)
	return dbOrm.UpdateByWhere(&UpLoadFile{}, "", "", where, args)
}

func GetUpLoadFileByUpLoadFileName(fileName string) *UpLoadFile {
	item := UpLoadFile{
		FileName: fileName,
	}
	_, err := dbOrm.GetByWhere(&item, "")
	if err == nil {
		return &item
	} else {
		return nil
	}
}
