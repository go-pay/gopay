package gopay

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"
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
	mu         sync.RWMutex
}

// 初始化微信客户端
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

// 提交付款码支付
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

// 统一下单
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

// 查询订单
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

// 关闭订单
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

// 撤销订单
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (w *WeChatClient) Reverse(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatReverseResponse, err error) {
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return nil, err
		}
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

// 申请退款
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (w *WeChatClient) Refund(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatRefundResponse, err error) {
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return nil, err
		}
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

// 查询退款
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

// 下载对账单
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

// 下载资金账单
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    貌似不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
func (w *WeChatClient) DownloadFundFlow(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return null, err
		}
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
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    貌似不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
func (w *WeChatClient) BatchQueryComment(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		body.Set("sign_type", SignType_HMAC_SHA256)
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return null, err
		}
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
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_1
func (w *WeChatClient) Transfer(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatTransfersResponse, err error) {
	body.Set("mch_appid", w.AppId)
	body.Set("mchid", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = wxBaseUrlCh + wxTransfers
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	body.Set("sign", getWeChatReleaseSign(w.ApiKey, SignType_MD5, body))

	httpClient := NewHttpClient().SetTLSConfig(tlsConfig).Type(TypeXML)
	if w.BaseURL != null {
		w.mu.RLock()
		url = w.BaseURL + wxTransfers
		w.mu.RUnlock()
	}
	wxRsp = new(WeChatTransfersResponse)
	res, errs := httpClient.Post(url).SendString(generateXml(body)).EndStruct(wxRsp)
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
func (w *WeChatClient) EntrustPublic(body BodyMap) (bs []byte, err error) {
	bs, err = w.doWeChat(body, wxEntrustPublic)

	return nil, nil
}

// 向微信发送请求
func (w *WeChatClient) doWeChat(body BodyMap, path string, tlsConfig ...*tls.Config) (bs []byte, err error) {
	body.Set("appid", w.AppId)
	body.Set("mch_id", w.MchId)
	var (
		url = wxBaseUrlCh + path
	)
	if body.Get("sign") == null {
		var sign string
		if !w.IsProd {
			body.Set("sign_type", SignType_MD5)
			sign, err = getWeChatSignBoxSign(w.MchId, w.ApiKey, body)
			if err != nil {
				return nil, err
			}
		} else {
			sign = getWeChatReleaseSign(w.ApiKey, body.Get("sign_type"), body)
		}
		body.Set("sign", sign)
	}

	httpClient := NewHttpClient()
	if w.IsProd && tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig[0])
	}

	if w.BaseURL != null {
		w.mu.RLock()
		url = w.BaseURL + path
		w.mu.RUnlock()
	}

	res, bs, errs := httpClient.Type(TypeXML).Post(url).SendString(generateXml(body)).EndBytes()
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
