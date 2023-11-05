package main

import (
	"fmt"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"github.com/edwardzhanged/novel-go/app/router"
	"gorm.io/gorm"
)

func main() {

	if err := conf.NewConfig(); err != nil {
		panic(err)
	}
	initDB(conf.GbGorm)
	router.InitRouter()

}

func initDB(db *gorm.DB) {
	// 自动迁移表结构
	fmt.Println("init db")
	err := db.AutoMigrate(model.User{})
	if err != nil {
		panic("failed to migrate table")
	}
}
