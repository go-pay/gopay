package gopay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"github.com/parnurzeal/gorequest"
	"sort"
	"strings"
)

//本地通过支付参数计算Sign值
func getLocalSign(apiKey string, signType string, body BodyMap) (sign string) {
	signStr := sortWeChatSignParams(apiKey, body)
	//fmt.Println("signStr:", signStr)
	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(apiKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//获取根据Key排序后的请求参数字符串
func sortWeChatSignParams(apiKey string, body BodyMap) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(k)
		buffer.WriteString("=")

		valueStr := convert2String(body[k])
		buffer.WriteString(valueStr)

		buffer.WriteString("&")
	}
	buffer.WriteString("key=")
	buffer.WriteString(apiKey)
	return buffer.String()
}

//从微信提供的接口获取：SandboxSignKey
func getSanBoxSign(mchId, nonceStr, apiKey, signType string) (key string, err error) {
	body := make(BodyMap)
	body.Set("mch_id", mchId)
	body.Set("nonce_str", nonceStr)

	//计算沙箱参数Sign
	sanboxSign := getLocalSign(apiKey, signType, body)
	//沙箱环境：获取key后，重新计算Sign
	key, err = getSanBoxSignKey(mchId, nonceStr, sanboxSign)
	if err != nil {
		return "", err
	}
	return
}

//从微信提供的接口获取：SandboxSignkey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	reqXml := generateXml(reqs)
	//fmt.Println("req:::", reqXml)
	_, byteList, errorList := gorequest.New().
		Post(wxURL_SanBox_GetSignKey).
		Type("xml").
		SendString(reqXml).EndBytes()
	if len(errorList) > 0 {
		return "", errorList[0]
	}
	keyResponse := new(getSignKeyResponse)
	err = xml.Unmarshal(byteList, keyResponse)
	if err != nil {
		return "", err
	}
	if keyResponse.ReturnCode == "FAIL" {
		return "", errors.New(keyResponse.Retmsg)
	}
	return keyResponse.SandboxSignkey, nil
}

//生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")

	for k, v := range bm {
		buffer.WriteString("<")
		buffer.WriteString(k)
		buffer.WriteString("><![CDATA[")

		valueStr := convert2String(v)
		buffer.WriteString(valueStr)

		buffer.WriteString("]]></")
		buffer.WriteString(k)
		buffer.WriteString(">")
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
