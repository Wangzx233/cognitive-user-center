package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// Decrypt rsa解
func Decrypt(cipherText []byte, fileName string) []byte {
	//1.打开私钥文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}
	//2.pem decode
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//3.解密数据
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}
