package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GbViper *viper.Viper
	GbGorm  *gorm.DB
)

func NewConfig() error {
	GbViper = viper.New()
	GbViper.SetConfigFile("config.yaml")
	GbViper.SetConfigType("yaml")
	err := GbViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	GbViper.WatchConfig()
	// 从 Viper 中获取 MySQL 连接信息
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		GbViper.GetString("mysql.user"),
		GbViper.GetString("mysql.password"), GbViper.GetString("mysql.host"),
		GbViper.GetInt("mysql.port"), GbViper.GetString("mysql.database"))
	fmt.Println(mysqlDsn)
	if GbGorm, err = gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{}); err != nil {
		return err
	}
	return nil
}
