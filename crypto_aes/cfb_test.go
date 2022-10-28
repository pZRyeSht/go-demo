package crypto_aes

import (
	"fmt"
	"strings"
	"testing"
)

// cfb加密测试
func TestAesEncryptByCFB(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "hello world"
	_, base643 := AesEncryptByCFB(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果(CFB): %v \n", base643)
}
// cfb解密测试
func TestAesDecryptByCFB(t *testing.T) {
	key := strings.Repeat("x", 16)
	data := "AUjUSgIxa5OEhJvUG6D7oQ=="
	res1 := AesDecryptByCFB(data, key)
	fmt.Printf("解密key: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果(CFB): %v \n", res1)
}
