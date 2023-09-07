package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VictimRouter struct{}

func (s *VictimRouter) InitVictimRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	VictimRouter := Router.Group("victim").Use(middleware.OperationRecord())
	victimRouter := v1.ApiGroupApp.SystemApiGroup.SystemVictim
	victimTxRouter := v1.ApiGroupApp.SystemApiGroup.SystemVictimTx
	{
		VictimRouter.POST("createVictim", victimRouter.CreateVictim)               // 创建Victim
		VictimRouter.POST("createVictimTx", victimTxRouter.CreateVictimTx)         // 创建VictimTx
		VictimRouter.POST("statVictimTx", victimTxRouter.StatVictimTx)             // 创建VictimTx
		VictimRouter.POST("deleteVictim", victimRouter.DeleteVictim)               // 删除Victim
		VictimRouter.POST("getVictimById", victimRouter.GetVictimById)             // 获取单条Victim消息
		VictimRouter.POST("updateVictim", victimRouter.UpdateVictim)               // 更新Victim
		VictimRouter.DELETE("deleteVictimsByIds", victimRouter.DeleteVictimsByIds) // 删除选中Victim
		VictimRouter.POST("getAllVictims", victimRouter.GetAllVictims)             // 获取所有Victim
		VictimRouter.POST("getVictimList", victimRouter.GetVictimList)             // 获取Victim列表
	}

}
