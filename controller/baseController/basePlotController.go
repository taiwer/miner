package baseController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"goplot/common/dbOrm"
	"goplot/common/monitor/monitorServer"
	"net/http"
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
	err := c.ShouldBindBodyWith(param, binding.JSON)
	zap.L().Info("Index",
		zap.Any("request", c.Request.Form),
		zap.Any("param", param),
		zap.Error(err),
	)
	return param
}

func (s *BasePlotController) getPlotClient(param *PloterControllerParam) (*monitorServer.Client, error) {
	//if param.PloterName != "" && param.UserName != "" {
	if param.PloterName != "" {
		client := monitorServer.MonitorClientMgrObj.GetClientByPloterName(param.PloterName)
		if client == nil {
			return nil, fmt.Errorf("not find %s online", param.PloterName)
		}
		return client, nil
	} else {
		return nil, fmt.Errorf("PloterName:%s UserName:%s faild", param.PloterName, param.UserName)
	}
}

func (s *BasePlotController) runNewPlot(c *gin.Context) {
	param := s.ParseParam(c)
	if param.PloterName != "" {
		plotClient, err := s.getPlotClient(param)
		if plotClient != nil {
			s.RespMessage(c, http.StatusOK, 200, "完毕")
		} else {
			s.RespFail(c, http.StatusOK, 0, err.Error())
		}
	}
}

func (s *BasePlotController) GetDataList(c *gin.Context) {
	param := s.ParseParam(c)
	if param.PloterName != "" {
		plotClient, err := s.getPlotClient(param)
		if plotClient != nil {
			switch param.DataName {
			case "get_job_list":
				var page dbOrm.Page
				jobList := plotClient.GetJobList()
				page.Rows = jobList
				page.Total = int64(len(jobList))
				s.RespData(c, http.StatusOK, 200, page)
			default:
				s.RespData(c, http.StatusOK, 200, "err default "+param.DataName)
			}
		} else {
			//var page dbOrm.Page
			//jobList := []interface{}{}
			//page.Rows = jobList
			//page.Total = int64(len(jobList))
			//s.RespData(c, http.StatusOK, 200, page)
			s.RespFail(c, http.StatusOK, 0, fmt.Sprintf("获取对象失败:%v", err))
		}
	}
}

func (s *BasePlotController) GetDataListText(c *gin.Context) {
	param := s.ParseParam(c)
	if param.PloterName != "" {
		plotPc, err := s.getPlotClient(param)
		if plotPc != nil {
			switch param.DataName {
			case "getNpcListText":
				s.RespData(c, http.StatusOK, 200, "gameObj.GetNpcListOfText()")
			default:
				s.RespData(c, http.StatusOK, 200, "err default "+param.DataName)
			}
		} else {
			s.RespFail(c, http.StatusOK, 0, err.Error())
		}
	}
}

func (s *BasePlotController) DoKillPlot(c *gin.Context) {
	param := s.ParseParam(c)
	if param.PloterName != "" {
		plotPc, err := s.getPlotClient(param)
		if plotPc != nil {
			item := struct {
				PlotId string `json:"plot_id"`
			}{}
			err := c.ShouldBindBodyWith(&item, binding.JSON)
			if err != nil {
				s.RespFail(c, http.StatusOK, 0, err.Error())
			}
			if err := plotPc.SendDelPlot(item.PlotId); err == nil {
				s.RespData(c, http.StatusOK, 200, "delete plot success")
			} else {
				s.RespFail(c, http.StatusOK, 0, fmt.Sprintf("delete plot faild %s", err.Error()))
			}
		} else {
			s.RespFail(c, http.StatusOK, 0, err.Error())
		}
	}
}

func (s *BasePlotController) DoNewPlot(c *gin.Context) {
	param := s.ParseParam(c)
	if param.PloterName != "" {
		plotPc, err := s.getPlotClient(param)
		if plotPc != nil {
			if err := plotPc.SendCreatePlot(); err == nil {
				s.RespData(c, http.StatusOK, 200, "delete plot success")
			} else {
				s.RespFail(c, http.StatusOK, 0, fmt.Sprintf("delete plot faild %s", err.Error()))
			}
		} else {
			s.RespFail(c, http.StatusOK, 0, err.Error())
		}
	}
}

func (s *BasePlotController) DoRunShell(c *gin.Context) {
	param := s.ParseParam(c)
	if param.PloterName != "" {
		plotPc, err := s.getPlotClient(param)
		if plotPc != nil {
			item := struct {
				ShellCommand string `json:"shell_command"`
			}{}
			err := c.ShouldBindBodyWith(&item, binding.JSON)
			if err != nil {
				s.RespFail(c, http.StatusOK, 0, err.Error())
			}
			if err := plotPc.SendRunShell(item.ShellCommand); err == nil {
				s.RespData(c, http.StatusOK, 200, "run shell success")
			} else {
				s.RespFail(c, http.StatusOK, 0, fmt.Sprintf("run shell faild %s", err.Error()))
			}
		} else {
			s.RespFail(c, http.StatusOK, 0, err.Error())
		}
	}
}
