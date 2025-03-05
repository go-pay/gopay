package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建订阅计划（CreateBillingPlan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_create
func (c *Client) CreateBillingPlan(ctx context.Context, bm gopay.BodyMap) (ppRsp *CreateBillingRsp, err error) {
	if err = bm.CheckEmptyError("product_id", "billing_cycles"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, planCreate)
	if err != nil {
		return nil, err
	}
	ppRsp = &CreateBillingRsp{Code: Success}
	ppRsp.Response = new(BillingDetail)
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

// 订阅计划列表（ListBillingPlan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_list
func (c *Client) ListBillingPlan(ctx context.Context, bm gopay.BodyMap) (ppRsp *BillingListRsp, err error) {
	uri := planList + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &BillingListRsp{Code: Success}
	ppRsp.Response = new(BillingPlans)
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

// 订阅计划详情（CatalogsProductDetails）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_get
func (c *Client) BillingPlanDetails(ctx context.Context, PlanID string, bm gopay.BodyMap) (ppRsp *PlanDetailRsp, err error) {
	if PlanID == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	uri := fmt.Sprintf(planDetail, PlanID) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &PlanDetailRsp{Code: Success}
	ppRsp.Response = new(PlanDetail)
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

// 更新订阅计划（UpdateBillingPlan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_patch
func (c *Client) UpdateBillingPlan(ctx context.Context, PlanID string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if PlanID == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	url := fmt.Sprintf(planUpdate, PlanID)
	res, bs, err := c.doPayPalPatch(ctx, patchs, url)
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

// 创建订阅（CreateBillingSubscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_create
func (c *Client) CreateBillingSubscription(ctx context.Context, bm gopay.BodyMap) (ppRsp *CreateSubscriptionRsp, err error) {
	if err = bm.CheckEmptyError("plan_id"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, subscriptionCreate)
	if err != nil {
		return nil, err
	}
	ppRsp = &CreateSubscriptionRsp{Code: Success}
	ppRsp.Response = new(SubscriptionDetail)
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

// 订阅详情（SubscriptionDetails）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get
func (c *Client) SubscriptionDetails(ctx context.Context, SubscriptionID string, bm gopay.BodyMap) (ppRsp *SubscriptionDetailRsp, err error) {
	if SubscriptionID == gopay.NULL {
		return nil, errors.New("subscription_id is empty")
	}
	uri := fmt.Sprintf(subscriptionDetail, SubscriptionID) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &SubscriptionDetailRsp{Code: Success}
	ppRsp.Response = new(SubscriptionDetail)
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

// 暂停订阅（SuspendSubscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_suspend
func (c *Client) SuspendSubscription(ctx context.Context, SubscriptionID string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if err = bm.CheckEmptyError("reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionSuspend, SubscriptionID)
	res, bs, err := c.doPayPalPost(ctx, bm, uri)
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

// 取消订阅（CancelSubscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_cancel
func (c *Client) CancelSubscription(ctx context.Context, SubscriptionID string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if err = bm.CheckEmptyError("reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionCancel, SubscriptionID)
	res, bs, err := c.doPayPalPost(ctx, bm, uri)
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

// 激活订阅（ActivateSubscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_activate
func (c *Client) ActivateSubscription(ctx context.Context, SubscriptionID string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if err = bm.CheckEmptyError("reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionActivate, SubscriptionID)
	res, bs, err := c.doPayPalPost(ctx, bm, uri)
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

// List transactions for subscription（ListTransSubscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_transactions
func (c *Client) ListTransSubscription(ctx context.Context, SubscriptionID string, bm gopay.BodyMap) (ppRsp *ListTransSubscriptionRsp, err error) {
	uri := fmt.Sprintf(subscriptionTransactions, SubscriptionID) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &ListTransSubscriptionRsp{Code: Success}
	ppRsp.Response = new(TransSubscription)
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
