package crypto_aes

import (
	"fmt"
	"strings"
	"testing"
)

// cbc加密
func TestAesEncryptByCBC(t *testing.T) {
	key := strings.Repeat("x", 16)
	fmt.Printf("key: %v\n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	text := "hello world"
	fmt.Printf("data 数据: %v \n", text)
	encrypt := AesEncryptByCBC(text, key)
	fmt.Printf("加密结果: %v \n", encrypt)
}

// cbc解密
func TestAesDecryptByCBC(t *testing.T) {
	key := strings.Repeat("x", 16)
	fmt.Printf("key: %v\n", key)
	fmt.Printf("解密key长度: %v \n", len(key))
	text := "DBE7+U2DiYbeIsHtFt7a3w=="
	fmt.Printf("cipher 数据: %v \n", text)
	decrypt := AesDecryptByCBC(text, key)
	fmt.Printf("解密结果: %v \n", decrypt)
}
