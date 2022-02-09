/*
	微信小程序
	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend
*/

package wechat

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	xaes "github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// Deprecated
// 推荐使用：github.com/go-pay/wechat-sdk
// Code2Session 获取微信小程序用户的OpenId、SessionKey、UnionId
//	appId:APPID
//	appSecret:AppSecret
//	wxCode:小程序调用wx.login 获取的code
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func Code2Session(ctx context.Context, appId, appSecret, wxCode string) (sessionRsp *Code2SessionRsp, err error) {
	sessionRsp = new(Code2SessionRsp)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + wxCode + "&grant_type=authorization_code"
	_, err = xhttp.NewClient().Get(url).EndStruct(ctx, sessionRsp)
	if err != nil {
		return nil, err
	}
	return sessionRsp, nil
}

// Deprecated
// 推荐使用：github.com/go-pay/wechat-sdk
// GetAppletAccessToken 获取微信小程序全局唯一后台接口调用凭据(AccessToken:157字符)
//	appId:APPID
//	appSecret:AppSecret
//	获取access_token文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func GetAppletAccessToken(ctx context.Context, appId, appSecret string) (accessToken *AccessToken, err error) {
	accessToken = new(AccessToken)
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret
	_, err = xhttp.NewClient().Get(url).EndStruct(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

// Deprecated
// 推荐使用：github.com/go-pay/wechat-sdk
// DecryptOpenDataToStruct 解密开放数据到结构体
//	encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//	iv：加密算法的初始向量，小程序获取到
//	sessionKey：会话密钥，通过  gopay.Code2Session() 方法获取到
//	beanPtr：需要解析到的结构体指针，操作完后，声明的结构体会被赋值
//	文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func DecryptOpenDataToStruct(encryptedData, iv, sessionKey string, beanPtr interface{}) (err error) {
	if encryptedData == util.NULL || iv == util.NULL || sessionKey == util.NULL {
		return errors.New("input params can not null")
	}
	var (
		cipherText, aesKey, ivKey, plainText []byte
		block                                cipher.Block
		blockMode                            cipher.BlockMode
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
	if block, err = aes.NewCipher(aesKey); err != nil {
		return fmt.Errorf("aes.NewCipher：%w", err)
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	plainText = make([]byte, 0, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	if len(plainText) > 0 {
		plainText = xaes.PKCS7UnPadding(plainText)
	}
	if err = json.Unmarshal(plainText, beanPtr); err != nil {
		return fmt.Errorf("json.Marshal(%s)：%w", string(plainText), err)
	}
	return
}

// Deprecated
// 推荐使用：github.com/go-pay/wechat-sdk
// GetAppletPaidUnionId 微信小程序用户支付完成后，获取该用户的 UnionId，无需用户授权。
//	accessToken：接口调用凭据
//	openId：用户的OpenID
//	transactionId：微信支付订单号
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func GetAppletPaidUnionId(ctx context.Context, accessToken, openId, transactionId string) (unionId *PaidUnionId, err error) {
	unionId = new(PaidUnionId)
	url := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=" + accessToken + "&openid=" + openId + "&transaction_id=" + transactionId
	_, err = xhttp.NewClient().Get(url).EndStruct(ctx, unionId)
	if err != nil {
		return nil, err
	}
	return unionId, nil
}
