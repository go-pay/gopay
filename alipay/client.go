package alipay

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xrsa"
)

type Client struct {
	AppId              string
	LocationName       string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	SignType           string
	AppAuthToken       string
	IsProd             bool
	privateKey         *rsa.PrivateKey
	aliPayPublicKey    *rsa.PublicKey // 支付宝证书公钥内容 alipayCertPublicKey_RSA2.crt
	autoSign           bool
	DebugSwitch        gopay.DebugSwitch
	location           *time.Location
	mu                 sync.RWMutex
}

// 初始化支付宝客户端
//	注意：如果使用支付宝公钥证书验签，请设置 支付宝根证书SN（client.SetAlipayRootCertSN()）、应用公钥证书SN（client.SetAppCertSN()）
//	appId：应用ID
//	privateKey：应用私钥，支持PKCS1和PKCS8
//	isProd：是否是正式环境
func NewClient(appId, privateKey string, isProd bool) (client *Client, err error) {
	key := xrsa.FormatAlipayPrivateKey(privateKey)
	priKey, err := xpem.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	client = &Client{
		AppId:       appId,
		privateKey:  priKey,
		IsProd:      isProd,
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}

// 开启请求完自动验签功能（默认不开启，推荐开启，只支持证书模式）
//	注意：只支持证书模式
//	alipayPublicKeyContent：支付宝公钥证书文件内容[]byte
func (a *Client) AutoVerifySign(alipayPublicKeyContent []byte) {
	pubKey, err := xpem.DecodePublicKey(alipayPublicKeyContent)
	if err != nil {
		xlog.Errorf("AutoVerifySign(%s),err:%+v", alipayPublicKeyContent, err)
	}
	if pubKey != nil {
		a.aliPayPublicKey = pubKey
		a.autoSign = true
	}
}

// Deprecated
//	推荐使用 PostAliPayAPISelfV2()
//	示例：请参考 client_test.go 的 TestClient_PostAliPayAPISelf() 方法
func (a *Client) PostAliPayAPISelf(bm gopay.BodyMap, method string, aliRsp interface{}) (err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, method); err != nil {
		return err
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return err
	}
	return nil
}

// Deprecated
//	推荐使用 RequestParam()
func (a *Client) GetRequestSignParam(bm gopay.BodyMap, method string) (string, error) {
	return a.RequestParam(bm, method)
}

// RequestParam 获取支付宝完整请求参数包含签名
//	注意：biz_content 需要自行通过bm.SetBodyMap()设置，不设置则没有此参数
func (a *Client) RequestParam(bm gopay.BodyMap, method string) (string, error) {
	var (
		bodyBs []byte
		err    error
		sign   string
	)
	// check if there is biz_content
	bz := bm.GetInterface("biz_content")
	if bzBody, ok := bz.(gopay.BodyMap); ok {
		if bodyBs, err = json.Marshal(bzBody); err != nil {
			return "", fmt.Errorf("json.Marshal(%v)：%w", bzBody, err)
		}
		bm.Set("biz_content", string(bodyBs))
	}
	bm.Set("method", method)

	// check public parameter
	a.checkPublicParam(bm)

	// check sign
	if bm.GetString("sign") == "" {
		sign, err = GetRsaSign(bm, bm.GetString("sign_type"), a.privateKey)
		if err != nil {
			return "", fmt.Errorf("GetRsaSign Error: %v", err)
		}
		bm.Set("sign", sign)
	}

	if a.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Alipay_Request: %s", req)
	}
	return bm.EncodeURLParams(), nil
}

// PostAliPayAPISelfV2 支付宝接口自行实现方法
//	注意：biz_content 需要自行通过bm.SetBodyMap()设置，不设置则没有此参数
//	示例：请参考 client_test.go 的 TestClient_PostAliPayAPISelf() 方法
func (a *Client) PostAliPayAPISelfV2(bm gopay.BodyMap, method string, aliRsp interface{}) (err error) {
	var (
		bs, bodyBs []byte
	)
	// check if there is biz_content
	bz := bm.GetInterface("biz_content")
	if bzBody, ok := bz.(gopay.BodyMap); ok {
		if bodyBs, err = json.Marshal(bzBody); err != nil {
			return fmt.Errorf("json.Marshal(%v)：%w", bzBody, err)
		}
		bm.Set("biz_content", string(bodyBs))
	}

	if bs, err = a.doAliPaySelf(bm, method); err != nil {
		return err
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return err
	}
	return nil
}

// 向支付宝发送自定义请求
func (a *Client) doAliPaySelf(bm gopay.BodyMap, method string) (bs []byte, err error) {
	var (
		url, sign string
	)
	bm.Set("method", method)
	// check public parameter
	a.checkPublicParam(bm)
	// check sign
	if bm.GetString("sign") == "" {
		sign, err = GetRsaSign(bm, bm.GetString("sign_type"), a.privateKey)
		if err != nil {
			return nil, fmt.Errorf("GetRsaSign Error: %v", err)
		}
		bm.Set("sign", sign)
	}
	if a.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Alipay_Request: %s", req)
	}

	httpClient := xhttp.NewClient()
	if a.IsProd {
		url = baseUrlUtf8
	} else {
		url = sandboxBaseUrlUtf8
	}
	res, bs, errs := httpClient.Type(xhttp.TypeForm).Post(url).SendString(bm.EncodeURLParams()).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// 向支付宝发送请求
func (a *Client) doAliPay(bm gopay.BodyMap, method string, authToken ...string) (bs []byte, err error) {
	var (
		bodyStr, url string
		bodyBs       []byte
		aat          string
	)
	if bm != nil {
		aat = bm.GetString("app_auth_token")
		bm.Remove("app_auth_token")
		if bodyBs, err = json.Marshal(bm); err != nil {
			return nil, fmt.Errorf("json.Marshal：%w", err)
		}
		bodyStr = string(bodyBs)
	}

	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", a.AppId)
	pubBody.Set("method", method)
	pubBody.Set("format", "JSON")
	if a.AppCertSN != util.NULL {
		pubBody.Set("app_cert_sn", a.AppCertSN)
	}
	if a.AliPayRootCertSN != util.NULL {
		pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}
	if a.ReturnUrl != util.NULL {
		pubBody.Set("return_url", a.ReturnUrl)
	}
	pubBody.Set("charset", "utf-8")
	if a.Charset != util.NULL {
		pubBody.Set("charset", a.Charset)
	}
	pubBody.Set("sign_type", RSA2)
	if a.SignType != util.NULL {
		pubBody.Set("sign_type", a.SignType)
	}
	pubBody.Set("timestamp", time.Now().Format(util.TimeLayout))
	if a.LocationName != util.NULL && a.location != nil {
		pubBody.Set("timestamp", time.Now().In(a.location).Format(util.TimeLayout))
	}
	pubBody.Set("version", "1.0")
	if a.NotifyUrl != util.NULL {
		pubBody.Set("notify_url", a.NotifyUrl)
	}
	if a.AppAuthToken != util.NULL {
		pubBody.Set("app_auth_token", a.AppAuthToken)
	}
	if aat != util.NULL {
		pubBody.Set("app_auth_token", aat)
	}
	if method == "alipay.user.info.share" {
		pubBody.Set("auth_token", authToken[0])
	}

	if bodyStr != util.NULL {
		pubBody.Set("biz_content", bodyStr)
	}
	sign, err := GetRsaSign(pubBody, pubBody.GetString("sign_type"), a.privateKey)
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %v", err)
	}
	pubBody.Set("sign", sign)
	if a.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(pubBody)
		xlog.Debugf("Alipay_Request: %s", req)
	}
	param := pubBody.EncodeURLParams()
	switch method {
	case "alipay.trade.app.pay", "alipay.fund.auth.order.app.freeze":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		if !a.IsProd {
			return []byte(sandboxBaseUrl + "?" + param), nil
		}
		return []byte(baseUrl + "?" + param), nil
	default:
		httpClient := xhttp.NewClient()
		if a.IsProd {
			url = baseUrlUtf8
		} else {
			url = sandboxBaseUrlUtf8
		}
		res, bs, errs := httpClient.Type(xhttp.TypeForm).Post(url).SendString(param).EndBytes()
		if len(errs) > 0 {
			return nil, errs[0]
		}
		if a.DebugSwitch == gopay.DebugOn {
			xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
		}
		return bs, nil
	}
}

// 公共参数检查
func (a *Client) checkPublicParam(bm gopay.BodyMap) {
	bm.Set("format", "JSON")

	if bm.GetString("app_id") == "" && a.AppId != util.NULL {
		bm.Set("app_id", a.AppId)
	}
	if bm.GetString("app_cert_sn") == "" && a.AppCertSN != util.NULL {
		bm.Set("app_cert_sn", a.AppCertSN)
	}
	if bm.GetString("alipay_root_cert_sn") == "" && a.AliPayRootCertSN != util.NULL {
		bm.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}
	if bm.GetString("return_url") == "" && a.ReturnUrl != util.NULL {
		bm.Set("return_url", a.ReturnUrl)
	}
	bm.Set("charset", "utf-8")
	if bm.GetString("charset") == "" && a.Charset != util.NULL {
		bm.Set("charset", a.Charset)
	}
	bm.Set("sign_type", RSA2)
	if bm.GetString("sign_type") == "" && a.SignType != util.NULL {
		bm.Set("sign_type", a.SignType)
	}
	bm.Set("timestamp", time.Now().Format(util.TimeLayout))
	if a.LocationName != util.NULL && a.location != nil {
		bm.Set("timestamp", time.Now().In(a.location).Format(util.TimeLayout))
	}
	bm.Set("version", "1.0")
	if bm.GetString("notify_url") == "" && a.NotifyUrl != util.NULL {
		bm.Set("notify_url", a.NotifyUrl)
	}
	if bm.GetString("app_auth_token") == "" && a.AppAuthToken != util.NULL {
		bm.Set("app_auth_token", a.AppAuthToken)
	}
}
