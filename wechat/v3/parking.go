package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 查询车牌服务开通信息
// Code = 0 is success
// bm：query 参数
func (c *ClientV3) V3VehicleParkingQuery(ctx context.Context, bm gopay.BodyMap) (*VehicleParkingQueryRsp, error) {
	uri := v3VehicleParkingQuery + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &VehicleParkingQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(VehicleParkingQuery)
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

// 创建停车入场
// Code = 0 is success
// bm：body 参数
func (c *ClientV3) V3VehicleParkingIn(ctx context.Context, bm gopay.BodyMap) (wxRsp *VehicleParkingInRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3VehicleParkingIn, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3VehicleParkingIn, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &VehicleParkingInRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(VehicleParkingIn)
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

// 扣费受理
// Code = 0 is success
// bm：body 参数
func (c *ClientV3) V3VehicleParkingFee(ctx context.Context, bm gopay.BodyMap) (wxRsp *VehicleParkingFeeRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3VehicleParkingFee, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3VehicleParkingFee, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &VehicleParkingFeeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(VehicleParkingFee)
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

// 查询订单
// Code = 0 is success
// bm：query 参数
func (c *ClientV3) V3VehicleParkingOrder(ctx context.Context, outTradeNo string, bm gopay.BodyMap) (*VehicleParkingOrderRsp, error) {
	uri := fmt.Sprintf(v3VehicleParkingOrder, outTradeNo) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &VehicleParkingOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(VehicleParkingOrder)
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
