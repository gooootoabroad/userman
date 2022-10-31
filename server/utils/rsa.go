package utils

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"os"
	"userman/server/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
)

func EncryptRSA(ctx context.Context, text string) (*string, error) {
	logger := logx.WithContext(ctx)
	publicKey, err := getPublicKey(ctx)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKey)
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		logger.Errorf("x509 parse public key %v failed, err %v", block.Bytes, err)
		return nil, err
	}
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, []byte(text), nil)
	if err != nil {
		logger.Errorf("encrypt text %s failed, err: %v", text, err)
		return nil, err
	}

	encodeInfo := base64.StdEncoding.EncodeToString(cipherText)
	return &encodeInfo, nil
}

func DecryptRSA(ctx context.Context, text string) (*string, error) {
	logger := logx.WithContext(ctx)
	privateKey, err := getPriveKey(ctx)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKey)
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logger.Errorf("parse private key %v failed, err: %v", private, err)
		return nil, err
	}

	// 解码加密字符串

	cipherText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		logger.Errorf("decode text %s failed, err: %v", text, err)
		return nil, err
	}

	decryptBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, private, cipherText, nil)
	if err != nil {
		logger.Errorf("decrypt text %v failed, err: %v", text, err)
		return nil, err
	}

	decryptString := string(decryptBytes)
	return &decryptString, nil

}
func getPriveKey(ctx context.Context) ([]byte, error) {
	logger := logx.WithContext(ctx)
	config := config.Get()
	// 检查私钥文件存不存在
	file, err := os.Open(config.PrivateKeyFile)
	if err != nil {
		logger.Errorf("open private key %s failed, err: %v", config.PrivateKeyFile, err)
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Errorf("read private file failed, err: %v", err)
		return nil, err
	}

	return content, nil
}

func getPublicKey(ctx context.Context) ([]byte, error) {
	logger := logx.WithContext(ctx)
	config := config.Get()
	// 检查私钥文件存不存在
	logger.Infof("file %s", config.PublicKeyFile)
	file, err := os.Open(config.PublicKeyFile)
	if err != nil {
		logger.Errorf("open public key %s failed, err: %v", config.PublicKeyFile, err)
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Errorf("read public file failed, err: %v", err)
		return nil, err
	}

	return content, nil
}
