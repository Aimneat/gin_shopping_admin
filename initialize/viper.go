package initialize

import (
	"fmt"
	"os"

	"gin-shop-admin/config"

	"github.com/spf13/viper"
)

type AllConfig struct {
	Server     config.Server
	Datasource config.Datasource
	Jwt        config.Jwt
}

var TotalConfig AllConfig

func ViperSetup() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s /n", err))
	}
	viper.Unmarshal(&TotalConfig)
}
