package wechat

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"reflect"
	"strings"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
)

// VerifySign 微信同步返回参数验签或异步通知参数验签
// ApiKey：API秘钥值
// signType：签名类型（调用API方法时填写的类型）
// bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq，推荐通 BodyMap 验签
// 返回参数ok：是否验签通过
// 返回参数err：其他错误信息，不要根据 error 是否为空来判断验签正确与否，需再单独判断返回的 ok
func VerifySign(apiKey, signType string, bean any) (ok bool, err error) {
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	kind := reflect.ValueOf(bean).Kind()
	if kind == reflect.Map {
		bm := bean.(gopay.BodyMap)
		bodySign := bm.GetString("sign")
		bm.Remove("sign")
		return GetReleaseSign(apiKey, signType, bm) == bodySign, nil
	}

	bs, err := json.Marshal(bean)
	if err != nil {
		return false, fmt.Errorf("json.Marshal(%s): %w", string(bs), err)
	}
	bm := make(gopay.BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return false, fmt.Errorf("json.Marshal(%s): %w", string(bs), err)
	}
	bodySign := bm.GetString("sign")
	bm.Remove("sign")
	return GetReleaseSign(apiKey, signType, bm) == bodySign, nil
}

// GetMiniPaySign JSAPI支付，统一下单获取支付参数后，再次计算出小程序用的paySign
// appId：APPID
// nonceStr：随即字符串
// packages：统一下单成功后拼接得到的值
// signType：签名类型
// timeStamp：时间
// ApiKey：API秘钥值
// 微信小程序支付API：https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html
// 微信小程序支付PaySign计算文档：https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=3
func GetMiniPaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
	var (
		buffer strings.Builder
		h      hash.Hash
	)
	buffer.WriteString("appId=")
	buffer.WriteString(appId)
	buffer.WriteString("&nonceStr=")
	buffer.WriteString(nonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString(packages)
	buffer.WriteString("&signType=")
	buffer.WriteString(signType)
	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timeStamp)
	buffer.WriteString("&key=")
	buffer.WriteString(apiKey)
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(buffer.String()))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// Deprecated
// 微信内H5支付官方文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
func GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
	return GetJsapiPaySign(appId, nonceStr, packages, signType, timeStamp, apiKey)
}

// GetJsapiPaySign JSAPI调起支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
// appId：APPID
// nonceStr：随即字符串
// packages：统一下单成功后拼接得到的值
// signType：签名类型
// timeStamp：时间
// ApiKey：API秘钥值
// 文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
func GetJsapiPaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
	var (
		buffer strings.Builder
		h      hash.Hash
	)
	buffer.WriteString("appId=")
	buffer.WriteString(appId)
	buffer.WriteString("&nonceStr=")
	buffer.WriteString(nonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString(packages)
	buffer.WriteString("&signType=")
	buffer.WriteString(signType)
	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timeStamp)
	buffer.WriteString("&key=")
	buffer.WriteString(apiKey)
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(buffer.String()))
	paySign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// GetAppPaySign APP支付，统一下单获取支付参数后，再次计算APP支付所需要的的sign
// appId：APPID
// partnerid：partnerid
// nonceStr：随即字符串
// prepayId：统一下单成功后得到的值
// signType：此处签名方式，务必与统一下单时用的签名方式一致
// timeStamp：时间
// ApiKey：API秘钥值
// 文档：https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12&index=2
func GetAppPaySign(appid, partnerid, noncestr, prepayid, signType, timestamp, apiKey string) (paySign string) {
	var (
		buffer strings.Builder
		h      hash.Hash
	)
	buffer.WriteString("appid=")
	buffer.WriteString(appid)
	buffer.WriteString("&noncestr=")
	buffer.WriteString(noncestr)
	buffer.WriteString("&package=Sign=WXPay")
	buffer.WriteString("&partnerid=")
	buffer.WriteString(partnerid)
	buffer.WriteString("&prepayid=")
	buffer.WriteString(prepayid)
	buffer.WriteString("&timestamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&key=")
	buffer.WriteString(apiKey)
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(buffer.String()))
	paySign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// Deprecated
// GetParamSign 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
// 注意：BodyMap中如无 sign_type 参数，默认赋值 sign_type 为 MD5
// appId：应用ID
// mchId：商户ID
// ApiKey：API秘钥值
// 返回参数 sign：通过Appid、MchId、ApiKey和BodyMap中的参数计算出的Sign值
func GetParamSign(appId, mchId, apiKey string, bm gopay.BodyMap) (sign string) {
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	var (
		signType string
		h        hash.Hash
	)
	signType = bm.GetString("sign_type")
	if signType == gopay.NULL {
		bm.Set("sign_type", SignType_MD5)
	}
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(bm.EncodeWeChatSignParams(apiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// Deprecated
// GetSanBoxParamSign 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
// 注意：沙箱环境默认 sign_type 为 MD5
// appId：应用ID
// mchId：商户ID
// ApiKey：API秘钥值
// 返回参数 sign：通过Appid、MchId、ApiKey和BodyMap中的参数计算出的Sign值
func GetSanBoxParamSign(ctx context.Context, appId, mchId, apiKey string, bm gopay.BodyMap) (sign string, err error) {
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	bm.Set("sign_type", SignType_MD5)
	bm.Set("total_fee", 101)
	var (
		sandBoxApiKey string
		hashMd5       hash.Hash
	)
	if sandBoxApiKey, err = getSanBoxKey(ctx, mchId, util.RandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	hashMd5 = md5.New()
	hashMd5.Write([]byte(bm.EncodeWeChatSignParams(sandBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(hashMd5.Sum(nil)))
	return
}
