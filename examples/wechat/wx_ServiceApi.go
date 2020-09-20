package wechat

import (
	"fmt"
	"net/http"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
)

func Code2Session() {
	// 获取微信用户的OpenId、SessionKey、UnionId
	//    appId:APPID
	//    appSecret:AppSecret
	//    wxCode:小程序调用wx.login 获取的code
	userIdRsp, err := wechat.Code2Session("AppID", "APPSecret", "011EZg6p0VO47n1p2W4p0mle6p0EZg6u")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("SessionKey:", userIdRsp.SessionKey)
	fmt.Println("OpenID:", userIdRsp.Openid)
	fmt.Println("UnionID:", userIdRsp.Unionid)
	fmt.Println("Errcode:", userIdRsp.Errcode)
	fmt.Println("Errmsg:", userIdRsp.Errmsg)
}

func GetAppWeChatLoginAccessToken() {
	accessToken, err := wechat.GetOauth2AccessToken("AppID", "AppSecret", "code")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("AccessToken:", accessToken.AccessToken)
	fmt.Println("Openid:", accessToken.Openid)
	fmt.Println("Unionid:", accessToken.Unionid)
	fmt.Println("Scope:", accessToken.Scope)
	fmt.Println("ExpiresIn:", accessToken.ExpiresIn)
	fmt.Println("Errcode:", accessToken.Errcode)
	fmt.Println("Errmsg:", accessToken.Errmsg)
}

func RefreshAppWeChatLoginAccessToken() {
	accessToken, err := wechat.RefreshOauth2AccessToken("AppID", "refreshToken")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("AccessToken:", accessToken.AccessToken)
	fmt.Println("Openid:", accessToken.Openid)
	fmt.Println("Scope:", accessToken.Scope)
	fmt.Println("ExpiresIn:", accessToken.ExpiresIn)
	fmt.Println("Errcode:", accessToken.Errcode)
	fmt.Println("Errmsg:", accessToken.Errmsg)
}

func GetWeChatAppletAccessToken() {
	// 获取小程序全局唯一后台接口调用凭据(AccessToken:157字符)
	//    appId:APPID
	//    appSecret:AppSecret
	accessToken, err := wechat.GetAppletAccessToken("AppID", "AppSecret")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("AccessToken:", accessToken.AccessToken)
	fmt.Println("ExpiresIn:", accessToken.ExpiresIn)
	fmt.Println("Errcode:", accessToken.Errcode)
	fmt.Println("Errmsg:", accessToken.Errmsg)
}

func GetWeChatAppletPaidUnionId() {
	accessToken := "21_3puo3mxoK6Ry2bR7Dh-qdn41wUP1wClwke8Zhf9a_i39jfWRq9rhNJZZZHaOt_Yad-Gp6u9_46dGW0RbIMz3nANInRI3m-1glvCnGW47v63sjYWV1iyTKOHGwDVxEv78kY-0OfkmkiIveVqAZCZaAAAQTQ"
	// 用户支付完成后，获取该用户的 UnionId，无需用户授权。
	//    accessToken：接口调用凭据
	//    openId：用户的OpenID
	//    transactionId：微信支付订单号
	rsp, err := wechat.GetAppletPaidUnionId(accessToken, "o0Df70MSI4Ygv2KQ2cLnoMN5CXI8", "4200000326201905256499385970")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("Unionid:", rsp.Unionid)
	fmt.Println("Errcode:", rsp.Errcode)
	fmt.Println("Errmsg:", rsp.Errmsg)
}

func GetWeChatUserInfo() {
	accessToken := "21_3puo3mxoK6Ry2bR7Dh-qdn41wUP1wClwke8Zhf9a_i39jfWRq9rhNJZZZHaOt_Yad-Gp6u9_46dGW0RbIMz3nANInRI3m-1glvCnGW47v63sjYWV1iyTKOHGwDVxEv78kY-0OfkmkiIveVqAZCZaAAAQTQ"
	// 获取用户基本信息(UnionID机制)(微信公众号)
	//    accessToken：接口调用凭据
	//    openId：用户的OpenID
	//    lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
	userInfo, err := wechat.GetUserInfo(accessToken, "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("userInfo:", *userInfo)
}

func GetWeChatUserInfoOpen() {
	accessToken := "21_3puo3mxoK6Ry2bR7Dh-qdn41wUP1wClwke8Zhf9a_i39jfWRq9rhNJZZZHaOt_Yad-Gp6u9_46dGW0RbIMz3nANInRI3m-1glvCnGW47v63sjYWV1iyTKOHGwDVxEv78kY-0OfkmkiIveVqAZCZaAAAQTQ"

	userInfo, err := wechat.GetUserInfoOpen(accessToken, "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("userInfo:", *userInfo)
}

func DecryptWeChatOpenDataToStruct() {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="

	// 微信小程序，手机号
	phone := new(wechat.UserPhone)
	// 解密开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    iv:加密算法的初始向量
	//    sessionKey:会话密钥
	//    beanPtr:需要解析到的结构体指针
	err := wechat.DecryptOpenDataToStruct(data, iv, session, phone)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("PhoneNumber:", phone.PhoneNumber)
	fmt.Println("PurePhoneNumber:", phone.PurePhoneNumber)
	fmt.Println("CountryCode:", phone.CountryCode)
	fmt.Println("Watermark:", phone.Watermark)

	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

	// 微信小程序 用户信息
	userInfo := new(wechat.AppletUserInfo)

	err = wechat.DecryptOpenDataToStruct(encryptedData, iv2, sessionKey, userInfo)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("NickName:", userInfo.NickName)
	fmt.Println("AvatarUrl:", userInfo.AvatarUrl)
	fmt.Println("Country:", userInfo.Country)
	fmt.Println("Province:", userInfo.Province)
	fmt.Println("City:", userInfo.City)
	fmt.Println("Gender:", userInfo.Gender)
	fmt.Println("OpenId:", userInfo.OpenId)
	fmt.Println("UnionId:", userInfo.UnionId)
	fmt.Println("Watermark:", userInfo.Watermark)
}

func DecryptWeChatOpenDataToBodyMap() {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="

	// 解密开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    iv:加密算法的初始向量
	//    sessionKey:会话密钥
	bm, err := wechat.DecryptOpenDataToBodyMap(data, iv, session)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("WeChatUserPhone:", bm)
}

func GetOpenIdByAuthCode() {
	// 授权码查询openid
	//    appId:APPID
	//    mchId:商户号
	//    apiKey:ApiKey
	//    authCode:用户授权码
	//    nonceStr:随即字符串
	openIdRsp, err := wechat.GetOpenIdByAuthCode("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", "135127679952609396", gotil.GetRandomString(32))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("ReturnCode:", openIdRsp.ReturnCode)
	fmt.Println("ReturnMsg:", openIdRsp.ReturnMsg)
	fmt.Println("ResultCode:", openIdRsp.ResultCode)
	fmt.Println("Appid:", openIdRsp.Appid)
	fmt.Println("MchId:", openIdRsp.MchId)
	fmt.Println("NonceStr:", openIdRsp.NonceStr)
	fmt.Println("err_code:", openIdRsp.ErrCode)
	fmt.Println("Openid:", openIdRsp.Openid)
}

// 解析notify参数、验签、返回数据到微信
func ParseWeChatNotifyAndVerifyWeChatSign(req *http.Request) string {
	rsp := new(wechat.NotifyResponse)

	// 解析参数
	bodyMap, err := wechat.ParseNotifyToBodyMap(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("bodyMap:", bodyMap)

	ok, err := wechat.VerifySign("GFDS8j98rewnmgl45wHTt980jg543abc", wechat.SignType_MD5, bodyMap)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("微信验签是否通过:", ok)

	rsp.ReturnCode = gopay.SUCCESS
	rsp.ReturnMsg = "OK"
	return rsp.ToXmlString()
}

// 解析微信退款异步通知的参数，解析出来的 req_info 是加密数据，需解密
func ParseWeChatRefundNotify(req *http.Request) string {
	rsp := new(wechat.NotifyResponse)
	// 解析参数
	notifyReq, err := wechat.ParseRefundNotify(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("notifyReq:", *notifyReq)
	// 退款通知无sign，不用验签

	// 解密退款异步通知的加密数据
	refundNotify, err := wechat.DecryptRefundNotifyReqInfo(notifyReq.ReqInfo, "GFDS8j98rewnmgl45wHTt980jg543abc")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("refundNotify:", *refundNotify)

	// 或者

	bodyMap, err := wechat.ParseNotifyToBodyMap(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("bodyMap:", bodyMap)

	// 解密退款异步通知的加密数据
	refundNotify2, err := wechat.DecryptRefundNotifyReqInfo(bodyMap.Get("req_info"), "GFDS8j98rewnmgl45wHTt980jg543abc")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("refundNotify:", *refundNotify2)

	// 返回微信
	rsp.ReturnCode = gopay.SUCCESS
	rsp.ReturnMsg = "OK"
	return rsp.ToXmlString()
}
