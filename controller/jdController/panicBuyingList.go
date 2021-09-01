package jdController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/codes"
	"github.com/taiwer/miner/common/dbJd"
	"github.com/taiwer/miner/controller/baseController"
	"github.com/taiwer/miner/models/httpMaster"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type PanicBuyingListController struct {
	baseController.BaseController
}

func (s *PanicBuyingListController) Command(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	if v, ok := args["cmd"]; ok {
		jdUsername := s.GetUserName(c)
		if v, ok := args["jd_user_name"]; ok {
			jdUsername = v.(string)
		}
		cmd, _ := v.(string)
		switch cmd {
		case "get_PanicBuying_list":
			s.RespData(c, http.StatusOK, 0, jdUsername)
		default:
			s.RespData(c, http.StatusOK, 0, "default cmd:"+cmd)
		}
	} else {
		s.RespData(c, http.StatusOK, 0, "cmd is null")
	}
}

func (s *PanicBuyingListController) Create(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	args["user_name"] = s.GetUserName(c)
	ret, _ := httpMaster.WdbOprPanicBuyingList("add", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespData(c, http.StatusOK, 200, "success")
	} else {
		s.RespData(c, http.StatusOK, 200, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *PanicBuyingListController) Del(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Del",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.WdbOprPanicBuyingList("del", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespData(c, http.StatusOK, 200, "success")
	} else {
		s.RespData(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *PanicBuyingListController) Update(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Update",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.WdbOprPanicBuyingList("mod", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespMessage(c, http.StatusOK, 200, "success")
	} else {
		s.RespMessage(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *PanicBuyingListController) List(c *gin.Context) {
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
	list := dbJd.GetPanicBuyingListWithPage(limit, offset, sort, sortOrder, fiter)
	s.RespData(c, http.StatusOK, code, list)
}
