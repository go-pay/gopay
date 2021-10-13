package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cedarwu/gopay"
	"github.com/cedarwu/gopay/pkg/util"
)

// （服务商模式）APP下单API
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_2_1.shtml
func (c *ClientV3) V3PartnerTransactionApp(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerPayApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiPartnerPayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商模式）JSAPI/小程序下单API
//	Code = 0 is success
//	服务商JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_1.shtml
//	服务商小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_1.shtml
func (c *ClientV3) V3PartnerTransactionJsapi(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiPartnerJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商模式）Native下单API
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
func (c *ClientV3) V3PartnerTransactionNative(bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerNative, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiPartnerNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Native)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商模式）H5下单API
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func (c *ClientV3) V3PartnerTransactionH5(bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiPartnerH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(H5Url)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商模式）查询订单API
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_2.shtml
func (c *ClientV3) V3PartnerQueryOrder(orderNoType OrderNoType, orderNo string, bm gopay.BodyMap) (wxRsp *PartnerQueryOrderRsp, err error) {
	var uri string
	if bm.GetString("sp_mchid") == gopay.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	switch orderNoType {
	case TransactionId:
		uri = fmt.Sprintf(v3ApiPartnerQueryOrderTransactionId, orderNo) + "?" + bm.EncodeURLParams()
	case OutTradeNo:
		uri = fmt.Sprintf(v3ApiPartnerQueryOrderOutTradeNo, orderNo) + "?" + bm.EncodeURLParams()
	default:
		return nil, errors.New("unsupported order number type")
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &PartnerQueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerQueryOrder)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商模式）关单API
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_3.shtml
func (c *ClientV3) V3PartnerCloseOrder(subMchid, tradeNo string) (wxRsp *CloseOrderRsp, err error) {
	url := fmt.Sprintf(v3ApiPartnerCloseOrder, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("sp_mchid", c.Mchid)
	bm.Set("sub_mchid", subMchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CloseOrderRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
