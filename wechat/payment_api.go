/*
	微信支付
	文档：https://pay.weixin.qq.com/wiki/doc/api/index.html
*/

package wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/iGoogle-ink/gopay"
	xaes "github.com/iGoogle-ink/gopay/pkg/aes"
	"github.com/iGoogle-ink/gopay/pkg/util"
	"github.com/iGoogle-ink/gopay/pkg/xhttp"
)

// GetParamSign 获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
//	注意：BodyMap中如无 sign_type 参数，默认赋值 sign_type 为 MD5
//	appId：应用ID
//	mchId：商户ID
//	ApiKey：API秘钥值
//	返回参数 sign：通过Appid、MchId、ApiKey和BodyMap中的参数计算出的Sign值
func GetParamSign(appId, mchId, apiKey string, bm gopay.BodyMap) (sign string) {
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	var (
		signType string
		h        hash.Hash
	)
	signType = bm.GetString("sign_type")
	if signType == util.NULL {
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

// GetSanBoxParamSign 获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
//	注意：沙箱环境默认 sign_type 为 MD5
//	appId：应用ID
//	mchId：商户ID
//	ApiKey：API秘钥值
//	返回参数 sign：通过Appid、MchId、ApiKey和BodyMap中的参数计算出的Sign值
func GetSanBoxParamSign(appId, mchId, apiKey string, bm gopay.BodyMap) (sign string, err error) {
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	bm.Set("sign_type", SignType_MD5)
	bm.Set("total_fee", 101)
	var (
		sandBoxApiKey string
		hashMd5       hash.Hash
	)
	if sandBoxApiKey, err = getSanBoxKey(mchId, util.GetRandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	hashMd5 = md5.New()
	hashMd5.Write([]byte(bm.EncodeWeChatSignParams(sandBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(hashMd5.Sum(nil)))
	return
}

// ParseNotifyToBodyMap 解析微信支付异步通知的结果到BodyMap（推荐）
//	req：*http.Request
//	返回参数bm：Notify请求的参数
//	返回参数err：错误信息
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := ioutil.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll：%w", err)
	}
	bm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &bm); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return
}

// Deprecated
// 推荐使用 ParseNotifyToBodyMap
func ParseNotify(req *http.Request) (notifyReq *NotifyRequest, err error) {
	notifyReq = new(NotifyRequest)
	err = xml.NewDecoder(req.Body).Decode(notifyReq)
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("xml.NewDecoder.Decode：%w", err)
	}
	return
}

// ParseRefundNotify 解析微信退款异步通知的参数
//	req：*http.Request
//	返回参数notifyReq：Notify请求的参数
//	返回参数err：错误信息
func ParseRefundNotify(req *http.Request) (notifyReq *RefundNotifyRequest, err error) {
	notifyReq = new(RefundNotifyRequest)
	err = xml.NewDecoder(req.Body).Decode(notifyReq)
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("xml.NewDecoder.Decode：%w", err)
	}
	return
}

// DecryptRefundNotifyReqInfo 解密微信退款异步通知的加密数据
//	reqInfo：gopay.ParseRefundNotify() 方法获取的加密数据 req_info
//	apiKey：API秘钥值
//	返回参数refundNotify：RefundNotify请求的加密数据
//	返回参数err：错误信息
//	文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
func DecryptRefundNotifyReqInfo(reqInfo, apiKey string) (refundNotify *RefundNotify, err error) {
	if reqInfo == util.NULL || apiKey == util.NULL {
		return nil, errors.New("reqInfo or apiKey is null")
	}
	var (
		encryptionB, bs []byte
		block           cipher.Block
		blockSize       int
	)
	if encryptionB, err = base64.StdEncoding.DecodeString(reqInfo); err != nil {
		return nil, err
	}
	h := md5.New()
	h.Write([]byte(apiKey))
	key := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	if len(encryptionB)%aes.BlockSize != 0 {
		return nil, errors.New("encryptedData is error")
	}
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return nil, err
	}
	blockSize = block.BlockSize()

	err = func(dst, src []byte) error {
		if len(src)%blockSize != 0 {
			return errors.New("crypto/cipher: input not full blocks")
		}
		if len(dst) < len(src) {
			return errors.New("crypto/cipher: output smaller than input")
		}
		for len(src) > 0 {
			block.Decrypt(dst, src[:blockSize])
			src = src[blockSize:]
			dst = dst[blockSize:]
		}
		return nil
	}(encryptionB, encryptionB)
	if err != nil {
		return nil, err
	}

	bs = xaes.PKCS7UnPadding(encryptionB)
	refundNotify = new(RefundNotify)
	if err = xml.Unmarshal(bs, refundNotify); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return
}

// VerifySign 微信同步返回参数验签或异步通知参数验签
//	ApiKey：API秘钥值
//	signType：签名类型（调用API方法时填写的类型）
//	bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq，推荐通 BodyMap 验签
//	返回参数ok：是否验签通过
//	返回参数err：其他错误信息，不要根据 error 是否为空来判断验签正确与否，需再单独判断返回的 ok
func VerifySign(apiKey, signType string, bean interface{}) (ok bool, err error) {
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	kind := reflect.ValueOf(bean).Kind()
	if kind == reflect.Map {
		bm := bean.(gopay.BodyMap)
		bodySign := bm.GetString("sign")
		bm.Remove("sign")
		return getReleaseSign(apiKey, signType, bm) == bodySign, nil
	}

	bs, err := json.Marshal(bean)
	if err != nil {
		return false, fmt.Errorf("json.Marshal(%s)：%w", string(bs), err)
	}
	bm := make(gopay.BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return false, fmt.Errorf("json.Marshal(%s)：%w", string(bs), err)
	}
	bodySign := bm.GetString("sign")
	bm.Remove("sign")
	return getReleaseSign(apiKey, signType, bm) == bodySign, nil
}

type NotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// ToXmlString 返回数据给微信
func (w *NotifyResponse) ToXmlString() (xmlStr string) {
	var buffer strings.Builder
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(w.ReturnCode)
	buffer.WriteString("]]></return_code>")
	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(w.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}

// GetMiniPaySign JSAPI支付，统一下单获取支付参数后，再次计算出小程序用的paySign
//	appId：APPID
//	nonceStr：随即字符串
//	packages：统一下单成功后拼接得到的值
//	signType：签名类型
//	timeStamp：时间
//	ApiKey：API秘钥值
//	微信小程序支付API：https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html
//	微信小程序支付PaySign计算文档：https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=3
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

// GetH5PaySign 微信内H5支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
//	appId：APPID
//	nonceStr：随即字符串
//	packages：统一下单成功后拼接得到的值
//	signType：签名类型
//	timeStamp：时间
//	ApiKey：API秘钥值
//	微信内H5支付官方文档：https://pay.weixin.qq.com/wiki/doc/api/wxpay/ch/pay/OfficialPayMent/chapter5_5.shtml
func GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
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
//	appId：APPID
//	partnerid：partnerid
//	nonceStr：随即字符串
//	prepayId：统一下单成功后得到的值
//	signType：此处签名方式，务必与统一下单时用的签名方式一致
//	timeStamp：时间
//	ApiKey：API秘钥值
//	APP支付官方文档：https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12&index=2
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

// DecryptOpenDataToBodyMap 解密开放数据到 BodyMap
//	encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//	iv：加密算法的初始向量，小程序获取到
//	sessionKey：会话密钥，通过  gopay.Code2Session() 方法获取到
//	文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func DecryptOpenDataToBodyMap(encryptedData, iv, sessionKey string) (bm gopay.BodyMap, err error) {
	if encryptedData == util.NULL || iv == util.NULL || sessionKey == util.NULL {
		return nil, errors.New("input params can not null")
	}
	var (
		cipherText, aesKey, ivKey, plainText []byte
		block                                cipher.Block
		blockMode                            cipher.BlockMode
	)
	cipherText, _ = base64.StdEncoding.DecodeString(encryptedData)
	aesKey, _ = base64.StdEncoding.DecodeString(sessionKey)
	ivKey, _ = base64.StdEncoding.DecodeString(iv)
	if len(cipherText)%len(aesKey) != 0 {
		return nil, errors.New("encryptedData is error")
	}
	if block, err = aes.NewCipher(aesKey); err != nil {
		return nil, fmt.Errorf("aes.NewCipher：%w", err)
	} else {
		blockMode = cipher.NewCBCDecrypter(block, ivKey)
		plainText = make([]byte, len(cipherText))
		blockMode.CryptBlocks(plainText, cipherText)
		if len(plainText) > 0 {
			plainText = xaes.PKCS7UnPadding(plainText)
		}
		bm = make(gopay.BodyMap)
		if err = json.Unmarshal(plainText, &bm); err != nil {
			return nil, fmt.Errorf("json.Marshal(%s)：%w", string(plainText), err)
		}
		return
	}
}

// GetOpenIdByAuthCode 授权码查询openid(AccessToken:157字符)
//	appId:APPID
//	mchId:商户号
//	ApiKey:apiKey
//	authCode:用户授权码
//	nonceStr:随即字符串
//	文档：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9
func GetOpenIdByAuthCode(appId, mchId, apiKey, authCode, nonceStr string) (openIdRsp *OpenIdByAuthCodeRsp, err error) {
	var (
		url string
		bm  gopay.BodyMap
	)
	url = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	bm = make(gopay.BodyMap)
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	bm.Set("auth_code", authCode)
	bm.Set("nonce_str", nonceStr)
	bm.Set("sign", getReleaseSign(apiKey, SignType_MD5, bm))

	openIdRsp = new(OpenIdByAuthCodeRsp)
	_, errs := xhttp.NewClient().Type(xhttp.TypeXML).Post(url).SendString(GenerateXml(bm)).EndStruct(openIdRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return openIdRsp, nil
}
