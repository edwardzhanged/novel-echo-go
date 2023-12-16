package conf

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GbViper *viper.Viper
	GbGorm  *gorm.DB
	GbRedis *redis.Client
)

func Initialize() error {
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
	// 连接 MySQL 数据库
	if GbGorm, err = gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{}); err != nil {
		return err
	}
	// 连接 Redis 数据库
	GbRedis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", GbViper.GetString("redis.host"),
			GbViper.GetString("redis.port")),
		Password: GbViper.GetString("redis.password"),
	})
	return nil
}
