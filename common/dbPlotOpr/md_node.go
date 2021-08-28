package dbPlotOpr

import (
	"fmt"
	"github.com/taiwer/miner/common/dbOrm"
	"go.uber.org/zap"
	"strings"
	"time"
)

//节点表
type Node struct {
	Id                 int64     `json:"id"`
	Name               string    `gorm:"size:255;primary_key" json:"name"` //节点名称第一次为空 登录后由服务器发送一个新名称
	Descs              string    `gorm:"size:255;" json:"descs"`
	Ip                 string    `gorm:"size:255"  json:"ip"`                       //节点第一次登录时获取的节点ip
	UserName           string    `gorm:"size:255" json:"user_name"`                 //私有IP
	GroupId            uint32    `gorm:"default:0" json:"group_id"`                 //表示改节点所属分组，未知分组id为0
	Online             uint32    `gorm:"default:0" json:"online"`                   //记录该节点当前是否在线,0表示离线，1表示在线
	Version            string    `gorm:"size:255" json:"version"`                   //节点版本
	StartAt            time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP" json:"start_at"` //开始运行的时间戳  用于显示节点运行的时间
	UpdateProxyVersion string    `gorm:"size:255" json:"update_proxy_version"`      //更新版本
	UpdateProxyStatus  string    `gorm:"size:255" json:"update_proxy_status"`       //更新状态
	UpdateProxyErrInfo string    `gorm:"size:255" json:"update_proxy_err_info"`     //更新错误描述信息
	Off                bool      `gorm:"default:false" json:"off"`                  // 为true 节点将暂停转发数据

	PlotSize            uint32 `gorm:"default:0" json:"plot_size"`              //设置 任务大小默认32
	PlotThread          uint32 `gorm:"default:0" json:"plot_thread"`            //设置 任务线程 默认2
	PlotBucket          uint32 `gorm:"default:0" json:"plot_bucket"`            //设置 任务大小默认128
	PlotMemory          uint32 `gorm:"default:0" json:"plot_memory"`            //设置 任务内存使用量 4GB 4096
	PlotJobTmpSize      uint32 `gorm:"default:0" json:"plot_job_tmp_size"`      //设置 任务占用tmp空间大小 默认380Gb 此值可以控制单硬盘任务数量
	PlotInterval        uint32 `gorm:"default:0" json:"plot_interval"`          //设置 任务间隔时间
	PlotDstWorkNum      uint32 `gorm:"default:0" json:"plot_dst_work_num"`      //设置 同时工作的p盘磁盘个数
	PlotMaxJob          uint32 `gorm:"default:0" json:"plot_max_job"`           //设置 同时最大工作任务
	PlotKeyName         string `gorm:"size:255;" json:"plot_key_name"`          //key名称
	PlotFarmerPublicKey string `gorm:"size:255;" json:"plot_farmer_public_key"` //农民key
	PlotPoolPublicKey   string `gorm:"size:255;" json:"plot_pool_public_key"`   //矿池key

	RuingJobCount      uint32 `gorm:"-" json:"ruing_job_count"`
	PlotCountToday     uint32 `gorm:"-" json:"plot_count_today"`
	PlotCountYesterday uint32 `gorm:"-" json:"plot_count_yesterday"`
	InstallNodeShell   string `gorm:"-" json:"install_node_shell"`
	//wget "http://42.193.163.60:8002/static/download?fileName=installNode.sh" -O install.sh; chmod +x install.sh; ./install.sh
	dbOrm.TimeModel //时间模型
}

func (n *Node) TableName() string {
	return TABLE_HEADER + "nodes"
}

func (s *Node) InserTodb(args map[string]interface{}) (int64, error) {
	db := getDb()
	db = db.Create(s)
	return db.RowsAffected, db.Error
}

func (s *Node) Update() (int64, error) {
	db := getDb()
	db = db.Model(s).Updates(s)
	return db.RowsAffected, db.Error
}

//===============
//add
func AddNode(u *Node, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelNodeByName(name string) (int64, error) {
	where := fmt.Sprintf("name='%s'", name)
	return dbOrm.DelByWhere(&Node{}, where)
}

//del list
func DelNodeByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&Node{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetNodeByIp(ip string) (item *Node) {
	item = &Node{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("ip='%s'", ip)); err != nil {
		zap.L().Error("GetNodeById", zap.String("err", err.Error()))
	}
	return item
}

func GetNodeByName(name string) (item *Node) {
	item = &Node{}
	if n, err := dbOrm.GetByWhere(item, fmt.Sprintf("name='%s'", name)); err != nil {
		zap.L().Error("GetNodeByName", zap.String("err", err.Error()))
		return nil
	} else {
		if n == 0 {
			return nil
		}
	}
	return item
}

//get list
func GetNodeList(field ...string) (list []*Node) {
	if _, err := dbOrm.GetList(&Node{}, &list, field...); err != nil {
		zap.L().Error("GetNodelist", zap.String("err", err.Error()))
	}
	return list
}

func GetNodeListByWhere(where string, field ...string) (list []*Node) {
	if _, err := dbOrm.GetListByWhere(&Node{}, &list, where, field...); err != nil {
		zap.L().Error("GetNodelist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetNodeListByUpdateTime(updateTime int64) (list []*Node) {
	if _, err := dbOrm.GetListByUpdateTime(&Node{}, &list, updateTime); err != nil {
		zap.L().Error("GetNodelist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

func GetFreeNodeList() (list []*Node) {
	where := fmt.Sprintf("role=%d", 0)
	return GetNodeListByWhere(where)
}

//get  Page
func GetNodeListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []Node
	return dbOrm.GetListWithPage(&Node{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func GetNodeCountInfo(groupId int64) (all, onLine, offline int) {
	Node := GetNodeListByWhere(fmt.Sprintf("group_id=%d", groupId), "online")
	all = 0
	onLine = 0
	offline = 0
	for _, v := range Node {
		all++
		if v.Online == 1 {
			onLine++
		} else {
			offline++
		}
	}
	return all, onLine, offline
}

func UpdateNodeByName(name string, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("name='%s'", name)
	return dbOrm.UpdateByWhere(&Node{}, "", "", where, args)
}

func UpdateNodeByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&Node{}, "", "", where, args)
}

func SetNodeOnlineAll() {
	SetNodeOnline(true, nil...)
}

func SetNodeOfflineAll() {
	SetNodeOnline(false, nil...)
}

func SetNodeOnline(online bool, names ...string) (int64, error) {
	where := ""
	if len(names) > 0 {
		ips := strings.Join(names, "','")
		where = fmt.Sprintf("name in('%s')", ips)
	}
	args := make(map[string]interface{}, 0)
	if online {
		args["online"] = 1
	} else {
		args["online"] = 0
	}
	if n, err := dbOrm.UpdateByWhere(&Node{}, "", "", where, args); err == nil {
		zap.L().Info("SetNodeUnOnline", zap.Int64("num", n))
		return n, err
	} else {
		zap.L().Error("SetNodeUnOnline", zap.String("err", err.Error()))
		return n, err
	}
}
