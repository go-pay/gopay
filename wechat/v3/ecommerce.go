package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// 二级商户进件API
// 注意：本接口会提交一些敏感信息，需调用 client.V3EncryptText() 进行加密。部分图片参数，请先调用 client.V3MediaUploadImage() 上传，获取MediaId
// Code = 0 is success
func (c *ClientV3) V3EcommerceApply(ctx context.Context, bm gopay.BodyMap) (*EcommerceApplyRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceApply, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceApply, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceApplyRsp{Code: Success, SignInfo: si, Response: &EcommerceApply{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询申请状态API
// 注意：applyId 和 outRequestNo 二选一
// Code = 0 is success
func (c *ClientV3) V3EcommerceApplyStatus(ctx context.Context, applyId int64, outRequestNo string) (*EcommerceApplyStatusRsp, error) {
	if applyId == 0 && outRequestNo == gopay.NULL {
		return nil, fmt.Errorf("applyId[%d] and outRequestNo[%s] empty at the same time", applyId, outRequestNo)
	}
	var url string
	if applyId != 0 {
		url = fmt.Sprintf(v3EcommerceApplyQueryById, applyId)
	} else {
		url = fmt.Sprintf(v3EcommerceApplyQueryByNo, outRequestNo)
	}
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceApplyStatusRsp{Code: Success, SignInfo: si, Response: &EcommerceApplyStatus{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 请求分账API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShare(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceProfitShare, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceProfitShare, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShare{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询分账结果API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareQuery(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareQueryRsp, error) {
	uri := v3EcommerceProfitShareQuery + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareQueryRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShareQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 请求分账回退API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareReturn(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareReturnRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceProfitShareReturn, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceProfitShareReturn, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareReturnRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShareReturn{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询分账回退结果API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareReturnResult(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareReturnResultRsp, error) {
	uri := v3EcommerceProfitShareReturnResult + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareReturnResultRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShareReturn{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 完结分账API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareFinish(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareFinishRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceProfitShareFinish, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceProfitShareFinish, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareFinishRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShareFinish{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询订单剩余待分金额API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareUnsplitAmount(ctx context.Context, transactionId string) (*EcommerceProfitShareUnsplitAmountRsp, error) {
	url := fmt.Sprintf(v3EcommerceProfitShareUnsplitAmount, transactionId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareUnsplitAmountRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShareUnsplitAmount{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 添加分账接收方API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareAddReceiver(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareAddReceiverRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceProfitShareAddReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceProfitShareAddReceiver, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceProfitShareAddReceiverRsp{Code: Success, SignInfo: si, Response: new(EcommerceProfitShareReceiver)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 删除分账接收方API
// Code = 0 is success
func (c *ClientV3) V3EcommerceProfitShareDeleteReceiver(ctx context.Context, bm gopay.BodyMap) (*EcommerceProfitShareDeleteReceiverRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceProfitShareDeleteReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceProfitShareDeleteReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &EcommerceProfitShareDeleteReceiverRsp{Code: Success, SignInfo: si, Response: &EcommerceProfitShareReceiver{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 请求补差API
// Code = 0 is success
func (c *ClientV3) V3EcommerceSubsidies(ctx context.Context, bm gopay.BodyMap) (*EcommerceSubsidiesRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceSubsidies, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceSubsidies, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceSubsidiesRsp{Code: Success, SignInfo: si, Response: new(EcommerceSubsidies)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 请求补差回退API
// Code = 0 is success
func (c *ClientV3) V3EcommerceSubsidiesReturn(ctx context.Context, bm gopay.BodyMap) (*EcommerceSubsidiesReturnRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceSubsidiesReturn, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceSubsidiesReturn, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceSubsidiesReturnRsp{Code: Success, SignInfo: si, Response: new(EcommerceSubsidiesReturn)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 取消补差API
// Code = 0 is success
func (c *ClientV3) V3EcommerceSubsidiesCancel(ctx context.Context, bm gopay.BodyMap) (*EcommerceSubsidiesCancelRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceSubsidiesCancel, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceSubsidiesCancel, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceSubsidiesCancelRsp{Code: Success, SignInfo: si, Response: new(EcommerceSubsidiesCancel)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}
