package crypto_aes

import "bytes"

// complement 补码
func complement(originByte []byte, blockSize int) []byte {
	// 计算补码长度
	padding := blockSize - len(originByte)%blockSize
	// 生成补码
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 追加补码
	return append(originByte, padText...)
}

// decipher 解码
func decipher(originDataByte []byte) []byte {
	length := len(originDataByte)
	code := int(originDataByte[length-1])
	return originDataByte[:(length - code)]
}
