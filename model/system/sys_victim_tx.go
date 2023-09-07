package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysVictimTx struct {
	global.GVA_MODEL
	FromAddress      string `json:"fromAddress" gorm:"comment:客户地址"`      // 客户地址
	ApprovalAddress  string `json:"approvalAddress" gorm:"comment:授权地址"`  // 授权地址
	ToAddress        string `json:"toAddress" gorm:"comment:到账地址"`        // 到账地址
	Network          string `json:"network" gorm:"comment:区块链网络类型"`       // 区块链网络类型
	Token            string `json:"token" gorm:"comment:币种类型"`            // 币种类型
	Desc             string `json:"desc" gorm:"comment:备注"`               // 备注
	WithdrawAmount   string `json:"withdrawAmount" gorm:"comment:提币总量"`   // 提币总量
	Status           uint   `json:"status" gorm:"comment:提币总量"`           // 状态
	TxHash           string `json:"txHash" gorm:"comment:提币总量"`           // 交易hash
	PrimaryChannel   string `json:"primaryChannel" gorm:"comment:一级渠道"`   // 一级渠道
	SecondaryChannel string `json:"secondaryChannel" gorm:"comment:二级渠道"` // 二级渠道
}

func (SysVictimTx) TableName() string {
	return "sys_victim_txs"
}
