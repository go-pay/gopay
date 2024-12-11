/*
@Author: wzy
@Time: 2022/6/7
*/
package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建批量支出（Create batch payout）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_post
func (c *Client) CreateBatchPayout(ctx context.Context, bm gopay.BodyMap) (ppRsp *CreateBatchPayoutRsp, err error) {
	if err = bm.CheckEmptyError("items", "sender_batch_header"); nil != err {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, createBatchPayout)
	if nil != err {
		return nil, err
	}
	ppRsp = &CreateBatchPayoutRsp{Code: Success}
	ppRsp.Response = new(BatchPayout)
	if err = json.Unmarshal(bs, ppRsp.Response); nil != err {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 批量支出详情（Show payout batch details）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get
func (c *Client) ShowPayoutBatchDetails(ctx context.Context, payoutBatchId string, bm gopay.BodyMap) (ppRsp *PayoutBatchDetailRsp, err error) {
	if payoutBatchId == gopay.NULL {
		return nil, errors.New("payout_batch_id is empty")
	}
	uri := fmt.Sprintf(showPayoutBatchDetail, payoutBatchId) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &PayoutBatchDetailRsp{Code: Success}
	ppRsp.Response = new(PayoutBatchDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 批量支出项目详情（Show Payout Item Details）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
func (c Client) ShowPayoutItemDetails(ctx context.Context, payoutItemId string) (ppRsp *PayoutItemDetailRsp, err error) {
	if payoutItemId == gopay.NULL {
		return nil, errors.New("payout_item_id is empty")
	}
	url := fmt.Sprintf(showPayoutItemDetail, payoutItemId)
	res, bs, err := c.doPayPalGet(ctx, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &PayoutItemDetailRsp{Code: Success}
	ppRsp.Response = new(PayoutItemDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 取消批量支付中收款人无PayPal账号的项目（Cancel Unclaimed Payout Item）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_cancel
func (c Client) CancelUnclaimedPayoutItem(ctx context.Context, payoutItemId string) (ppRsp *CancelUnclaimedPayoutItemRsp, err error) {
	if payoutItemId == gopay.NULL {
		return nil, errors.New("payout_item_id is empty")
	}
	url := fmt.Sprintf(cancelUnclaimedPayoutItem, payoutItemId)
	res, bs, err := c.doPayPalPost(ctx, nil, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &CancelUnclaimedPayoutItemRsp{Code: Success}
	ppRsp.Response = new(PayoutItemDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
