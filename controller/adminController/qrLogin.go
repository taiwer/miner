package adminController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/dbJd"
	"github.com/taiwer/miner/common/middleware/jwt"
	"github.com/taiwer/miner/common/middleware/models"
	"github.com/taiwer/miner/controller/baseController"
	"github.com/taiwer/miner/models/jdSeckill/jd"
	"github.com/taiwer/miner/models/jdSeckill/seckill"
	"go.uber.org/zap"
	"net/http"
)

//扫码登录
type QrLoginController struct {
	baseController.BaseController
}

//Logout 退出登录
func (s *QrLoginController) Command(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)
	zap.L().Info("Create",
		zap.Any("request", c.Request.Form),
		zap.Any("args", args),
	)
	if v, ok := args["cmd"]; ok {
		cmd, _ := v.(string)
		switch cmd {
		case "show":
			if wlfstkSmdl, err := jd.QrLogin(); err != nil {
				s.RespData(c, http.StatusOK, 0, err)
			} else {
				data := map[string]string{
					"wlfstkSmdl": wlfstkSmdl,
					"imgUrl":     fmt.Sprintf("http://127.0.0.1:8000/user/qr_login/get_img?wlfstkSmdl=%s.png", wlfstkSmdl),
				}
				s.RespData(c, http.StatusOK, 200, data)
			}
		case "get_tick":
			if wlfstkSmdl, ok := args["wlfstkSmdl"]; ok {
				if code, msg, ticket, err := jd.QrcodeTicket(wlfstkSmdl.(string)); err != nil {
					s.RespData(c, http.StatusOK, 0, err)
				} else {
					data := map[string]interface{}{
						"code":   code,
						"msg":    msg,
						"ticket": ticket,
					}
					s.RespData(c, http.StatusOK, 200, data)
				}
			}
		case "get_token":
			if ticker, ok := args["tick"]; ok {
				if thor, unick, err := jd.TicketInfo(ticker.(string)); err != nil {
					data := map[string]interface{}{
						"thor":  thor,
						"unick": unick,
						"msg":   err.Error(),
					}
					s.RespData(c, http.StatusOK, 200, data)
				} else {
					userName := "Jd_" + unick
					jd.JdUserListObj.Add(userName, thor)
					user := dbJd.JdUser{
						Username: userName,
						Password: "",
						Nickname: unick,
						Token:    thor,
					}
					if dbJd.GetJdUserByName(user.Username) != nil {
						dbJd.UpdateJdUserByName(user.Username, map[string]interface{}{"token": user.Token})
					} else {
						if _, err := dbJd.AddJdUser(&user, nil); err != nil {

						}
					}
					var myjwt = (&jwt.JWT{}).GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
					myjwt.SetLoginSuccessResponse(c, &models.UserRole{UserName: userName, UserRoles: nil})
				}
			}
		default:
			s.RespData(c, http.StatusOK, 0, "default cmd:"+cmd)
		}
	}
}

//Logout 退出登录
func (s *QrLoginController) GetImg(c *gin.Context) {
	args := make(map[string]interface{})
	c.ShouldBind(&args)

	wlfstkSmdl := c.Request.Form.Get("wlfstkSmdl")
	if wlfstkSmdl != "" {
		//if v, ok := args["wlfstkSmdl"]; ok {
		//wlfstkSmdl, _ := v.(string)
		destFile := seckill.QrImgPath + "/" + wlfstkSmdl
		zap.L().Info("DownLoad File " + destFile)
		c.File(destFile)
	}
}
