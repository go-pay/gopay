package chinaunionpay

import (
	"math/rand"
	"time"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

/**
@param dataLen: 数据长度，base64膨胀率1.33,返回的string长度是dataLen的1.33倍
*/
func NewRandomBase64(strLen int) string {
	ba := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		ba[i] = encodeURL[rand.Int()%len(encodeURL)]
	}

	return string(ba)
}

func NewRandomHex(strLen int) string {
	ba := make([]byte, strLen)
	baseStr := []byte("0123456789abcdef")
	for i := 0; i < strLen; i++ {
		ba[i] = baseStr[rand.Int()%len(baseStr)]
	}

	return string(ba)
}
