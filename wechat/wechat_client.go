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
func (w *Client) Micropay(body gopay.BodyMap) (wxRsp *MicropayResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxMicropay)
	} else {
		bs, err = w.doWeChat(body, wxSandboxMicropay)
	}
	if err != nil {
		return
	}
	wxRsp = new(MicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 统一下单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (w *Client) UnifiedOrder(body gopay.BodyMap) (wxRsp *UnifiedOrderResponse, err error) {
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
	wxRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 查询订单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (w *Client) QueryOrder(body gopay.BodyMap) (wxRsp *QueryOrderResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxOrderquery)
	} else {
		bs, err = w.doWeChat(body, wxSandboxOrderquery)
	}
	if err != nil {
		return
	}
	wxRsp = new(QueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 关闭订单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (w *Client) CloseOrder(body gopay.BodyMap) (wxRsp *CloseOrderResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxCloseorder)
	} else {
		bs, err = w.doWeChat(body, wxSandboxCloseorder)
	}
	if err != nil {
		return
	}
	wxRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 撤销订单
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (w *Client) Reverse(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *ReverseResponse, err error) {
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
	wxRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 申请退款
//    注意：如已使用client.AddCertFilePath()或client.AddCertFileByte()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传空字符串 ""，如方法需单独使用证书，则传证书Path
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (w *Client) Refund(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *RefundResponse, err error) {
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
	wxRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 查询退款
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (w *Client) QueryRefund(body gopay.BodyMap) (wxRsp *QueryRefundResponse, err error) {
	var bs []byte
	if w.IsProd {
		bs, err = w.doWeChat(body, wxRefundquery)
	} else {
		bs, err = w.doWeChat(body, wxSandboxRefundquery)
	}
	if err != nil {
		return
	}
	wxRsp = new(QueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%s", err.Error())
	}
	return
}

// 下载对账单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (w *Client) DownloadBill(body gopay.BodyMap) (wxRsp string, err error) {
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
func (w *Client) DownloadFundFlow(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return gopay.NULL, err
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
func (w *Client) BatchQueryComment(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		body.Set("sign_type", SignType_HMAC_SHA256)
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
			return gopay.NULL, err
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
func (w *Client) Transfer(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *TransfersResponse, err error) {
	body.Set("mch_appid", w.AppId)
	body.Set("mchid", w.MchId)
	var (
		tlsConfig *tls.Config
		url       = wxBaseUrlCh + wxTransfers
	)
	if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	body.Set("sign", getReleaseSign(w.ApiKey, SignType_MD5, body))

	httpClient := gopay.NewHttpClient().SetTLSConfig(tlsConfig).Type(gopay.TypeXML)
	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + wxTransfers
		w.mu.RUnlock()
	}
	wxRsp = new(TransfersResponse)
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
func (w *Client) EntrustPublic(body gopay.BodyMap) (bs []byte, err error) {
	bs, err = w.doWeChat(body, wxEntrustPublic)

	return nil, nil
}

// 向微信发送请求
func (w *Client) doWeChat(body gopay.BodyMap, path string, tlsConfig ...*tls.Config) (bs []byte, err error) {
	var url = wxBaseUrlCh + path
	body.Set("appid", w.AppId)
	body.Set("mch_id", w.MchId)

	if body.Get("sign") == gopay.NULL {
		var sign string
		if !w.IsProd {
			body.Set("sign_type", SignType_MD5)
			sign, err = getSignBoxSign(w.MchId, w.ApiKey, body)
			if err != nil {
				return nil, err
			}
		} else {
			sign = getReleaseSign(w.ApiKey, body.Get("sign_type"), body)
		}
		body.Set("sign", sign)
	}

	httpClient := gopay.NewHttpClient()
	if w.IsProd && tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig[0])
	}

	if w.BaseURL != gopay.NULL {
		w.mu.RLock()
		url = w.BaseURL + path
		w.mu.RUnlock()
	}

	res, bs, errs := httpClient.Type(gopay.TypeXML).Post(url).SendString(generateXml(body)).EndBytes()
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
