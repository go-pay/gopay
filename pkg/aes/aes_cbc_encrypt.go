package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

// 加密后的Bytes数组
func CBCEncryptData(originData, key []byte) ([]byte, error) {
	return encrypt(originData, key)
}

// 加密后的Bytes数组
func CBCEncryptIvData(originData, key, iv []byte) ([]byte, error) {
	return encryptIv(originData, key, iv)
}

func encrypt(originData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	originData = PKCS7Padding(originData, blockSize)
	secretData := make([]byte, len(originData))
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}

func encryptIv(originData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCEncrypter(block, iv[:block.BlockSize()])

	originData = PKCS7Padding(originData, block.BlockSize())
	secretData := make([]byte, len(originData))
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}
