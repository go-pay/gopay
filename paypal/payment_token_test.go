package paypal

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestPaymentSetupTokenCreate(t *testing.T) {
	bm := make(gopay.BodyMap)

	bm.SetBodyMap("payment_source", func(b1 gopay.BodyMap) {
		b1.SetBodyMap("paypal", func(b3 gopay.BodyMap) {
			b3.SetBodyMap("experience_context", func(b4 gopay.BodyMap) {
				b4.Set("shipping_preference", "SET_PROVIDED_ADDRESS").
					Set("payment_method_preference", "IMMEDIATE_PAYMENT_REQUIRED").
					Set("brand_name", "gopay").
					Set("return_url", "https://example.com/returnUrl").
					Set("cancel_url", "https://example.com/cancelUrl")
			}).
				Set("permit_multiple_payment_tokens", false).
				Set("usage_pattern", "IMMEDIATE").
				Set("usage_type", "MERCHANT").
				Set("customer_type", "CONSUMER")
		})
	})
	bm.SetBodyMap("customer", func(b2 gopay.BodyMap) {
		b2.Set("id", "1234567890")
	})

	xlog.Debug("bm: ", bm.JsonBody())

	ppRsp, err := client.CreateSetupToken(ctx, bm)
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

func TestRetrieveSetupToken(t *testing.T) {
	ppRsp, err := client.RetrieveSetupToken(ctx, "5CS813092M1570432")
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
}

func TestListAllPaymentTokens(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("customer_id", "1234567890").
		Set("total_required", true)

	ppRsp, err := client.ListAllPaymentTokens(ctx, bm)
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
}

func TestRetrievePaymentToken(t *testing.T) {
	ppRsp, err := client.RetrievePaymentToken(ctx, "5CS813092M1570432")
	if err != nil {
		t.Error(err)
		return
	}

	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestCreatePaymentToken(t *testing.T) {
	bm := make(gopay.BodyMap)

	bm.SetBodyMap("payment_source", func(b1 gopay.BodyMap) {
		b1.SetBodyMap("token", func(b2 gopay.BodyMap) {
			b2.Set("id", "5CS813092M1570432").
				Set("type", "SETUP_TOKEN")
		})
	})

	ppRsp, err := client.CreatePaymentToken(ctx, bm)
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
}

func TestDeletePaymentToken(t *testing.T) {
	ppRsp, err := client.DeletePaymentToken(ctx, "5CS813092M1570432")
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
}
