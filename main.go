package main

import (
	"fmt"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"github.com/edwardzhanged/novel-go/app/router"
	"gorm.io/gorm"
)

func main() {
	if err := conf.Initialize(); err != nil {
		panic(err)
	}
	initDB(conf.GbGorm)
	router.InitRouter()

}

func initDB(db *gorm.DB) {
	// 自动迁移表结构
	fmt.Println("init db")

	err := db.AutoMigrate(
		model.UserInfo{},

		// Author
		model.AuthorCode{},
		model.AuthorInfo{},
		model.AuthorIncome{},
		model.AuthorIncomeDetail{},
		// Book
		model.BookCategory{},
		model.BookInfo{},
		model.BookChapter{},
		model.BookContent{},
		model.BookComment{},
		model.BookCommentReply{},
		// Home
		model.HomeBook{},
		model.HomeFriendLink{},
		// News
		model.NewsCategory{},
		model.NewsInfo{},
		model.NewsContent{},
		// Pay
		model.PayAlipay{},
		model.PayWechat{},
		// System
		model.SysUser{},
		model.SysRole{},
		model.SysMenu{},
		model.SysUserRole{},
		model.SysRoleMenu{},
		model.SysLog{},
		// User
		model.UserFeedback{},
		model.UserBookshelf{},
		model.UserReadHistory{},
		model.UserConsumeLog{},
		model.UserPayLog{},
	)
	if err != nil {
		panic("failed to migrate table")
	}
}

// TODO: 2021/7/21
//  1. MIDDLEWARE调研,日志
//  2. 实现login
