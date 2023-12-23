package wechat

import (
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func GetAppWeChatLoginAccessToken() {
	accessToken, err := wechat.GetOauth2AccessToken(ctx, "AppID", "AppSecret", "code")
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Debug("AccessToken:", accessToken.AccessToken)
	xlog.Debug("Openid:", accessToken.Openid)
	xlog.Debug("Unionid:", accessToken.Unionid)
	xlog.Debug("Scope:", accessToken.Scope)
	xlog.Debug("ExpiresIn:", accessToken.ExpiresIn)
	xlog.Debug("Errcode:", accessToken.Errcode)
	xlog.Debug("Errmsg:", accessToken.Errmsg)
}

func RefreshAppWeChatLoginAccessToken() {
	accessToken, err := wechat.RefreshOauth2AccessToken(ctx, "AppID", "refreshToken")
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Debug("AccessToken:", accessToken.AccessToken)
	xlog.Debug("Openid:", accessToken.Openid)
	xlog.Debug("Scope:", accessToken.Scope)
	xlog.Debug("ExpiresIn:", accessToken.ExpiresIn)
	xlog.Debug("Errcode:", accessToken.Errcode)
	xlog.Debug("Errmsg:", accessToken.Errmsg)
}

func GetOpenIdByAuthCode() {
	// 授权码查询openid
	//    appId:APPID
	//    mchId:商户号
	//    apiKey:ApiKey
	//    authCode:用户授权码
	//    nonceStr:随即字符串
	openIdRsp, err := wechat.GetOpenIdByAuthCode(ctx, "wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", "135127679952609396", util.RandomString(32))
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Debug("ReturnCode:", openIdRsp.ReturnCode)
	xlog.Debug("ReturnMsg:", openIdRsp.ReturnMsg)
	xlog.Debug("ResultCode:", openIdRsp.ResultCode)
	xlog.Debug("Appid:", openIdRsp.Appid)
	xlog.Debug("MchId:", openIdRsp.MchId)
	xlog.Debug("NonceStr:", openIdRsp.NonceStr)
	xlog.Debug("err_code:", openIdRsp.ErrCode)
	xlog.Debug("Openid:", openIdRsp.Openid)
}

// 解析notify参数、验签、返回数据到微信
func ParseWeChatNotifyAndVerifyWeChatSign(req *http.Request) string {
	rsp := new(wechat.NotifyResponse)

	// 解析参数
	bodyMap, err := wechat.ParseNotifyToBodyMap(req)
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("bodyMap:", bodyMap)

	ok, err := wechat.VerifySign("GFDS8j98rewnmgl45wHTt980jg543abc", wechat.SignType_MD5, bodyMap)
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("微信验签是否通过:", ok)

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
		xlog.Debug("err:", err)
	}
	xlog.Debug("notifyReq:", *notifyReq)
	// 退款通知无sign，不用验签

	// 解密退款异步通知的加密数据
	refundNotify, err := wechat.DecryptRefundNotifyReqInfo(notifyReq.ReqInfo, "GFDS8j98rewnmgl45wHTt980jg543abc")
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("refundNotify:", *refundNotify)

	// 或者

	bodyMap, err := wechat.ParseNotifyToBodyMap(req)
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("bodyMap:", bodyMap)

	// 解密退款异步通知的加密数据
	refundNotify2, err := wechat.DecryptRefundNotifyReqInfo(bodyMap.GetString("req_info"), "GFDS8j98rewnmgl45wHTt980jg543abc")
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("refundNotify:", *refundNotify2)

	// 返回微信
	rsp.ReturnCode = gopay.SUCCESS
	rsp.ReturnMsg = "OK"
	return rsp.ToXmlString()
}
