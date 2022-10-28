package crypto_aes

import (
	"fmt"
	"strings"
	"testing"
)

// 加密测试
func TestECBEncrypt(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "hello world"
	s := AesEncryptByECB(data, key)
	fmt.Printf("加密密钥: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果: %v \n", s)
}

// 解密测试
func TestECBDecrypt(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "7JD8hYTVGySjnRlKS/bTMQ=="
	s := AesDecryptByECB(data, key)
	fmt.Printf("解密密钥: %v \n", key)
	fmt.Printf("解密key长度: %v \n", len(key))
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", s)
}
