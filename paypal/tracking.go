package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// UpdateTracker 更新或取消订单的物流信息
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_trackers_patch
// 仅支持 op=replace/add；status 仅允许 replace 到 CANCELLED
// 成功状态码：204 No Content
func (c *Client) UpdateTracker(ctx context.Context, orderId, trackerId string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if trackerId == gopay.NULL {
		return nil, errors.New("tracker_id is empty")
	}
	url := fmt.Sprintf(updateTracker, orderId, trackerId)
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
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
