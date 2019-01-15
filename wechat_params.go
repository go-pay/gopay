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
	"strconv"
	"strings"
)

//获取请求支付的参数
func (w *WeChatPayParams) getRequestBody(appId, mchId string, params *WeChatPayParams) (reqs requestBody) {
	reqs = make(requestBody)
	reqs.Set("appid", appId)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", params.NonceStr)
	reqs.Set("body", params.Body)
	reqs.Set("out_trade_no", params.OutTradeNo)
	reqs.Set("total_fee", strconv.Itoa(params.TotalFee))
	reqs.Set("spbill_create_ip", params.SpbillCreateIp)
	reqs.Set("notify_url", params.NotifyUrl)
	reqs.Set("trade_type", params.TradeType)

	if params.DeviceInfo != "" {
		reqs.Set("device_info", params.DeviceInfo)
	}
	if params.SignType != "" {
		reqs.Set("sign_type", params.SignType)
	} else {
		reqs.Set("sign_type", "MD5")
	}
	if params.Detail != "" {
		reqs.Set("detail", params.Detail)
	}
	if params.Attach != "" {
		reqs.Set("attach", params.Attach)
	}
	if params.FeeType != "" {
		reqs.Set("fee_type", params.FeeType)
	}
	if params.TimeStart != "" {
		reqs.Set("time_start", params.TimeStart)
	}
	if params.TimeExpire != "" {
		reqs.Set("time_expire", params.TimeExpire)
	}
	if params.GoodsTag != "" {
		reqs.Set("goods_tag", params.GoodsTag)
	}
	if params.ProductId != "" {
		reqs.Set("product_id", params.ProductId)
	}
	if params.LimitPay != "" {
		reqs.Set("limit_pay", params.LimitPay)
	}
	if params.Openid != "" {
		reqs.Set("openid", params.Openid)
	}
	if params.Receipt != "" {
		reqs.Set("receipt", params.Receipt)
	}
	if params.SceneInfo != "" {
		//marshal, _ := json.Marshal(params.SceneInfo)
		//reqs.Set("scene_info", string(marshal))
		reqs.Set("scene_info", params.SceneInfo)
	}
	return reqs
}

//获取SanBox秘钥
func (w *WeChatPayParams) getSanBoxSignKey(mchId, nonceStr, secretKey, signType string) (key string, err error) {
	body := make(requestBody)
	body.Set("mch_id", mchId)
	body.Set("nonce_str", nonceStr)

	//计算沙箱参数Sign
	sanboxSign := getSign(secretKey, signType, body)
	//沙箱环境：获取key后，重新计算Sign
	key, err = getSanBoxSignKey(mchId, nonceStr, sanboxSign)
	if err != nil {
		return "", err
	}
	return
}

//获取Sign签名和请求支付的参数
func getSign(secretKey string, signType string, body requestBody) (sign string) {

	signStr := getSignString(secretKey, body)
	//fmt.Println("signStr:", signStr)
	var hashSign []byte
	if signType == WX_SignType_MD5 {

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

//获取根据Key排序后的请求参数字符串
func getSignString(secretKey string, body requestBody) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(body[k])
		buffer.WriteString("&")
	}
	buffer.WriteString("key=")
	buffer.WriteString(secretKey)
	return buffer.String()
}

//获取SanboxKey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(requestBody)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	reqXml := generateXml(reqs)
	//fmt.Println("req:::", reqXml)
	_, byteList, errorList := gorequest.New().
		Post(wxURL_sanbox_getsignkey).
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
