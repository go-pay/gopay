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
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"reflect"
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
	fmt.Println(url)
	agent := HttpAgent()
	_, _, errs := agent.Get(url).EndStruct(userInfo)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return userInfo, nil
	}
}
