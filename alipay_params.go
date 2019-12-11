package gopay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"net/url"
)

//	AppId      string `json:"app_id"`      //支付宝分配给开发者的应用ID
//	Method     string `json:"method"`      //接口名称
//	Format     string `json:"format"`      //仅支持 JSON
//	ReturnUrl  string `json:"return_url"`  //HTTP/HTTPS开头字符串
//	Charset    string `json:"charset"`     //请求使用的编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
//	SignType   string `json:"sign_type"`   //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
//	Sign       string `json:"sign"`        //商户请求参数的签名串
//	Timestamp  string `json:"timestamp"`   //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
//	Version    string `json:"version"`     //调用的接口版本，固定为：1.0
//	NotifyUrl  string `json:"notify_url"`  //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
//	BizContent string `json:"biz_content"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档

type OpenApiRoyaltyDetailInfoPojo struct {
	RoyaltyType  string `json:"royalty_type,omitempty"`
	TransOut     string `json:"trans_out,omitempty"`
	TransOutType string `json:"trans_out_type,omitempty"`
	TransInType  string `json:"trans_in_type,omitempty"`
	TransIn      string `json:"trans_in"`
	Amount       string `json:"amount,omitempty"`
	Desc         string `json:"desc,omitempty"`
}

// 设置 应用公钥证书SN
//    appCertSN：应用公钥证书SN，通过 gopay.GetCertSN() 获取
func (a *AliPayClient) SetAppCertSN(appCertSN string) (client *AliPayClient) {
	a.mu.Lock()
	a.AppCertSN = appCertSN
	a.mu.Unlock()
	return a
}

// 设置 支付宝公钥证书SN
//    aliPayPublicCertSN：支付宝公钥证书SN，通过 gopay.GetCertSN() 获取
func (a *AliPayClient) SetAliPayPublicCertSN(aliPayPublicCertSN string) (client *AliPayClient) {
	a.mu.Lock()
	a.AliPayPublicCertSN = aliPayPublicCertSN
	a.mu.Unlock()
	return a
}

// 设置 支付宝CA根证书SN
//    aliPayRootCertSN：支付宝CA根证书SN，通过 gopay.GetRootCertSN() 获取
func (a *AliPayClient) SetAliPayRootCertSN(aliPayRootCertSN string) (client *AliPayClient) {
	a.mu.Lock()
	a.AliPayRootCertSN = aliPayRootCertSN
	a.mu.Unlock()
	return a
}

// 设置 app_cert_sn、alipay_root_cert_sn、alipay_cert_sn 通过应用公钥证书路径
//    appCertPath：应用公钥证书路径
//    aliPayRootCertPath：支付宝根证书文件路径
//    aliPayPublicCertPath：支付宝公钥证书文件路径
func (a *AliPayClient) SetCertSnByPath(appCertPath, aliPayRootCertPath, aliPayPublicCertPath string) (err error) {
	appCertSn, err := GetCertSN(appCertPath)
	if err != nil {
		return fmt.Errorf("get app_cert_sn return err, but alse return alipay client. err: %v", err)
	}
	rootCertSn, err := GetRootCertSN(aliPayRootCertPath)
	if err != nil {
		return fmt.Errorf("get alipay_root_cert_sn return err, but alse return alipay client. err: %v", err)
	}
	publicCertSn, err := GetCertSN(aliPayPublicCertPath)
	if err != nil {
		return fmt.Errorf("get alipay_cert_sn return err, but alse return alipay client. err: %v", err)
	}
	a.mu.Lock()
	a.AppCertSN = appCertSn
	a.AliPayRootCertSN = rootCertSn
	a.AliPayPublicCertSN = publicCertSn
	a.mu.Unlock()
	return nil
}

// 设置支付后的ReturnUrl
func (a *AliPayClient) SetReturnUrl(url string) (client *AliPayClient) {
	a.mu.Lock()
	a.ReturnUrl = url
	a.mu.Unlock()
	return a
}

// 设置支付宝服务器主动通知商户服务器里指定的页面http/https路径。
func (a *AliPayClient) SetNotifyUrl(url string) (client *AliPayClient) {
	a.mu.Lock()
	a.NotifyUrl = url
	a.mu.Unlock()
	return a
}

// 设置编码格式，如utf-8,gbk,gb2312等，默认推荐使用 utf-8
func (a *AliPayClient) SetCharset(charset string) (client *AliPayClient) {
	a.mu.Lock()
	if charset == null {
		a.Charset = "utf-8"
	} else {
		a.Charset = charset
	}
	a.mu.Unlock()
	return a
}

// 设置签名算法类型，目前支持RSA2和RSA，默认推荐使用 RSA2
func (a *AliPayClient) SetSignType(signType string) (client *AliPayClient) {
	a.mu.Lock()
	if signType == null {
		a.SignType = "RSA2"
	} else {
		a.SignType = signType
	}
	a.mu.Unlock()
	return a
}

// 设置应用授权
func (a *AliPayClient) SetAppAuthToken(appAuthToken string) (client *AliPayClient) {
	a.mu.Lock()
	a.AppAuthToken = appAuthToken
	a.mu.Unlock()
	return a
}

// 设置用户信息授权
func (a *AliPayClient) SetAuthToken(authToken string) (client *AliPayClient) {
	a.mu.Lock()
	a.AuthToken = authToken
	a.mu.Unlock()
	return a
}

// 获取参数签名
func getRsaSign(bm BodyMap, signType, privateKey string) (sign string, err error) {
	var (
		block          *pem.Block
		h              hash.Hash
		key            *rsa.PrivateKey
		hashs          crypto.Hash
		encryptedBytes []byte
	)

	if block, _ = pem.Decode([]byte(privateKey)); block == nil {
		return null, errors.New("pem.Decode：privateKey decode error")
	}
	if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}
	switch signType {
	case "RSA":
		h = sha1.New()
		hashs = crypto.SHA1
	case "RSA2":
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	if _, err = h.Write([]byte(bm.EncodeAliPaySignParams())); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, key, hashs, h.Sum(nil)); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// 格式化请求URL参数
func FormatAliPayURLParam(body BodyMap) (urlParam string) {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}
