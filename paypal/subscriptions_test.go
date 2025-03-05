package paypal

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestClient_CreateBillingPlan(t *testing.T) {
	var billingCycles []*BillingCycles
	var billingCycle = &BillingCycles{
		Frequency: &Frequency{
			IntervalUnit:  "MONTH",
			IntervalCount: 1,
		},
		TenureType:  "REGULAR",
		Sequence:    1,
		TotalCycles: 0,
		PricingScheme: &PricingScheme{
			FixedPrice: &FixedPrice{
				Value:        "101",
				CurrencyCode: "USD",
			},
		},
	}
	billingCycles = append(billingCycles, billingCycle)

	// 创建 PayPal 支付订单
	bm := make(gopay.BodyMap)
	bm.Set("product_id", "PROD-10J947659N0823244").
		Set("name", "gopay").
		Set("billing_cycles", billingCycles).
		Set("description", "Monthly subscription for premium users").
		Set("payment_preferences", &PaymentPreferences{
			AutoBillOutstanding:     true,
			SetupFeeFailureAction:   "CONTINUE",
			PaymentFailureThreshold: 3,
		})

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.CreateBillingPlan(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestListBillingPlan(t *testing.T) {
	ppRsp, err := client.ListBillingPlan(ctx, nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Plans {
		xlog.Debugf("ppRsp.Response.Item: %+v", v)
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestBillingPlanDetail(t *testing.T) {
	ppRsp, err := client.BillingPlanDetails(ctx, "P-4A621926UG9673307M7D3JIA", nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestUpdateBillingPlan(t *testing.T) {
	var ps []*Patch
	item := &Patch{
		Op:    "replace",
		Path:  "/name", // reference_id is yourself set when create order
		Value: "Updated Video Streaming Service Plan",
	}

	ps = append(ps, item)

	ppRsp, err := client.UpdateBillingPlan(ctx, "P-4A621926UG9673307M7D3JIA", ps)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
}
