package dbPlotOpr

import (
	"strconv"
)

const TABLE_HEADER = "cl_"

const CHART_FORMAT = "%x:%x;"

func ChartFormat(index int, value uint32) string {
	return strconv.FormatInt(int64(index), 36) + ":" + strconv.FormatInt(int64(value), 36) + ";"
	//return fmt.Sprintf(CHART_FORMAT, index, value)
}
