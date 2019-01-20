package gopay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

//本地通过支付参数计算Sign值
func getLocalSign(secretKey string, signType string, body BodyMap) (sign string) {
	signStr := sortSignParams(secretKey, body)
	//fmt.Println("signStr:", signStr)
	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(secretKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//从微信提供的接口获取：SandboxSignKey
func getSanBoxSign(mchId, nonceStr, secretKey, signType string) (key string, err error) {
	body := make(BodyMap)
	body.Set("mch_id", mchId)
	body.Set("nonce_str", nonceStr)

	//计算沙箱参数Sign
	sanboxSign := getLocalSign(secretKey, signType, body)
	//沙箱环境：获取key后，重新计算Sign
	key, err = getSanBoxSignKey(mchId, nonceStr, sanboxSign)
	if err != nil {
		return "", err
	}
	return
}

//生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")

	for k, v := range bm {
		buffer.WriteString("<")
		buffer.WriteString(k)
		buffer.WriteString("><![CDATA[")
		value, ok := v.(int)
		if ok {
			value := strconv.Itoa(value)
			buffer.WriteString(value)
		} else {
			buffer.WriteString(v.(string))
		}
		buffer.WriteString("]]></")
		buffer.WriteString(k)
		buffer.WriteString(">")
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
