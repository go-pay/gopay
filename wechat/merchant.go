package wechat

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// 企业付款（企业向微信用户个人付款）
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (w *Client) Transfer(ctx context.Context, bm gopay.BodyMap) (wxRsp *TransfersResponse, err error) {
	if err = bm.CheckEmptyError("nonce_str", "partner_trade_no", "openid", "check_name", "amount", "desc", "spbill_create_ip"); err != nil {
		return nil, err
	}
	bm.Set("mch_appid", w.AppId)
	bm.Set("mchid", w.MchId)
	var (
		url = baseUrlCh + transfers
	)
	bm.Set("sign", w.getReleaseSign(w.ApiKey, SignType_MD5, bm))

	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + transfers
		w.mu.RUnlock()
	}
	req := GenerateXml(bm)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request: %s", req)
	}
	res, bs, err := w.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Response: %d, %s", res.StatusCode, string(bs))

	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(TransfersResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 查询企业付款
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此方法未支持沙箱环境，默认正式环境
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (w *Client) GetTransferInfo(ctx context.Context, bm gopay.BodyMap) (wxRsp *TransfersInfoResponse, err error) {
	if err = bm.CheckEmptyError("nonce_str", "partner_trade_no"); err != nil {
		return nil, err
	}
	bm.Set("appid", w.AppId)
	bm.Set("mch_id", w.MchId)
	var (
		url = baseUrlCh + getTransferInfo
	)
	bm.Set("sign", w.getReleaseSign(w.ApiKey, SignType_MD5, bm))

	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + getTransferInfo
		w.mu.RUnlock()
	}
	req := GenerateXml(bm)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request: %s", req)
	}
	res, bs, err := w.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(TransfersInfoResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 企业付款到银行卡API（正式）
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
// 注意：enc_bank_no、enc_true_name 两参数，开发者需自行获取RSA公钥，加密后再 Set 到 BodyMap，参考 client_test.go 里的 TestClient_PayBank() 方法
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
// RSA加密文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7
// 银行编码查看地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4&index=5
func (w *Client) PayBank(ctx context.Context, bm gopay.BodyMap) (wxRsp *PayBankResponse, err error) {
	if err = bm.CheckEmptyError("partner_trade_no", "nonce_str", "enc_bank_no", "enc_true_name", "bank_code", "amount"); err != nil {
		return nil, err
	}
	bm.Set("mch_id", w.MchId)
	var (
		url = baseUrlCh + payBank
	)
	bm.Set("sign", w.getReleaseSign(w.ApiKey, SignType_MD5, bm))

	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + payBank
		w.mu.RUnlock()
	}
	req := GenerateXml(bm)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request: %s", req)
	}
	res, bs, err := w.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(PayBankResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 查询企业付款到银行卡API（正式）
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3
func (w *Client) QueryBank(ctx context.Context, bm gopay.BodyMap) (wxRsp *QueryBankResponse, err error) {
	if err = bm.CheckEmptyError("nonce_str", "partner_trade_no"); err != nil {
		return nil, err
	}
	bm.Set("mch_id", w.MchId)
	var (
		url = baseUrlCh + queryBank
	)
	bm.Set("sign", w.getReleaseSign(w.ApiKey, SignType_MD5, bm))

	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + queryBank
		w.mu.RUnlock()
	}
	req := GenerateXml(bm)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request: %s", req)
	}
	res, bs, err := w.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(QueryBankResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 获取RSA加密公钥API（正式）
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4
func (w *Client) GetRSAPublicKey(ctx context.Context, bm gopay.BodyMap) (wxRsp *RSAPublicKeyResponse, err error) {
	if err = bm.CheckEmptyError("nonce_str", "sign_type"); err != nil {
		return nil, err
	}
	bm.Set("mch_id", w.MchId)
	var (
		url = getPublicKey
	)
	bm.Set("sign", w.getReleaseSign(w.ApiKey, bm.GetString("sign_type"), bm))

	req := GenerateXml(bm)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request: %s", req)
	}
	res, bs, err := w.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(RSAPublicKeyResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 请求单次分账
// 单次分账请求按照传入的分账接收方账号和资金进行分账，
// 同时会将订单剩余的待分账金额解冻给本商户。
// 故操作成功后，订单不能再进行分账，也不能进行分账完结。
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
func (w *Client) ProfitSharing(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingResponse, err error) {
	return w.profitSharing(ctx, bm, profitSharing)
}

// 请求多次分账
// 微信订单支付成功后，商户发起分账请求，将结算后的钱分到分账接收方。多次分账请求仅会按照传入的分账接收方进行分账，不会对剩余的金额进行任何操作。
// 故操作成功后，在待分账金额不等于零时，订单依旧能够再次进行分账。
// 多次分账，可以将本商户作为分账接收方直接传入，实现释放资金给本商户的功能
// 对同一笔订单最多能发起20次多次分账请求
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
func (w *Client) MultiProfitSharing(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingResponse, err error) {
	return w.profitSharing(ctx, bm, multiProfitSharing)
}

func (w *Client) profitSharing(ctx context.Context, bm gopay.BodyMap, uri string) (wxRsp *ProfitSharingResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "transaction_id", "out_order_no", "receivers")
	if err != nil {
		return nil, err
	}

	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPostTLS(ctx, bm, uri)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 查询分账结果
// 发起分账请求后，可调用此接口查询分账结果；发起分账完结请求后，可调用此接口查询分账完结的执行结果。
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3
func (w *Client) ProfitSharingQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingQueryResponse, err error) {
	err = bm.CheckEmptyError("transaction_id", "out_order_no", "nonce_str")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bm.Set("mch_id", w.MchId)
	if bm.GetString("sign") == gopay.NULL {
		sign := w.getReleaseSign(w.ApiKey, bm.GetString("sign_type"), bm)
		bm.Set("sign", sign)
	}
	bs, err := w.doProdPostPure(ctx, bm, profitSharingQuery)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingQueryResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 添加分账接收方
// 商户发起添加分账接收方请求，后续可通过发起分账请求将结算后的钱分到该分账接收方。
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4
func (w *Client) ProfitSharingAddReceiver(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingAddReceiverResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "receiver")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPost(ctx, bm, profitSharingAddReceiver)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingAddReceiverResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 删除分账接收方
// 商户发起删除分账接收方请求，删除后不支持将结算后的钱分到该分账接收方
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5
func (w *Client) ProfitSharingRemoveReceiver(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingAddReceiverResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "receiver")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPost(ctx, bm, profitSharingRemoveReceiver)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingAddReceiverResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 完结分账
// 1、不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给本商户
// 2、调用多次分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
// 3、已调用请求单次分账后，剩余待分账金额为零，不需要再调用此接口。
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6
func (w *Client) ProfitSharingFinish(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "transaction_id", "out_order_no", "description")
	if err != nil {
		return nil, err
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPostTLS(ctx, bm, profitSharingFinish)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 服务商可通过调用此接口查询订单剩余待分金额
// 接口频率：30QPS
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_10&index=7
func (w *Client) ProfitSharingOrderAmountQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingOrderAmountQueryResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "transaction_id", "nonce_str")
	if err != nil {
		return nil, err
	}

	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPostTLS(ctx, bm, profitSharingOrderAmountQuery)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingOrderAmountQueryResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 服务商可以查询子商户设置的允许服务商分账的最大比例
// 接口频率：30QPS
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_10&index=7
func (w *Client) ProfitSharingMerchantRatioQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingMerchantRatioQuery, err error) {
	err = bm.CheckEmptyError("mch_id", "nonce_str")
	if err != nil {
		return nil, err
	}
	if (bm.GetString("sub_mch_id") == gopay.NULL) && (bm.GetString("brand_mch_id") == gopay.NULL) {
		return nil, errors.New("param sub_mch_id and brand_mch_id can not be null at the same time")
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPostTLS(ctx, bm, profitSharingMerchantRatioQuery)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingMerchantRatioQuery)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 分账回退
// 对订单进行退款时，如果订单已经分账，可以先调用此接口将指定的金额从分账接收方（仅限商户类型的分账接收方）回退给本商户，然后再退款。
// 回退以原分账请求为依据，可以对分给分账接收方的金额进行多次回退，只要满足累计回退不超过该请求中分给接收方的金额。
// 此接口采用同步处理模式，即在接收到商户请求后，会实时返回处理结果
// 此功能需要接收方在商户平台-交易中心-分账-分账接收设置下，开启同意分账回退后，才能使用。
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7
func (w *Client) ProfitSharingReturn(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingReturnResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_return_no", "return_account_type", "return_account", "return_amount", "description")
	if err != nil {
		return nil, err
	}

	if (bm.GetString("order_id") == gopay.NULL) && (bm.GetString("out_order_no") == gopay.NULL) {
		return nil, errors.New("param order_id and out_order_no can not be null at the same time")
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPostTLS(ctx, bm, profitSharingReturn)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingReturnResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 回退结果查询
// 商户需要核实回退结果，可调用此接口查询回退结果。
// 如果分账回退接口返回状态为处理中，可调用此接口查询回退结果
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8
func (w *Client) ProfitSharingReturnQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *ProfitSharingReturnResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_return_no")
	if err != nil {
		return nil, err
	}

	if (bm.GetString("order_id") == gopay.NULL) && (bm.GetString("out_order_no") == gopay.NULL) {
		return nil, errors.New("param order_id and out_order_no can not be null at the same time")
	}
	// 设置签名类型，官方文档此接口只支持 HMAC_SHA256
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bs, err := w.doProdPost(ctx, bm, profitSharingReturnQuery)
	if err != nil {
		return nil, err
	}
	wxRsp = new(ProfitSharingReturnResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}
