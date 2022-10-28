package crypto_aes

import (
	"fmt"
	"strings"
	"testing"
)

// ctr解码测试
func TestAesEncryptByCTR(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "hello world"
	hex, base64 := AesEncryptByCTR(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果(hex): %v \n", hex)
	fmt.Printf("加密结果(base64): %v \n", base64)
}

func TestAesDecryptByCTR(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "AUjUSgIxa5OEhJvUG6D7oQ=="
	res := AesDecryptByCTR(data, key)
	fmt.Printf("解密key: %v \n", key)
	fmt.Printf("解密key长度: %v \n", len(key))
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", res)
}
