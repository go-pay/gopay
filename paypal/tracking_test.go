package paypal

import (
	"testing"

	"github.com/go-pay/xlog"

	"github.com/go-pay/gopay"
)

func TestAddTrackingNumber(t *testing.T) {
	var items []*ShipItem
	var item = &ShipItem{
		Name:     "T-Shirt",
		Quantity: 1,
		Sku:      "sku02",
		Url:      "https://www.example.com/example.jpg",
		ImageUrl: "https://www.example.com/example",
	}
	items = append(items, item)

	bm := make(gopay.BodyMap)
	bm.Set("capture_id", "1DW71051X94135205").
		Set("tracking_number", "UJ639398620YP").
		Set("carrier", "YANWEN").
		Set("items", items)

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.AddTrackingNumber(ctx, "3TH70640XJ5862H", bm)
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
