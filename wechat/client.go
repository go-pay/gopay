package wechat

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xhttp"
)

type Client struct {
	AppId       string
	MchId       string
	ApiKey      string
	BaseURL     string
	IsProd      bool
	certificate tls.Certificate
	certPool    *x509.CertPool
	mu          sync.RWMutex
}

// 初始化微信客户端
//    appId：应用ID
//    mchId：商户ID
//    ApiKey：API秘钥值
//    IsProd：是否是正式环境
func NewClient(appId, mchId, apiKey string, isProd bool) (client *Client) {
	return &Client{
		AppId:  appId,
		MchId:  mchId,
		ApiKey: apiKey,
		IsProd: isProd}
}

// 提交付款码支付
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
func (w *Client) Micropay(bm gopay.BodyMap) (wxRsp *MicropayResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "auth_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, microPay, nil)
	} else {
		bm.Set("total_fee", 1)
		bs, err = w.doSanBoxPost(bm, sandboxMicroPay)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(MicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 授权码查询openid（正式）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9
func (w *Client) AuthCodeToOpenId(bm gopay.BodyMap) (wxRsp *AuthCodeToOpenIdResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "auth_code")
	if err != nil {
		return nil, err
	}

	bs, err := w.doProdPost(bm, authCodeToOpenid, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(AuthCodeToOpenIdResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 统一下单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (w *Client) UnifiedOrder(bm gopay.BodyMap) (wxRsp *UnifiedOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, unifiedOrder, nil)
	} else {
		bm.Set("total_fee", 101)
		bs, err = w.doSanBoxPost(bm, sandboxUnifiedOrder)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 查询订单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (w *Client) QueryOrder(bm gopay.BodyMap) (wxRsp *QueryOrderResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, nil, err
	}
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("transaction_id") == gotil.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, orderQuery, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxOrderQuery)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(QueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s)：%w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s)：%w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 关闭订单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (w *Client) CloseOrder(bm gopay.BodyMap) (wxRsp *CloseOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, closeOrder, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxCloseOrder)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 撤销订单
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (w *Client) Reverse(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ReverseResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return nil, err
		}
		bs, err = w.doProdPost(bm, reverse, tlsConfig)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxReverse)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 申请退款
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (w *Client) Refund(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *RefundResponse, resBm gopay.BodyMap, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "total_fee", "refund_fee")
	if err != nil {
		return nil, nil, err
	}
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("transaction_id") == gotil.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return nil, nil, err
		}
		bs, err = w.doProdPost(bm, refund, tlsConfig)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxRefund)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s)：%w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s)：%w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 查询退款
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (w *Client) QueryRefund(bm gopay.BodyMap) (wxRsp *QueryRefundResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, nil, err
	}
	if bm.Get("refund_id") == gotil.NULL && bm.Get("out_refund_no") == gotil.NULL && bm.Get("transaction_id") == gotil.NULL && bm.Get("out_trade_no") == gotil.NULL {
		return nil, nil, errors.New("refund_id, out_refund_no, out_trade_no, transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, refundQuery, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxRefundQuery)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(QueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s)：%w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s)：%w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 下载对账单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (w *Client) DownloadBill(bm gopay.BodyMap) (wxRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "bill_type")
	if err != nil {
		return gotil.NULL, err
	}
	billType := bm.Get("bill_type")
	if billType != "ALL" && billType != "SUCCESS" && billType != "REFUND" && billType != "RECHARGE_REFUND" {
		return gotil.NULL, errors.New("bill_type error, please reference: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, downloadBill, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxDownloadBill)
	}
	if err != nil {
		return gotil.NULL, err
	}
	return string(bs), nil
}

// 下载资金账单（正式）
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    貌似不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
func (w *Client) DownloadFundFlow(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp string, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return gotil.NULL, err
	}
	err = bm.CheckEmptyError("nonce_str", "bill_date", "account_type")
	if err != nil {
		return gotil.NULL, err
	}
	accountType := bm.Get("account_type")
	if accountType != "Basic" && accountType != "Operation" && accountType != "Fees" {
		return gotil.NULL, errors.New("account_type error, please reference: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7")
	}
	bm.Set("sign_type", SignType_HMAC_SHA256)
	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return gotil.NULL, err
	}
	bs, err := w.doProdPost(bm, downloadFundFlow, tlsConfig)
	if err != nil {
		return gotil.NULL, err
	}
	wxRsp = string(bs)
	return
}

// 交易保障
//    文档地址：（JSAPI）https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_8&index=9
//    文档地址：（付款码）https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_14&index=8
//    文档地址：（Native）https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_8&index=9
//    文档地址：（APP）https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_8&index=10
//    文档地址：（H5）https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_8&index=9
//    文档地址：（微信小程序）https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_8&index=9
func (w *Client) Report(bm gopay.BodyMap) (wxRsp *ReportResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "interface_url", "execute_time", "return_code", "return_msg", "result_code", "user_ip")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, report, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxReport)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(ReportResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 拉取订单评价数据（正式）
//    注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    貌似不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
func (w *Client) BatchQueryComment(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp string, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return gotil.NULL, err
	}
	err = bm.CheckEmptyError("nonce_str", "begin_time", "end_time", "offset")
	if err != nil {
		return gotil.NULL, err
	}
	bm.Set("sign_type", SignType_HMAC_SHA256)
	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return gotil.NULL, err
	}
	bs, err := w.doProdPost(bm, batchQueryComment, tlsConfig)
	if err != nil {
		return gotil.NULL, err
	}
	return string(bs), nil
}

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
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
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

// 公众号纯签约（正式）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_1&index=1
func (w *Client) EntrustPublic(bm gopay.BodyMap) (wxRsp *EntrustPublicResponse, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdGet(bm, entrustPublic, SignType_MD5)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustPublicResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// APP纯签约-预签约接口-获取预签约ID（正式）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_5&index=2
func (w *Client) EntrustAppPre(bm gopay.BodyMap) (wxRsp *EntrustAppPreResponse, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, entrustApp, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustAppPreResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// H5纯签约（正式）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_16&index=4
func (w *Client) EntrustH5(bm gopay.BodyMap) (wxRsp *EntrustH5Response, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp", "clientip")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdGet(bm, entrustH5, SignType_HMAC_SHA256)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustH5Response)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 支付中签约（正式）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_13&index=5
func (w *Client) EntrustPaying(bm gopay.BodyMap) (wxRsp *EntrustPayingResponse, err error) {
	err = bm.CheckEmptyError("contract_mchid", "contract_appid",
		"out_trade_no", "nonce_str", "body", "notify_url", "total_fee",
		"spbill_create_ip", "trade_type", "plan_id", "contract_code",
		"request_serial", "contract_display_account", "contract_notify_url")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, entrustPaying, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustPayingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// ProfitSharing 请求单次分账
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
// 单次分账请求按照传入的分账接收方账号和资金进行分账，
// 同时会将订单剩余的待分账金额解冻给本商户。
// 故操作成功后，订单不能再进行分账，也不能进行分账完结。
func (w *Client) ProfitSharing(bm gopay.BodyMap) (wxRsp *ProfitSharingResponse, err error) {
	return w.profitSharing(bm, profitSharing)
}

// MultiProfitSharing 请求多次分账
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
// 微信订单支付成功后，商户发起分账请求，将结算后的钱分到分账接收方。多次分账请求仅会按照传入的分账接收方进行分账，不会对剩余的金额进行任何操作。
// 故操作成功后，在待分账金额不等于零时，订单依旧能够再次进行分账。
// 多次分账，可以将本商户作为分账接收方直接传入，实现释放资金给本商户的功能
// 对同一笔订单最多能发起20次多次分账请求
func (w *Client) MultiProfitSharing(bm gopay.BodyMap) (wxRsp *ProfitSharingResponse, err error) {
	return w.profitSharing(bm, multiProfitSharing)
}

func (w *Client) profitSharing(bm gopay.BodyMap, uri string) (wxRsp *ProfitSharingResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "appid", "nonce_str", "transaction_id", "out_order_no", "receivers")
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
	//检查每个分账接收者的必传属性
	for _, r := range arr {
		err = r.CheckEmptyError("type", "account", "amount", "description")
		if err != nil {
			return nil, err
		}
	}
	bs, err := w.doProdPost(bm, uri, nil)
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
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3
// 发起分账请求后，可调用此接口查询分账结果；发起分账完结请求后，可调用此接口查询分账完结的执行结果。
func (w *Client) ProfitSharingQuery(bm gopay.BodyMap) (wxRsp *ProfitSharingQueryResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "transaction_id", "out_order_no", "nonce_str")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, profitSharingQuery, nil)
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
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4
// 商户发起添加分账接收方请求，后续可通过发起分账请求将结算后的钱分到该分账接收方。
func (w *Client) ProfitSharingAddReceiver(bm gopay.BodyMap) (wxRsp *ProfitSharingAddReceiverResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "appid", "nonce_str", "receiver")
	if err != nil {
		return nil, err
	}
	//输入参数 接收方
	r, err := bm.GetBodyMap("receiver")
	if err != nil {
		return nil, err
	}
	err = r.CheckEmptyError("type", "account", "relation_type")
	if err != nil {
		return nil, err
	}
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
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5
// 商户发起删除分账接收方请求，删除后不支持将结算后的钱分到该分账接收方
func (w *Client) ProfitSharingRemoveReceiver(bm gopay.BodyMap) (wxRsp *ProfitSharingAddReceiverResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "appid", "nonce_str", "receiver")
	if err != nil {
		return nil, err
	}
	//输入参数 接收方
	r, err := bm.GetBodyMap("receiver")
	if err != nil {
		return nil, err
	}
	err = r.CheckEmptyError("type", "account")
	if err != nil {
		return nil, err
	}
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
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6
// 1、不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给本商户
// 2、调用多次分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
// 3、已调用请求单次分账后，剩余待分账金额为零，不需要再调用此接口。
func (w *Client) ProfitSharingFinish(bm gopay.BodyMap) (wxRsp *ProfitSharingResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "appid", "nonce_str", "transaction_id", "out_order_no", "description")
	if err != nil {
		return nil, err
	}

	bs, err := w.doProdPost(bm, profitSharingFinish, nil)
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
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7
// 对订单进行退款时，如果订单已经分账，可以先调用此接口将指定的金额从分账接收方（仅限商户类型的分账接收方）回退给本商户，然后再退款。
// 回退以原分账请求为依据，可以对分给分账接收方的金额进行多次回退，只要满足累计回退不超过该请求中分给接收方的金额。
// 此接口采用同步处理模式，即在接收到商户请求后，会实时返回处理结果
// 此功能需要接收方在商户平台-交易中心-分账-分账接收设置下，开启同意分账回退后，才能使用。
func (w *Client) ProfitSharingReturn(bm gopay.BodyMap) (wxRsp *ProfitSharingReturnResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "appid", "nonce_str", "out_return_no", "return_account_type", "return_account", "return_amount", "description")
	if err != nil {
		return nil, err
	}

	if (bm.Get("order_id") == gotil.NULL) && (bm.Get("out_order_no") == gotil.NULL) {
		return nil, errors.New("param order_id and out_order_no can not be null at the same time")
	}

	bs, err := w.doProdPost(bm, profitSharingReturn, nil)
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
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8
// 商户需要核实回退结果，可调用此接口查询回退结果。
// 如果分账回退接口返回状态为处理中，可调用此接口查询回退结果
func (w *Client) ProfitSharingReturnQuery(bm gopay.BodyMap) (wxRsp *ProfitSharingReturnResponse, err error) {
	err = bm.CheckEmptyError("mch_id", "appid", "nonce_str", "out_return_no")
	if err != nil {
		return nil, err
	}

	if (bm.Get("order_id") == gotil.NULL) && (bm.Get("out_order_no") == gotil.NULL) {
		return nil, errors.New("param order_id and out_order_no can not be null at the same time")
	}

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

// doSanBoxPost sanbox环境post请求
func (w *Client) doSanBoxPost(bm gopay.BodyMap, path string) (bs []byte, err error) {
	var url = baseUrlCh + path
	w.mu.RLock()
	defer w.mu.RUnlock()
	bm.Set("appid", w.AppId)
	bm.Set("mch_id", w.MchId)

	if bm.Get("sign") == gotil.NULL {
		bm.Set("sign_type", SignType_MD5)
		sign, err := getSignBoxSign(w.MchId, w.ApiKey, bm)
		if err != nil {
			return nil, err
		}
		bm.Set("sign", sign)
	}

	if w.BaseURL != gotil.NULL {
		url = w.BaseURL + path
	}
	res, bs, errs := xhttp.NewClient().Type(xhttp.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") || strings.Contains(string(bs), "html") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

// Post请求、正式
func (w *Client) doProdPost(bm gopay.BodyMap, path string, tlsConfig *tls.Config) (bs []byte, err error) {
	var url = baseUrlCh + path
	w.mu.RLock()
	defer w.mu.RUnlock()
	bm.Set("appid", w.AppId)
	bm.Set("mch_id", w.MchId)

	if bm.Get("sign") == gotil.NULL {
		sign := getReleaseSign(w.ApiKey, bm.Get("sign_type"), bm)
		bm.Set("sign", sign)
	}

	httpClient := xhttp.NewClient()
	if w.IsProd && tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}
	if w.BaseURL != gotil.NULL {
		url = w.BaseURL + path
	}
	res, bs, errs := httpClient.Type(xhttp.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") || strings.Contains(string(bs), "html") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}

// Get请求、正式
func (w *Client) doProdGet(bm gopay.BodyMap, path, signType string) (bs []byte, err error) {
	var url = baseUrlCh + path
	w.mu.RLock()
	defer w.mu.RUnlock()
	bm.Set("appid", w.AppId)
	bm.Set("mch_id", w.MchId)
	bm.Remove("sign")
	sign := getReleaseSign(w.ApiKey, signType, bm)
	bm.Set("sign", sign)

	if w.BaseURL != gotil.NULL {
		w.mu.RLock()
		url = w.BaseURL + path
		w.mu.RUnlock()
	}
	param := bm.EncodeGetParams()
	url = url + "?" + param
	res, bs, errs := xhttp.NewClient().Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") || strings.Contains(string(bs), "html") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
