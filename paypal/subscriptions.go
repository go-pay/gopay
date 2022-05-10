package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

//  创建订阅计划（CreateBillingPlan）
//	Code = 0 is success
//	文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_create
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
