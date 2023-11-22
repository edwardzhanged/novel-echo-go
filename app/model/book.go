package model

import "gorm.io/gorm"

type BookInfo struct {
	gorm.Model
	WorkDirection string `gorm:"type:varchar(20);not null;comment:作品方向"`
	BookName      string `gorm:"type:varchar(20);not null;comment:书名"`
	AuthorID      int    `gorm:"type:int;not null;comment:作者ID"`
	Score         int    `gorm:"type:int;not null;comment:评分"`
}
