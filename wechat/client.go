package wechat

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/iGoogle-ink/gopay"
)

type Client struct {
	AppId      string
	MchId      string
	ApiKey     string
	BaseURL    string
	CertFile   []byte
	KeyFile    []byte
	Pkcs12File []byte
	IsProd     bool
	mu         sync.RWMutex
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
		bs, err = w.doWeChat(bm, wxMicropay, nil)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxMicropay, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(MicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
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
		bs, err = w.doWeChat(bm, wxUnifiedorder, nil)
	} else {
		bm.Set("total_fee", 101)
		bs, err = w.doWeChat(bm, wxSandboxUnifiedorder, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
}

// 查询订单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (w *Client) QueryOrder(bm gopay.BodyMap) (wxRsp *QueryOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("transaction_id") == gopay.NULL {
		return nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(bm, wxOrderquery, nil)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxOrderquery, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(QueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
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
		bs, err = w.doWeChat(bm, wxCloseorder, nil)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxCloseorder, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
}

// 撤销订单
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (w *Client) Reverse(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *ReverseResponse, err error) {
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
		bs, err = w.doWeChat(bm, wxReverse, tlsConfig)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxReverse, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
}

// 申请退款
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (w *Client) Refund(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *RefundResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "total_fee", "refund_fee")
	if err != nil {
		return nil, err
	}
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("transaction_id") == gopay.NULL {
		return nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return nil, err
		}
		bs, err = w.doWeChat(bm, wxRefund, tlsConfig)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxRefund, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
}

// 查询退款
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (w *Client) QueryRefund(bm gopay.BodyMap) (wxRsp *QueryRefundResponse, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, err
	}
	if bm.Get("refund_id") == gopay.NULL && bm.Get("out_refund_no") == gopay.NULL && bm.Get("transaction_id") == gopay.NULL && bm.Get("out_trade_no") == gopay.NULL {
		return nil, errors.New("refund_id, out_refund_no, out_trade_no, transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(bm, wxRefundquery, nil)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxRefundquery, nil)
	}
	if err != nil {
		return
	}
	wxRsp = new(QueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return
}

// 下载对账单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (w *Client) DownloadBill(bm gopay.BodyMap) (wxRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "bill_type")
	if err != nil {
		return gopay.NULL, err
	}
	billType := bm.Get("bill_type")
	if billType != "ALL" && billType != "SUCCESS" && billType != "REFUND" && billType != "RECHARGE_REFUND" {
		return gopay.NULL, errors.New("bill_type error, please reference: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(bm, wxDownloadbill, nil)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxDownloadbill, nil)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

// 下载资金账单
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    貌似不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
func (w *Client) DownloadFundFlow(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "bill_date", "account_type")
	if err != nil {
		return gopay.NULL, err
	}
	accountType := bm.Get("account_type")
	if accountType != "Basic" && accountType != "Operation" && accountType != "Fees" {
		return gopay.NULL, errors.New("account_type error, please reference: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7")
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	bm.Set("sign_type", SignType_HMAC_SHA256)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return gopay.NULL, err
		}
		bs, err = w.doWeChat(bm, wxDownloadfundflow, tlsConfig)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxDownloadfundflow, nil)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

// 拉取订单评价数据
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    貌似不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
func (w *Client) BatchQueryComment(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	err = bm.CheckEmptyError("nonce_str", "begin_time", "end_time", "offset")
	if err != nil {
		return gopay.NULL, err
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	bm.Set("sign_type", SignType_HMAC_SHA256)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return gopay.NULL, err
		}
		bs, err = w.doWeChat(bm, wxBatchquerycomment, tlsConfig)
	} else {
		bs, err = w.doWeChat(bm, wxSandboxBatchquerycomment, nil)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

// 企业向微信用户个人付款
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，否则，3证书Path均不可空
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (w *Client) Transfer(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *TransfersResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "partner_trade_no", "openid", "check_name", "amount", "desc", "spbill_create_ip")
	if err != nil {
		return nil, err
	}
	bm.Set("mch_appid", w.AppId)
	bm.Set("mchid", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = wxBaseUrlCh + wxTransfers
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	bm.Set("sign", getReleaseSign(w.ApiKey, SignType_MD5, bm))

	httpClient := gopay.NewHttpClient().SetTLSConfig(tlsConfig).Type(gopay.TypeXML)
	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + wxTransfers
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

// 公众号纯签约（未完成）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_1&index=1
func (w *Client) EntrustPublic(bm gopay.BodyMap) (bs []byte, err error) {
	bs, err = w.doWeChat(bm, wxEntrustPublic, nil)

	return nil, nil
}

// 向微信发送请求
func (w *Client) doWeChat(bm gopay.BodyMap, path string, tlsConfig *tls.Config) (bs []byte, err error) {
	var url = wxBaseUrlCh + path
	bm.Set("appid", w.AppId)
	bm.Set("mch_id", w.MchId)

	if bm.Get("sign") == gopay.NULL {
		var sign string
		if !w.IsProd {
			bm.Set("sign_type", SignType_MD5)
			sign, err = getSignBoxSign(w.MchId, w.ApiKey, bm)
			if err != nil {
				return nil, err
			}
		} else {
			sign = getReleaseSign(w.ApiKey, bm.Get("sign_type"), bm)
		}
		bm.Set("sign", sign)
	}

	httpClient := gopay.NewHttpClient()
	if w.IsProd && tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}

	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + path
		w.mu.RUnlock()
	}

	res, bs, errs := httpClient.Type(gopay.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
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
