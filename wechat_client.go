package gopay

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type WeChatClient struct {
	AppId      string
	MchId      string
	ApiKey     string
	BaseURL    string
	CertFile   []byte
	KeyFile    []byte
	Pkcs12File []byte
	IsProd     bool
}

// 初始化微信客户端 ok
//    appId：应用ID
//    mchId：商户ID
//    ApiKey：API秘钥值
//    IsProd：是否是正式环境
func NewWeChatClient(appId, mchId, apiKey string, isProd bool) (client *WeChatClient) {
	return &WeChatClient{
		AppId:  appId,
		MchId:  mchId,
		ApiKey: apiKey,
		IsProd: isProd}
}

// 提交付款码支付 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
func (w *WeChatClient) Micropay(body BodyMap) (wxRsp *WeChatMicropayResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxMicropay)
	} else {
		bs, err = w.doWeChat(body, wxSandboxMicropay)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatMicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 统一下单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (w *WeChatClient) UnifiedOrder(body BodyMap) (wxRsp *WeChatUnifiedOrderResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxUnifiedorder)
	} else {
		body.Set("total_fee", 101)
		bs, err = w.doWeChat(body, wxSandboxUnifiedorder)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatUnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 查询订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (w *WeChatClient) QueryOrder(body BodyMap) (wxRsp *WeChatQueryOrderResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxOrderquery)
	} else {
		bs, err = w.doWeChat(body, wxSandboxOrderquery)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatQueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 关闭订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (w *WeChatClient) CloseOrder(body BodyMap) (wxRsp *WeChatCloseOrderResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxCloseorder)
	} else {
		bs, err = w.doWeChat(body, wxSandboxCloseorder)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatCloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 撤销订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (w *WeChatClient) Reverse(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatReverseResponse, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.IsProd {
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%s", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wxReverse, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wxSandboxReverse)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 申请退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (w *WeChatClient) Refund(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatRefundResponse, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.IsProd {
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%s", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wxRefund, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wxSandboxRefund)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 查询退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (w *WeChatClient) QueryRefund(body BodyMap) (wxRsp *WeChatQueryRefundResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxRefundquery)
	} else {
		bs, err = w.doWeChat(body, wxSandboxRefundquery)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatQueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 下载对账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (w *WeChatClient) DownloadBill(body BodyMap) (wxRsp string, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxDownloadbill)
	} else {
		bs, err = w.doWeChat(body, wxSandboxDownloadbill)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

// 下载资金账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (w *WeChatClient) DownloadFundFlow(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.IsProd {
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return null, fmt.Errorf("ioutil.ReadFile：%s", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return null, fmt.Errorf("tls.LoadX509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wxDownloadfundflow, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wxSandboxDownloadfundflow)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

// 拉取订单评价数据
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (w *WeChatClient) BatchQueryComment(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.IsProd {
		body.Set("sign_type", SignType_HMAC_SHA256)
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return null, fmt.Errorf("ioutil.ReadFile：%s", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return null, fmt.Errorf("tls.LoadX509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wxBatchquerycomment, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wxSandboxBatchquerycomment)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

// 企业向微信用户个人付款
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_1
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
func (w *WeChatClient) Transfer(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatTransfersResponse, err error) {
	body.Set("mch_appid", w.AppId)
	body.Set("mchid", w.MchId)
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
		agent       *gorequest.SuperAgent
		errs        []error
		res         gorequest.Response
	)
	pkcsPool = x509.NewCertPool()
	if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile：%s", err.Error())
	}
	pkcsPool.AppendCertsFromPEM(pkcs)
	if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
		return nil, fmt.Errorf("tls.LoadX509KeyPair：%s", err.Error())
	}
	tlsConfig = &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            pkcsPool,
		InsecureSkipVerify: true}
	body.Set("sign", getWeChatReleaseSign(w.ApiKey, SignType_MD5, body))
	agent = HttpAgent().TLSClientConfig(tlsConfig)
	if w.BaseURL != null {
		agent.Post(w.BaseURL + wxTransfers)
	} else {
		agent.Post(wxBaseUrlCh + wxTransfers)
	}
	if res, bs, errs = agent.Type("xml").SendString(generateXml(body)).EndBytes(); len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	wxRsp = new(WeChatTransfersResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 公众号纯签约（未完成）
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_1&index=1
func (w *WeChatClient) EntrustPublic(body BodyMap) (bs []byte, err error) {
	bs, err = w.doWeChat(body, wxEntrustPublic)

	return nil, nil
}

// 向微信发送请求 ok
func (w *WeChatClient) doWeChat(body BodyMap, path string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	body.Set("appid", w.AppId)
	body.Set("mch_id", w.MchId)
	var (
		sign string
		errs []error
		res  gorequest.Response
	)
	if body.Get("sign") != null {
		goto GoRequest
	}
	if !w.IsProd {
		body.Set("sign_type", SignType_MD5)
		if sign, err = getWeChatSignBoxSign(w.MchId, w.ApiKey, body); err != nil {
			return
		}
	} else {
		sign = getWeChatReleaseSign(w.ApiKey, body.Get("sign_type"), body)
	}
	body.Set("sign", sign)
GoRequest:
	agent := HttpAgent()
	if w.IsProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}
	if w.BaseURL != null {
		agent.Post(w.BaseURL + path)
	} else {
		agent.Post(wxBaseUrlCh + path)
	}
	if res, bytes, errs = agent.Type("xml").SendString(generateXml(body)).EndBytes(); len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bytes), "HTML") {
		return nil, errors.New(string(bytes))
	}
	return
}
