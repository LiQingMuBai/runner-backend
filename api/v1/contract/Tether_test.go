package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func Test_Print(t *testing.T) {

	//var y *big.Int
	//
	//y.SetInt64(1000000)
	//
	//var x *big.Int
	//
	//result := x.Div(x.SetInt64(647742484), y.SetInt64(1000000))
	//
	//log.Println(result.Int64())

	var balance int64 = 2110000

	log.Println(float64(balance) / 1000000)

}

func Test_TransferAToB(t *testing.T) {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/dab126a4e1f444569c8f517a42cddda2")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	store, err := NewTetherTokenTransactor(common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Storage contract: %v", err)
	}

	privateKey, err := crypto.HexToECDSA("56eaadc26990a5836ea890ff7a262703a4793fb60cfdb51af66c7ea17443b25b")
	if err != nil {
		log.Fatal(err)
	}

	nonce, _ := conn.NonceAt(context.Background(), common.HexToAddress("0xFF8C78F54235D04f88378f98f5ec6Fa68802f1b4"), nil)
	gasPrice, _ := conn.SuggestGasPrice(context.Background())
	//用哪条链，就用那个id
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	auth.GasLimit = uint64(300000)
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.GasPrice = gasPrice
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	log.Println("Nonce: ", auth.Nonce)
	store2, err := NewTetherTokenCaller(common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Storage contract: %v", err)
	}
	amount, err := store2.BalanceOf(nil, common.HexToAddress("0xE3DB91F3B07F282f9890aE52e7fC0Bd2Adee2C74"))
	log.Println(amount)
	//// Call the store() function

	tx, err := store.TransferFrom(auth, common.HexToAddress("0xE3DB91F3B07F282f9890aE52e7fC0Bd2Adee2C74"), common.HexToAddress("0xFF8C78F54235D04f88378f98f5ec6Fa68802f1b4"), big.NewInt(3000000))

	if err != nil {
		log.Println(err)
	}
	if err == nil {
		log.Println(tx.Hash())
	}
}
func TestTetherTokenFilterer_WatchApproval(t *testing.T) {

	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/dab126a4e1f444569c8f517a42cddda2")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	store, err := NewTetherTokenCaller(common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Storage contract: %v", err)
	}
	amount, err := store.BalanceOf(nil, common.HexToAddress("0xE3DB91F3B07F282f9890aE52e7fC0Bd2Adee2C74"))
	log.Println(amount)
	approvalAmount, err := store.Allowed(nil, common.HexToAddress("0xE3DB91F3B07F282f9890aE52e7fC0Bd2Adee2C74"), common.HexToAddress("0xFF8C78F54235D04f88378f98f5ec6Fa68802f1b4"))
	log.Println(approvalAmount)

	log.Println(approvalAmount.BitLen())
	if approvalAmount.BitLen() > 0 {
		log.Println("approval ulimit")
	} else {

		log.Println("111111111111111111")
	}
	//// Create an authorized transactor and call the store function
	//nonce, _ := conn.NonceAt(context.Background(), common.HexToAddress("你私钥对应的账户地址"), nil)
	//gasPrice, _ := conn.SuggestGasPrice(context.Background())
	////用哪条链，就用那个id
	//auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, big.NewInt(5))
	//auth.GasLimit = uint64(300000)
	//auth.Nonce = new(big.Int).SetUint64(nonce)
	//auth.GasPrice = gasPrice
	//if err != nil {
	//	log.Fatalf("Failed to create authorized transactor: %v", err)
	//}
	//// Call the store() function
	//tx, err := store.Store(auth, big.NewInt(420))
	//if err != nil {
	//	log.Fatalf("Failed to update value: %v", err)
	//}
	//fmt.Printf("Update pending: 0x%x\n", tx.Hash())

}
