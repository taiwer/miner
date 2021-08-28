package dbPlotOpr

import (
	"fmt"
	"github.com/taiwer/miner/common/dbOrm"
	"go.uber.org/zap"
)

//节点表
type PlotDisk struct {
	Id           int64  `json:"id"`
	Device       string `gorm:"size:255" json:"device"`          //磁盘的挂载设备名称
	Mount        string `gorm:"size:255;" json:"mount"`          //节点名称磁盘挂载点
	NodeId       string `gorm:"size:255;" json:"node_id"`        //节点名称
	Descs        string `gorm:"size:255;" json:"descs"`          //描述
	Online       bool   `gorm:"default:0" json:"online"`         //记录该节点当前是否在线,0表示离线，1表示在线
	Size         uint64 `gorm:"default:0" json:"size"`           //磁盘空间大小
	FreeSize     uint64 `gorm:"default:0" json:"free_size"`      //磁盘空间大小
	UsedPercent  uint32 `gorm:"default:0" json:"used_percent"`   //磁盘使用百分比
	PlotCount    uint32 `gorm:"default:0" json:"plot_count"`     //磁盘空间大小
	PlotMaxCount uint32 `gorm:"default:0" json:"plot_max_count"` //任务容量

	Priority            uint32 `gorm:"default:0" json:"priority"`               //设置 优先级
	DiskType            string `gorm:"size:255;" json:"disk_type"`              //磁盘类型
	PlotOff             bool   `gorm:"default:false" json:"plot_off"`           // 为true 节点将暂停转发数据
	PlotSize            uint32 `gorm:"default:0" json:"plot_size"`              //设置 任务大小默认32
	PlotThread          uint32 `gorm:"default:0" json:"plot_thread"`            //设置 任务线程 默认2
	PlotBucket          uint32 `gorm:"default:0" json:"plot_bucket"`            //设置 任务大小默认128
	PlotMemory          uint32 `gorm:"default:0" json:"plot_memory"`            //设置 任务内存使用量 4GB 4096
	PlotJobTmpSize      uint32 `gorm:"default:0" json:"plot_job_tmp_size"`      //设置 任务占用tmp空间大小 默认380Gb 此值可以控制单硬盘任务数量
	PlotInterval        uint32 `gorm:"default:0" json:"plot_interval"`          //设置 任务间隔时间
	PlotMaxJob          uint32 `gorm:"default:0" json:"plot_max_job"`           //设置 同时最大工作任务
	PlotFarmerPublicKey string `gorm:"size:255;" json:"plot_farmer_public_key"` //农民key
	PlotPoolPublicKey   string `gorm:"size:255;" json:"plot_pool_public_key"`   //矿池key
	PlotCountToday      uint32 `gorm:"-" json:"plot_count_today"`
	PlotCountYesterday  uint32 `gorm:"-" json:"plot_count_yesterday"`
	dbOrm.TimeModel            //时间模型
}

func (n *PlotDisk) TableName() string {
	return TABLE_HEADER + "plot_disks"
}

func (s *PlotDisk) InserTodb(args map[string]interface{}) (int64, error) {
	db := getDb()
	db.Create(s)
	return db.RowsAffected, db.Error
}

func (s *PlotDisk) InserOrUpdate(args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("mount='%s' and node_id='%s'", s.Mount, s.NodeId)
	if count, _ := dbOrm.GetCount(s, where); count == 0 {
		return s.InserTodb(args)
	} else {
		newargs := make(map[string]interface{}, 10)
		newargs["size"] = s.Size
		newargs["free_size"] = s.FreeSize
		newargs["used_percent"] = s.UsedPercent
		newargs["plot_count"] = s.PlotCount
		newargs["plot_max_count"] = s.PlotMaxCount
		RowsAffected, err := dbOrm.UpdateByWhere(s, "", "", where, newargs)
		return RowsAffected, err
	}
}

//===============
//add
func AddPlotDisk(u *PlotDisk, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelPlotDiskById(Id int64) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.DelByWhere(&PlotDisk{}, where)
}

//del list
func DelPlotDiskByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&PlotDisk{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetPlotDiskByIp(ip string) (item *PlotDisk) {
	item = &PlotDisk{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("ip='%s'", ip)); err != nil {
		zap.L().Error("GetPlotDiskById", zap.String("err", err.Error()))
	}
	return item
}

func GetPlotDiskByName(name string) (item *PlotDisk) {
	item = &PlotDisk{}
	if n, err := dbOrm.GetByWhere(item, fmt.Sprintf("name='%s'", name)); err != nil {
		zap.L().Error("GetPlotDiskByName", zap.String("err", err.Error()))
		return nil
	} else {
		if n == 0 {
			return nil
		}
	}
	return item
}

//get list
func GetPlotDiskList(field ...string) (list []*PlotDisk) {
	if _, err := dbOrm.GetList(&PlotDisk{}, &list, field...); err != nil {
		zap.L().Error("GetPlotDisklist", zap.String("err", err.Error()))
	}
	return list
}

func GetPlotDiskListByWhere(where string, field ...string) (list []*PlotDisk) {
	if _, err := dbOrm.GetListByWhere(&PlotDisk{}, &list, where, field...); err != nil {
		zap.L().Error("GetPlotDisklist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetPlotDiskListByUpdateTime(updateTime int64) (list []*PlotDisk) {
	if _, err := dbOrm.GetListByUpdateTime(&PlotDisk{}, &list, updateTime); err != nil {
		zap.L().Error("GetPlotDisklist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

func GetFreePlotDiskList() (list []*PlotDisk) {
	where := fmt.Sprintf("role=%d", 0)
	return GetPlotDiskListByWhere(where)
}

//get  Page
func GetPlotDiskListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []PlotDisk
	return dbOrm.GetListWithPage(&PlotDisk{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdatePlotDiskById(id int64, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("id=%d", id)
	return dbOrm.UpdateByWhere(&PlotDisk{}, "", "", where, args)
}

func UpdatePlotDiskByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&PlotDisk{}, "", "", where, args)
}
