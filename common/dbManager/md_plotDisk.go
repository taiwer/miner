package dbManager

import "go.uber.org/zap"

type PlotDisk struct {
	CommonItem
	Id     int64
	Mount  string `gorm:"size:255;primary_key" json:"mount"` //节点名称磁盘挂载点
	NodeId string `gorm:"size:255;primary_key" json:"name"`  //节点名称
	Descs  string `gorm:"size:255;" json:"descs"`            //描述

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
}

func NewPlotDisk() *PlotDisk {
	return &PlotDisk{}
}

func (s *PlotDisk) GetPK() interface{} {
	return s.Id
}

func (s *PlotDisk) SetValue(obj interface{}) bool {
	ret, changeData := s.CommonItem.SetValue(s, obj)
	if changeData != nil {
		zap.L().Info("PlotDisk SetValue changed",
			zap.Any("Pk", s.GetPK()),
			zap.Any("changeData", changeData),
		)
	}
	return ret
}
