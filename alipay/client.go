package alipay

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
	"github.com/iGoogle-ink/gopay/pkg/xhttp"
	"github.com/iGoogle-ink/gopay/pkg/xlog"
)

type Client struct {
	AppId              string
	PrivateKeyType     PKCSType
	PrivateKey         string
	LocationName       string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	SignType           string
	AppAuthToken       string
	AuthToken          string
	IsProd             bool
	DebugSwitch        gopay.DebugSwitch
	location           *time.Location
	mu                 sync.RWMutex
}

// 初始化支付宝客户端
//	注意：如果使用支付宝公钥证书验签，请设置 支付宝根证书SN（client.SetAlipayRootCertSN()）、应用公钥证书SN（client.SetAppCertSN()）
//	appId：应用ID
//	privateKey：应用私钥，支持PKCS1和PKCS8
//	isProd：是否是正式环境
func NewClient(appId, privateKey string, isProd bool) (client *Client) {
	return &Client{
		AppId:       appId,
		PrivateKey:  privateKey,
		IsProd:      isProd,
		DebugSwitch: gopay.DebugOff,
	}
}

// PostAliPayAPISelf 支付宝接口自行实现方法
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

// 向支付宝发送请求
func (a *Client) doAliPay(bm gopay.BodyMap, method string) (bs []byte, err error) {
	var (
		bodyStr, url string
		bodyBs       []byte
	)
	if bm != nil {
		if bodyBs, err = json.Marshal(bm); err != nil {
			return nil, fmt.Errorf("json.Marshal：%w", err)
		}
		bodyStr = string(bodyBs)
	}
	pubBody := make(gopay.BodyMap)
	func() {
		a.mu.RLock()
		defer a.mu.RUnlock()

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
		if a.AuthToken != util.NULL {
			pubBody.Set("auth_token", a.AuthToken)
		}
	}()

	if bodyStr != util.NULL {
		pubBody.Set("biz_content", bodyStr)
	}
	sign, err := GetRsaSign(pubBody, pubBody.Get("sign_type"), a.PrivateKeyType, a.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %v", err)
	}
	pubBody.Set("sign", sign)
	if a.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(pubBody)
		xlog.Debugf("Alipay_Request: %s", req)
	}
	param := FormatURLParam(pubBody)
	switch method {
	case "alipay.trade.app.pay":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		if !a.IsProd {
			return []byte(sandboxBaseUrl + "?" + param), nil
		}
		return []byte(baseUrl + "?" + param), nil
	default:
		httpClient := xhttp.NewClient()
		if !a.IsProd {
			url = sandboxBaseUrlUtf8
		} else {
			url = baseUrlUtf8
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

func getSignData(bs []byte) (signData string) {
	str := string(bs)
	indexStart := strings.Index(str, `":`)
	indexEnd := strings.Index(str, `,"sign"`)
	signData = str[indexStart+2 : indexEnd]
	return
}
