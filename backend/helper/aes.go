package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const SECRETKEY = "KoR7JcF947b1HMcD"

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

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
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
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func EncryptPrimaryKey(pk string) (secret string, err error) {
	tm := time.Now().Unix()
	rand.Seed(tm * 2)
	pk = fmt.Sprintf("%s|%s|%d", RandomString(rand.Intn(10)), pk, tm)
	result, err := AesEncrypt([]byte(pk), []byte(SECRETKEY))
	if err != nil {
		return
	}
	secret = Quote(base64.StdEncoding.EncodeToString(result))
	return
}

func DecryptPrimaryKey(secret string) (pk string, err error) {
	pks, err := AesDecrypt([]byte(secret), []byte(SECRETKEY))
	m := bytes.SplitAfterN(pks, []byte("|"), 3)
	if len(m) == 3 {
		if tm, err := strconv.Atoi(string(m[2])); err == nil && int64(tm)+3600 < time.Now().Unix() {
			pk = string(m[1])
			return pk, err
		}
	}
	err = errors.New("wrong format secret string")
	return
}
