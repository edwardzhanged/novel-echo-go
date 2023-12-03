package model

import (
	"gorm.io/gorm"
	"time"
)

type PayAlipay struct {
	gorm.Model
	OutTradeNo    string    `gorm:"type:varchar(64);not null" json:"out_trade_no" comment:"商户订单号"`
	TradeNo       string    `gorm:"type:varchar(64);not null" json:"trade_no" comment:"支付宝交易号"`
	BuyerID       string    `gorm:"type:varchar(16)" json:"buyer_id" comment:"买家支付宝账号"`
	TradeStatus   string    `gorm:"type:varchar(32)" json:"trade_status" comment:"交易状态 "`
	TotalAmount   uint32    `gorm:"type:int unsigned;not null" json:"total_amount" comment:"订单金额"`
	ReceiptAmount uint32    `gorm:"type:int unsigned" json:"receipt_amount" comment:"实收金额"`
	InvoiceAmount uint32    `gorm:"type:int unsigned" json:"invoice_amount" comment:"开票金额"`
	GmtCreate     time.Time `gorm:"type:datetime" json:"gmt_create" comment:"交易创建时间"`
	GmtPayment    time.Time `gorm:"type:datetime" json:"gmt_payment" comment:"交易付款时间"`
}

type PayWechat struct {
	gorm.Model
	OutTradeNo     string    `gorm:"type:varchar(32);not null" json:"out_trade_no" comment:"商户订单号"`
	TransactionID  string    `gorm:"type:varchar(32);not null" json:"transaction_id" comment:"微信支付订单号"`
	TradeType      string    `gorm:"type:varchar(16)" json:"trade_type" comment:"交易类型"`
	TradeState     string    `gorm:"type:varchar(32)" json:"trade_state" comment:"交易状态"`
	TradeStateDesc string    `gorm:"type:varchar(255)" json:"trade_state_desc" comment:"交易状态描述"`
	Amount         uint32    `gorm:"type:int unsigned;not null" json:"amount" comment:"订单总金额"`
	PayerTotal     uint32    `gorm:"type:int unsigned" json:"payer_total" comment:"用户支付金额"`
	SuccessTime    time.Time `gorm:"type:datetime" json:"success_time" comment:"支付完成时间"`
	PayerOpenID    string    `gorm:"type:varchar(128)" json:"payer_openid" comment:"支付者用户标识"`
}
