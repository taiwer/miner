package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/codes"
	jwt "github.com/taiwer/miner/common/middleware/jwt"
	"github.com/taiwer/miner/common/myutils"
	"github.com/taiwer/miner/common/rbacModel"
	"github.com/taiwer/miner/controller/baseController"
	"github.com/taiwer/miner/models/page"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

//User 注入IUserService
type User struct {
	baseController.BaseController
}

//GetUserInfo 根据token获取用户信息
func (s *User) GetUserInfo(c *gin.Context) {
	zap.L().Info("GetUserInfo")
	roles := jwt.ExtractClaims(c)
	if v, ok := roles["userName"]; ok {
		userName := v.(string)
		code := codes.SUCCESS
		userRoles := s.GetUserRoles(c) //用户角色
		data := page.User{Roles: userRoles, Introduction: "", Name: userName}
		s.RespData(c, http.StatusOK, code, &data)
		zap.L().Info("GetUserInfo",
			zap.Any("rspData", data),
		)
	} else {
		s.RespData(c, http.StatusOK, 0, nil)
	}
}

//Logout 退出登录
func (s *User) Logout(c *gin.Context) {
	s.RespOk(c, http.StatusOK, codes.SUCCESS)
}

//Logout 退出登录
func (s *User) SignUp(c *gin.Context) {
	var param struct {
		UserName string `json:"user_name"`
		Passwd   string `json:"passwd"`
		RePasswd string `json:"re_passwd"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBind(&param); err != nil {
		s.RespMessage(c, http.StatusOK, 0, "missing Username or Password or email")
		return
	}
	zap.L().Info("Index",
		zap.Any("request", c.Request.Form),
		zap.Any("param", param),
	)
	if param.Passwd != param.RePasswd {
		s.RespMessage(c, http.StatusOK, 0, "确认密码错误")
		return
	}
	username := param.UserName
	if user := rbacModel.GetUserByName(username); user != nil {
		s.RespMessage(c, http.StatusOK, 0, "账号已经存在")
		return
	}
	user := rbacModel.User{Username: param.UserName,
		Password: myutils.Pwdhash(param.Passwd),
		Email:    param.Email,
		Nickname: param.UserName,
	}
	if _, err := rbacModel.AddUser(&user, nil); err != nil {
		s.RespMessage(c, http.StatusOK, 0, err.Error())
		return
	}
	s.RespOk(c, http.StatusOK, codes.SUCCESS)
}

//GetUsers 获取用户信息
func (s *User) GetUsers(c *gin.Context) {
	code := codes.SUCCESS
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	sort := c.Query("sort")
	sortOrder := c.Query("sortOrder")
	search := c.Query("search")
	searchFiter := ""
	if search != "" {
		searchFiter = fmt.Sprintf("username LIKE '%%%s%%' or qq LIKE '%%%s%%' or descs LIKE '%%%s%%'", search, search, search)
	}
	fiter := ""
	if searchFiter != "" {
		fiter = fmt.Sprintf("(%s) AND (%s)", fiter, searchFiter)
	}
	data := rbacModel.GetUserListWithPage(pageSize, offset, sort, sortOrder, fiter)
	s.RespData(c, http.StatusOK, code, data)
}

//AddUser 新建用户
func (s *User) AddUser(c *gin.Context) {
	code := codes.InvalidParams
	//err := c.Bind(&user)
	username := c.Query("username")
	password := c.Query("password")
	if username != "" && password != "" {
		user := &rbacModel.User{}
		user.Username = username
		user.Password = password
		if rbacModel.GetUserByName(username) == nil {
			if _, err := rbacModel.AddUser(user, nil); err != nil {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			code = codes.ErrExistUser
		}
	}
	s.RespOk(c, http.StatusOK, code)
}

//UpdateUser 修改用户
func (s *User) UpdateUser(c *gin.Context) {
	code := codes.InvalidParams
	c.Request.ParseForm()
	args := make(map[string]interface{}, 0)
	for k, v := range c.Request.Form {
		switch k {
		case "qq":
			args["qq"] = v[0]
		default:

		}
	}
	username := c.Query("username")
	if username != "" {
		if _, err := rbacModel.UpdateUserByName(username, args); err == nil {
			code = codes.SUCCESS
		} else {
			code = codes.ERROR
		}
	}
	s.RespOk(c, http.StatusOK, code)
}

//DeleteUser 删除用户
func (s *User) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if _, err := rbacModel.DelUserById(int64(id)); err != nil {
		code = codes.ERROR
		s.RespFail(c, http.StatusOK, code, "不允许删除admin账号!")
	} else {
		s.RespOk(c, http.StatusOK, code)
	}
}
