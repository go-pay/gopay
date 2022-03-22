package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AES-CBC 加密数据
func CBCEncrypt(originData, key, iv []byte) ([]byte, error) {
	return cbcEncrypt(originData, key, iv)
}

// AES-CBC 解密数据
func CBCDecrypt(secretData, key, iv []byte) ([]byte, error) {
	return cbcDecrypt(secretData, key, iv)
}

func cbcEncrypt(originData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originData = PKCS7Padding(originData, block.BlockSize())
	secretData := make([]byte, len(originData))
	blockMode := cipher.NewCBCEncrypter(block, iv[:block.BlockSize()])
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}

func cbcDecrypt(secretData, key, iv []byte) (originByte []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originByte = make([]byte, len(secretData))
	blockMode := cipher.NewCBCDecrypter(block, iv[:block.BlockSize()])
	blockMode.CryptBlocks(originByte, secretData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return PKCS7UnPadding(originByte), nil
}
