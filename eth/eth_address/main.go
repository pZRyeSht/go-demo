package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
)

func main() {
	// 创建私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKey:", privateKey)
	// 私钥导出
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 公钥派生
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKeyBytes:", publicKeyBytes)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("pri:", hexutil.Encode(privateKeyBytes)[2:])
	fmt.Println("publicKeyBytes Encode:", hexutil.Encode(publicKeyBytes))
	fmt.Println("pub:", hexutil.Encode(publicKeyBytes)[4:])
	fmt.Println("addr:", hexutil.Encode(hash.Sum(nil)[12:]))
	// get client
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	// get balance
	account := common.HexToAddress(crypto.PubkeyToAddress(*publicKeyECDSA).Hex())
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance:", balance)
}
