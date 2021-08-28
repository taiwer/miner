package dbManager

import "go.uber.org/zap"

type Node struct {
	CommonItem
	Name                string `gorm:"size:255;primary_key" json:"name"`        //节点名称第一次为空 登录后由服务器发送一个新名称
	Ip                  string `gorm:"size:255"  json:"ip"`                     //节点第一次登录时获取的节点ip
	UserName            string `gorm:"size:255" json:"user_name"`               //私有IP
	PlotSize            uint32 `gorm:"default:0" json:"plot_size"`              //设置 任务大小默认32
	PlotThread          uint32 `gorm:"default:0" json:"plot_thread"`            //设置 任务线程 默认2
	PlotBucket          uint32 `gorm:"default:0" json:"plot_bucket"`            //设置 任务大小默认128
	PlotMemory          uint32 `gorm:"default:0" json:"plot_memory"`            //设置 任务内存使用量 4GB 4096
	PlotJobTmpSize      uint32 `gorm:"default:0" json:"plot_job_tmp_size"`      //设置 任务占用tmp空间大小 默认380Gb 此值可以控制单硬盘任务数量
	PlotInterval        uint32 `gorm:"default:0" json:"plot_interval"`          //设置 任务间隔时间
	PlotDstWorkNum      uint32 `gorm:"default:0" json:"plot_dst_work_num"`      //设置 同时工作的p盘磁盘个数
	PlotMaxJob          uint32 `gorm:"default:0" json:"plot_max_job"`           //设置 同时工作的p盘磁盘个数
	PlotKeyName         string `gorm:"size:255;" json:"plot_key_name"`          //key名称
	PlotFarmerPublicKey string `gorm:"size:255;" json:"plot_farmer_public_key"` //农民key
	PlotPoolPublicKey   string `gorm:"size:255;" json:"plot_pool_public_key"`   //矿池key
}

func NewNode() *Node {
	return &Node{}
}

func (s *Node) GetPK() interface{} {
	return s.Name
}

func (s *Node) SetValue(obj interface{}) bool {
	ret, changeData := s.CommonItem.SetValue(s, obj)
	if changeData != nil {
		zap.L().Info("Node SetValue changed",
			zap.Any("Pk", s.GetPK()),
			zap.Any("changeData", changeData),
		)
	}
	return ret
}
