package jwt

import (
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/taiwer/miner/common/middleware/models"
	"github.com/taiwer/miner/common/myutils"
	"github.com/taiwer/miner/common/rbacModel"
	"github.com/taiwer/miner/common/settings"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/taiwer/miner/common/codes"
	"github.com/taiwer/miner/common/helper"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

// JWT 注入IService
type JWT struct {
}

//app 程序配置

//GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit(jwtAuthorizator IAuthorizator) (authMiddleware *GinJWTMiddleware) {
	app := settings.Conf.App
	authMiddleware, err := New(&GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 5,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: app.IdentityKey,
		PayloadFunc: func(data interface{}) MapClaims {
			if v, ok := data.(*models.UserRole); ok {
				//get roles from username
				v.UserRoles = nil
				jsonRole, _ := json.Marshal(v.UserRoles)
				//maps the claims in the JWT
				return MapClaims{
					"userName":  v.UserName,
					"userRoles": helper.B2S(jsonRole),
				}
			}
			return MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			roles := ExtractClaims(c)
			//extracts identity from roles
			jsonRole := roles["userRoles"].(string)
			var userRoles []*models.Role
			json.Unmarshal(helper.S2B(jsonRole), &userRoles)
			//Set the identity
			return &models.UserRole{
				UserName:  roles["userName"].(string),
				UserRoles: userRoles,
			}
		},
		//登录接口
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals models.User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password
			if user := rbacModel.GetUserByName(username); user != nil {
				has := myutils.Pwdhash(password)
				if user.Password == password || user.Password == has {
					return &models.UserRole{
						UserName: username,
					}, nil
				}
			}
			return nil, ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator.HandleAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc:      time.Now,
		DisabledAbort: true,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

//NoRouteHandler 404 handler
func NoRouteHandler(c *gin.Context) {
	code := codes.PageNotFound
	c.JSON(404, gin.H{"code": code, "message": codes.GetMsg(code)})
}
