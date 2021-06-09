package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func TestClient_SystemOauthToken(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "3a06216ac8f84b8c93507bb9774bWX11")

	// 发起请求
	aliRsp, err := client.SystemOauthToken(bm)
	if err != nil {
		xlog.Errorf("client.SystemOauthToken(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp:", aliRsp.Response.AccessToken)
	xlog.Debug("aliRsp:", aliRsp.SignData)
}

func TestClient_OpenAuthTokenApp(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "866185490c4e40efa9f71efea6766X02")

	// 发起请求
	aliRsp, err := client.OpenAuthTokenApp(bm)
	if err != nil {
		xlog.Errorf("client.OpenAuthTokenApp(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserInfoAuth(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 接口权限值，目前只支持auth_user和auth_base两个值。具体说明看文档介绍
	bm.Set("scopes", []string{"auth_user"})
	bm.Set("state", "init")

	// 发起请求
	aliRsp, err := client.UserInfoAuth(bm)
	if err != nil {
		xlog.Errorf("client.UserInfoAuth(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserInfoShare(t *testing.T) {
	// 发起请求
	aliRsp, err := client.UserInfoShare()
	if err != nil {
		xlog.Errorf("client.UserInfoShare(),error:%+v", err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

	// 同步返回验签
	ok, err := VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		xlog.Errorf("client.VerifySyncSign(%s,%s,%s),error:%+v", aliPayPublicKey, aliRsp.SignData, aliRsp.Sign, err)
		return
	}
	xlog.Debug("ok:", ok)
}
