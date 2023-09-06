package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 申请退款API
// Code = 0 is success
func (c *ClientV3) V3Refund(ctx context.Context, bm gopay.BodyMap) (wxRsp *RefundRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3DomesticRefund, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3DomesticRefund, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundOrderResponse)
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

// 查询单笔退款API
// 注意：商户查询时，bm 可传 nil；服务商时，传相应query参数
// Code = 0 is success
func (c *ClientV3) V3RefundQuery(ctx context.Context, outRefundNo string, bm gopay.BodyMap) (wxRsp *RefundQueryRsp, err error) {
	uri := fmt.Sprintf(v3DomesticRefundQuery, outRefundNo)
	if bm != nil {
		uri = fmt.Sprintf(v3DomesticRefundQuery, outRefundNo) + "?" + bm.EncodeURLParams()
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundQueryResponse)
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

// 申请退款API
// Code = 0 is success
func (c *ClientV3) V3EcommerceRefund(ctx context.Context, bm gopay.BodyMap) (wxRsp *EcommerceRefundRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3CommerceRefund, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CommerceRefund, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EcommerceRefundRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceRefund)
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

// 通过微信支付退款单号查询退款API
// Code = 0 is success
func (c *ClientV3) V3EcommerceRefundQueryById(ctx context.Context, refundId string, bm gopay.BodyMap) (wxRsp *EcommerceRefundQueryRsp, err error) {
	uri := fmt.Sprintf(v3CommerceRefundQueryById, refundId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EcommerceRefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceRefundQuery)
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

// 通过商户退款单号查询退款API
// Code = 0 is success
func (c *ClientV3) V3EcommerceRefundQueryByNo(ctx context.Context, outRefundNo string, bm gopay.BodyMap) (wxRsp *EcommerceRefundQueryRsp, err error) {
	uri := fmt.Sprintf(v3CommerceRefundQueryByNo, outRefundNo) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EcommerceRefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceRefundQuery)
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

// 垫付退款回补API
// Code = 0 is success
func (c *ClientV3) V3EcommerceRefundAdvance(ctx context.Context, refundId string, bm gopay.BodyMap) (wxRsp *EcommerceRefundAdvanceRsp, err error) {
	url := fmt.Sprintf(v3CommerceRefundAdvance, refundId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EcommerceRefundAdvanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceRefundAdvance)
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

// 查询垫付回补结果API
// Code = 0 is success
func (c *ClientV3) V3EcommerceRefundAdvanceResult(ctx context.Context, refundId string, bm gopay.BodyMap) (wxRsp *EcommerceRefundAdvanceRsp, err error) {
	uri := fmt.Sprintf(v3CommerceRefundAdvanceResult, refundId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EcommerceRefundAdvanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceRefundAdvance)
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
