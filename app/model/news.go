package model

import (
	"gorm.io/gorm"
)

type NewsCategory struct {
	gorm.Model
	Name      string     `gorm:"type:varchar(20);not null" json:"name" comment:"类别名"`
	Sort      uint8      `gorm:"type:tinyint unsigned;not null;default:10" json:"sort" comment:"排序"`
	NewsInfos []NewsInfo `gorm:"foreignKey:CategoryID"`
}

type NewsInfo struct {
	gorm.Model
	CategoryID   uint64 `gorm:"type:bigint unsigned;not null" json:"category_id" comment:"类别ID"`
	CategoryName string `gorm:"type:varchar(50);not null" json:"category_name" comment:"类别名"`
	SourceName   string `gorm:"type:varchar(50);not null" json:"source_name" comment:"新闻来源"`
	Title        string `gorm:"type:varchar(100);not null" json:"title" comment:"新闻标题"`
}

type NewsContent struct {
	gorm.Model
	NewsID   uint64   `gorm:"type:bigint unsigned;not null" json:"news_id" comment:"新闻ID"`
	NewsInfo NewsInfo `gorm:"foreignKey:NewsID"`
	Content  string   `gorm:"type:mediumtext;not null" json:"content" comment:"新闻内容"`
}
