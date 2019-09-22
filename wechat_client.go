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
	AppId   string
	MchId   string
	apiKey  string
	baseURL string
	isProd  bool
}

//初始化微信客户端 ok
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
func NewWeChatClient(appId, mchId, apiKey string, isProd bool) (client *WeChatClient) {
	return &WeChatClient{
		AppId:  appId,
		MchId:  mchId,
		apiKey: apiKey,
		isProd: isProd}
}

//提交付款码支付 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
func (w *WeChatClient) Micropay(body BodyMap) (wxRsp *WeChatMicropayResponse, err error) {
	var bs []byte
	if w.isProd {
		bs, err = w.doWeChat(body, wx_Micropay)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_Micropay)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatMicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//统一下单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (w *WeChatClient) UnifiedOrder(body BodyMap) (wxRsp *WeChatUnifiedOrderResponse, err error) {
	var bs []byte
	if w.isProd {
		bs, err = w.doWeChat(body, wx_UnifiedOrder)
	} else {
		body.Set("total_fee", 101)
		bs, err = w.doWeChat(body, wx_SanBox_UnifiedOrder)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatUnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//查询订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (w *WeChatClient) QueryOrder(body BodyMap) (wxRsp *WeChatQueryOrderResponse, err error) {
	var bs []byte
	if w.isProd {
		bs, err = w.doWeChat(body, wx_OrderQuery)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_OrderQuery)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatQueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//关闭订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (w *WeChatClient) CloseOrder(body BodyMap) (wxRsp *WeChatCloseOrderResponse, err error) {
	var bs []byte
	if w.isProd {
		bs, err = w.doWeChat(body, wx_CloseOrder)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_CloseOrder)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatCloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//撤销订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (w *WeChatClient) Reverse(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatReverseResponse, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.isProd {
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wx_Reverse, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_Reverse)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//申请退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (w *WeChatClient) Refund(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatRefundResponse, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.isProd {
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wx_Refund, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_Refund)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//查询退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (w *WeChatClient) QueryRefund(body BodyMap) (wxRsp *WeChatQueryRefundResponse, err error) {
	var bs []byte
	if w.isProd {
		bs, err = w.doWeChat(body, wx_RefundQuery)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_RefundQuery)
	}
	if err != nil {
		return
	}
	wxRsp = new(WeChatQueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//下载对账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (w *WeChatClient) DownloadBill(body BodyMap) (wxRsp string, err error) {
	var bs []byte
	if w.isProd {
		bs, err = w.doWeChat(body, wx_DownloadBill)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_DownloadBill)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

//下载资金账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (w *WeChatClient) DownloadFundFlow(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.isProd {
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return null, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return null, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wx_DownloadFundFlow, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_DownloadFundFlow)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

//拉取订单评价数据
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (w *WeChatClient) BatchQueryComment(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs, pkcs    []byte
		pkcsPool    *x509.CertPool
		certificate tls.Certificate
		tlsConfig   *tls.Config
	)
	if w.isProd {
		body.Set("sign_type", SignType_HMAC_SHA256)
		pkcsPool = x509.NewCertPool()
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return null, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return null, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		bs, err = w.doWeChat(body, wx_BatchQueryComment, tlsConfig)
	} else {
		bs, err = w.doWeChat(body, wx_SanBox_BatchQueryComment)
	}
	if err != nil {
		return
	}
	wxRsp = string(bs)
	return
}

//企业向微信用户个人付款
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
		return nil, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
	}
	pkcsPool.AppendCertsFromPEM(pkcs)
	if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
		return nil, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
	}
	tlsConfig = &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            pkcsPool,
		InsecureSkipVerify: true}
	body.Set("sign", getWeChatReleaseSign(w.apiKey, SignType_MD5, body))
	agent = HttpAgent().TLSClientConfig(tlsConfig)
	if w.baseURL != null {
		agent.Post(w.baseURL + wx_Transfers)
	} else {
		agent.Post(wx_base_url_ch + wx_Transfers)
	}
	if res, bs, errs = agent.Type("xml").SendString(generateXml(body)).EndBytes(); len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %v", res.StatusCode)
	}
	wxRsp = new(WeChatTransfersResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return
}

//向微信发送请求 ok
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
	if !w.isProd {
		body.Set("sign_type", SignType_MD5)
		if sign, err = getWeChatSignBoxSign(w.MchId, w.apiKey, body); err != nil {
			return
		}
	} else {
		sign = getWeChatReleaseSign(w.apiKey, body.Get("sign_type"), body)
	}
	body.Set("sign", sign)
GoRequest:
	agent := HttpAgent()
	if w.isProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}
	if w.baseURL != null {
		agent.Post(w.baseURL + path)
	} else {
		agent.Post(wx_base_url_ch + path)
	}
	if res, bytes, errs = agent.Type("xml").SendString(generateXml(body)).EndBytes(); len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %v", res.StatusCode)
	}
	if strings.Contains(string(bytes), "HTML") {
		return nil, errors.New(string(bytes))
	}
	return
}
