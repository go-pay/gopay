package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// SendConsumptionInformation Send Consumption Information
// Doc: https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (c *Client) SendConsumptionInformation(ctx context.Context, transactionId string, bm gopay.BodyMap) (rsp *TransactionHistoryRsp, err error) {
	path := fmt.Sprintf(sendConsumptionInformation, transactionId) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doRequestGet(ctx, path)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.stauts_coud = %d", res.StatusCode)
	}
	rsp = &TransactionHistoryRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return rsp, nil
}
