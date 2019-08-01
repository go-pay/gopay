//==================================
//  * Name：Jerry
//  * DateTime：2019/5/6 13:16
//  * Desc：
//==================================
package gopay

import (
	"bytes"
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
	"github.com/parnurzeal/gorequest"
	"net/http"
	"reflect"
	"strings"
)

//解析支付完成后的Notify信息
func ParseNotifyResult(req *http.Request) (notifyRsp *WeChatNotifyRequest, err error) {
	notifyRsp = new(WeChatNotifyRequest)
	defer req.Body.Close()
	err = xml.NewDecoder(req.Body).Decode(notifyRsp)
	if err != nil {
		return nil, err
	}
	return
}

type WeChatNotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

//返回数据给微信
func (this *WeChatNotifyResponse) ToXmlString() (xmlStr string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(this.ReturnCode)
	buffer.WriteString("]]></return_code>")

	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(this.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}

//支付通知的签名验证和参数签名后的Sign
//    apiKey：API秘钥值
//    signType：签名类型 MD5 或 HMAC-SHA256（默认请填写 MD5）
//    notifyRsp：利用 gopay.ParseNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数sign：根据参数计算的sign值，非微信返回参数中的Sign
func VerifyPayResultSign(apiKey string, signType string, notifyRsp *WeChatNotifyRequest) (ok bool, sign string) {
	body := make(BodyMap)
	body.Set("return_code", notifyRsp.ReturnCode)
	body.Set("return_msg", notifyRsp.ReturnMsg)
	body.Set("appid", notifyRsp.Appid)
	body.Set("mch_id", notifyRsp.MchId)
	body.Set("device_info", notifyRsp.DeviceInfo)
	body.Set("nonce_str", notifyRsp.NonceStr)
	body.Set("sign_type", notifyRsp.SignType)
	body.Set("result_code", notifyRsp.ResultCode)
	body.Set("err_code", notifyRsp.ErrCode)
	body.Set("err_code_des", notifyRsp.ErrCodeDes)
	body.Set("openid", notifyRsp.Openid)
	body.Set("is_subscribe", notifyRsp.IsSubscribe)
	body.Set("trade_type", notifyRsp.TradeType)
	body.Set("bank_type", notifyRsp.BankType)
	body.Set("total_fee", notifyRsp.TotalFee)
	body.Set("settlement_total_fee", notifyRsp.SettlementTotalFee)
	body.Set("fee_type", notifyRsp.FeeType)
	body.Set("cash_fee", notifyRsp.CashFee)
	body.Set("cash_fee_type", notifyRsp.CashFeeType)
	body.Set("coupon_fee", notifyRsp.CouponFee)
	body.Set("coupon_count", notifyRsp.CouponCount)
	body.Set("coupon_type_0", notifyRsp.CouponType0)
	body.Set("coupon_id_0", notifyRsp.CouponId0)
	body.Set("coupon_fee_0", notifyRsp.CouponFee0)
	body.Set("transaction_id", notifyRsp.TransactionId)
	body.Set("out_trade_no", notifyRsp.OutTradeNo)
	body.Set("attach", notifyRsp.Attach)
	body.Set("time_end", notifyRsp.TimeEnd)

	newBody := make(BodyMap)
	for key := range body {
		vStr := body.Get(key)
		if vStr != null && vStr != "0" {
			newBody.Set(key, vStr)
		}
	}

	sign = getLocalSign(apiKey, signType, newBody)

	ok = sign == notifyRsp.Sign
	return
}

//JSAPI支付，统一下单获取支付参数后，再次计算出小程序用的paySign
//    appId：APPID
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名类型
//    timeStamp：时间
//    apiKey：API秘钥值
func GetMiniPaySign(appId, nonceStr, prepayId, signType, timeStamp, apiKey string) (paySign string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("appId=")
	buffer.WriteString(appId)

	buffer.WriteString("&nonceStr=")
	buffer.WriteString(nonceStr)

	buffer.WriteString("&package=")
	buffer.WriteString(prepayId)

	buffer.WriteString("&signType=")
	buffer.WriteString(signType)

	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timeStamp)

	buffer.WriteString("&key=")
	buffer.WriteString(apiKey)

	signStr := buffer.String()

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
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//JSAPI支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
//    appId：APPID
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名类型
//    timeStamp：时间
//    apiKey：API秘钥值
func GetH5PaySign(appId, nonceStr, prepayId, signType, timeStamp, apiKey string) (paySign string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("appId=")
	buffer.WriteString(appId)

	buffer.WriteString("&nonceStr=")
	buffer.WriteString(nonceStr)

	buffer.WriteString("&package=")
	buffer.WriteString(prepayId)

	buffer.WriteString("&signType=")
	buffer.WriteString(signType)

	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timeStamp)

	buffer.WriteString("&key=")
	buffer.WriteString(apiKey)

	signStr := buffer.String()

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
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//APP支付，统一下单获取支付参数后，再次计算APP支付所需要的的sign
//    appId：APPID
//    partnerid：partnerid
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：此处签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
func GetAppPaySign(appid, partnerid, noncestr, prepayid, signType, timestamp, apiKey string) (paySign string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("appid=")
	buffer.WriteString(appid)

	buffer.WriteString("&nonceStr=")
	buffer.WriteString(noncestr)

	buffer.WriteString("&package=Sign=WXPay")

	buffer.WriteString("&partnerid=")
	buffer.WriteString(partnerid)

	buffer.WriteString("&prepayid=")
	buffer.WriteString(prepayid)

	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timestamp)

	buffer.WriteString("&key=")
	buffer.WriteString(apiKey)

	signStr := buffer.String()

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
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//解密开放数据
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    iv:加密算法的初始向量
//    sessionKey:会话密钥
//    beanPtr:需要解析到的结构体指针
func DecryptOpenDataToStruct(encryptedData, iv, sessionKey string, beanPtr interface{}) (err error) {
	//验证参数类型
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入beanPtr类型必须是以指针形式")
	}
	//验证interface{}类型
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	aesKey, _ := base64.StdEncoding.DecodeString(sessionKey)
	ivKey, _ := base64.StdEncoding.DecodeString(iv)
	cipherText, _ := base64.StdEncoding.DecodeString(encryptedData)

	if len(cipherText)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	//fmt.Println("cipherText:", cipherText)
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}
	//解密
	blockMode := cipher.NewCBCDecrypter(block, ivKey)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	//fmt.Println("plainText1:", plainText)
	plainText = PKCS7UnPadding(plainText)
	//fmt.Println("plainText:", plainText)
	//解析
	err = json.Unmarshal(plainText, beanPtr)
	if err != nil {
		return err
	}
	return nil
}

//获取微信用户的OpenId、SessionKey、UnionId
//    appId:APPID
//    appSecret:AppSecret
//    wxCode:小程序调用wx.login 获取的code
func Code2Session(appId, appSecret, wxCode string) (sessionRsp *Code2SessionRsp, err error) {
	sessionRsp = new(Code2SessionRsp)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + wxCode + "&grant_type=authorization_code"
	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(sessionRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return sessionRsp, nil
	}
}

//获取小程序全局唯一后台接口调用凭据(AccessToken:157字符)
//    appId:APPID
//    appSecret:AppSecret
func GetAccessToken(appId, appSecret string) (accessToken *AccessToken, err error) {
	accessToken = new(AccessToken)
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret

	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(accessToken)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return accessToken, nil
	}
}

//授权码查询openid(AccessToken:157字符)
//    appId:APPID
//    mchId:商户号
//    apiKey:ApiKey
//    authCode:用户授权码
//    nonceStr:随即字符串
func GetOpenIdByAuthCode(appId, mchId, apiKey, authCode, nonceStr string) (openIdRsp *OpenIdByAuthCodeRsp, err error) {

	url := "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	body := make(BodyMap)
	body.Set("appid", appId)
	body.Set("mch_id", mchId)
	body.Set("auth_code", authCode)
	body.Set("nonce_str", nonceStr)
	sign := getLocalSign(apiKey, SignType_MD5, body)

	body.Set("sign", sign)
	reqXML := generateXml(body)
	//===============发起请求===================
	agent := gorequest.New()
	agent.Post(url)
	agent.Type("xml")
	agent.SendString(reqXML)
	_, bs, errs := agent.EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	openIdRsp = new(OpenIdByAuthCodeRsp)
	err = xml.Unmarshal(bs, openIdRsp)
	if err != nil {
		return nil, err
	}
	return openIdRsp, nil
}

//用户支付完成后，获取该用户的 UnionId，无需用户授权。
//    accessToken：接口调用凭据
//    openId：用户的OpenID
//    transactionId：微信支付订单号
func GetPaidUnionId(accessToken, openId, transactionId string) (unionId *PaidUnionId, err error) {
	unionId = new(PaidUnionId)
	url := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=" + accessToken + "&openid=" + openId + "&transaction_id=" + transactionId

	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(unionId)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return unionId, nil
	}
}

//获取用户基本信息(UnionID机制)
//    accessToken：接口调用凭据
//    openId：用户的OpenID
//    lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
func GetWeChatUserInfo(accessToken, openId string, lang ...string) (userInfo *WeChatUserInfo, err error) {
	userInfo = new(WeChatUserInfo)
	var url string
	if len(lang) > 0 {
		url = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=" + lang[0]
	} else {
		url = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=zh_CN"
	}
	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(userInfo)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return userInfo, nil
	}
}
