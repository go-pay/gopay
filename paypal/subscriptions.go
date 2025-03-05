package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建计划（Create plan）
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
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	ppRsp.Response = new(BillingDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return ppRsp, nil
}

// 计划列表（List plans）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_list
func (c *Client) PlanList(ctx context.Context, bm gopay.BodyMap) (ppRsp *PlanListRsp, err error) {
	uri := planList + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &PlanListRsp{Code: Success}
	ppRsp.Response = new(BillingPlan)
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

// 计划详情（Show plan details）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_get
func (c *Client) PlanDetails(ctx context.Context, planId string) (ppRsp *PlanDetailRsp, err error) {
	if planId == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	uri := fmt.Sprintf(planDetail, planId)
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

// 更新计划（Update plan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_patch
func (c *Client) PlanUpdate(ctx context.Context, planId string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if planId == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	url := fmt.Sprintf(planUpdate, planId)
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

// 激活计划（Activate plan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_activate
func (c *Client) PlanActivate(ctx context.Context, planId string) (ppRsp *EmptyRsp, err error) {
	if planId == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	url := fmt.Sprintf(planActivate, planId)
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

// 停用计划（Deactivate plan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_activate
func (c *Client) PlanDeactivate(ctx context.Context, planId string) (ppRsp *EmptyRsp, err error) {
	if planId == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	url := fmt.Sprintf(planDeactivate, planId)
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

// 更新计划价格（Update pricing）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_activate
func (c *Client) PlanUpdatePrice(ctx context.Context, planId string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if planId == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	if err = bm.CheckEmptyError("pricing_schemes"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(planUpdatePrice, planId)
	res, bs, err := c.doPayPalPost(ctx, bm, url)
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

// 创建订阅（Create subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_create
func (c *Client) SubscriptionCreate(ctx context.Context, bm gopay.BodyMap) (ppRsp *SubscriptionCreateRsp, err error) {
	if err = bm.CheckEmptyError("plan_id"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, subscriptionCreate)
	if err != nil {
		return nil, err
	}
	ppRsp = &SubscriptionCreateRsp{Code: Success}
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

// 订阅详情（Show subscription details）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get
func (c *Client) SubscriptionDetails(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *SubscriptionDetailRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscription_id is empty")
	}
	uri := fmt.Sprintf(subscriptionDetail, subscriptionId) + "?" + bm.EncodeURLParams()
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

// 更新订阅（Update subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_patch
func (c *Client) SubscriptionUpdate(ctx context.Context, subscriptionId string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscriptionId is empty")
	}
	url := fmt.Sprintf(subscriptionUpdate, subscriptionId)
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

// 修改计划或订阅数量（Revise plan or quantity of subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_revise
func (c *Client) SubscriptionRevise(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *SubscriptionReviseRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscriptionId is empty")
	}
	uri := fmt.Sprintf(subscriptionRevise, subscriptionId)
	res, bs, err := c.doPayPalPost(ctx, bm, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &SubscriptionReviseRsp{Code: Success}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	ppRsp.Response = new(ReviseSubscription)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return ppRsp, nil
}

// 暂停订阅（Suspend subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_suspend
func (c *Client) SubscriptionSuspend(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscriptionId is empty")
	}
	if err = bm.CheckEmptyError("reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionSuspend, subscriptionId)
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

// 取消订阅（Cancel subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_cancel
func (c *Client) SubscriptionCancel(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscriptionId is empty")
	}
	if err = bm.CheckEmptyError("reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionCancel, subscriptionId)
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

// 激活订阅（Activate subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_activate
func (c *Client) SubscriptionActivate(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscriptionId is empty")
	}
	if err = bm.CheckEmptyError("reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionActivate, subscriptionId)
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

// 订阅时获取授权付款（Capture authorized payment on subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_capture
func (c *Client) SubscriptionCapture(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if subscriptionId == gopay.NULL {
		return nil, errors.New("subscriptionId is empty")
	}
	if err = bm.CheckEmptyError("note", "capture_type", "amount"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(subscriptionCapture, subscriptionId)
	res, bs, err := c.doPayPalPost(ctx, bm, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &EmptyRsp{Code: Success}

	if res.StatusCode != http.StatusAccepted {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 订阅的交易列表（List transactions for subscription）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_transactions
func (c *Client) SubscriptionTransactionList(ctx context.Context, subscriptionId string, bm gopay.BodyMap) (ppRsp *SubscriptionTransactionListRsp, err error) {
	uri := fmt.Sprintf(subscriptionTransactions, subscriptionId) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &SubscriptionTransactionListRsp{Code: Success}
	ppRsp.Response = new(SubscriptionTransaction)
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
