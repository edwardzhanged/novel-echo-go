package model

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	Username         string            `gorm:"type:varchar(50);not null" json:"username" comment:"登录名"`
	Password         string            `gorm:"type:varchar(100);not null" json:"password" comment:"登录密码-加密"`
	Salt             string            `gorm:"type:varchar(8);not null" json:"salt" comment:"加密盐值"`
	NickName         string            `gorm:"type:varchar(50)" json:"nick_name" comment:"昵称"`
	UserPhoto        string            `gorm:"type:varchar(100)" json:"user_photo" comment:"用户头像"`
	UserSex          uint8             `gorm:"type:tinyint unsigned" json:"user_sex" comment:"用户性别 0-男 1-女"`
	AccountBalance   uint64            `gorm:"type:bigint unsigned;not null;default:0" json:"account_balance" comment:"账户余额"`
	Status           uint8             `gorm:"type:tinyint unsigned;not null;default:0" json:"status" comment:"用户状态"`
	UserFeedbacks    []UserFeedback    `gorm:"foreignKey:UserID"`
	UserPayLogs      []UserPayLog      `gorm:"foreignKey:UserID"`
	UserBookshelves  []UserBookshelf   `gorm:"foreignKey:UserID"`
	UserConsumeLogs  []UserConsumeLog  `gorm:"foreignKey:UserID"`
	UserReadHistorys []UserReadHistory `gorm:"foreignKey:UserID"`
	BookComments     []BookComment     `gorm:"foreignKey:UserID"`
	AuthorInfo       AuthorInfo        `gorm:"foreignKey:UserID"`
}

type UserFeedback struct {
	gorm.Model
	UserID  uint64 `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"反馈用户id"`
	Content string `gorm:"type:varchar(512);not null" json:"content" comment:"反馈内容"`
}

type UserBookshelf struct {
	gorm.Model
	UserID       uint64 `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"用户ID"`
	BookID       uint64 `gorm:"type:bigint unsigned;not null" json:"book_id" comment:"小说ID"`
	PreContentID uint64 `gorm:"type:bigint unsigned" json:"pre_content_id" comment:"上一次阅读的章节内容表ID"`
}

type UserReadHistory struct {
	gorm.Model
	UserID       uint64 `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"用户ID"`
	BookID       uint64 `gorm:"type:bigint unsigned;not null" json:"book_id" comment:"小说ID"`
	PreContentID uint64 `gorm:"type:bigint unsigned;not null" json:"pre_content_id" comment:"上一次阅读的章节内容表ID"`
}

type UserConsumeLog struct {
	gorm.Model
	UserID       uint64 `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"消费用户ID"`
	Amount       uint32 `gorm:"type:int unsigned;not null" json:"amount" comment:"消费使用的金额，单位：屋币"`
	ProductType  uint8  `gorm:"type:tinyint unsigned;not null;default:0" json:"product_type" comment:"消费商品类型 0-小说VIP章节"`
	ProductID    uint64 `gorm:"type:bigint unsigned" json:"product_id" comment:"消费的的商品ID，例如：章节ID"`
	ProductName  string `gorm:"type:varchar(50)" json:"product_name" comment:"消费的的商品名，例如：章节名"`
	ProductValue uint32 `gorm:"type:int unsigned" json:"product_value" comment:"消费的的商品值，例如：1"`
}

type UserPayLog struct {
	gorm.Model
	UserID       uint64    `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"充值用户ID"`
	PayChannel   uint8     `gorm:"type:tinyint unsigned;not null;default:1" json:"pay_channel" comment:"充值方式"`
	OutTradeNo   string    `gorm:"type:varchar(64);not null" json:"out_trade_no" comment:"商户订单号"`
	PayAlipay    PayAlipay `gorm:"foreignKey:OutTradeNo"`
	PayWechat    PayWechat `gorm:"foreignKey:OutTradeNo"`
	Amount       uint32    `gorm:"type:int unsigned;not null" json:"amount" comment:"充值金额，单位：分"`
	ProductType  uint8     `gorm:"type:tinyint unsigned;not null;default:0" json:"product_type" comment:"充值商品类型"`
	ProductID    uint64    `gorm:"type:bigint unsigned" json:"product_id" comment:"充值商品ID"`
	ProductName  string    `gorm:"type:varchar(255);not null" json:"product_name" comment:"充值商品名，示例值：屋币"`
	ProductValue uint32    `gorm:"type:int unsigned" json:"product_value" comment:"充值商品值"`
	PayTime      time.Time `gorm:"type:datetime;not null" json:"pay_time" comment:"充值时间"`
}
