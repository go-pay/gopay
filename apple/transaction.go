package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// GetTransactionHistory
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func GetTransactionHistory(ctx context.Context, signConfig *SignConfig, originalTransactionId string, bm gopay.BodyMap, sandbox bool) (rsp *TransactionHistoryRsp, err error) {
	uri := hostUrl + fmt.Sprintf(getTransactionHistory, originalTransactionId) + "?" + bm.EncodeURLParams()
	if sandbox {
		uri = sandBoxHostUrl + fmt.Sprintf(getTransactionHistory, originalTransactionId) + "?" + bm.EncodeURLParams()
	}
	token, err := generatingToken(ctx, signConfig)
	if err != nil {
		return nil, err
	}
	cli := xhttp.NewClient()
	cli.Header.Set("Authorization", "Bearer "+token)
	res, bs, err := cli.Type(xhttp.TypeJSON).Get(uri).EndBytes(ctx)
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
	return
}
