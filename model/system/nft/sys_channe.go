package nft

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysChannel struct {
	global.GVA_MODEL
	ChannelCode      uint   `json:"channel_code" gorm:"comment:渠道编码"`      // 渠道编码
	CollectionAmount string `json:"collection_amount" gorm:"comment:归集金额"` // 归集金额
}

func (SysChannel) TableName() string {
	return "sys_channel"
}
