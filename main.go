package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	//	"github.com/axgle/mahonia"
)

func main() {
	var content string = "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"
	var test string = "123456"
	var key string = "94397012af00cc495d8e9f6a4bdc5a34"
	//bcontent := []byte(content)
	decodeBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeBytes)
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(keyBytes)
	fmt.Println(len(keyBytes))
	res, err := Dncrypt(content, keyBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	entest, err := Encrypt([]byte(test), []byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(entest))
	testres, err := Dncrypt(entest, []byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(testres))
}
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上
*/

//使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	rawData = PKCS7Padding(rawData, blockSize)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	//block大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func AesCBCDncrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	//解填充
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

func Encrypt(rawData, key []byte) (string, error) {
	data, err := AesCBCEncrypt(rawData, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Dncrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := AesCBCDncrypt(data, key)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}
