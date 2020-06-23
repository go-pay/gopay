package qq

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/goutil"
	"github.com/iGoogle-ink/goutil/xhttp"
)

type Client struct {
	MchId       string
	ApiKey      string
	IsProd      bool
	certificate tls.Certificate
	certPool    *x509.CertPool
	mu          sync.RWMutex
}

// 初始化QQ客户端（正式环境）
//    mchId：商户ID
//    ApiKey：API秘钥值
func NewClient(mchId, apiKey string) (client *Client) {
	if mchId != goutil.NULL && apiKey != goutil.NULL {
		return &Client{
			MchId:  mchId,
			ApiKey: apiKey,
		}
	}
	return nil
}

// 提交付款码支付
//    文档地址：https://qpay.qq.com/buss/wiki/1/1122
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
//    文档地址：https://qpay.qq.com/buss/wiki/1/1125
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
//    文档地址：https://qpay.qq.com/buss/wiki/38/1203
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
//    文档地址：https://qpay.qq.com/buss/wiki/38/1205
func (q *Client) OrderQuery(bm gopay.BodyMap) (qqRsp *OrderQueryResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.Get("out_trade_no") == goutil.NULL && bm.Get("transaction_id") == goutil.NULL {
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
//    文档地址：https://qpay.qq.com/buss/wiki/38/1206
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
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    文档地址：https://qpay.qq.com/buss/wiki/38/1207
func (q *Client) Refund(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (qqRsp *RefundResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "refund_fee", "op_user_id", "op_user_passwd")
	if err != nil {
		return nil, err
	}
	if bm.Get("out_trade_no") == goutil.NULL && bm.Get("transaction_id") == goutil.NULL {
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
//    文档地址：https://qpay.qq.com/buss/wiki/38/1208
func (q *Client) RefundQuery(bm gopay.BodyMap) (qqRsp *RefundQueryResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.Get("refund_id") == goutil.NULL && bm.Get("out_refund_no") == goutil.NULL && bm.Get("transaction_id") == goutil.NULL && bm.Get("out_trade_no") == goutil.NULL {
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
//    文档地址：https://qpay.qq.com/buss/wiki/38/1209
func (q *Client) StatementDown(bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "bill_type")
	if err != nil {
		return goutil.NULL, err
	}
	billType := bm.Get("bill_type")
	if billType != "ALL" && billType != "SUCCESS" && billType != "REFUND" && billType != "RECHAR" {
		return goutil.NULL, errors.New("bill_type error, please reference: https://qpay.qq.com/buss/wiki/38/1209")
	}
	bs, err := q.doQQ(bm, statementDown, nil)
	if err != nil {
		return goutil.NULL, err
	}
	return string(bs), nil
}

// 资金账单
//    文档地址：https://qpay.qq.com/buss/wiki/38/3089
func (q *Client) AccRoll(bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "acc_type")
	if err != nil {
		return goutil.NULL, err
	}
	accType := bm.Get("acc_type")
	if accType != "CASH" && accType != "MARKETING" {
		return goutil.NULL, errors.New("acc_type error, please reference: https://qpay.qq.com/buss/wiki/38/3089")
	}
	bs, err := q.doQQ(bm, accRoll, nil)
	if err != nil {
		return goutil.NULL, err
	}
	return string(bs), nil
}

// 向QQ发送请求
func (q *Client) doQQ(bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {
	bm.Set("mch_id", q.MchId)
	if bm.Get("fee_type") == goutil.NULL {
		bm.Set("fee_type", "CNY")
	}

	if bm.Get("sign") == goutil.NULL {
		var sign string
		sign = getReleaseSign(q.ApiKey, bm.Get("sign_type"), bm)
		bm.Set("sign", sign)
	}

	httpClient := xhttp.NewHttpClient()
	if tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}

	res, bs, errs := httpClient.Type(xhttp.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
