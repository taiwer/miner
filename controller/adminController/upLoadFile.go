package adminController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/codes"
	"github.com/taiwer/miner/common/dbPlotOpr"
	"github.com/taiwer/miner/controller/baseController"
	"github.com/taiwer/miner/models/httpMaster"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UpLoadFileController struct {
	baseController.BaseController
}

func (s *UpLoadFileController) Command(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	if v, ok := args["cmd"]; ok {
		cmd, _ := v.(string)
		switch cmd {
		case "get_game_server_version_select":
			list := dbPlotOpr.GetUpLoadFileListByWhere("file_class='.gz'")
			s.RespData(c, http.StatusOK, 200, list)
		default:
			s.RespData(c, http.StatusOK, 0, "default cmd:"+cmd)
		}
	} else {
		s.RespData(c, http.StatusOK, 0, "cmd is null")
	}
}

func (s *UpLoadFileController) Create(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.HttpRequest("upLoadFile", "add", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespData(c, http.StatusOK, 200, "success")
	} else {
		s.RespData(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *UpLoadFileController) Del(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Del",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.HttpRequest("upLoadFile", "del", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespData(c, http.StatusOK, 200, "success")
	} else {
		s.RespData(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *UpLoadFileController) Update(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Update",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	ret, _ := httpMaster.HttpRequest("upLoadFile", "mod", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespMessage(c, http.StatusOK, 200, "success")
	} else {
		s.RespMessage(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}

func (s *UpLoadFileController) List(c *gin.Context) {
	code := codes.SUCCESS
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	sort := c.Query("sort")
	sortOrder := c.Query("sortOrder")
	search := c.Query("search")
	fiter := c.Query("fiter")
	searchFiter := ""
	if search != "" {
		searchFiter = fmt.Sprintf("username LIKE '%%%s%%' or qq LIKE '%%%s%%' or descs LIKE '%%%s%%'", search, search, search)
	}
	if searchFiter != "" {
		if fiter != "" {
			fiter = fmt.Sprintf("(%s) AND (%s)", fiter, searchFiter)
		} else {
			fiter = searchFiter
		}
	}
	list := dbPlotOpr.GetUpLoadFileListWithPage(limit, offset, sort, sortOrder, fiter)
	//rows := list.Rows.(*[]dbMyGameOpr.Account)
	//for i := 0; i < len(*rows); i++ {
	//	item := &(*rows)[i]
	//	item.NetPackNum, _ = dbMyGameOpr.GetNetProtocolPackCount(item.Account)
	//}
	s.RespData(c, http.StatusOK, code, list)
}

func (s *UpLoadFileController) ListSelect(c *gin.Context) {
	code := codes.SUCCESS
	FileClass := c.Query("file_class")
	var list []*dbPlotOpr.UpLoadFile
	if FileClass != "" {
		list = dbPlotOpr.GetUpLoadFileListByWhere(fmt.Sprintf("file_class='%s'", FileClass))
	} else {
		list = dbPlotOpr.GetUpLoadFileList()
	}
	type selectRow struct {
		Id       int64  `json:"id"`
		FileName string `json:"file_name"`
	}
	data := make([]selectRow, 0)
	for _, v := range list {
		row := selectRow{
			Id:       v.Id,
			FileName: v.FileName,
		}
		data = append(data, row)
	}
	s.RespData(c, http.StatusOK, code, data)
}

func (s *UpLoadFileController) UpLoad(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		s.RespMessage(c, http.StatusOK, 0, "请求失败")
		return
	}
	filePath := s.GetUpLoadFilePath()
	os.MkdirAll(filePath, os.ModePerm)
	saveName := file.Filename
	if err := c.SaveUploadedFile(file, filePath+saveName); err != nil {
		s.RespMessage(c, http.StatusOK, 0, "保存文件失败"+err.Error())
		return
	}
	zap.L().Info(fmt.Sprintf("success fileName:%s fileSize:%d", file.Filename, file.Size))
	args := make(map[string]interface{})
	args["file_name"] = saveName
	args["file_class"] = path.Ext(saveName)
	ret, _ := httpMaster.HttpRequest("upLoadFile", "add", args, "", c.Request.RemoteAddr)
	if ret == "0" {
		s.RespMessage(c, http.StatusOK, 200, "success")
	} else {
		s.RespMessage(c, http.StatusOK, 0, fmt.Sprintf("ret=%s", ret))
	}
}
