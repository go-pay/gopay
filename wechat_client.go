package gopay

import (
	"encoding/xml"
	"github.com/parnurzeal/gorequest"
)

type weChatClient struct {
	AppId     string
	MchId     string
	secretKey string
	isProd    bool
}

//初始化微信客户端 ok
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
func NewWeChatClient(appId, mchId, secretKey string, isProd bool) *weChatClient {
	client := new(weChatClient)
	client.AppId = appId
	client.MchId = mchId
	client.secretKey = secretKey
	client.isProd = isProd
	return client
}

//提交付款码支付 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
func (this *weChatClient) Micropay(body BodyMap) (wxRsp *WeChatMicropayResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_Micropay)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_Micropay)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatMicropayResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//统一下单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (this *weChatClient) UnifiedOrder(body BodyMap) (wxRsp *WeChatUnifiedOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_UnifiedOrder)
		if err != nil {
			return nil, err
		}
	} else {
		body.Set("total_fee", 101)
		bytes, err = this.doWeChat(body, wxURL_SanBox_UnifiedOrder)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatUnifiedOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (this *weChatClient) QueryOrder(body BodyMap) (wxRsp *WeChatQueryOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_OrderQuery)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_OrderQuery)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatQueryOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//关闭订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (this *weChatClient) CloseOrder(body BodyMap) (wxRsp *WeChatCloseOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_CloseOrder)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_CloseOrder)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatCloseOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//撤销订单
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (this *weChatClient) Reverse(body BodyMap) (wxRsp *WeChatReverseResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_Reverse)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_Reverse)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatReverseResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//申请退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (this *weChatClient) Refund(body BodyMap) (wxRsp *WeChatRefundResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_Refund)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_Refund)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatRefundResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (this *weChatClient) QueryRefund(body BodyMap) (wxRsp *WeChatQueryRefundResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_RefundQuery)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_RefundQuery)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatQueryRefundResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//下载对账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (this *weChatClient) DownloadBill(body BodyMap) (wxRsp string, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_DownloadBill)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_DownloadBill)
	}
	wxRsp = string(bytes)
	if err != nil {
		return wxRsp, err
	}
	return wxRsp, nil
}

//下载资金账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (this *weChatClient) DownloadFundFlow(body BodyMap) (wxRsp string, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		bytes, err = this.doWeChat(body, wxURL_DownloadFundFlow)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_DownloadFundFlow)
	}
	wxRsp = string(bytes)
	if err != nil {
		return wxRsp, err
	}
	return wxRsp, nil
}

//拉取订单评价数据
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (this *weChatClient) BatchQueryComment(body BodyMap) (wxRsp string, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		body.Set("sign_type", SignType_HMAC_SHA256)
		bytes, err = this.doWeChat(body, wxURL_BatchQueryComment)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_BatchQueryComment)
	}

	wxRsp = string(bytes)
	if err != nil {
		return wxRsp, err
	}
	return wxRsp, nil
}

//向微信发送请求 ok
func (this *weChatClient) doWeChat(body BodyMap, url string) (bytes []byte, err error) {
	var sign string
	body.Set("appid", this.AppId)
	body.Set("mch_id", this.MchId)
	//===============生成参数===================
	if !this.isProd {
		//沙箱环境
		body.Set("sign_type", SignType_MD5)
		//从微信接口获取SanBoxSignKey
		key, err := getSanBoxSign(this.MchId, body.Get("nonce_str"), this.secretKey, body.Get("sign_type"))
		if err != nil {
			return nil, err
		}
		sign = getLocalSign(key, body.Get("sign_type"), body)
	} else {
		//正式环境
		//本地计算Sign
		sign = getLocalSign(this.secretKey, body.Get("sign_type"), body)
	}
	body.Set("sign", sign)
	reqXML := generateXml(body)
	//===============发起请求===================
	agent := gorequest.New()
	if this.isProd {
		agent.Post(url)
	} else {
		agent.Post(url)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	_, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bytes, nil
}
