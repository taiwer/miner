package baseController

import (
	"fmt"
	jwt "github.com/taiwer/miner/common/middleware/jwt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/codes"
)

//ResponseData 数据返回结构体
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ResponseSuccess 返回成功结构体
type ResponseSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//ResponseFail 返回成功结构体
type ResponseFail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

//RespData 数据返回
type BaseController struct {
}

func (s *BaseController) GetUserName(c *gin.Context) string {
	token := c.MustGet("JWT_PAYLOAD").(jwt.MapClaims)
	if username, ok := token["userName"]; ok {
		return username.(string)
	}
	return ""
}

//获取用户角色
func (s *BaseController) GetUserRoles(c *gin.Context) []string {
	userRoles := []string{} //规则
	userName := s.GetUserName(c)
	if userName == "admin" {
		userRoles = append(userRoles, "admin")
	}
	if strings.HasPrefix(userName, "Jd_") {
		userRoles = append(userRoles, "jdkill") //京东秒杀
	}
	userRoles = append(userRoles, "user") //默认
	return userRoles
}

func (s *BaseController) RespData(c *gin.Context, httpCode, code int, data interface{}) {
	resp := ResponseData{
		Code:    code,
		Message: codes.GetMsg(code),
		Data:    data,
	}
	s.RespJSON(c, httpCode, resp)
}

func (s *BaseController) RespMessage(c *gin.Context, httpCode, code int, message string) {
	resp := ResponseData{
		Code:    code,
		Message: message,
	}
	s.RespJSON(c, httpCode, resp)
}

//RespOk 返回操作成功
func (s *BaseController) RespOk(c *gin.Context, httpCode, code int) {
	resp := ResponseSuccess{
		Code:    code,
		Message: codes.GetMsg(code),
	}
	s.RespJSON(c, httpCode, resp)
}

//RespFail 返回操作失败
func (s *BaseController) RespFail(c *gin.Context, httpCode, code int, detail string) {
	resp := ResponseFail{
		Code:    code,
		Message: fmt.Sprintf("%s: %s", codes.GetMsg(code), detail),
		Detail:  detail,
	}
	s.RespJSON(c, httpCode, resp)
}

//RespJSON 返回JSON数据
func (s *BaseController) RespJSON(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
}

//GetPage 获取每页数量
func (s *BaseController) GetPage(c *gin.Context) (page, pagesize int) {
	page, _ = strconv.Atoi(c.Query("page"))
	pagesize, _ = strconv.Atoi(c.Query("limit"))
	if pagesize == 0 {
		pagesize = 10
	}
	if page == 0 {
		page = 1
	}
	return
}

func (s *BaseController) GetUpLoadFilePath() string {
	return "../up_load_file/"
}

func (s *BaseController) GetImgFilePath() string {
	return "../img/"
}
