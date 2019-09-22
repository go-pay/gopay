package gopay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"hash"
	"strings"
)

type Country int

//设置支付国家（默认：中国国内）
//    根据支付地区情况设置国家
//    country：<China：中国国内，China2：中国国内（冗灾方案），SoutheastAsia：东南亚，Other：其他国家>
func (w *WeChatClient) SetCountry(country Country) (client *WeChatClient) {
	switch country {
	case China:
		w.baseURL = wx_base_url_ch
	case China2:
		w.baseURL = wx_base_url_ch2
	case SoutheastAsia:
		w.baseURL = wx_base_url_hk
	case Other:
		w.baseURL = wx_base_url_us
	default:
		w.baseURL = wx_base_url_ch
	}
	return w
}

//获取微信支付正式环境Sign值
func getWeChatReleaseSign(apiKey string, signType string, bm BodyMap) (sign string) {
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(bm.EncodeWeChatSignParams(apiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

//获取微信支付沙箱环境Sign值
func getWeChatSignBoxSign(mchId, apiKey string, bm BodyMap) (sign string, err error) {
	var (
		sanBoxApiKey string
		h            hash.Hash
	)
	if sanBoxApiKey, err = getSanBoxKey(mchId, GetRandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	h = md5.New()
	h.Write([]byte(bm.EncodeWeChatSignParams(sanBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

//从微信提供的接口获取：SandboxSignKey
func getSanBoxKey(mchId, nonceStr, apiKey, signType string) (key string, err error) {
	body := make(BodyMap)
	body.Set("mch_id", mchId)
	body.Set("nonce_str", nonceStr)
	//沙箱环境：获取沙箱环境ApiKey
	if key, err = getSanBoxSignKey(mchId, nonceStr, getWeChatReleaseSign(apiKey, signType, body)); err != nil {
		return
	}
	return
}

//从微信提供的接口获取：SandboxSignkey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)
	var (
		byteList    []byte
		errorList   []error
		keyResponse *getSignKeyResponse
	)
	if _, byteList, errorList = HttpAgent().Post(wx_SanBox_GetSignKey).Type("xml").SendString(generateXml(reqs)).EndBytes(); len(errorList) > 0 {
		return null, errorList[0]
	}
	keyResponse = new(getSignKeyResponse)
	if err = xml.Unmarshal(byteList, keyResponse); err != nil {
		return
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
