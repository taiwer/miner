package dbManager

import (
	"fmt"
	"go.uber.org/zap"
)

func LogInfoUpdateItem(itemName string, id int64, updateTime int64, rUpdatetime int64) {
	zap.L().Info(fmt.Sprintf("Info dataBaseUpdate name:%s", itemName),
		zap.Int64("Id", id),
		zap.Int64("updateTime", updateTime),
		zap.Int64("rUpdatetime", rUpdatetime),
	)
}

func LogInfoAddItem(itemName string, item interface{}) {
	zap.L().Info(fmt.Sprintf("Info dataBaseAdd name:%s", itemName),
		zap.String("item", fmt.Sprintf("%v", item)),
	)
}
func LogInfo(msg string) {
	zap.L().Info(msg)
}
