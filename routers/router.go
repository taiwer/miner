package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/taiwer/miner/common/middleware/cors"
	"github.com/taiwer/miner/common/middleware/jwt"
	"github.com/taiwer/miner/common/settings"
	"github.com/taiwer/miner/controller"
	"github.com/taiwer/miner/controller/adminController"
	"github.com/taiwer/miner/controller/jdController"
	"github.com/taiwer/miner/controller/staticController"
)

//InitRouter 初始化Router
func InitRouter(conf *settings.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler()) //跨域
	r.Use(gin.Recovery())
	gin.SetMode(conf.RunMode)
	Configure(r)
	return r
}

//Configure 配置router
func Configure(r *gin.Engine) {
	//controller declare
	var cUser controller.User
	var cQrLogin adminController.QrLoginController
	var cGlobal adminController.GlobalListController

	var cDownLoad staticController.DownLoadController
	var cUpLoadFile adminController.UpLoadFileController

	var cJdMiaoshaList jdController.MiaoShaListController
	var cJdPanicBuyingList jdController.PanicBuyingListController
	//var tag cv1.Tag
	var myjwt jwt.JWT
	//inject declare
	var authMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler) //404页面
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/user/sign_up", cUser.SignUp)
	r.POST("/user/qr_login/command", cQrLogin.Command) //扫码登录
	r.GET("/user/qr_login/get_img", cQrLogin.GetImg)   //获取登录二维码图片
	r.GET("/static/download", cDownLoad.DownLoad)
	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/info", cUser.GetUserInfo)
		userAPI.POST("/logout", cUser.Logout)
	}
	var adminMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AdminAuthorizator{})
	apiv1 := r.Group("/api/v1")
	//使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	apiv1.Use(adminMiddleware.MiddlewareFunc())
	{
		//vue获取table信息
		//apiv1.GET("/table/list", article.GetTables)
		apiv1.GET("/user/list", cUser.GetUsers)
		apiv1.POST("/user", cUser.AddUser)
		apiv1.PUT("/user", cUser.UpdateUser)
		apiv1.DELETE("/user/:id", cUser.DeleteUser)

		apiv1.POST("/global/command", cGlobal.Command)
		apiv1.POST("/global/create", cGlobal.Create)
		apiv1.DELETE("/global/del:id", cGlobal.Del)
		apiv1.PUT("/global/update", cGlobal.Update)
		apiv1.GET("/global/list", cGlobal.List)

		apiv1.POST("/up_load_file/command", cUpLoadFile.Command)
		apiv1.POST("/up_load_file/create", cUpLoadFile.Create)
		apiv1.DELETE("/up_load_file/del:id", cUpLoadFile.Del)
		apiv1.PUT("/up_load_file/update", cUpLoadFile.Update)
		apiv1.GET("/up_load_file/list", cUpLoadFile.List)
		apiv1.GET("/up_load_file/list_select", cUpLoadFile.ListSelect)
		apiv1.POST("/up_load_file/up_load", cUpLoadFile.UpLoad)

		apiv1.POST("/jd/miaosha_list/command", cJdMiaoshaList.Command)
		apiv1.POST("/jd/miaosha_list/create", cJdMiaoshaList.Create)
		apiv1.DELETE("/jd/miaosha_list/del:id", cJdMiaoshaList.Del)
		apiv1.PUT("/jd/miaosha_list/update", cJdMiaoshaList.Update)
		apiv1.GET("/jd/miaosha_list/list", cJdMiaoshaList.List)

		apiv1.POST("/jd/panic_buying_list/command", cJdPanicBuyingList.Command)
		apiv1.POST("/jd/panic_buying_list/create", cJdPanicBuyingList.Create)
		apiv1.DELETE("/jd/panic_buying_list/del:id", cJdPanicBuyingList.Del)
		apiv1.PUT("/jd/panic_buying_list/update", cJdPanicBuyingList.Update)
		apiv1.GET("/jd/panic_buying_list/list", cJdPanicBuyingList.List)
	}
}
