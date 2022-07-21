package qq

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/cedarwu/gopay"
	"github.com/cedarwu/gopay/pkg/util"
	"github.com/cedarwu/gopay/pkg/xhttp"
	"github.com/cedarwu/gopay/pkg/xlog"
)

type Client struct {
	MchId       string
	ApiKey      string
	IsProd      bool
	DebugSwitch gopay.DebugSwitch
	certificate *tls.Certificate
	mu          sync.RWMutex
}

// 初始化QQ客户端（正式环境）
//	mchId：商户ID
//	ApiKey：API秘钥值
func NewClient(mchId, apiKey string) (client *Client) {
	return &Client{
		MchId:       mchId,
		ApiKey:      apiKey,
		DebugSwitch: gopay.DebugOff,
	}
}

// 向QQ发送Post请求，对于本库未提供的QQ API，可自行实现，通过此方法发送请求
//	bm：请求参数的BodyMap
//	url：完整url地址，例如：https://qpay.qq.com/cgi-bin/pay/qpay_unified_order.cgi
//	tlsConfig：tls配置，如无需证书请求，传nil
func (q *Client) PostQQAPISelf(bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {
	return q.doQQ(bm, url, tlsConfig)
}

// 提交付款码支付
//	文档地址：https://qpay.qq.com/buss/wiki/1/1122
func (q *Client) MicroPay(bm gopay.BodyMap) (qqRsp *MicroPayResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "device_info", "auth_code")
	if err != nil {
		return nil, err
	}
	bm.Set("trade_type", TradeType_MicroPay)
	bs, err := q.doQQ(bm, microPay, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(MicroPayResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 撤销订单
//	文档地址：https://qpay.qq.com/buss/wiki/1/1125
func (q *Client) Reverse(bm gopay.BodyMap) (qqRsp *ReverseResponse, err error) {
	err = bm.CheckEmptyError("sub_mch_id", "nonce_str", "out_trade_no", "op_user_id", "op_user_passwd")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQ(bm, reverse, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 统一下单
//	文档地址：https://qpay.qq.com/buss/wiki/38/1203
func (q *Client) UnifiedOrder(bm gopay.BodyMap) (qqRsp *UnifiedOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "trade_type", "notify_url")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQ(bm, unifiedOrder, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 订单查询
//	文档地址：https://qpay.qq.com/buss/wiki/38/1205
func (q *Client) OrderQuery(bm gopay.BodyMap) (qqRsp *OrderQueryResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("transaction_id") == util.NULL {
		return nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	bs, err := q.doQQ(bm, orderQuery, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(OrderQueryResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 关闭订单
//	文档地址：https://qpay.qq.com/buss/wiki/38/1206
func (q *Client) CloseOrder(bm gopay.BodyMap) (qqRsp *CloseOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQ(bm, orderClose, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 申请退款
//	注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 nil，否则，3证书Path均不可空
//	文档地址：https://qpay.qq.com/buss/wiki/38/1207
func (q *Client) Refund(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (qqRsp *RefundResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "refund_fee", "op_user_id", "op_user_passwd")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("transaction_id") == util.NULL {
		return nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	tlsConfig, err := q.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQ(bm, refund, tlsConfig)
	if err != nil {
		return nil, err
	}
	qqRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 退款查询
//	文档地址：https://qpay.qq.com/buss/wiki/38/1208
func (q *Client) RefundQuery(bm gopay.BodyMap) (qqRsp *RefundQueryResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.GetString("refund_id") == util.NULL && bm.GetString("out_refund_no") == util.NULL && bm.GetString("transaction_id") == util.NULL && bm.GetString("out_trade_no") == util.NULL {
		return nil, errors.New("refund_id, out_refund_no, out_trade_no, transaction_id are not allowed to be null at the same time")
	}
	bs, err := q.doQQ(bm, refundQuery, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(RefundQueryResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// 交易账单
//	文档地址：https://qpay.qq.com/buss/wiki/38/1209
func (q *Client) StatementDown(bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "bill_type")
	if err != nil {
		return util.NULL, err
	}
	billType := bm.GetString("bill_type")
	if billType != "ALL" && billType != "SUCCESS" && billType != "REFUND" && billType != "RECHAR" {
		return util.NULL, errors.New("bill_type error, please reference: https://qpay.qq.com/buss/wiki/38/1209")
	}
	bs, err := q.doQQ(bm, statementDown, nil)
	if err != nil {
		return util.NULL, err
	}
	return string(bs), nil
}

// 资金账单
//	文档地址：https://qpay.qq.com/buss/wiki/38/3089
func (q *Client) AccRoll(bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "acc_type")
	if err != nil {
		return util.NULL, err
	}
	accType := bm.GetString("acc_type")
	if accType != "CASH" && accType != "MARKETING" {
		return util.NULL, errors.New("acc_type error, please reference: https://qpay.qq.com/buss/wiki/38/3089")
	}
	bs, err := q.doQQ(bm, accRoll, nil)
	if err != nil {
		return util.NULL, err
	}
	return string(bs), nil
}

// 向QQ发送请求
func (q *Client) doQQ(bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {

	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", q.MchId)
	}
	if bm.GetString("fee_type") == util.NULL {
		bm.Set("fee_type", "CNY")
	}

	if bm.GetString("sign") == util.NULL {
		sign := GetReleaseSign(q.ApiKey, bm.GetString("sign_type"), bm)
		bm.Set("sign", sign)
	}

	httpClient := xhttp.NewClient()
	if tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}
	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Request: %s", bm.JsonBody())
	}
	res, bs, errs := httpClient.Type(xhttp.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

// Get请求、正式
func (q *Client) doQQGet(bm gopay.BodyMap, url, signType string) (bs []byte, err error) {
	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", q.MchId)
	}
	bm.Remove("sign")
	sign := GetReleaseSign(q.ApiKey, signType, bm)
	bm.Set("sign", sign)

	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Request: %s", bm.JsonBody())
	}
	param := bm.EncodeURLParams()
	url = url + "?" + param
	res, bs, errs := xhttp.NewClient().Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") || strings.Contains(string(bs), "html") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

func (q *Client) doQQRed(bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {

	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", q.MchId)
	}
	if bm.GetString("sign") == util.NULL {
		sign := GetReleaseSign(q.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	httpClient := xhttp.NewClient()
	if tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}
	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Request: %s", bm.JsonBody())
	}
	res, bs, errs := httpClient.Type(xhttp.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
