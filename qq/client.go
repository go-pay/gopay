package qq

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"hash"
	"strings"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
)

type Client struct {
	MchId       string
	ApiKey      string
	IsProd      bool
	DebugSwitch gopay.DebugSwitch
	logger      xlog.XLogger
	mu          sync.RWMutex
	sha256Hash  hash.Hash
	md5Hash     hash.Hash
	hc          *xhttp.Client
	tlsHc       *xhttp.Client
}

// 初始化QQ客户端（正式环境）
// mchId：商户ID
// ApiKey：API秘钥值
func NewClient(mchId, apiKey string) (client *Client) {
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	return &Client{
		MchId:       mchId,
		ApiKey:      apiKey,
		DebugSwitch: gopay.DebugOff,
		logger:      logger,
		sha256Hash:  hmac.New(sha256.New, []byte(apiKey)),
		md5Hash:     md5.New(),
		hc:          xhttp.NewClient(),
		tlsHc:       xhttp.NewClient(),
	}
}

// SetBodySize 设置http response body size(MB)
func (q *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		q.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的xhttp.Client
func (q *Client) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		q.hc = client
	}
}

// SetTLSHttpClient 设置自定义的xhttp.Client
func (q *Client) SetTLSHttpClient(client *xhttp.Client) {
	if client != nil {
		q.tlsHc = client
	}
}

func (q *Client) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		q.logger = logger
	}
}

// 向QQ发送Post请求，对于本库未提供的QQ API，可自行实现，通过此方法发送请求
// bm：请求参数的BodyMap
// url：完整url地址，例如：https://qpay.qq.com/cgi-bin/pay/qpay_unified_order.cgi
// tlsConfig：tls配置，如无需证书请求，传nil
func (q *Client) PostQQAPISelf(ctx context.Context, bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", q.MchId)
	}
	if bm.GetString("fee_type") == gopay.NULL {
		bm.Set("fee_type", "CNY")
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := GetReleaseSign(q.ApiKey, bm.GetString("sign_type"), bm)
		bm.Set("sign", sign)
	}
	req := GenerateXml(bm)
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Request: %s", req)
	}
	httpClient := xhttp.NewClient()
	if q.IsProd && tlsConfig != nil {
		httpClient.SetHttpTLSConfig(tlsConfig)
	}
	res, bs, err := httpClient.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

// 提交付款码支付
// 文档地址：https://qpay.qq.com/buss/wiki/1/1122
func (q *Client) MicroPay(ctx context.Context, bm gopay.BodyMap) (qqRsp *MicroPayResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "device_info", "auth_code")
	if err != nil {
		return nil, err
	}
	bm.Set("trade_type", TradeType_MicroPay)
	bs, err := q.doQQPost(ctx, bm, microPay)
	if err != nil {
		return nil, err
	}
	qqRsp = new(MicroPayResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 撤销订单
// 文档地址：https://qpay.qq.com/buss/wiki/1/1125
func (q *Client) Reverse(ctx context.Context, bm gopay.BodyMap) (qqRsp *ReverseResponse, err error) {
	err = bm.CheckEmptyError("sub_mch_id", "nonce_str", "out_trade_no", "op_user_id", "op_user_passwd")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQPost(ctx, bm, reverse)
	if err != nil {
		return nil, err
	}
	qqRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 统一下单
// 文档地址：https://qpay.qq.com/buss/wiki/38/1203
func (q *Client) UnifiedOrder(ctx context.Context, bm gopay.BodyMap) (qqRsp *UnifiedOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "trade_type", "notify_url")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQPost(ctx, bm, unifiedOrder)
	if err != nil {
		return nil, err
	}
	qqRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 订单查询
// 文档地址：https://qpay.qq.com/buss/wiki/38/1205
func (q *Client) OrderQuery(ctx context.Context, bm gopay.BodyMap) (qqRsp *OrderQueryResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("transaction_id") == gopay.NULL {
		return nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	bs, err := q.doQQPost(ctx, bm, orderQuery)
	if err != nil {
		return nil, err
	}
	qqRsp = new(OrderQueryResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 关闭订单
// 文档地址：https://qpay.qq.com/buss/wiki/38/1206
func (q *Client) CloseOrder(ctx context.Context, bm gopay.BodyMap) (qqRsp *CloseOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQPost(ctx, bm, orderClose)
	if err != nil {
		return nil, err
	}
	qqRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 申请退款
// 注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 nil，否则，3证书Path均不可空
// 文档地址：https://qpay.qq.com/buss/wiki/38/1207
func (q *Client) Refund(ctx context.Context, bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath any) (qqRsp *RefundResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "refund_fee", "op_user_id", "op_user_passwd")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("transaction_id") == gopay.NULL {
		return nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	bs, err := q.doQQPostTLS(ctx, bm, refund)
	if err != nil {
		return nil, err
	}
	qqRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 退款查询
// 文档地址：https://qpay.qq.com/buss/wiki/38/1208
func (q *Client) RefundQuery(ctx context.Context, bm gopay.BodyMap) (qqRsp *RefundQueryResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.GetString("refund_id") == gopay.NULL && bm.GetString("out_refund_no") == gopay.NULL && bm.GetString("transaction_id") == gopay.NULL && bm.GetString("out_trade_no") == gopay.NULL {
		return nil, errors.New("refund_id, out_refund_no, out_trade_no, transaction_id are not allowed to be null at the same time")
	}
	bs, err := q.doQQPost(ctx, bm, refundQuery)
	if err != nil {
		return nil, err
	}
	qqRsp = new(RefundQueryResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// 交易账单
// 文档地址：https://qpay.qq.com/buss/wiki/38/1209
func (q *Client) StatementDown(ctx context.Context, bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "bill_type")
	if err != nil {
		return gopay.NULL, err
	}
	billType := bm.GetString("bill_type")
	if billType != "ALL" && billType != "SUCCESS" && billType != "REFUND" && billType != "RECHAR" {
		return gopay.NULL, errors.New("bill_type error, please reference: https://qpay.qq.com/buss/wiki/38/1209")
	}
	bs, err := q.doQQPost(ctx, bm, statementDown)
	if err != nil {
		return gopay.NULL, err
	}
	return string(bs), nil
}

// 资金账单
// 文档地址：https://qpay.qq.com/buss/wiki/38/3089
func (q *Client) AccRoll(ctx context.Context, bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "acc_type")
	if err != nil {
		return gopay.NULL, err
	}
	accType := bm.GetString("acc_type")
	if accType != "CASH" && accType != "MARKETING" {
		return gopay.NULL, errors.New("acc_type error, please reference: https://qpay.qq.com/buss/wiki/38/3089")
	}
	bs, err := q.doQQPost(ctx, bm, accRoll)
	if err != nil {
		return gopay.NULL, err
	}
	return string(bs), nil
}

// 向QQ发送请求
func (q *Client) doQQPost(ctx context.Context, bm gopay.BodyMap, url string) (bs []byte, err error) {
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", q.MchId)
	}
	if bm.GetString("fee_type") == gopay.NULL {
		bm.Set("fee_type", "CNY")
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := q.getReleaseSign(q.ApiKey, bm.GetString("sign_type"), bm)
		bm.Set("sign", sign)
	}
	req := GenerateXml(bm)
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Request: %s", req)
	}
	res, bs, err := q.hc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

// 向QQ发送请求 TLS
func (q *Client) doQQPostTLS(ctx context.Context, bm gopay.BodyMap, url string) (bs []byte, err error) {
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", q.MchId)
	}
	if bm.GetString("fee_type") == gopay.NULL {
		bm.Set("fee_type", "CNY")
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := q.getReleaseSign(q.ApiKey, bm.GetString("sign_type"), bm)
		bm.Set("sign", sign)
	}
	req := GenerateXml(bm)
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Request: %s", req)
	}
	res, bs, err := q.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Response: %d, %s", res.StatusCode, string(bs))
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
func (q *Client) doQQGet(ctx context.Context, bm gopay.BodyMap, url, signType string) (bs []byte, err error) {
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", q.MchId)
	}
	bm.Remove("sign")
	sign := q.getReleaseSign(q.ApiKey, signType, bm)
	bm.Set("sign", sign)
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Request: %s", bm.JsonBody())
	}
	param := bm.EncodeURLParams()
	uri := url + "?" + param
	res, bs, err := q.hc.Req(xhttp.TypeXML).Get(uri).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") || strings.Contains(string(bs), "html") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

func (q *Client) doQQRed(ctx context.Context, bm gopay.BodyMap, url string) (bs []byte, err error) {
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", q.MchId)
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := GetReleaseSign(q.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}
	req := GenerateXml(bm)
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Request: %s", req)
	}
	res, bs, err := q.tlsHc.Req(xhttp.TypeXML).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if q.DebugSwitch == gopay.DebugOn {
		q.logger.Debugf("QQ_Response: %d, %s", res.StatusCode, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
