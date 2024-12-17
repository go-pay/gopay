package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-pay/gopay"
)

// 创单结单合并API
// Code = 0 is success
// 注意：限制条件：【免确认订单模式】，用户已授权状态下，可调用该接口。
func (c *ClientV3) V3ScoreDirectComplete(ctx context.Context, bm gopay.BodyMap) (wxRsp *ScoreDirectCompleteRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScoreDirectComplete, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScoreDirectComplete, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreDirectCompleteRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreDirectComplete)
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

// 商户预授权API
// Code = 0 is success
func (c *ClientV3) V3ScorePermission(ctx context.Context, bm gopay.BodyMap) (wxRsp *ScorePermissionRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScorePermission, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScorePermission, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScorePermissionRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePermission)
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

// 查询用户授权记录（授权协议号）API
// Code = 0 is success
func (c *ClientV3) V3ScorePermissionQuery(ctx context.Context, authCode, serviceId string) (wxRsp *ScorePermissionQueryRsp, err error) {
	uri := fmt.Sprintf(v3ScorePermissionQuery, authCode) + "?service_id=" + serviceId
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScorePermissionQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePermissionQuery)
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

// 解除用户授权关系（授权协议号）API
// Code = 0 is success
func (c *ClientV3) V3ScorePermissionTerminate(ctx context.Context, authCode, serviceId, reason string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePermissionTerminate, authCode)
	bm := make(gopay.BodyMap)
	bm.Set("service_id", serviceId).
		Set("reason", reason)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询用户授权记录（openid）API
// Code = 0 is success
func (c *ClientV3) V3ScorePermissionOpenidQuery(ctx context.Context, appid, openid, serviceid string) (wxRsp *ScorePermissionOpenidQueryRsp, err error) {
	uri := fmt.Sprintf(v3ScorePermissionOpenidQuery, openid) + "?appid=" + appid + "&service_id=" + serviceid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScorePermissionOpenidQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScorePermissionOpenidQuery)
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

// 解除用户授权关系（openid）API
// Code = 0 is success
func (c *ClientV3) V3ScorePermissionOpenidTerminate(ctx context.Context, appid, openid, serviceid, reason string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ScorePermissionOpenidTerminate, openid)
	bm := make(gopay.BodyMap)
	bm.Set("service_id", serviceid).
		Set("appid", appid).
		Set("reason", reason)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 创建支付分订单API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderCreate(ctx context.Context, bm gopay.BodyMap) (wxRsp *ScoreOrderCreateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScoreOrderCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScoreOrderCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderCreate)
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

// V3ScoreOrderPartnerCreate 服务行模式创建支付分订单
// Code = 0 is success
// 微信文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter6_2_1.shtml
func (c *ClientV3) V3ScoreOrderPartnerCreate(ctx context.Context, bm gopay.BodyMap) (*ScoreOrderPartnerCreateRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ScoreOrderPartnerCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScoreOrderPartnerCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &ScoreOrderPartnerCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderPartnerCreate)
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

// 查询支付分订单API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderQuery(ctx context.Context, orderNoType OrderNoType, appid, orderNo, serviceid string) (wxRsp *ScoreOrderQueryRsp, err error) {
	var uri string
	switch orderNoType {
	case OutTradeNo:
		uri = v3ScoreOrderQuery + "?appid=" + appid + "&out_order_no=" + orderNo + "&service_id=" + serviceid
	case QueryId:
		uri = v3ScoreOrderQuery + "?appid=" + appid + "&query_id=" + orderNo + "&service_id=" + serviceid
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
	wxRsp = &ScoreOrderQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderQuery)
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

// V3ScoreOrderPartnerQuery 服务商模式查询支付分订单
// Code = 0 is success
// 微信文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter6_2_2.shtml
func (c *ClientV3) V3ScoreOrderPartnerQuery(ctx context.Context, orderNoType OrderNoType, orderNo, serviceid, subMchid string) (*ScoreOrderPartnerQueryRsp, error) {
	query := url.Values{}
	query.Set("service_id", serviceid)
	query.Set("sub_mchid", subMchid)
	switch orderNoType {
	case OutTradeNo:
		query.Set("out_order_no", orderNo)
	case QueryId:
		query.Set("query_id", orderNo)
	default:
		return nil, errors.New("unsupported order number type")
	}
	uri := fmt.Sprintf("%s?%s", v3ScoreOrderPartnerQuery, query.Encode())
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &ScoreOrderPartnerQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderPartnerQuery)
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

// 取消支付分订单API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderCancel(ctx context.Context, appid, tradeNo, serviceid, reason string) (wxRsp *ScoreOrderCancelRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderCancel, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("appid", appid).
		Set("service_id", serviceid).
		Set("reason", reason)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCancelRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderCancel)
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

// V3ScoreOrderPartnerCancel 服务商模式取消支付分订单
// Code = 0 is success
// 微信文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter6_2_3.shtml
func (c *ClientV3) V3ScoreOrderPartnerCancel(ctx context.Context, subMchid, tradeNo, serviceid, reason string) (*ScoreOrderPartnerCancelRsp, error) {
	path := fmt.Sprintf(v3ScoreOrderPartnerCancel, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("sub_mchid", subMchid).
		Set("service_id", serviceid).
		Set("reason", reason)
	authorization, err := c.authorization(MethodPost, path, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, path, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &ScoreOrderPartnerCancelRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 修改订单金额API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderModify(ctx context.Context, tradeNo string, bm gopay.BodyMap) (wxRsp *ScoreOrderModifyRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderModify, tradeNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderModifyRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderModify)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 完结支付分订单API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderComplete(ctx context.Context, tradeNo string, bm gopay.BodyMap) (wxRsp *ScoreOrderCompleteRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderComplete, tradeNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCompleteRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderComplete)
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

// V3ScoreOrderPartnerComplete 服务商模式完结支付分订单A
// Code = 0 is success
// 微信文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter6_2_5.shtml
func (c *ClientV3) V3ScoreOrderPartnerComplete(ctx context.Context, tradeNo string, bm gopay.BodyMap) (*ScoreOrderPartnerCompleteRsp, error) {
	path := fmt.Sprintf(v3ScoreOrderPartnerComplete, tradeNo)
	authorization, err := c.authorization(MethodPost, path, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, path, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &ScoreOrderPartnerCompleteRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户发起催收扣款API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderPay(ctx context.Context, appid, tradeNo, serviceid string) (wxRsp *ScoreOrderPayRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderPay, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("appid", appid).
		Set("service_id", serviceid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderPayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderPay)
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

// 同步服务订单信息API
// Code = 0 is success
func (c *ClientV3) V3ScoreOrderSync(ctx context.Context, tradeNo string, bm gopay.BodyMap) (wxRsp *ScoreOrderSyncRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderSync, tradeNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderSyncRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderSync)
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
