package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// GetTransactionHistory Get Transaction History
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func (c *Client) GetTransactionHistory(ctx context.Context, transactionId string, bm gopay.BodyMap) (rsp *TransactionHistoryRsp, err error) {
	path := fmt.Sprintf(getTransactionHistory, transactionId) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doRequestGet(ctx, path)
	if err != nil {
		return nil, err
	}
	rsp = &TransactionHistoryRsp{}
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

// GetTransactionInfo Get Transaction Info
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_info
func (c *Client) GetTransactionInfo(ctx context.Context, transactionId string) (rsp *TransactionInfoRsp, err error) {
	path := fmt.Sprintf(getTransactionInfo, transactionId)
	res, bs, err := c.doRequestGet(ctx, path)
	if err != nil {
		return nil, err
	}
	rsp = &TransactionInfoRsp{}
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
