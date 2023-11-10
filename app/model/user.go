package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid      string `gorm:"type:varchar(36);not null;comment:用户UUID;column:uid"`
	Nickname  string `gorm:"type:varchar(20);not null;comment:用户名"`
	Password  string `gorm:"type:varchar(100);not null;comment:用户密码"`
	Phone     string `gorm:"type:varchar(20);not null;uniqueIndex;comment:用户手机号"`
	UserPhoto string
	UserSex   string
}
