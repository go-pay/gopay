package paypal

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestPaymentAuthorizeDetail(t *testing.T) {
	ppRsp, err := client.PaymentAuthorizeDetail(ctx, "4X223967G91314611")
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

func TestPaymentReauthorize(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.SetBodyMap("amount", func(bm gopay.BodyMap) {
		bm.Set("currency_code", "USD").
			Set("value", "10.99")
	})

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.PaymentReauthorize(ctx, "4X223967G91314611", bm)
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

func TestPaymentAuthorizeVoid(t *testing.T) {
	ppRsp, err := client.PaymentAuthorizeVoid(ctx, "4X223967G91314611")
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

func TestPaymentAuthorizeCapture(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("invoice_id", "INVOICE-123").
		Set("final_capture", true).
		Set("note_to_payer", "If the ordered color is not available, we will substitute with a different color free of charge.").
		Set("soft_descriptor", "Bob's Custom Sweaters").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("currency_code", "USD").
				Set("value", "10.99")
		})

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.PaymentAuthorizeCapture(ctx, "4X223967G91314611", bm)
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

func TestPaymentCaptureDetail(t *testing.T) {
	ppRsp, err := client.PaymentCaptureDetail(ctx, "4X223967G91314611")
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

func TestPaymentCaptureRefund(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("invoice_id", "INVOICE-123").
		Set("note_to_payer", "Defective product").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("currency_code", "USD").
				Set("value", "10.99")
		})

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.PaymentCaptureRefund(ctx, "4X223967G91314611", bm)
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

func TestPaymentRefundDetail(t *testing.T) {
	ppRsp, err := client.PaymentRefundDetail(ctx, "4X223967G91314611")
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
