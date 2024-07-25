package paypal

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestCreateOrder(t *testing.T) {
	var pus []*PurchaseUnit
	var item = &PurchaseUnit{
		ReferenceId: util.RandomString(16),
		Amount: &Amount{
			CurrencyCode: "USD",
			Value:        "8",
		},
	}
	pus = append(pus, item)

	bm := make(gopay.BodyMap)
	// can be AUTHORIZE
	bm.Set("intent", "CAPTURE").
		Set("purchase_units", pus).
		SetBodyMap("application_context", func(b gopay.BodyMap) {
			b.Set("brand_name", "gopay").
				Set("locale", "en-PT").
				Set("return_url", "https://example.com/returnUrl").
				Set("cancel_url", "https://example.com/cancelUrl")
		})

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.CreateOrder(ctx, bm)
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

func TestOrderDetail(t *testing.T) {
	ppRsp, err := client.OrderDetail(ctx, "4X223967G91314611", nil)
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
	for _, v := range ppRsp.Response.PurchaseUnits {
		xlog.Debugf("ppRsp.Response.PurchaseUnit.ReferenceId: %+v", v.ReferenceId)
		xlog.Debugf("ppRsp.Response.PurchaseUnit.Amount: %+v", v.Amount)
		if v.Shipping != nil && v.Shipping.Address != nil {
			xlog.Debugf("ppRsp.Response.PurchaseUnit.Shipping.Address: %+v", v.Shipping.Address)
		}
		xlog.Debugf("ppRsp.Response.PurchaseUnit.Description: %+v", v.Description)
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestUpdateOrder(t *testing.T) {
	var ps []*Patch
	item := &Patch{
		Op:   "replace",
		Path: "/purchase_units/@reference_id=='default'/shipping/address", // reference_id is yourself set when create order
		Value: &Address{
			AddressLine1: "321 Townsend St",
			AddressLine2: "Floor 7",
			AdminArea1:   "San Francisco",
			AdminArea2:   "CA",
			PostalCode:   "94107",
			CountryCode:  "US",
		},
	}
	item2 := &Patch{
		Op:    "add",
		Path:  "/purchase_units/@reference_id=='default'/description",
		Value: "I am patch info",
	}
	ps = append(ps, item)
	ps = append(ps, item2)

	ppRsp, err := client.UpdateOrder(ctx, "4X223967G91314611", ps)
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

func TestOrderAuthorize(t *testing.T) {
	ppRsp, err := client.OrderAuthorize(ctx, "4X223967G91314611", nil)
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
	for _, v := range ppRsp.Response.PurchaseUnits {
		xlog.Debugf("ppRsp.Response.PurchaseUnit.ReferenceId: %+v", v.ReferenceId)
		xlog.Debugf("ppRsp.Response.PurchaseUnit.Amount: %+v", v.Amount)
		if v.Shipping != nil && v.Shipping.Address != nil {
			xlog.Debugf("ppRsp.Response.PurchaseUnit.Shipping.Address: %+v", v.Shipping.Address)
		}
		xlog.Debugf("ppRsp.Response.PurchaseUnit.Description: %+v", v.Description)
		if v.Payments != nil && v.Payments.Authorizations != nil {
			xlog.Debugf("ppRsp.Response.PurchaseUnit.Payments.Authorizations: %+v", v.Payments.Authorizations)
		}
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestOrderCapture(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.SetBodyMap("payment_source", func(b gopay.BodyMap) {
		b.SetBodyMap("token", func(b gopay.BodyMap) {
			b.Set("id", "The PayPal-generated ID for the token").
				Set("type", "BILLING_AGREEMENT")
		})
	})
	ppRsp, err := client.OrderCapture(ctx, "4X223967G91314611", bm)
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
	for _, v := range ppRsp.Response.PurchaseUnits {
		xlog.Debugf("ppRsp.Response.PurchaseUnit.ReferenceId: %+v", v.ReferenceId)
		xlog.Debugf("ppRsp.Response.PurchaseUnit.Amount: %+v", v.Amount)
		if v.Shipping != nil && v.Shipping.Address != nil {
			xlog.Debugf("ppRsp.Response.PurchaseUnit.Shipping.Address: %+v", v.Shipping.Address)
		}
		xlog.Debugf("ppRsp.Response.PurchaseUnit.Description: %+v", v.Description)
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}
