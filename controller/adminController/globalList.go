package adminController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goplot/common/codes"
	"goplot/common/dbPlotOpr"
	"goplot/controller/baseController"
	"goplot/models/httpMaster"
	"net/http"
	"strconv"
)

type GlobalListController struct {
	baseController.BaseController
}

func (s *GlobalListController) Command(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	if v, ok := args["cmd"]; ok {
		cmd, _ := v.(string)
		switch cmd {
		case "update_server_name":
			s.RespData(c, http.StatusOK, 200, "success")
		default:
			s.RespData(c, http.StatusOK, 0, "default cmd:"+cmd)
		}
	} else {
		s.RespData(c, http.StatusOK, 0, "cmd is null")
	}
}

func (s *GlobalListController) Create(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.HttpRequest("globalList", "add", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespData(c, http.StatusOK, 200, "success")
	} else {
		s.RespData(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *GlobalListController) Del(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Del",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.HttpRequest("globalList", "del", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespData(c, http.StatusOK, 200, "success")
	} else {
		s.RespData(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *GlobalListController) Update(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Update",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.HttpRequest("globalList", "mod", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespMessage(c, http.StatusOK, 200, "success")
	} else {
		s.RespMessage(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *GlobalListController) List(c *gin.Context) {
	code := codes.SUCCESS
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	sort := c.Query("sort")
	sortOrder := c.Query("sortOrder")
	search := c.Query("search")
	fiter := c.Query("fiter")
	searchFiter := ""
	if search != "" {
		searchFiter = fmt.Sprintf("name LIKE '%%%s%%' or text LIKE '%%%s%%'", search, search)
	}
	if searchFiter != "" {
		if fiter != "" {
			fiter = fmt.Sprintf("(%s) AND (%s)", fiter, searchFiter)
		} else {
			fiter = searchFiter
		}
	}
	list := dbPlotOpr.GetGlobalListWithPage(limit, offset, sort, sortOrder, fiter)
	s.RespData(c, http.StatusOK, code, list)
}
