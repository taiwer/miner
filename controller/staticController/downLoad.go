package staticController

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goplot/common/dbPlotOpr"
	"goplot/controller/baseController"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type DownLoadController struct {
	baseController.BaseController
}

func (s *DownLoadController) DownLoad(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Del",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	os := c.Request.Form.Get("os")
	//arch := c.Request.Form.Get("arch")
	fileName := c.Request.Form.Get("fileName")
	if fileName == "" {
		s.RespData(c, http.StatusNotFound, 0, "err fileName is null")
		return
	}
	if os == "Darwin" {
		fileName = strings.ReplaceAll(fileName, ".tar.gz", "_darwin_arm64.tar.gz")
	}
	if path.Ext(fileName) == ".sh" {
		destFile := s.GetUpLoadFilePath() + fileName
		if buf, err := ioutil.ReadFile(destFile); err == nil {
			bufs := strings.ReplaceAll(string(buf), "10.211.55.2", dbPlotOpr.GetGlobalText("WebHost"))
			c.Writer.Write([]byte(bufs))
		} else {
			s.RespData(c, http.StatusNotFound, 0, "404")
		}
		return
	}
	destFile := s.GetUpLoadFilePath() + fileName
	zap.L().Info("DownLoad File " + destFile)
	c.File(destFile)
}
