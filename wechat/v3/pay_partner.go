package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// （服务商、电商模式）APP下单API
// Code = 0 is success
func (c *ClientV3) V3PartnerTransactionApp(ctx context.Context, bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerPayApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerPayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商、电商模式）JSAPI/小程序下单API
// Code = 0 is success
func (c *ClientV3) V3PartnerTransactionJsapi(ctx context.Context, bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商、电商模式）Native下单API
// Code = 0 is success
func (c *ClientV3) V3PartnerTransactionNative(ctx context.Context, bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerNative, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Native)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商模式）H5下单API
// Code = 0 is success
func (c *ClientV3) V3PartnerTransactionH5(ctx context.Context, bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(H5Url)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商、电商模式）查询订单API
// Code = 0 is success
func (c *ClientV3) V3PartnerQueryOrder(ctx context.Context, orderNoType OrderNoType, orderNo string, bm gopay.BodyMap) (wxRsp *PartnerQueryOrderRsp, err error) {
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
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &PartnerQueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerQueryOrder)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// （服务商、电商模式）关单API
// Code = 0 is success
func (c *ClientV3) V3PartnerCloseOrder(ctx context.Context, tradeNo string, bm gopay.BodyMap) (wxRsp *CloseOrderRsp, err error) {
	url := fmt.Sprintf(v3ApiPartnerCloseOrder, tradeNo)
	if bm.GetString("sp_mchid") == gopay.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
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
