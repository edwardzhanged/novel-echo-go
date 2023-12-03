package model

import (
	"gorm.io/gorm"
)

type HomeBook struct {
	gorm.Model
	Type   uint8    `gorm:"type:tinyint unsigned;not null;default:0" json:"type" comment:"推荐类型 0-轮播图 1-顶部栏 2-本周强推 3-热门推荐 4-精品推荐"`
	Sort   uint8    `gorm:"type:tinyint unsigned;not null" json:"sort" comment:"推荐排序"`
	BookID uint64   `gorm:"type:bigint unsigned;not null;foreignKey:BookID" json:"book_id" comment:"推荐小说ID"`
	Book   BookInfo `json:"book"`
}

type HomeFriendLink struct {
	gorm.Model
	LinkName string `gorm:"type:varchar(50);not null" json:"link_name" comment:"链接名"`
	LinkURL  string `gorm:"type:varchar(100);not null" json:"link_url" comment:"链接url"`
	Sort     uint8  `gorm:"type:tinyint unsigned;not null;default:11" json:"sort" comment:"排序号"`
	IsOpen   uint8  `gorm:"type:tinyint unsigned;not null;default:1" json:"is_open" comment:"是否开启 0-不开启 1-开启"`
}
