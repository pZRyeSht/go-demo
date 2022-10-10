package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
)

func main() {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-------------------------------")
	prvKey, pubKey := GenerateRsaKey(1024)
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-------------------------------")
	var data = "789399988a38eb7e5e9b5462410020a5f2b4549928fa8bba4e5c29c1e9c6887d"
	fmt.Println("对消息进行签名操作...")
	signData := RsaSignWithSha256([]byte(data), prvKey)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		fmt.Println("签名信息验证成功！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-------------------------------")
	fmt.Println("加密前的数据：", data)
	ciphertext := RsaEncrypt([]byte(data), pubKey)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
}

// GenerateRsaKey 生成rsa的密钥对
func GenerateRsaKey(keySize int) (prvKey, pubKey []byte) {
	// ============ 私钥 ============
	// 1. 使用 rsa 中的 GenerateKey 方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize) // keySize (i.e. 1024)
	if err != nil {
		panic(err)
	}
	// 2. 通过 x509 标准将得到的 ras 私钥序列化为 ASN.1 的 DER 编码字符串
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 组织一个 pem.Block (base64编码)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY", // The type, taken from the preamble (i.e. "RSA PRIVATE KEY").
		Bytes: derStream,
	}
	// 4. pem 编码
	prvKey = pem.EncodeToMemory(block)
	// 5. 同时可选择写入文件
	//file, err := os.Create("private.pem")
	//if err != nil {
	//	panic(err)
	//}
	//_ = pem.Encode(file, block)
	//_ = file.Close()

	// ============ 公钥 ============
	// 1.私钥派生公钥
	publicKey := &privateKey.PublicKey
	// 2. 使用x509标准序列化
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey = pem.EncodeToMemory(block)
	//file, err = os.Create("public.pem")
	//if err != nil {
	//	panic(err)
	//}
	//_ = pem.Encode(file, block)
	//_ = file.Close()
	return prvKey, pubKey
}

// RsaSignWithSha256 签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

// RsaVerySignWithSha256 签名验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// RsaEncrypt 公钥加密
func RsaEncrypt(data, keyBytes []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// RsaDecrypt 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}

// SegmentRSAEncrypt 分段Rsa 公钥加密
func SegmentRSAEncrypt(plainText []byte, length int, pubK []byte) []byte {
	var ret = make([]string, 0)

	for i := 0; i < len(plainText); i++ {
		if len(plainText)-(length*(i+1)) > 0 {
			data := plainText[i*length : length*(i+1)]
			ret = append(ret, string(data))
		} else {
			data := plainText[len(ret)*length:]
			ret = append(ret, string(data))
			break
		}
	}
	var rsaEncryptRes string
	for _, v := range ret {
		rsaEncodeV := RsaEncrypt([]byte(v), pubK)

		encodeV := base64.StdEncoding.EncodeToString(rsaEncodeV)

		rsaEncryptRes = rsaEncryptRes + encodeV
	}
	return []byte(rsaEncryptRes)
}

// SegmentRSADecrypt 分段Rsa 私钥解密
func SegmentRSADecrypt(ciphertext []byte, splitLen int, priK []byte) []byte {
	var ret = make([]string, 0)
	for i := 0; i < len(ciphertext); i++ {
		if len(ciphertext)-(splitLen*(i+1)) > 0 {
			data := ciphertext[i*splitLen : splitLen*(i+1)]
			ret = append(ret, string(data))
		} else {
			data := ciphertext[len(ret)*splitLen:]
			ret = append(ret, string(data))
			break
		}
	}
	var rsaDecryptRes string
	for _, v := range ret {
		decodeV, _ := base64.StdEncoding.DecodeString(v)

		resDecodeV := RsaDecrypt(decodeV, priK)

		rsaDecryptRes = rsaDecryptRes + string(resDecodeV)
	}

	return []byte(rsaDecryptRes)
}
