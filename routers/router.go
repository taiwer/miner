package routers

import (
	"github.com/gin-gonic/gin"
	"goplot/common/middleware/cors"
	"goplot/common/middleware/jwt"
	"goplot/common/settings"
	"goplot/controller"
	"goplot/controller/adminController"
	"goplot/controller/plotController"
	"goplot/controller/staticController"
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
	var cGlobal adminController.GlobalListController
	var cPlotPc plotController.PlotPcController
	var cPlotPcUser plotController.PlotUserPcController
	var cPlotKey plotController.PlotKeyController
	var cPlotKeyUser plotController.PlotKeyUserController
	var cPlotDisk plotController.PlotDiskController

	var cDownLoad staticController.DownLoadController
	var cUpLoadFile adminController.UpLoadFileController
	//var tag cv1.Tag
	var myjwt jwt.JWT
	//inject declare
	var authMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler) //404页面
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/user/sign_up", cUser.SignUp)
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

		apiv1.POST("/ploter/command", cPlotPc.Command)
		apiv1.POST("/ploter/create", cPlotPc.Create)
		apiv1.DELETE("/ploter/del:id", cPlotPc.Del)
		apiv1.PUT("/ploter/update", cPlotPc.Update)
		apiv1.GET("/ploter/list", cPlotPc.List)

		apiv1.POST("/ploter_user/command", cPlotPcUser.Command)
		apiv1.POST("/ploter_user/create", cPlotPcUser.Create)
		apiv1.DELETE("/ploter_user/del:id", cPlotPcUser.Del)
		apiv1.PUT("/ploter_user/update", cPlotPcUser.Update)
		apiv1.GET("/ploter_user/list", cPlotPcUser.List)

		apiv1.POST("/plot_disk/command", cPlotDisk.Command)
		apiv1.POST("/plot_disk/create", cPlotDisk.Create)
		apiv1.DELETE("/plot_disk/del:id", cPlotDisk.Del)
		apiv1.PUT("/plot_disk/update", cPlotDisk.Update)
		apiv1.GET("/plot_disk/list", cPlotDisk.List)

		apiv1.POST("/plot_key/command", cPlotKey.Command)
		apiv1.POST("/plot_key/create", cPlotKey.Create)
		apiv1.DELETE("/plot_key/del:id", cPlotKey.Del)
		apiv1.PUT("/plot_key/update", cPlotKey.Update)
		apiv1.GET("/plot_key/list", cPlotKey.List)
		apiv1.GET("/plot_key/list_select", cPlotKey.ListSelect)

		apiv1.POST("/plot_key_user/command", cPlotKeyUser.Command)
		apiv1.POST("/plot_key_user/create", cPlotKeyUser.Create)
		apiv1.DELETE("/plot_key_user/del:id", cPlotKeyUser.Del)
		apiv1.PUT("/plot_key_user/update", cPlotKeyUser.Update)
		apiv1.GET("/plot_key_user/list", cPlotKeyUser.List)
		apiv1.GET("/plot_key_user/list_select", cPlotKeyUser.ListSelect)

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
	}
}
