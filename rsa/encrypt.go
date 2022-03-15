package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// Encrypt rsa加密
func Encrypt(plainText []byte, fileName string) []byte {
	//1.打开公钥文件
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
	file.Close()
	//2.pem decode
	block, _ := pem.Decode(buf)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	pubKey := publicKey.(*rsa.PublicKey)
	//3.使用公钥加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}
