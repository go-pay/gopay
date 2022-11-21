package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// GetAllSubscriptionStatuses
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func GetAllSubscriptionStatuses(ctx context.Context, signConfig *SignConfig, originalTransactionId string, sandbox bool) (rsp *AllSubscriptionStatusesRsp, err error) {
	uri := hostUrl + fmt.Sprintf(getAllSubscriptionStatuses, originalTransactionId)
	if sandbox {
		uri = sandBoxHostUrl + fmt.Sprintf(getAllSubscriptionStatuses, originalTransactionId)
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
	rsp = &AllSubscriptionStatusesRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return
}
