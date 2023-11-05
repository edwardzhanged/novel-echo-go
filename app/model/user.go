package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid     string `gorm:"type:varchar(36);not null;uniqueIndex;comment:用户UUID"`
	Username string `gorm:"type:varchar(20);not null;uniqueIndex;comment:用户名"`
	Password string
	Phone    string
}
