package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

const SECRETKEY = "KoR7JcF947b1HMcD"

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

func AesEncrypt(rawData string, key []byte) (string, error) {
	data, err := AesCBCEncrypt([]byte(rawData), key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func AesDecrypt(rawData string, key []byte) (string, error) {
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

func EncryptPrimaryKey(pk string) (secret string, err error) {
	tm := time.Now().Unix()
	mrand.Seed(tm * 2)
	pk = fmt.Sprintf("%s|%s|%d", RandomString(mrand.Intn(10)), pk, tm)
	secret, err = AesEncrypt(pk, []byte(SECRETKEY))
	if err != nil {
		return
	}
	return
}

func DecryptPrimaryKey(secret string) (pk string, err error) {
	pks, err := AesDecrypt(secret, []byte(SECRETKEY))
	m := strings.SplitN(pks, "|", 3)
	if len(m) == 3 {
		if tm, err := strconv.Atoi(m[2]); err == nil && int64(tm)+3600 > time.Now().Unix() {
			pk = m[1]
			return pk, err
		}
	}
	err = errors.New("wrong format secret string")
	return
}
