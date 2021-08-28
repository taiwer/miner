package baseController

import (
	"github.com/gin-gonic/gin"
)

type PloterControllerParam struct {
	PloterName string `json:"ploter_name"`
	UserName   string `json:"user_name"`
	DataName   string `json:"data_name"`
}

//RespData 数据返回
type BasePlotController struct {
	BaseController
}

func (s *BasePlotController) ParseParam(c *gin.Context) *PloterControllerParam {
	param := &PloterControllerParam{}
	return param
}

func (s *BasePlotController) runNewPlot(c *gin.Context) {

}

func (s *BasePlotController) GetDataList(c *gin.Context) {

}

func (s *BasePlotController) GetDataListText(c *gin.Context) {

}

func (s *BasePlotController) DoKillPlot(c *gin.Context) {

}

func (s *BasePlotController) DoNewPlot(c *gin.Context) {

}

func (s *BasePlotController) DoRunShell(c *gin.Context) {

}
