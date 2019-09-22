package gopay

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
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//获取微信支付所需参数里的Sign值（通过支付参数计算Sign值）
//    注意：BodyMap中如无 sign_type 参数，默认赋值 sign_type 为 MD5
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    返回参数 sign：通过Appid、MchId、ApiKey和BodyMap中的参数计算出的Sign值
func GetWeChatParamSign(appId, mchId, apiKey string, bm BodyMap) (sign string) {
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	var (
		signType string
		h        hash.Hash
	)
	signType = bm.Get("sign_type")
	if signType == null {
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

//获取微信支付沙箱环境所需参数里的Sign值（通过支付参数计算Sign值）
//    注意：沙箱环境默认 sign_type 为 MD5
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    返回参数 sign：通过Appid、MchId、ApiKey和BodyMap中的参数计算出的Sign值
func GetWeChatSanBoxParamSign(appId, mchId, apiKey string, bm BodyMap) (sign string, err error) {
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	bm.Set("sign_type", SignType_MD5)
	var (
		sanBoxApiKey string
		hashMd5      hash.Hash
	)
	if sanBoxApiKey, err = getSanBoxKey(mchId, GetRandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	hashMd5 = md5.New()
	hashMd5.Write([]byte(bm.EncodeWeChatSignParams(sanBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(hashMd5.Sum(nil)))
	return
}

//解析微信支付异步通知的结果到BodyMap
//    req：*http.Request
//    返回参数bm：Notify请求的参数
//    返回参数err：错误信息
func ParseWeChatNotifyResultToBodyMap(req *http.Request) (bm BodyMap, err error) {
	var bs []byte
	if bs, err = ioutil.ReadAll(req.Body); err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll：%v", err.Error())
	}
	bm = make(BodyMap)
	if err = xml.Unmarshal(bs, &bm); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//解析微信支付异步通知的参数
//    req：*http.Request
//    返回参数notifyReq：Notify请求的参数
//    返回参数err：错误信息
func ParseWeChatNotifyResult(req *http.Request) (notifyReq *WeChatNotifyRequest, err error) {
	notifyReq = new(WeChatNotifyRequest)
	if err = xml.NewDecoder(req.Body).Decode(notifyReq); err != nil {
		return nil, fmt.Errorf("xml.NewDecoder：%v", err.Error())
	}
	return
}

//验证微信支付异步通知的Sign值（Deprecated）
//    apiKey：API秘钥值
//    signType：签名类型 MD5 或 HMAC-SHA256（默认请填写 MD5）
//    notifyReq：利用 gopay.ParseWeChatNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数sign：根据参数计算的sign值，非微信返回参数中的Sign
func VerifyWeChatResultSign(apiKey, signType string, notifyReq *WeChatNotifyRequest) (ok bool, sign string) {
	body := make(BodyMap)
	body.Set("return_code", notifyReq.ReturnCode)
	body.Set("return_msg", notifyReq.ReturnMsg)
	body.Set("appid", notifyReq.Appid)
	body.Set("mch_id", notifyReq.MchId)
	body.Set("device_info", notifyReq.DeviceInfo)
	body.Set("nonce_str", notifyReq.NonceStr)
	body.Set("sign_type", notifyReq.SignType)
	body.Set("result_code", notifyReq.ResultCode)
	body.Set("err_code", notifyReq.ErrCode)
	body.Set("err_code_des", notifyReq.ErrCodeDes)
	body.Set("openid", notifyReq.Openid)
	body.Set("is_subscribe", notifyReq.IsSubscribe)
	body.Set("trade_type", notifyReq.TradeType)
	body.Set("bank_type", notifyReq.BankType)
	body.Set("total_fee", notifyReq.TotalFee)
	body.Set("settlement_total_fee", notifyReq.SettlementTotalFee)
	body.Set("fee_type", notifyReq.FeeType)
	body.Set("cash_fee", notifyReq.CashFee)
	body.Set("cash_fee_type", notifyReq.CashFeeType)
	body.Set("coupon_fee", notifyReq.CouponFee)
	body.Set("coupon_count", notifyReq.CouponCount)
	body.Set("coupon_type_0", notifyReq.CouponType0)
	body.Set("coupon_type_1", notifyReq.CouponType1)
	body.Set("coupon_id_0", notifyReq.CouponId0)
	body.Set("coupon_id_1", notifyReq.CouponId1)
	body.Set("coupon_fee_0", notifyReq.CouponFee0)
	body.Set("coupon_fee_1", notifyReq.CouponFee1)
	body.Set("transaction_id", notifyReq.TransactionId)
	body.Set("out_trade_no", notifyReq.OutTradeNo)
	body.Set("attach", notifyReq.Attach)
	body.Set("time_end", notifyReq.TimeEnd)
	newBody := make(BodyMap)
	for key := range body {
		vStr := body.Get(key)
		if vStr != null && vStr != "0" {
			newBody.Set(key, vStr)
		}
	}
	sign = getWeChatReleaseSign(apiKey, signType, newBody)
	ok = sign == notifyReq.Sign
	return
}

//微信同步返回参数验签或异步通知参数验签
//    apiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
func VerifyWeChatSign(apiKey, signType string, bean interface{}) (ok bool, err error) {
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	var (
		bm       BodyMap
		bs       []byte
		kind     reflect.Kind
		bodySign string
	)
	kind = reflect.ValueOf(bean).Kind()
	if kind == reflect.Map {
		bm = bean.(BodyMap)
		goto Verify
	}
	if bs, err = json.Marshal(bean); err != nil {
		return false, fmt.Errorf("json.Marshal：%v", err.Error())
	}
	bm = make(BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return false, fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
Verify:
	bodySign = bm.Get("sign")
	bm.Remove("sign")
	return getWeChatReleaseSign(apiKey, signType, bm) == bodySign, nil
}

type WeChatNotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

//返回数据给微信
func (w *WeChatNotifyResponse) ToXmlString() (xmlStr string) {
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

//JSAPI支付，统一下单获取支付参数后，再次计算出小程序用的paySign
//    appId：APPID
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：签名类型
//    timeStamp：时间
//    apiKey：API秘钥值
//    微信小程序支付API：https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html
func GetMiniPaySign(appId, nonceStr, prepayId, signType, timeStamp, apiKey string) (paySign string) {
	var (
		buffer strings.Builder
		h      hash.Hash
	)
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
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(buffer.String()))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

//微信内H5支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
//    appId：APPID
//    nonceStr：随即字符串
//    packages：统一下单成功后拼接得到的值
//    signType：签名类型
//    timeStamp：时间
//    apiKey：API秘钥值
//    微信内H5支付官方文档：https://pay.weixin.qq.com/wiki/doc/api/external/jsapi.php?chapter=7_7&index=6
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

//APP支付，统一下单获取支付参数后，再次计算APP支付所需要的的sign
//    appId：APPID
//    partnerid：partnerid
//    nonceStr：随即字符串
//    prepayId：统一下单成功后得到的值
//    signType：此处签名方式，务必与统一下单时用的签名方式一致
//    timeStamp：时间
//    apiKey：API秘钥值
//    APP支付官方文档：https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12
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

//解密开放数据到结构体
//    encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//    iv：加密算法的初始向量，小程序获取到
//    sessionKey：会话密钥，通过  gopay.Code2Session() 方法获取到
//    beanPtr：需要解析到的结构体指针，操作完后，声明的结构体会被赋值
//    文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func DecryptWeChatOpenDataToStruct(encryptedData, iv, sessionKey string, beanPtr interface{}) (err error) {
	var (
		cipherText []byte
		aesKey     []byte
		ivKey      []byte
		block      cipher.Block
		blockMode  cipher.BlockMode
		plainText  []byte
	)
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入beanPtr类型必须是以指针形式")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	cipherText, _ = base64.StdEncoding.DecodeString(encryptedData)
	aesKey, _ = base64.StdEncoding.DecodeString(sessionKey)
	ivKey, _ = base64.StdEncoding.DecodeString(iv)
	if len(cipherText)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	block, err = aes.NewCipher(aesKey)
	if err != nil {
		return fmt.Errorf("aes.NewCipher：%v", err.Error())
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	if len(plainText) > 0 {
		plainText = PKCS7UnPadding(plainText)
	}
	if err = json.Unmarshal(plainText, beanPtr); err != nil {
		return fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
	return
}

//解密开放数据到 BodyMap
//    encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//    iv：加密算法的初始向量，小程序获取到
//    sessionKey：会话密钥，通过  gopay.Code2Session() 方法获取到
//    文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func DecryptWeChatOpenDataToBodyMap(encryptedData, iv, sessionKey string) (bm BodyMap, err error) {
	var (
		cipherText []byte
		aesKey     []byte
		ivKey      []byte
		block      cipher.Block
		blockMode  cipher.BlockMode
		plainText  []byte
	)
	cipherText, _ = base64.StdEncoding.DecodeString(encryptedData)
	aesKey, _ = base64.StdEncoding.DecodeString(sessionKey)
	ivKey, _ = base64.StdEncoding.DecodeString(iv)
	if len(cipherText)%len(aesKey) != 0 {
		return nil, errors.New("encryptedData is error")
	}
	if block, err = aes.NewCipher(aesKey); err != nil {
		return nil, fmt.Errorf("aes.NewCipher：%v", err.Error())
	} else {
		blockMode = cipher.NewCBCDecrypter(block, ivKey)
		plainText = make([]byte, len(cipherText))
		blockMode.CryptBlocks(plainText, cipherText)
		if len(plainText) > 0 {
			plainText = PKCS7UnPadding(plainText)
		}
		bm = make(BodyMap)
		if err = json.Unmarshal(plainText, &bm); err != nil {
			return nil, fmt.Errorf("json.Unmarshal：%v", err.Error())
		}
		return
	}
}

//App应用微信第三方登录，code换取access_token
//    appId：应用唯一标识，在微信开放平台提交应用审核通过后获得
//    appSecret：应用密钥AppSecret，在微信开放平台提交应用审核通过后获得
//    code：App用户换取access_token的code
func GetAppWeChatLoginAccessToken(appId, appSecret, code string) (accessToken *AppWeChatLoginAccessToken, err error) {
	accessToken = new(AppWeChatLoginAccessToken)
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appId + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"
	if _, _, errs := HttpAgent().Get(url).EndStruct(accessToken); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}

//刷新App应用微信第三方登录后，获取的 access_token
//    appId：应用唯一标识，在微信开放平台提交应用审核通过后获得
//    appSecret：应用密钥AppSecret，在微信开放平台提交应用审核通过后获得
//    code：App用户换取access_token的code
func RefreshAppWeChatLoginAccessToken(appId, refreshToken string) (accessToken *RefreshAppWeChatLoginAccessTokenRsp, err error) {
	accessToken = new(RefreshAppWeChatLoginAccessTokenRsp)
	url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + appId + "&grant_type=refresh_token&refresh_token=" + refreshToken
	if _, _, errs := HttpAgent().Get(url).EndStruct(accessToken); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}

//获取微信小程序用户的OpenId、SessionKey、UnionId
//    appId:APPID
//    appSecret:AppSecret
//    wxCode:小程序调用wx.login 获取的code
//    文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func Code2Session(appId, appSecret, wxCode string) (sessionRsp *Code2SessionRsp, err error) {
	sessionRsp = new(Code2SessionRsp)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + wxCode + "&grant_type=authorization_code"
	if _, _, errs := HttpAgent().Get(url).EndStruct(sessionRsp); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}

//获取微信小程序全局唯一后台接口调用凭据(AccessToken:157字符)
//    appId:APPID
//    appSecret:AppSecret
//    获取access_token文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func GetWeChatAppletAccessToken(appId, appSecret string) (accessToken *AccessToken, err error) {
	accessToken = new(AccessToken)
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret
	if _, _, errs := HttpAgent().Get(url).EndStruct(accessToken); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}

//授权码查询openid(AccessToken:157字符)
//    appId:APPID
//    mchId:商户号
//    apiKey:ApiKey
//    authCode:用户授权码
//    nonceStr:随即字符串
//    文档：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9
func GetOpenIdByAuthCode(appId, mchId, apiKey, authCode, nonceStr string) (openIdRsp *OpenIdByAuthCodeRsp, err error) {
	var (
		url  string
		bm   BodyMap
		bs   []byte
		errs []error
	)
	url = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	bm = make(BodyMap)
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	bm.Set("auth_code", authCode)
	bm.Set("nonce_str", nonceStr)
	bm.Set("sign", getWeChatReleaseSign(apiKey, SignType_MD5, bm))
	if _, bs, errs = HttpAgent().Post(url).Type("xml").SendString(generateXml(bm)).EndBytes(); len(errs) > 0 {
		return nil, errs[0]
	}
	openIdRsp = new(OpenIdByAuthCodeRsp)
	if err = xml.Unmarshal(bs, openIdRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//微信小程序用户支付完成后，获取该用户的 UnionId，无需用户授权。
//    accessToken：接口调用凭据
//    openId：用户的OpenID
//    transactionId：微信支付订单号
//    文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func GetWeChatAppletPaidUnionId(accessToken, openId, transactionId string) (unionId *PaidUnionId, err error) {
	unionId = new(PaidUnionId)
	url := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=" + accessToken + "&openid=" + openId + "&transaction_id=" + transactionId
	if _, _, errs := HttpAgent().Get(url).EndStruct(unionId); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}

//获取用户基本信息(UnionID机制)
//    accessToken：接口调用凭据
//    openId：用户的OpenID
//    lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
//    获取用户基本信息(UnionID机制)文档：https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839
func GetWeChatUserInfo(accessToken, openId string, lang ...string) (userInfo *WeChatUserInfo, err error) {
	userInfo = new(WeChatUserInfo)
	var url string
	if len(lang) > 0 {
		url = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=" + lang[0]
	} else {
		url = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=zh_CN"
	}
	if _, _, errs := HttpAgent().Get(url).EndStruct(userInfo); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}
