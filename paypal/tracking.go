package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// AddTrackingNumber 添加物流单号
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_track_create
func (c *Client) AddTrackingNumber(ctx context.Context, orderId string, bm gopay.BodyMap) (ppRsp *AddTrackingNumberRsp, err error) {
	if err = bm.CheckEmptyError("tracking_number", "carrier", "capture_id"); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(addTrackingNumber, orderId)
	res, bs, err := c.doPayPalPost(ctx, bm, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &AddTrackingNumberRsp{Code: Success}
	ppRsp.Response = new(OrderDetail)
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
