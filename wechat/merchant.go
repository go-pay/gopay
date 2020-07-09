package wechat

import (
	"crypto/tls"
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xhttp"
)

// 企业付款（企业向微信用户个人付款）
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (w *Client) Transfer(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *TransfersResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	if err = bm.CheckEmptyError("nonce_str", "partner_trade_no", "openid", "check_name", "amount", "desc", "spbill_create_ip"); err != nil {
		return nil, err
	}
	bm.Set("mch_appid", w.AppId)
	bm.Set("mchid", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = baseUrlCh + transfers
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	bm.Set("sign", getReleaseSign(w.ApiKey, SignType_MD5, bm))

	httpClient := xhttp.NewClient().SetTLSConfig(tlsConfig).Type(xhttp.TypeXML)
	if w.BaseURL != gotil.NULL {
		w.mu.RLock()
		url = w.BaseURL + transfers
		w.mu.RUnlock()
	}
	wxRsp = new(TransfersResponse)
	res, errs := httpClient.Post(url).SendString(generateXml(bm)).EndStruct(wxRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return wxRsp, nil
}

// 查询企业付款
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//    注意：此方法未支持沙箱环境，默认正式环境
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (w *Client) GetTransferInfo(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *TransfersInfoResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	if err = bm.CheckEmptyError("nonce_str", "partner_trade_no"); err != nil {
		return nil, err
	}
	bm.Set("appid", w.AppId)
	bm.Set("mch_id", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = baseUrlCh + getTransferInfo
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	bm.Set("sign", getReleaseSign(w.ApiKey, SignType_MD5, bm))

	httpClient := xhttp.NewClient().SetTLSConfig(tlsConfig).Type(xhttp.TypeXML)
	if w.BaseURL != gotil.NULL {
		w.mu.RLock()
		url = w.BaseURL + getTransferInfo
		w.mu.RUnlock()
	}
	wxRsp = new(TransfersInfoResponse)
	res, errs := httpClient.Post(url).SendString(generateXml(bm)).EndStruct(wxRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return wxRsp, nil
}

// 企业付款到银行卡API
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
//    注意：enc_bank_no、enc_true_name 两参数，开发者需自行获取RSA公钥，加密后再 Set 到 BodyMap，参考 client_test.go 里的 TestClient_PayBank() 方法
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
//    RSA加密文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7
//    银行编码查看地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4&index=5
func (w *Client) PayBank(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *PayBankResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	if err = bm.CheckEmptyError("partner_trade_no", "nonce_str", "enc_bank_no", "enc_true_name", "bank_code", "amount"); err != nil {
		return nil, err
	}
	bm.Set("mch_id", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = baseUrlCh + payBank
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	bm.Set("sign", getReleaseSign(w.ApiKey, SignType_MD5, bm))

	httpClient := xhttp.NewClient().SetTLSConfig(tlsConfig).Type(xhttp.TypeXML)
	if w.BaseURL != gotil.NULL {
		w.mu.RLock()
		url = w.BaseURL + payBank
		w.mu.RUnlock()
	}
	wxRsp = new(PayBankResponse)
	res, errs := httpClient.Post(url).SendString(generateXml(bm)).EndStruct(wxRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return wxRsp, nil
}

// 查询企业付款到银行卡API
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//    注意：此方法未支持沙箱环境，默认正式环境
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3
func (w *Client) QueryBank(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *QueryBankResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	if err = bm.CheckEmptyError("nonce_str", "partner_trade_no"); err != nil {
		return nil, err
	}
	bm.Set("mch_id", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = baseUrlCh + queryBank
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	bm.Set("sign", getReleaseSign(w.ApiKey, SignType_MD5, bm))

	httpClient := xhttp.NewClient().SetTLSConfig(tlsConfig).Type(xhttp.TypeXML)
	if w.BaseURL != gotil.NULL {
		w.mu.RLock()
		url = w.BaseURL + queryBank
		w.mu.RUnlock()
	}
	wxRsp = new(QueryBankResponse)
	res, errs := httpClient.Post(url).SendString(generateXml(bm)).EndStruct(wxRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return wxRsp, nil
}

// 获取RSA加密公钥API
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//    注意：此方法未支持沙箱环境，默认正式环境
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4
func (w *Client) GetRSAPublicKey(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *RSAPublicKeyResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	if err = bm.CheckEmptyError("nonce_str", "sign_type"); err != nil {
		return nil, err
	}
	bm.Set("mch_id", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = getPublicKey
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	bm.Set("sign", getReleaseSign(w.ApiKey, bm.Get("sign_type"), bm))

	httpClient := xhttp.NewClient().SetTLSConfig(tlsConfig).Type(xhttp.TypeXML)
	wxRsp = new(RSAPublicKeyResponse)
	res, errs := httpClient.Post(url).SendString(generateXml(bm)).EndStruct(wxRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return wxRsp, nil
}
