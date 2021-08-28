package dbPlotOpr

import (
	"fmt"
	"go.uber.org/zap"
)

func inserNode() {
	fmt.Println("insert Node ...")
	groups := [0]Node{}
	for _, v := range groups {
		if _, err := AddNode(&v, nil); err != nil {
			zap.L().Error("AddNode", zap.String("err", err.Error()))
		}
	}
	fmt.Println("insert Node end", len(groups))
}

func inserNodeGroup() {
	fmt.Println("inserNodeGroup ...")
	groups := [2]NodeGroup{
		{Id: 1, Name: "默认服务器"},
		{Id: 2, Name: "默认转发服务器"},
	}
	for _, v := range groups {
		if _, err := AddNodeGroup(&v, nil); err != nil {
			zap.L().Error("NodeGroup", zap.String("err", err.Error()))
		}
	}
	fmt.Println("inserNodeGroup end", len(groups))
}
