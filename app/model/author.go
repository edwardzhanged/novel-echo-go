package model

import (
	"gorm.io/gorm"
	"time"
)

type AuthorCode struct {
	gorm.Model
	InviteCode   string     `gorm:"type:varchar(100);not null" json:"invite_code" comment:"邀请码"`
	ValidityTime time.Time  `gorm:"type:datetime;not null" json:"validity_time" comment:"有效时间"`
	IsUsed       uint8      `gorm:"type:tinyint unsigned;not null;default:0" json:"is_used" comment:"是否使用过"`
	AuthorInfo   AuthorInfo `gorm:"foreignKey:InviteCodeID"`
}

type AuthorInfo struct {
	gorm.Model
	UserID              uint64               `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"用户ID"`
	InviteCodeID        string               `gorm:"type:varchar(20);not null" json:"invite_code_id" comment:"邀请码ID"`
	PenName             string               `gorm:"type:varchar(20);not null" json:"pen_name" comment:"笔名"`
	TelPhone            string               `gorm:"type:varchar(20)" json:"tel_phone" comment:"手机号码"`
	ChatAccount         string               `gorm:"type:varchar(50)" json:"chat_account" comment:"QQ或微信账号"`
	Email               string               `gorm:"type:varchar(50)" json:"email" comment:"电子邮箱"`
	WorkDirection       uint8                `gorm:"type:tinyint unsigned" json:"work_direction" comment:"作品方向"`
	Status              uint8                `gorm:"type:tinyint unsigned;not null;default:0" json:"status" comment:"用户状态 "`
	AuthorIncomes       []AuthorIncome       `gorm:"foreignKey:AuthorID"`
	AuthorIncomeDetails []AuthorIncomeDetail `gorm:"foreignKey:AuthorID"`
	BookInfos           []BookInfo           `gorm:"foreignKey:AuthorID"`
}

type AuthorIncome struct {
	gorm.Model
	AuthorID       uint64 `gorm:"type:bigint unsigned;not null" json:"author_id" comment:"作家ID"`
	BookID         uint64 `gorm:"type:bigint unsigned;not null" json:"book_id" comment:"小说ID"`
	IncomeMonth    string `gorm:"type:date;not null" json:"income_month" comment:"收入月份"`
	PreTaxIncome   uint32 `gorm:"type:int unsigned;not null;default:0" json:"pre_tax_income" comment:"税前收入，单位：分"`
	AfterTaxIncome uint32 `gorm:"type:int unsigned;not null;default:0" json:"after_tax_income" comment:"税后收入，单位：分"`
	PayStatus      uint8  `gorm:"type:tinyint unsigned;not null;default:0" json:"pay_status" comment:"支付状态 0-待支付 1-已支付"`
	ConfirmStatus  uint8  `gorm:"type:tinyint unsigned;not null;default:0" json:"confirm_status" comment:"稿费确认状态 0-待确认 1-已确认"`
	Detail         string `gorm:"type:varchar(255)" json:"detail" comment:"详情"`
}

type AuthorIncomeDetail struct {
	gorm.Model
	AuthorID      uint64 `gorm:"type:bigint unsigned;not null" json:"author_id" comment:"作家ID"`
	BookID        uint64 `gorm:"type:bigint unsigned;not null;default:0" json:"book_id" comment:"小说ID 0表示全部作品"`
	IncomeDate    string `gorm:"type:date;not null" json:"income_date" comment:"收入日期"`
	IncomeAccount uint32 `gorm:"type:int unsigned;not null;default:0" json:"income_account" comment:"订阅总额，单位：分"`
	IncomeCount   uint32 `gorm:"type:int unsigned;not null;default:0" json:"income_count" comment:"订阅次数"`
	IncomeNumber  uint32 `gorm:"type:int unsigned;not null;default:0" json:"income_number" comment:"订阅人数"`
}
