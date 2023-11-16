package alipay

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xrsa"
)

type Client struct {
	AppId              string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	SignType           string
	AppAuthToken       string
	IsProd             bool
	aesKey             string // biz_content 加密的 AES KEY
	ivKey              []byte
	privateKey         *rsa.PrivateKey
	aliPayPublicKey    *rsa.PublicKey // 支付宝证书公钥内容 alipayPublicCert.crt
	autoSign           bool
	DebugSwitch        gopay.DebugSwitch
	location           *time.Location
	hc                 *xhttp.Client
}

// 初始化支付宝客户端
// 注意：如果使用支付宝公钥证书验签，请使用 client.SetCertSnByContent() 或 client.SetCertSnByPath() 设置 应用公钥证书、支付宝公钥证书、支付宝根证书
// appid：应用ID
// privateKey：应用私钥，支持PKCS1和PKCS8
// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
func NewClient(appid, privateKey string, isProd bool) (client *Client, err error) {
	if appid == util.NULL || privateKey == util.NULL {
		return nil, gopay.MissAlipayInitParamErr
	}
	key := xrsa.FormatAlipayPrivateKey(privateKey)
	priKey, err := xpem.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	client = &Client{
		AppId:       appid,
		Charset:     UTF8,
		SignType:    RSA2,
		IsProd:      isProd,
		privateKey:  priKey,
		DebugSwitch: gopay.DebugOff,
		hc:          xhttp.NewClient(),
	}
	return client, nil
}

// 开启请求完自动验签功能（默认不开启，推荐开启，只支持证书模式）
// 注意：只支持证书模式
// alipayPublicKeyContent：支付宝公钥证书文件内容[]byte
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

// SetBodySize 设置http response body size(MB)
func (a *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		a.hc.SetBodySize(sizeMB)
	}
}

// SetAESKey 设置 biz_content 的AES加密key，设置此参数默认开启 biz_content 参数加密
func (a *Client) SetAESKey(aesKey string) {
	a.aesKey = aesKey
	a.ivKey = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

// Deprecated
// 推荐使用 RequestParam()
func (a *Client) GetRequestSignParam(bm gopay.BodyMap, method string) (string, error) {
	return a.RequestParam(bm, method)
}

// RequestParam 获取支付宝完整请求参数包含签名
// 注意：biz_content 需要自行通过bm.SetBodyMap()设置，不设置则没有此参数
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
		sign, err = a.getRsaSign(bm, bm.GetString("sign_type"))
		if err != nil {
			return "", fmt.Errorf("GetRsaSign Error: %w", err)
		}
		bm.Set("sign", sign)
	}

	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request: %s", bm.JsonBody())
	}
	return bm.EncodeURLParams(), nil
}

// 公共参数处理
func (a *Client) pubParamsHandle(bm gopay.BodyMap, method, bizContent string, authToken ...string) (param string, err error) {
	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", a.AppId).
		Set("method", method).
		Set("format", "JSON").
		Set("charset", a.Charset).
		Set("sign_type", a.SignType).
		Set("version", "1.0").
		Set("timestamp", time.Now().Format(util.TimeLayout))

	// version
	if version := bm.GetString("version"); version != util.NULL {
		pubBody.Set("version", version)
	}
	if a.AppCertSN != util.NULL {
		pubBody.Set("app_cert_sn", a.AppCertSN)
	}
	if a.AliPayRootCertSN != util.NULL {
		pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}
	// return_url
	if a.ReturnUrl != util.NULL {
		pubBody.Set("return_url", a.ReturnUrl)
	}
	if returnUrl := bm.GetString("return_url"); returnUrl != util.NULL {
		pubBody.Set("return_url", returnUrl)
	}
	if a.location != nil {
		pubBody.Set("timestamp", time.Now().In(a.location).Format(util.TimeLayout))
	}
	// notify_url
	if a.NotifyUrl != util.NULL {
		pubBody.Set("notify_url", a.NotifyUrl)
	}
	if notifyUrl := bm.GetString("notify_url"); notifyUrl != util.NULL {
		pubBody.Set("notify_url", notifyUrl)
	}
	// default use app_auth_token
	if a.AppAuthToken != util.NULL {
		pubBody.Set("app_auth_token", a.AppAuthToken)
	}
	// if user set app_auth_token in body_map, use this
	if aat := bm.GetString("app_auth_token"); aat != util.NULL {
		pubBody.Set("app_auth_token", aat)
	}
	if len(authToken) > 0 {
		pubBody.Set("auth_token", authToken[0])
	}
	if bizContent != util.NULL {
		if a.aesKey == util.NULL {
			pubBody.Set("biz_content", bizContent)
		} else {
			// AES Encrypt biz_content
			encryptBizContent, err := a.encryptBizContent(bizContent)
			if err != nil {
				return "", fmt.Errorf("EncryptBizContent Error: %w", err)
			}
			if a.DebugSwitch == gopay.DebugOn {
				xlog.Debugf("Alipay_Origin_BizContent: %s", bizContent)
				xlog.Debugf("Alipay_Encrypt_BizContent: %s", encryptBizContent)
			}
			pubBody.Set("biz_content", encryptBizContent)
		}
	}
	// sign
	sign, err := a.getRsaSign(pubBody, pubBody.GetString("sign_type"))
	if err != nil {
		return "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	pubBody.Set("sign", sign)
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request: %s", pubBody.JsonBody())
	}
	param = pubBody.EncodeURLParams()
	return
}

// 公共参数检查
func (a *Client) checkPublicParam(bm gopay.BodyMap) {
	bm.Set("format", "JSON").
		Set("charset", a.Charset).
		Set("sign_type", a.SignType).
		Set("version", "1.0").
		Set("timestamp", time.Now().Format(util.TimeLayout))

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
	if a.location != nil {
		bm.Set("timestamp", time.Now().In(a.location).Format(util.TimeLayout))
	}
	if bm.GetString("notify_url") == "" && a.NotifyUrl != util.NULL {
		bm.Set("notify_url", a.NotifyUrl)
	}
	if bm.GetString("app_auth_token") == "" && a.AppAuthToken != util.NULL {
		bm.Set("app_auth_token", a.AppAuthToken)
	}
}

func (a *Client) encryptBizContent(originData string) (string, error) {
	encryptData, err := aes.CBCEncrypt([]byte(originData), []byte(a.aesKey), a.ivKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptData), nil
}
