package alipay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cedarwu/gopay"
	"github.com/cedarwu/gopay/pkg/util"
)

// alipay.user.info.auth(用户登陆授权)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.user.info.auth
func (a *Client) UserInfoAuth(bm gopay.BodyMap) (aliRsp *UserInfoAuthResponse, err error) {
	err = bm.CheckEmptyError("scopes", "state")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.info.auth"); err != nil {
		return nil, err
	}
	if strings.Contains(string(bs), "<head>") {
		return nil, errors.New(string(bs))
	}
	aliRsp = new(UserInfoAuthResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.system.oauth.token(换取授权访问令牌)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token
func (a *Client) SystemOauthToken(bm gopay.BodyMap) (aliRsp *SystemOauthTokenResponse, err error) {
	if bm.GetString("code") == util.NULL && bm.GetString("refresh_token") == util.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("grant_type")
	if err != nil {
		return nil, err
	}

	if a.AppCertSN != util.NULL {
		bm.Set("app_cert_sn", a.AppCertSN)
	}
	if a.AliPayRootCertSN != util.NULL {
		bm.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}

	var bs []byte
	if bs, err = systemOauthToken(a.AppId, a.privateKey, bm, "alipay.system.oauth.token", a.IsProd, a.SignType); err != nil {
		return nil, err
	}
	aliRsp = new(SystemOauthTokenResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.ErrorResponse != nil {
		info := aliRsp.ErrorResponse
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.open.auth.token.app(换取应用授权令牌)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.auth.token.app
func (a *Client) OpenAuthTokenApp(bm gopay.BodyMap) (aliRsp *OpenAuthTokenAppResponse, err error) {
	if bm.GetString("code") == util.NULL && bm.GetString("refresh_token") == util.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("grant_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.open.auth.token.app"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAuthTokenAppResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.open.app.alipaycert.download(应用支付宝公钥证书下载)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.app.alipaycert.download
func (a *Client) PublicCertDownload(bm gopay.BodyMap) (aliRsp *PublicCertDownloadRsp, err error) {
	err = bm.CheckEmptyError("alipay_cert_sn")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.open.app.alipaycert.download"); err != nil {
		return nil, err
	}
	aliRsp = new(PublicCertDownloadRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	certBs, err := base64.StdEncoding.DecodeString(aliRsp.Response.AlipayCertContent)
	if err != nil {
		return nil, fmt.Errorf("AlipayCertContent(%s)_DecodeErr:%+v", aliRsp.Response.AlipayCertContent, err)
	}
	aliRsp.Response.AlipayCertContent = string(certBs)
	return aliRsp, nil
}
