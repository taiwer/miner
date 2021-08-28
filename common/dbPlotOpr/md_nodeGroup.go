package dbPlotOpr

import (
	"fmt"
	"go.uber.org/zap"
	"goplot/common/dbOrm"
	"math/rand"
	//"github.com/astaxie/beego/validation"
)

//节点表
type NodeGroup struct {
	Id int64
	// 0表示为分组
	Name            string `gorm:"size:255"`  //组名称
	NodePortBegin   int64  `gorm:"default:0"` //
	NodePortEnd     int64  `gorm:"default:0"` //
	Role            int64  `gorm:"default:0"` //节点的角色:0，空闲；1表示转发节点；2代理节点
	Del             int64  `gorm:"default:0"` //
	dbOrm.TimeModel        //时间模型
	//Group   *Group `gorm:"rel(fk)"`

	NodeAllCount     int `gorm:"-"`
	NodeOnLineCount  int `gorm:"-"`
	NodeOffLineCount int `gorm:"-"`
}

func (n *NodeGroup) TableName() string {
	return TABLE_HEADER + "node_groups"
}

func (s *NodeGroup) ParamFormNodeGroup() {
	all, onLine, offLine := GetNodeCountInfo(s.Id)
	s.NodeAllCount = all
	s.NodeOnLineCount = onLine
	s.NodeOffLineCount = offLine
}

func NewFreeNodeGroup(role int64) (list []*NodeGroup) {
	AllNodeGroup := &NodeGroup{
		Id: -1,
	}
	freeGroup := &NodeGroup{
		Id: 0,
	}
	if role == 1 {
		AllNodeGroup.Name = "全部节点转发分组"
		AllNodeGroup.Role = 1
		freeGroup.Name = "空闲转发分组"
		freeGroup.Role = 1
	} else if role == 2 {
		AllNodeGroup.Name = "全部节点代理分组"
		AllNodeGroup.Role = 2
		freeGroup.Name = "空闲代理分组"
		freeGroup.Role = 2
	}

	list = append(list, AllNodeGroup) //添加空闲分组
	list = append(list, freeGroup)    //添加空闲分组
	return list
}

//===============
//add
func AddNodeGroup(u *NodeGroup, args map[string]interface{}) (int64, error) {
	return dbOrm.Add(u)
}

func DelNodeGroupById(Id int64) (int64, error) {
	where := fmt.Sprintf("id=%d", Id)
	return dbOrm.DelByWhere(&NodeGroup{}, where)
}

//del list
func DelNodeGroupByIds(Ids []int64) (int64, error) {
	where := ""
	for _, v := range Ids {
		if where == "" {
			where = fmt.Sprintf("id=%d", v)
		} else {
			where = fmt.Sprintf("%s or id=%d", where, v)
		}
	}
	if where != "" {
		return dbOrm.DelByWhere(&NodeGroup{}, where)
	}
	return 0, fmt.Errorf("Err where == ''")
}

func GetNodeGroupById(id int64) (item *NodeGroup) {
	item = &NodeGroup{}
	if _, err := dbOrm.GetByWhere(item, fmt.Sprintf("id=%d", id)); err != nil {
		zap.L().Error("GetNodeGroupById", zap.String("err", err.Error()))
	}
	return item
}

//get list
func GetNodeGroupList(role int64) (list []*NodeGroup) {
	where := "del=0"
	if role != 0 {
		where = fmt.Sprintf("%s and role=%d", where, role)
	}
	if _, err := dbOrm.GetListByWhere(&NodeGroup{}, &list, where); err != nil {
		zap.L().Error("GetNodeGrouplist", zap.String("err", err.Error()))
	}
	return list
}

func GetNodeGroupListByWhere(where string) (list []*NodeGroup) {
	if _, err := dbOrm.GetListByWhere(&NodeGroup{}, &list, where); err != nil {
		zap.L().Error("GetNodeGrouplist",
			zap.String("where", where),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//获取实例 通过更新时间
func GetNodeGroupListByUpdateTime(updateTime int64) (list []*NodeGroup) {
	if _, err := dbOrm.GetListByUpdateTime(&NodeGroup{}, &list, updateTime); err != nil {
		zap.L().Error("GetNodeGrouplist",
			zap.Int64("updateTime", updateTime),
			zap.String("err", err.Error()),
		)
	}
	return list
}

//get  Page
func GetNodeGroupListWithPage(pageSize int, offset int, sort string, sortOrder string, where string) (Page dbOrm.Page) {
	var vList []*NodeGroup
	return dbOrm.GetListWithPage(&NodeGroup{}, &vList, pageSize, offset, sort, sortOrder, where)
}

func UpdateNodeGroupById(id int64, args map[string]interface{}) (int64, error) {
	where := fmt.Sprintf("id=%d", id)
	return dbOrm.UpdateByWhere(&NodeGroup{}, "", "", where, args)
}

func UpdateNodeGroupByWhere(where string, args map[string]interface{}) (int64, error) {
	return dbOrm.UpdateByWhere(&NodeGroup{}, "", "", where, args)
}

func (s *NodeGroup) GetRandPort() int64 {
	if s.NodePortEnd > s.NodePortBegin {
		return rand.Int63n(s.NodePortEnd-s.NodePortBegin) + s.NodePortBegin
	} else {
		return 5000
	}
}
