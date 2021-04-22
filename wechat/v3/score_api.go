package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
)

// 创建支付分订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_14.shtml
func (c *ClientV3) V3ScoreOrderCreate(bm gopay.BodyMap) (wxRsp *ScoreOrderCreateRsp, err error) {
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}
	authorization, err := c.authorization(MethodPost, v3ScoreOrderCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ScoreOrderCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderCreate)
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

// 查询支付分订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_15.shtml
func (c *ClientV3) V3ScoreOrderQuery(orderNoType OrderNoType, orderNo, serviceId string) (wxRsp *ScoreOrderQueryRsp, err error) {
	var uri string
	switch orderNoType {
	case OutTradeNo:
		uri = v3ScoreOrderQuery + "?appid=" + c.Appid + "&out_order_no=" + orderNo + "&service_id=" + serviceId
	case QueryId:
		uri = v3ScoreOrderQuery + "?appid=" + c.Appid + "&query_id=" + orderNo + "&service_id=" + serviceId
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
	wxRsp = &ScoreOrderQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderQuery)
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

// 取消支付分订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_16.shtml
func (c *ClientV3) V3ScoreOrderCancel(tradeNo, serviceId, reason string) (wxRsp *ScoreOrderCancelRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderCancel, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("appid", c.Appid)
	bm.Set("service_id", serviceId)
	bm.Set("reason", reason)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCancelRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderCancel)
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

// 修改订单金额API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_17.shtml
func (c *ClientV3) V3ScoreOrderModify(tradeNo string, bm gopay.BodyMap) (wxRsp *ScoreOrderModifyRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderModify, tradeNo)
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderModifyRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderModify)
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

// 完结支付分订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_18.shtml
func (c *ClientV3) V3ScoreOrderComplete(tradeNo string, bm gopay.BodyMap) (wxRsp *ScoreOrderCompleteRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderComplete, tradeNo)
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCompleteRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderComplete)
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

// 商户发起催收扣款API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_19.shtml
func (c *ClientV3) V3ScoreOrderPay(tradeNo, serviceId string) (wxRsp *ScoreOrderPayRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderPay, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("appid", c.Appid)
	bm.Set("service_id", serviceId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderPayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderPay)
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

// 同步服务订单信息API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_20.shtml
func (c *ClientV3) V3ScoreOrderSync(tradeNo string, bm gopay.BodyMap) (wxRsp *ScoreOrderSyncRsp, err error) {
	url := fmt.Sprintf(v3ScoreOrderSync, tradeNo)
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderSyncRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderSync)
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
