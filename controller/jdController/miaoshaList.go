package jdController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/codes"
	"github.com/taiwer/miner/common/dbPlotOpr"
	"github.com/taiwer/miner/controller/baseController"
	"github.com/taiwer/miner/models/httpMaster"
	"github.com/taiwer/miner/models/jdSeckill/jd"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type MiaoShaListController struct {
	baseController.BaseController
}

func (s *MiaoShaListController) Command(c *gin.Context) {
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
		case "get_miaosha_list":
			list := jd.GetMiaoshaList(jdUsername)
			s.RespData(c, http.StatusOK, 200, list)
		case "get_item_info":
			skuId := args["skuId"]
			if skuId != nil {
				item := jd.GetItemInfo(jdUsername, skuId.(string))
				s.RespData(c, http.StatusOK, 200, item)
			}
		case "add_item_cart":
			skuId := args["skuId"]
			if skuId != nil {
				if err := jd.AddItemToCart(jdUsername, skuId.(string), 1); err == nil {
					s.RespData(c, http.StatusOK, 200, "ok")
				} else {
					s.RespData(c, http.StatusOK, 200, err.Error())
				}

			}
		case "get_cart_list":
			if item, err := jd.GetCartList(jdUsername); err != nil {
				s.RespData(c, http.StatusOK, 0, err.Error())
			} else {
				obj := make(map[string]interface{})
				json.Unmarshal([]byte(item), &obj)
				s.RespData(c, http.StatusOK, 200, obj)
			}
		case "cart_check_single":
			Id := args["id"].(string)
			num := args["num"].(float64)
			if item, err := jd.CartCheckSingle(jdUsername, Id, int(num)); err != nil {
				s.RespData(c, http.StatusOK, 0, err.Error())
			} else {
				obj := make(map[string]interface{})
				json.Unmarshal([]byte(item), &obj)
				s.RespData(c, http.StatusOK, 200, obj)
			}
		case "cart_uncheck_single":
			Id := args["id"].(string)
			num := args["num"].(float64)
			if item, err := jd.CartUnCheckSingle(jdUsername, Id, int(num)); err != nil {
				s.RespData(c, http.StatusOK, 0, err.Error())
			} else {
				obj := make(map[string]interface{})
				json.Unmarshal([]byte(item), &obj)
				s.RespData(c, http.StatusOK, 200, obj)
			}
		case "cart_uncheck_all":
			if item, err := jd.CartUnCheckAll(jdUsername); err != nil {
				s.RespData(c, http.StatusOK, 0, err.Error())
			} else {
				obj := make(map[string]interface{})
				json.Unmarshal([]byte(item), &obj)
				s.RespData(c, http.StatusOK, 200, obj)
			}
		case "get_order_info":
			if needPay, err := jd.GetOrderInfo(jdUsername); err != nil {
				s.RespData(c, http.StatusOK, 0, err.Error())
			} else {
				s.RespData(c, http.StatusOK, 200, needPay)
			}
		case "submit_order":
			if item, err := jd.SubmitOrder(jdUsername); err != nil {
				s.RespData(c, http.StatusOK, 0, err.Error())
			} else {
				obj := make(map[string]interface{})
				json.Unmarshal([]byte(item), &obj)
				s.RespData(c, http.StatusOK, 200, obj)
			}
		default:
			s.RespData(c, http.StatusOK, 0, "default cmd:"+cmd)
		}
	} else {
		s.RespData(c, http.StatusOK, 0, "cmd is null")
	}
}

func (s *MiaoShaListController) Create(c *gin.Context) {
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

func (s *MiaoShaListController) Del(c *gin.Context) {
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

func (s *MiaoShaListController) Update(c *gin.Context) {
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

func (s *MiaoShaListController) List(c *gin.Context) {
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
