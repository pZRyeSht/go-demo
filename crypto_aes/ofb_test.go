package crypto_aes

import (
	"fmt"
	"strings"
	"testing"
)

// ofb加密测试
func TestECBEncryptOFB(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "hello world"
	_, base643 := AesEncryptByCFB(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果(CFB): %v \n", base643)
}

// ofb解密测试
func TestECBDecryptOFB(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "7JD8hYTVGySjnRlKS/bTMQ=="
	s := AesDecryptByOFB(data, key)
	fmt.Printf("解密密钥: %v \n", key)
	fmt.Printf("解密key长度: %v \n", len(key))
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", s)
}
