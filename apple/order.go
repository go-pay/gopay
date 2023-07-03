package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// LookUpOrderId Look Up Order ID
// Doc: https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
func (c *Client) LookUpOrderId(ctx context.Context, orderId string) (rsp *LookUpOrderIdRsp, err error) {
	path := fmt.Sprintf(lookUpOrderID, orderId)
	res, bs, err := c.doRequestGet(ctx, path)
	if err != nil {
		return nil, err
	}
	rsp = &LookUpOrderIdRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return rsp, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode == http.StatusOK {
		return rsp, nil
	}
	if err = statusCodeErrCheck(rsp.StatusCodeErr); err != nil {
		return rsp, err
	}
	return rsp, nil
}
