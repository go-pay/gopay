package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 请求分账API
// 微信会在接到请求后立刻返回请求接收结果，分账结果需要自行调用查询接口来获取
// Code = 0 is success
func (c *ClientV3) V3ProfitShareOrder(ctx context.Context, bm gopay.BodyMap) (*ProfitShareOrderRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareOrder, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ProfitShareOrder, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareOrder)
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

// 查询分账结果API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareOrderQuery(ctx context.Context, orderNo string, bm gopay.BodyMap) (*ProfitShareOrderQueryRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareQuery, orderNo) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareOrderQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareOrderQuery)
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

// 请求分账回退API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareReturn(ctx context.Context, bm gopay.BodyMap) (*ProfitShareReturnRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareReturn, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ProfitShareReturn, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareReturnRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareReturn)
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

// 查询分账回退结果API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareReturnResult(ctx context.Context, returnNo string, bm gopay.BodyMap) (*ProfitShareReturnResultRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareReturnResult, returnNo) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareReturnResultRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareReturnResult)
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

// 解冻剩余资金API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareOrderUnfreeze(ctx context.Context, bm gopay.BodyMap) (*ProfitShareOrderUnfreezeRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareUnfreeze, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ProfitShareUnfreeze, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareOrderUnfreezeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareOrderUnfreeze)
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

// 查询剩余待分金额API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareUnsplitAmount(ctx context.Context, transId string) (*ProfitShareUnsplitAmountRsp, error) {
	url := fmt.Sprintf(v3ProfitShareUnsplitAmount, transId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareUnsplitAmountRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareUnsplitAmount)
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

// 查询最大分账比例API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareMerchantConfigs(ctx context.Context, subMchId string) (*ProfitShareMerchantConfigsRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareMerchantConfigs, subMchId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareMerchantConfigsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareMerchantConfigs)
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

// 新增分账接收方API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareAddReceiver(ctx context.Context, bm gopay.BodyMap) (*ProfitShareAddReceiverRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareAddReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ProfitShareAddReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareAddReceiverRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareAddReceiver)
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

// 删除分账接收方API
// Code = 0 is success
func (c *ClientV3) V3ProfitShareDeleteReceiver(ctx context.Context, bm gopay.BodyMap) (*ProfitShareDeleteReceiverRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareDeleteReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ProfitShareDeleteReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareDeleteReceiverRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareDeleteReceiver)
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

// 申请分账账单
// Code = 0 is success
func (c *ClientV3) V3ProfitShareBills(ctx context.Context, bm gopay.BodyMap) (*ProfitShareBillsRsp, error) {
	uri := v3ProfitShareBills + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareBillsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareBills)
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
