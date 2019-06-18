//==================================
//  * Name：Jerry
//  * DateTime：2019/6/18 19:24
//  * Desc：
//==================================
package gopay

import "bytes"

//格式化秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")

	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen

	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

//格式化秘钥
func FormatAliPayPublicKey(publickKey string) (pKey string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")

	rawLen := 64
	keyLen := len(publickKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen

	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publickKey[start:])
		} else {
			buffer.WriteString(publickKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}
