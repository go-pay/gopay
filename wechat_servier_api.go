//==================================
//  * Name：Jerry
//  * DateTime：2019/5/6 13:16
//  * Desc：
//==================================
package gopay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"github.com/parnurzeal/gorequest"
	"strings"
)

func HttpAgent() (agent *gorequest.SuperAgent) {
	agent = gorequest.New()
	agent.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return
}

//JSAPI支付，支付参数后，再次计算出小程序用的paySign
func GetMiniPaySign(appId, nonceStr, prepayId, signType, timeStamp, secretKey string) (paySign string) {
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
	buffer.WriteString(secretKey)

	signStr := buffer.String()

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
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//JSAPI支付，支付参数后，再次计算出微信内H5支付需要用的paySign
func GetH5PaySign(appId, nonceStr, prepayId, signType, timeStamp, secretKey string) (paySign string) {
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
	buffer.WriteString(secretKey)

	signStr := buffer.String()

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
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//获取微信用户的OpenId、SessionKey、UnionId
//    appId:APPID
//    appSecret:AppSecret
//    wxCode:小程序调用wx.login 获取的code
func Code2Session(appId, appSecret, wxCode string) (userRsp *Code2SessionRsp, err error) {
	userRsp = new(Code2SessionRsp)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + wxCode + "&grant_type=authorization_code"
	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(userRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return userRsp, nil
	}
}

//获取小程序全局唯一后台接口调用凭据(AccessToken:157字符)
//    appId:APPID
//    appSecret:AppSecret
func GetAccessToken(appId, appSecret string) (rsp *GetAccessTokenRsp, err error) {
	rsp = new(GetAccessTokenRsp)
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret

	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(rsp)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return rsp, nil
	}
}

//用户支付完成后，获取该用户的 UnionId，无需用户授权。
//    accessToken：接口调用凭据
//    openId：用户的OpenID
//    transactionId：微信支付订单号
func GetPaidUnionId(accessToken, openId, transactionId string) (rsp *GetPaidUnionIdRsp, err error) {
	rsp = new(GetPaidUnionIdRsp)
	url := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=" + accessToken + "&openid=" + openId + "&transaction_id=" + transactionId

	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(rsp)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return rsp, nil
	}
}
