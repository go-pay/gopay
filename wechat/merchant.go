package wechat

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xhttp"
	"github.com/iGoogle-ink/gotil/xlog"
)

// 企业付款（企业向微信用户个人付款）
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//	注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
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
	if w.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_Request: %s", req)
	}
	res, bs, errs := httpClient.Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if w.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(TransfersResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 查询企业付款
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//	注意：此方法未支持沙箱环境，默认正式环境
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
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
	if w.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_Request: %s", req)
	}
	res, bs, errs := httpClient.Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if w.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(TransfersInfoResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 企业付款到银行卡API
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//	注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
//	注意：enc_bank_no、enc_true_name 两参数，开发者需自行获取RSA公钥，加密后再 Set 到 BodyMap，参考 client_test.go 里的 TestClient_PayBank() 方法
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
//	RSA加密文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7
//	银行编码查看地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4&index=5
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
	if w.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_Request: %s", req)
	}
	res, bs, errs := httpClient.Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if w.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(PayBankResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 查询企业付款到银行卡API
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//	注意：此方法未支持沙箱环境，默认正式环境
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3
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
	if w.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_Request: %s", req)
	}
	res, bs, errs := httpClient.Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if w.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(QueryBankResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 获取RSA加密公钥API
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则3证书Path均不可为nil（string类型）
//	注意：此方法未支持沙箱环境，默认正式环境
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4
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
	if w.DebugSwitch == gopay.DebugOn {
		req, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_Request: %s", req)
	}
	res, bs, errs := httpClient.Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if w.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(RSAPublicKeyResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharing 请求单次分账
//	单次分账请求按照传入的分账接收方账号和资金进行分账，
//	同时会将订单剩余的待分账金额解冻给本商户。
//	故操作成功后，订单不能再进行分账，也不能进行分账完结。
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
func (w *Client) ProfitSharing(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ProfitSharingResponse, err error) {
	return w.profitSharing(bm, profitSharing, certFilePath, keyFilePath, pkcs12FilePath)
}

// MultiProfitSharing 请求多次分账
//	微信订单支付成功后，商户发起分账请求，将结算后的钱分到分账接收方。多次分账请求仅会按照传入的分账接收方进行分账，不会对剩余的金额进行任何操作。
//	故操作成功后，在待分账金额不等于零时，订单依旧能够再次进行分账。
//	多次分账，可以将本商户作为分账接收方直接传入，实现释放资金给本商户的功能
//	对同一笔订单最多能发起20次多次分账请求
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
func (w *Client) MultiProfitSharing(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ProfitSharingResponse, err error) {
	return w.profitSharing(bm, multiProfitSharing, certFilePath, keyFilePath, pkcs12FilePath)
}

func (w *Client) profitSharing(bm gopay.BodyMap, uri string, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ProfitSharingResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "transaction_id", "out_order_no", "receivers")
	if err != nil {
		return nil, err
	}

	arr, err := bm.GetArrayBodyMap("receivers")
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, errors.New("receivers is empty")
	}
	// 检查每个分账接收者的必传属性
	for _, r := range arr {
		err = r.CheckEmptyError("type", "account", "amount", "description")
		if err != nil {
			return nil, err
		}
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, uri, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharingQuery 查询分账结果
//	发起分账请求后，可调用此接口查询分账结果；发起分账完结请求后，可调用此接口查询分账完结的执行结果。
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3
func (w *Client) ProfitSharingQuery(bm gopay.BodyMap) (wxRsp *ProfitSharingQueryResponse, err error) {
	err = bm.CheckEmptyError("transaction_id", "out_order_no", "nonce_str")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	func() {
		w.mu.RLock()
		defer w.mu.RUnlock()
		bm.Set("mch_id", w.MchId)
		if bm.Get("sign") == gotil.NULL {
			sign := getReleaseSign(w.ApiKey, bm.Get("sign_type"), bm)
			bm.Set("sign", sign)
		}
	}()
	bs, err := w.doProdPostPure(bm, profitSharingQuery, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingQueryResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharingAddReceiver 添加分账接收方
//	商户发起添加分账接收方请求，后续可通过发起分账请求将结算后的钱分到该分账接收方。
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4
func (w *Client) ProfitSharingAddReceiver(bm gopay.BodyMap) (wxRsp *ProfitSharingAddReceiverResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "receiver")
	if err != nil {
		return nil, err
	}
	// 输入参数 接收方
	r, err := bm.GetBodyMap("receiver")
	if err != nil {
		return nil, err
	}
	err = r.CheckEmptyError("type", "account", "relation_type")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPost(bm, profitSharingAddReceiver, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingAddReceiverResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharingRemoveReceiver 删除分账接收方
//	商户发起删除分账接收方请求，删除后不支持将结算后的钱分到该分账接收方
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5
func (w *Client) ProfitSharingRemoveReceiver(bm gopay.BodyMap) (wxRsp *ProfitSharingAddReceiverResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "receiver")
	if err != nil {
		return nil, err
	}
	// 输入参数 接收方
	r, err := bm.GetBodyMap("receiver")
	if err != nil {
		return nil, err
	}
	err = r.CheckEmptyError("type", "account")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPost(bm, profitSharingRemoveReceiver, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingAddReceiverResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharingFinish 完结分账
//	1、不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给本商户
//	2、调用多次分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
//	3、已调用请求单次分账后，剩余待分账金额为零，不需要再调用此接口。
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6
func (w *Client) ProfitSharingFinish(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ProfitSharingResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "transaction_id", "out_order_no", "description")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, profitSharingFinish, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharingReturn 分账回退
//	对订单进行退款时，如果订单已经分账，可以先调用此接口将指定的金额从分账接收方（仅限商户类型的分账接收方）回退给本商户，然后再退款。
//	回退以原分账请求为依据，可以对分给分账接收方的金额进行多次回退，只要满足累计回退不超过该请求中分给接收方的金额。
//	此接口采用同步处理模式，即在接收到商户请求后，会实时返回处理结果
//	此功能需要接收方在商户平台-交易中心-分账-分账接收设置下，开启同意分账回退后，才能使用。
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7
func (w *Client) ProfitSharingReturn(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ProfitSharingReturnResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_return_no", "return_account_type", "return_account", "return_amount", "description")
	if err != nil {
		return nil, err
	}

	if (bm.Get("order_id") == gotil.NULL) && (bm.Get("out_order_no") == gotil.NULL) {
		return nil, errors.New("param order_id and out_order_no can not be null at the same time")
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, profitSharingReturn, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingReturnResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharingReturnQuery 回退结果查询
//	商户需要核实回退结果，可调用此接口查询回退结果。
//	如果分账回退接口返回状态为处理中，可调用此接口查询回退结果
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8
func (w *Client) ProfitSharingReturnQuery(bm gopay.BodyMap) (wxRsp *ProfitSharingReturnResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_return_no")
	if err != nil {
		return nil, err
	}

	if (bm.Get("order_id") == gotil.NULL) && (bm.Get("out_order_no") == gotil.NULL) {
		return nil, errors.New("param order_id and out_order_no can not be null at the same time")
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPost(bm, profitSharingReturnQuery, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingReturnResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
