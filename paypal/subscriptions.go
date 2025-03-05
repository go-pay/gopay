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
	res, bs, err := c.doPayPalPost(ctx, bm, subscriptionCreate)
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

// 订阅列表（ListBillingPlan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_list
func (c *Client) ListBillingPlan(ctx context.Context, bm gopay.BodyMap) (ppRsp *BillingListRsp, err error) {
	uri := subscriptionList + "?" + bm.EncodeURLParams()
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

// 订阅详情（CatalogsProductDetails）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_get
func (c *Client) BillingPlanDetails(ctx context.Context, PlanID string, bm gopay.BodyMap) (ppRsp *PlanDetailRsp, err error) {
	if PlanID == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	uri := fmt.Sprintf(subscriptionDetail, PlanID) + "?" + bm.EncodeURLParams()
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

// 更新订阅（UpdateBillingPlan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_patch
func (c *Client) UpdateBillingPlan(ctx context.Context, PlanID string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if PlanID == gopay.NULL {
		return nil, errors.New("plan_id is empty")
	}
	url := fmt.Sprintf(subscriptionUpdate, PlanID)
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
