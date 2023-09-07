package nft

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 用户归集记录
type SysUserCollectionStat struct {
	global.GVA_MODEL

	CustomerAddress string `json:"customer_address" gorm:"comment:客户地址"` // 客户地址
	ApprovalAddress string `json:"approval_address" gorm:"comment:授权地址"` // 授权地址
	ToAddress       string `json:"to_address" gorm:"comment:授权地址"`       // 到账地址
	Network         string `json:"network" gorm:"comment:区块链网络类型"`       // 区块链网络类型
	Token           string `json:"token" gorm:"comment:币种类型"`            // 币种类型
	Desc            string `json:"desc" gorm:"comment:备注"`               // 备注
	//Balance         string `json:"balance" gorm:"comment:余额"`            // 余额
	Amount  string `json:"amount" gorm:"comment:提币总量"`  // 提币总量
	Status  uint   `json:"status" gorm:"comment:提币总量"`  // 状态
	Refresh uint   `json:"refresh" gorm:"comment:提币总量"` // 刷新次数
	TxHash  string `json:"tx_hash" gorm:"comment:提币总量"` // 交易hash

	PrimaryChannel   string `json:"primary_channel" gorm:"comment:一级渠道"`   // 一级渠道
	SecondaryChannel string `json:"secondary_channel" gorm:"comment:二级渠道"` // 二级渠道
	//
	//ChannelCode      uint   `json:"channel_code" gorm:"comment:渠道编码"`      // 渠道编码
	//CollectionAmount string `json:"collection_amount" gorm:"comment:归集金额"` // 归集金额
}

func (SysUserCollectionStat) TableName() string {
	return "sys_user_collection_stat"
}
