package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// GetRefundHistory Get Refund History
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
func (c *Client) GetRefundHistory(ctx context.Context, transactionId, revision string) (rsp *RefundHistoryRsp, err error) {
	path := fmt.Sprintf(getRefundHistory, transactionId) + "?revision=" + revision
	res, bs, err := c.doRequestGet(ctx, path)
	if err != nil {
		return nil, err
	}
	rsp = &RefundHistoryRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode == http.StatusOK {
		return rsp, nil
	}
	if err = statusCodeErrCheck(rsp.StatusCodeErr); err != nil {
		return rsp, err
	}
	return rsp, nil
}
