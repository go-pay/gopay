package gopay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"strings"
)

type Country int

//设置支付国家（默认：中国国内）
//    根据支付地区情况设置国家
//    country：<China：中国国内，China2：中国国内（冗灾方案），SoutheastAsia：东南亚，Other：其他国家>
func (this *weChatClient) SetCountry(country Country) (client *weChatClient) {
	switch country {
	case China:
		this.baseURL = wx_base_url_ch
	case China2:
		this.baseURL = wx_base_url_ch2
	case SoutheastAsia:
		this.baseURL = wx_base_url_hk
	case Other:
		this.baseURL = wx_base_url_us
	default:
		this.baseURL = wx_base_url_ch
	}
	return this
}

//本地通过支付参数计算Sign值
func getLocalSign(apiKey string, signType string, bm BodyMap) (sign string) {
	signStr := bm.EncodeWeChatSignParams(apiKey)
	//fmt.Println("signStr:", signStr)
	var hashSign []byte
	if signType == SignType_HMAC_SHA256 {
		hash := hmac.New(sha256.New, []byte(apiKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
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
		return null, err
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
	_, byteList, errorList := HttpAgent().
		Post(wx_SanBox_GetSignKey).
		Type("xml").
		SendString(reqXml).EndBytes()
	if len(errorList) > 0 {
		return null, errorList[0]
	}
	keyResponse := new(getSignKeyResponse)
	err = xml.Unmarshal(byteList, keyResponse)
	if err != nil {
		return null, err
	}
	if keyResponse.ReturnCode == "FAIL" {
		return null, errors.New(keyResponse.ReturnMsg)
	}
	return keyResponse.SandboxSignkey, nil
}

//生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	var buffer strings.Builder
	buffer.WriteString("<xml>")

	for key := range bm {
		buffer.WriteByte('<')
		buffer.WriteString(key)
		buffer.WriteString("><![CDATA[")
		buffer.WriteString(bm.Get(key))
		buffer.WriteString("]]></")
		buffer.WriteString(key)
		buffer.WriteByte('>')
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
