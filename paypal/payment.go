package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 支付授权详情（Show details for authorized payment）
//	Code = 0 is success
//	文档：https://developer.paypal.com/api/payments/v2/#authorizations_get
func (c *Client) PaymentAuthorizeDetail(ctx context.Context, authorizationId string) (ppRsp *PaymentAuthorizeDetailRsp, err error) {
	if authorizationId == gopay.NULL {
		return nil, errors.New("authorization_id is empty")
	}
	url := fmt.Sprintf(paymentAuthorizeDetail, authorizationId)
	res, bs, err := c.doPayPalGet(ctx, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentAuthorizeDetailRsp{Code: Success}
	ppRsp.Response = new(PaymentAuthorizeDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 重新授权支付授权（Reauthorize authorized payment）
//	Note：This request is currently not supported for Partner use cases.
//	文档：https://developer.paypal.com/api/payments/v2/#authorizations_reauthorize
func (c *Client) PaymentReauthorize(ctx context.Context, authorizationId string, bm gopay.BodyMap) (ppRsp *PaymentReauthorizeRsp, err error) {
	if authorizationId == gopay.NULL {
		return nil, errors.New("authorization_id is empty")
	}
	url := fmt.Sprintf(paymentReauthorize, authorizationId)
	res, bs, err := c.doPayPalPost(ctx, bm, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentReauthorizeRsp{Code: Success}
	ppRsp.Response = new(PaymentAuthorizeDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 作废支付授权（Void authorized payment）
//	Code = 0 is success
//	文档：https://developer.paypal.com/api/payments/v2/#authorizations_void
func (c *Client) PaymentAuthorizeVoid(ctx context.Context, authorizationId string) (ppRsp *EmptyRsp, err error) {
	if authorizationId == gopay.NULL {
		return nil, errors.New("authorization_id is empty")
	}
	url := fmt.Sprintf(paymentAuthorizeVoid, authorizationId)
	res, bs, err := c.doPayPalPost(ctx, nil, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &EmptyRsp{Code: Success}
	if res.StatusCode != http.StatusNoContent {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 支付授权捕获（Capture authorized payment）
//	Code = 0 is success
//	文档：https://developer.paypal.com/api/payments/v2/#authorizations_capture
func (c *Client) PaymentAuthorizeCapture(ctx context.Context, authorizationId string, bm gopay.BodyMap) (ppRsp *PaymentAuthorizeCaptureRsp, err error) {
	if authorizationId == gopay.NULL {
		return nil, errors.New("authorization_id is empty")
	}
	url := fmt.Sprintf(paymentAuthorizeCapture, authorizationId)
	res, bs, err := c.doPayPalPost(ctx, bm, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentAuthorizeCaptureRsp{Code: Success}
	ppRsp.Response = new(PaymentAuthorizeCapture)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 支付捕获详情（Show captured payment details）
//	Code = 0 is success
//	文档：https://developer.paypal.com/api/payments/v2/#captures_get
func (c *Client) PaymentCaptureDetail(ctx context.Context, captureId string) (ppRsp *PaymentCaptureDetailRsp, err error) {
	if captureId == gopay.NULL {
		return nil, errors.New("capture_id is empty")
	}
	url := fmt.Sprintf(paymentCaptureDetail, captureId)
	res, bs, err := c.doPayPalGet(ctx, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentCaptureDetailRsp{Code: Success}
	ppRsp.Response = new(PaymentAuthorizeCapture)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 支付捕获退款（Refund captured payment）
//	Code = 0 is success
//	文档：https://developer.paypal.com/api/payments/v2/#captures_refund
func (c *Client) PaymentCaptureRefund(ctx context.Context, captureId string, bm gopay.BodyMap) (ppRsp *PaymentCaptureRefundRsp, err error) {
	if captureId == gopay.NULL {
		return nil, errors.New("capture_id is empty")
	}
	url := fmt.Sprintf(paymentCaptureRefund, captureId)
	res, bs, err := c.doPayPalPost(ctx, bm, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentCaptureRefundRsp{Code: Success}
	ppRsp.Response = new(PaymentCaptureRefund)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 支付退款详情（Show refund details）
//	Code = 0 is success
//	文档：https://developer.paypal.com/api/payments/v2/#refunds_get
func (c *Client) PaymentRefundDetail(ctx context.Context, refundId string) (ppRsp *PaymentRefundDetailRsp, err error) {
	if refundId == gopay.NULL {
		return nil, errors.New("refund_id is empty")
	}
	url := fmt.Sprintf(paymentRefundDetail, refundId)
	res, bs, err := c.doPayPalGet(ctx, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentRefundDetailRsp{Code: Success}
	ppRsp.Response = new(PaymentCaptureRefund)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
