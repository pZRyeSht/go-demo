package crypto_aes

import (
	"fmt"
	"strings"
	"testing"
)

func TestAesGCM(t *testing.T) {
	key := strings.Repeat("x",16)
	data := "hello world"
	// 加密
	gcm := AesEncryptByGCM(data, key)
	fmt.Printf("密钥key: %s \n",key)
	fmt.Printf("加密数据: %s \n",data)
	fmt.Printf("加密结果: %s \n",gcm)
	// 解密
	byGCM := AesDecryptByGCM(gcm, key)
	fmt.Printf("解密结果: %s \n",byGCM)
}
