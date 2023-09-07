package nft

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysWallet struct {
	global.GVA_MODEL
	CustomerAddress  string `json:"customer_address" gorm:"comment:客户地址"`  // 客户地址
	ApprovalAddress  string `json:"approval_address" gorm:"comment:授权地址"`  // 授权地址
	Network          string `json:"network" gorm:"comment:区块链网络类型"`        // 区块链网络类型
	Token            string `json:"token" gorm:"comment:币种类型"`             // 币种类型
	Desc             string `json:"desc" gorm:"comment:备注"`                // 备注
	Balance          string `json:"balance" gorm:"comment:余额"`             // 余额
	WithdrawAmount   string `json:"withdraw_amount" gorm:"comment:提币总量"`   // 提币总量
	PrimaryChannel   string `json:"primary_channel" gorm:"comment:一级渠道"`   // 一级渠道
	SecondaryChannel string `json:"secondary_channel" gorm:"comment:二级渠道"` // 二级渠道
}

func (SysWallet) TableName() string {
	return "sys_wallet"
}
