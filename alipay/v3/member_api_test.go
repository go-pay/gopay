package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

func TestClient_SystemOauthToken(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code").
		Set("code", "309cd66a1c024e4b9d89f7241db4UA18").
		Set(HeaderAppAuthToken, "202504BBcb0db76aff98487fb3fd2b733a474X87")

	// 发起请求
	aliRsp, err := client.SystemOauthToken(ctx, bm)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
	xlog.Debugf("aliRsp: %s", js.MarshalString(aliRsp))
}

func TestClient_UserInfoShare(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("auth_token", "authbseB27c501c0804548b09e1f121ae0835X18").
		Set(HeaderAppAuthToken, "202504BB06476f1dc64f465997935ecf5072eX88")

	// 发起请求
	aliRsp, err := client.UserInfoShare(ctx, bm)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
	xlog.Debug("aliRsp:", js.MarshalString(aliRsp))
}

func TestClient_UserAuthRelationshipQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("scopes", "auth_user").
		Set("open_id", "018OacbttSLyJtNfdPbDOcaGoo-ncctDVT45IdYxaUsmIY8")
	// 发起请求
	aliRsp, err := client.UserAuthRelationshipQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
	xlog.Debug("aliRsp:", js.MarshalString(aliRsp))
}
