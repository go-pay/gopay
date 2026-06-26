package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// CreatePaymentToken creates a payment token.
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payment-tokens/v3/#payment-tokens_create
func (c *Client) CreatePaymentToken(ctx context.Context, bm gopay.BodyMap) (ppRsp *PaymentTokenCreateRsp, err error) {
	if err = bm.CheckEmptyError("payment_source"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, createPaymentToken)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentTokenCreateRsp{Code: Success}
	ppRsp.Response = new(PaymentMethodDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, err
}

// ListAllPaymentTokens lists all payment tokens.
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payment-tokens/v3/#customer_payment-tokens_get
func (c *Client) ListAllPaymentTokens(ctx context.Context, query gopay.BodyMap) (ppRsp *PaymentTokenListRsp, err error) {
	uri := paymentTokenList + "?" + query.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentTokenListRsp{Code: Success}
	ppRsp.Response = new(PaymentTokensList)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, err
}

// RetrievePaymentToken retrieves a payment token.
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payment-tokens/v3/#payment-tokens_get
func (c *Client) RetrievePaymentToken(ctx context.Context, id string) (ppRsp *PaymentTokenDetailRsp, err error) {
	if id == gopay.NULL {
		return nil, errors.New("id is empty")
	}
	uri := fmt.Sprintf(retrievePaymentToken, id)
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentTokenDetailRsp{Code: Success}
	ppRsp.Response = new(PaymentMethodDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, err
}

// DeletePaymentToken deletes a payment token.
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payment-tokens/v3/#payment-tokens_delete
func (c *Client) DeletePaymentToken(ctx context.Context, id string) (ppRsp *EmptyRsp, err error) {
	if id == gopay.NULL {
		return nil, errors.New("id is empty")
	}
	uri := fmt.Sprintf(deletePaymentToken, id)
	res, bs, err := c.doPayPalDelete(ctx, uri)
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
	return ppRsp, err
}

// CreateSetupToken creates a setup token.
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payment-tokens/v3/#setup-tokens_create
func (c *Client) CreateSetupToken(ctx context.Context, bm gopay.BodyMap) (ppRsp *PaymentSetupTokenCreateRsp, err error) {
	if err = bm.CheckEmptyError("payment_source"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, createSetupToken)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentSetupTokenCreateRsp{Code: Success}
	ppRsp.Response = new(PaymentSetupTokenDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, err
}

// RetrieveSetupToken retrieves a setup token.
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payment-tokens/v3/#setup-tokens_get
func (c *Client) RetrieveSetupToken(ctx context.Context, id string) (ppRsp *PaymentSetupTokenCreateRsp, err error) {
	if id == gopay.NULL {
		return nil, errors.New("id is empty")
	}
	uri := fmt.Sprintf(retrieveSetupToken, id)
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &PaymentSetupTokenCreateRsp{Code: Success}
	ppRsp.Response = new(PaymentSetupTokenDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, err
}
