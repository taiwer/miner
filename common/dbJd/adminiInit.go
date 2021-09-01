package dbJd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/taiwer/miner/common/dbOrm"
	"github.com/taiwer/miner/common/myutils"
	"log"
	"strings"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
)

func getDb() *gorm.DB {
	return dbOrm.GetDb("")
}

func getLogDb() *gorm.DB {
	return dbOrm.GetDb("db_log")
}

func Syncdb(_force bool, skipTab ...string) {
	if _force {
		//db.DropTable(&Account{})
	}
	AutoMigrate()
	if _force {
		InserInitData(true)
		fmt.Println("database init is complete.\nPlease restart the application")
	}
}

type Inmodel interface {
	GetDbNameSpace() string
}

func dropTable(model Inmodel) {
	db := dbOrm.GetDb(model.GetDbNameSpace())
	if db != nil {
		db.DropTable(model)
	}
}

func autoMigrate(model Inmodel) {
	db := dbOrm.GetDb(model.GetDbNameSpace())
	if db != nil {
		db.Debug().AutoMigrate(model)
	}
}

//自动构建数据表
func AutoMigrate() {

	autoMigrate(&JdUser{})
	autoMigrate(&PanicBuying{})

}

func CreateTrigger(db *gorm.DB) {
	if db == nil {
		db = getDb()
	}
	createTrigger(db, (&JdUser{}).TableName(), "delete_name", "name")
	createTrigger(db, (&PanicBuying{}).TableName(), "delete_name", "name")
}

func InserInitData(_force bool) {
	if _force {
		CreatePROCEDURE(nil)
		CreateTrigger(nil) //穿件触发事件
		fmt.Println("database init is complete.\nPlease restart the application")
	}
}

func CreatePROCEDURE(db *gorm.DB) {
	if db == nil {
		db = getDb()
	}
	if strList, err := myutils.GetFileContentAsStringLines("conf/createProcedure.sql"); err == nil {
		sql := ""
		for i := 0; i < len(strList); i++ {
			if strList[i] != "--" {
				sql = fmt.Sprintf("%s\n%s", sql, strList[i])
				continue
			}
			if sql != "" {
				db = db.Exec(string(sql))
				fmt.Println("========+++++++++++=======", db.RowsAffected, db.Error)
				sql = ""
			}
		}
		if sql != "" {
			db = db.Exec(string(sql))
			fmt.Println("========+++++++++++=======", db.RowsAffected, db.Error)
			sql = ""
		}
	} else {
		log.Println("Err", err)
	}
}
func createTrigger(db *gorm.DB, tableName string, deleteIdFieldName, deleteField string) {
	triggerInsert := fmt.Sprintf("CREATE TRIGGER `%s_oninsert` AFTER INSERT ON `%s` FOR EACH ROW ", tableName, tableName)
	triggerUpdate := fmt.Sprintf("CREATE TRIGGER `%s_onupdate` AFTER UPDATE ON `%s` FOR EACH ROW ", tableName, tableName)
	triggerDelete := fmt.Sprintf("CREATE TRIGGER `%s_ondelete` AFTER DELETE ON `%s` FOR EACH ROW ", tableName, tableName)
	updateLogSqlList := make([]string, 0)
	updateLogSqlList = append(updateLogSqlList, fmt.Sprintf("insert into cl_update_logs (update_table,tag) VALUES ('%s',1) ON DUPLICATE KEY UPDATE tag=tag+1;", tableName))
	//triggerInsert
	{
		sql := fmt.Sprintf("%s begin %s end", triggerInsert, strings.Join(updateLogSqlList, "\r"))
		fmt.Println(sql)
		dbl := db.Exec(sql)
		fmt.Println("========+++++++++++=======", dbl.RowsAffected, dbl.Error)
	}
	//triggerUpdate
	{
		sql := fmt.Sprintf("%s begin %s end", triggerUpdate, strings.Join(updateLogSqlList, "\r"))
		fmt.Println(sql)
		dbl := db.Exec(sql)
		fmt.Println("========+++++++++++=======", dbl.RowsAffected, dbl.Error)
	}
	//triggerDelete
	if deleteIdFieldName != "" {
		deleteLogSqlList := make([]string, 0)
		deleteLogSqlList = append(deleteLogSqlList, fmt.Sprintf("insert into cl_delete_logs (delete_table,%s) VALUES ('%s',old.%s);", deleteIdFieldName, tableName, deleteField))
		sql := fmt.Sprintf("%s begin %s end", triggerDelete, strings.Join(deleteLogSqlList, "\r"))
		fmt.Println(sql)
		dbl := db.Exec(sql)
		fmt.Println("========+++++++++++=======", dbl.RowsAffected, dbl.Error)
	}
}
