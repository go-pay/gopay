package paypal

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
)

var (
	client   *Client
	ctx      = context.Background()
	err      error
	Clientid = ""
	Secret   = ""
)

func TestMain(m *testing.M) {
	client, err = NewClient(Clientid, Secret, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	xlog.Debugf("Appid: %s", client.Appid)
	xlog.Debugf("AccessToken: %s", client.AccessToken)
	xlog.Debugf("ExpiresIn: %d", client.ExpiresIn)
	os.Exit(m.Run())
}

func TestBasicAuth(t *testing.T) {
	uname := "jerry"
	passwd := "12346"
	auth := base64.StdEncoding.EncodeToString([]byte(uname + ":" + passwd))
	xlog.Debugf("Basic %s", auth)
}

func TestCreateOrder(t *testing.T) {
	var pus []*PurchaseUnit
	var item = &PurchaseUnit{
		ReferenceId: util.GetRandomString(16),
		Amount: &Amount{
			CurrencyCode: "USD",
			Value:        "10",
		},
		Shipping: &Shipping{
			Name: &Name{FullName: "Fumingming"},
			Type: "PICKUP_IN_PERSON",
			Address: &Address{
				AddressLine1: "123 Townsend St",
				AddressLine2: "Floor 6",
				AdminArea1:   "San Francisco",
				AdminArea2:   "CA",
				PostalCode:   "94107",
				CountryCode:  "US",
			},
		},
	}
	pus = append(pus, item)

	bm := make(gopay.BodyMap)
	bm.Set("intent", "CAPTURE").
		Set("purchase_units", pus)

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.CreateOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
}

func TestOrderDetail(t *testing.T) {
	ppRsp, err := client.OrderDetail(ctx, "9W792304W36262048", nil)
	if err != nil {
		xlog.Error(err)
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

	ppRsp, err := client.UpdateOrder(ctx, "9W792304W36262048", ps)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
	xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
}
