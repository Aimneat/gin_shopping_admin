package main

import (
	"fmt"
	"gin-shop-admin/initialize"
	"gin-shop-admin/router"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	initialize.GormSetup()
// 	initialize.ViperSetup()
// }

func main() {

	initialize.ViperSetup()
	initialize.Db = initialize.GormSetup()
	if initialize.Db != nil {
		initialize.MysqlTables(initialize.Db) // 初始化表
		// 程序结束前关闭数据库链接
		db := initialize.Db.DB()
		defer db.Close()
	}

	gin.SetMode(initialize.TotalConfig.Server.RunMode)
	endPoint := fmt.Sprintf("%s:%s", initialize.TotalConfig.Server.Host, initialize.TotalConfig.Server.Port)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      router.InitRouter(),
		ReadTimeout:  initialize.TotalConfig.Server.ReadTimeout,
		WriteTimeout: initialize.TotalConfig.Server.WriteTimeout,
		// ReadTimeout:    10000000,
		// WriteTimeout:   10000000,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
