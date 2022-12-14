package initialize

import (
	"fmt"
	"gin-shop-admin/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

var Db *gorm.DB

func GormSetup() *gorm.DB {
	// var err error

	driverName := TotalConfig.Datasource.DriverName
	db, err := gorm.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		TotalConfig.Datasource.Username,
		TotalConfig.Datasource.Password,
		TotalConfig.Datasource.Host,
		TotalConfig.Datasource.Port,
		TotalConfig.Datasource.Database,
		TotalConfig.Datasource.Charset,
		// setting.TotalConfig.Datasource.Loc,
	))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	if TotalConfig.Server.RunMode == "debug" {
		db.LogMode(true)
	}

	return db
}

// MysqlTables
//@author:
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Rights{},
		&models.Categories{},
		&models.Roles{},
	)
	// if err != nil {
	// 	// global.GSA_LOG.Error("register table failed", zap.Any("err", err))
	// 	os.Exit(0)
	// }
}

func GetDB() *gorm.DB {
	return Db
}
