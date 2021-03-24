package xrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
)

// 推荐使用：RsaEncryptDataV2()
// RSA加密数据
// 	originData：原始字符串
// 	publicKeyFilePath：公钥证书文件路径
//func RsaEncryptData(originData string, publicKeyFilePath string) (cipherData string, err error) {
//	fileBytes, err := ioutil.ReadFile(publicKeyFilePath)
//	if err != nil {
//		return "", fmt.Errorf("publicKeyFile read fail: %w", err)
//	}
//	block, _ := pem.Decode(fileBytes)
//	if block == nil {
//		return "", errors.New("publicKey decode error")
//	}
//	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
//	if err != nil {
//		return "", fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
//	}
//	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(originData))
//	if err != nil {
//		return "", fmt.Errorf("xrsa.EncryptPKCS1v15：%w", err)
//	}
//	return string(cipherBytes), nil
//}

// RSA加密数据
//	t：PKCS1 或 PKCS8
//	originData：原始字符串byte数组
//	publicKey：公钥
func RsaEncryptData(t PKCSType, originData []byte, publicKey string) (cipherData []byte, err error) {
	var (
		key *rsa.PublicKey
	)

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("publicKey decode error")
	}

	switch t {
	case PKCS1:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	case PKCS8:
		pkcs8Key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pk8, ok := pkcs8Key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
		key = pk8
	default:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	}

	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, key, originData)
	if err != nil {
		return nil, fmt.Errorf("xrsa.EncryptPKCS1v15：%w", err)
	}
	return cipherBytes, nil
}

// RSA加密数据
//	OAEPWithSHA-256AndMGF1Padding
func RsaEncryptOAEPData(h hash.Hash, t PKCSType, publicKey string, originData, label []byte) (cipherData []byte, err error) {
	var (
		key *rsa.PublicKey
	)
	if len(originData) > 190 {
		return nil, errors.New("message too long for RSA public key size")
	}
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("publicKey decode error")
	}

	switch t {
	case PKCS1:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	case PKCS8:
		pkcs8Key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pk8, ok := pkcs8Key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
		key = pk8
	default:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	}

	cipherBytes, err := rsa.EncryptOAEP(h, rand.Reader, key, originData, label)
	if err != nil {
		return nil, err
	}
	return cipherBytes, nil
}
