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

//获取请求支付的参数
func getRequestBody(appId, mchId string, params *WeChatPayParams) (reqs BodyMap) {
	reqs = make(BodyMap)
	reqs.Set("appid", appId)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", params.NonceStr)
	if params.Body != "" {
		reqs.Set("body", params.Body)
	}
	if params.OutTradeNo != "" {
		reqs.Set("out_trade_no", params.OutTradeNo)
	}
	if params.TotalFee != -1 {
		reqs.Set("total_fee", strconv.Itoa(params.TotalFee))
	}
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
		reqs.Set("scene_info", params.SceneInfo)
	}
	if params.TransactionId != "" {
		reqs.Set("transaction_id", params.TransactionId)
	}
	return reqs
}

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
