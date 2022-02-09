package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// 解密数据的Bytes数组
func CBCDecryptData(secretData, key []byte) ([]byte, error) {
	return decrypt(secretData, key)
}

// 解密数据的Bytes数组
func CBCDecryptIvData(secretData, key, iv []byte) ([]byte, error) {
	return decryptIv(secretData, key, iv)
}

func decrypt(secretData, key []byte) (originByte []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originByte = make([]byte, len(secretData))
	blockMode.CryptBlocks(originByte, secretData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return PKCS7UnPadding(originByte), nil
}

func decryptIv(secretData, key, iv []byte) (originByte []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv[:block.BlockSize()])

	originByte = make([]byte, len(secretData))
	blockMode.CryptBlocks(originByte, secretData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return PKCS7UnPadding(originByte), nil
}
