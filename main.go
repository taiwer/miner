package main

import (
	"fmt"
	"github.com/taiwer/miner/common/dbManager"
	"github.com/taiwer/miner/common/dbOrm"
	"github.com/taiwer/miner/common/dbPlotOpr"
	"github.com/taiwer/miner/common/logger"
	"github.com/taiwer/miner/common/rbacModel"
	"github.com/taiwer/miner/common/settings"
	"github.com/taiwer/miner/models/jdSeckill"
	"github.com/taiwer/miner/routers"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//启动调试10000端口 http://localhost:10000/debug/pprof/heap
	//设置日志输出格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//1. 加载配置
	if err := settings.Init(); err != nil {
		log.Printf("init settings faild err:%v\n", err)
		return
	}
	log.Printf("Conf:%+v", settings.Conf)
	//2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.RunMode); err != nil {
		log.Printf("init settings faild err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	//3. 连接数据库
	zap.L().Info("连接数据库...")
	if err := dbOrm.Init("", settings.Conf.MySQLConfig); err != nil {
		zap.L().Error("init mysql faild", zap.String("err", err.Error()))
		return
	} else {
		defer dbOrm.DbClose()
		//网站数据库建表
		if len(os.Args) >= 2 {
			switch os.Args[1] {
			case "syncdb":
				rbacModel.Syncdb(true)
				dbPlotOpr.Syncdb(true)
			case "migrate":
				rbacModel.Syncdb(false)
				//自动构建数据表
				dbPlotOpr.Syncdb(false)
			default:
			}
		}

		//dbPlotOpr.Syncdb(false)
		//dbPlotOpr.CreateTrigger(nil)
		//rbacModel.Syncdb(true)
	}
	go dbManager.DataBaseMgrObj.Run() //加载数据库数据
	//启动网站接口服务
	router := routers.InitRouter(settings.Conf.App)
	conf := settings.Conf.HttpServer
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	go s.ListenAndServe()
	jdSeckill.RunSeckill()
	select {}
	log.Println("exit")
}
