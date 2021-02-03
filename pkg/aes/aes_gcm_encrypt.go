package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/iGoogle-ink/gopay/pkg/util"
)

func GCMEncrypt(originText, additional, key []byte) (nonce []byte, cipherText []byte, err error) {
	return gcmEncrypt(originText, additional, key)
}

func gcmEncrypt(originText, additional, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	nonce := []byte(util.GetRandomString(12))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, fmt.Errorf("cipher.NewGCM(),error:%w", err)
	}

	cipherBytes := gcm.Seal(nil, nonce, originText, additional)

	return nonce, cipherBytes, nil
}
