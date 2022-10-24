package chinaunionpay

import (
	"bytes"
	"fmt"
	"github.com/go-pay/gopay/pkg/xlog"
	"net/http"
)

type BaseClient struct {
	RequestId  string
	HttpClient *http.Client

	RequestBody  []byte
	ResponseBody []byte
}

func NewBaseClient(requestId string) *BaseClient {
	c := new(BaseClient)
	c.HttpClient = GetDefaultHTTPClient()
	c.RequestId = requestId
	return c
}

// Deprecated
func (r *BaseClient) SetRequestId(requestId string) {
	r.RequestId = requestId
}

func (r *BaseClient) SetRequestBody(body []byte) {
	body = r.TrimSpace(body)
	xlog.Debug(r.RequestId,"请求通道BODY:%s", body)
	r.RequestBody = body
}

func (r *BaseClient) SetResponseBody(body []byte) {
	body = r.TrimSpace(body)
	xlog.Debug(r.RequestId,"通道应答BODY:%s", body)
	r.ResponseBody = body
}

func (r *BaseClient) GetRequestBody() string {
	return string(r.RequestBody)
}

func (r *BaseClient) GetResponseBody() string {
	return string(r.ResponseBody)
}

func (r *BaseClient) TrimSpace(body []byte) []byte {
	body = bytes.TrimSpace(body)
	body = bytes.Replace(body, []byte("\r"), []byte(""), -1)
	body = bytes.Replace(body, []byte("\n"), []byte(""), -1)
	return body
}

func (r *BaseClient) GetGateWayError(gatewayCode, gatewayMsg string) error {
	return fmt.Errorf(r.GetGateWayStatus(gatewayCode, gatewayMsg))
}

func (r *BaseClient) GetGateWayStatus(gatewayCode, gatewayMsg string) string {
	return fmt.Sprintf("gatewayCode:%s gatewayMsg:%s", gatewayCode, gatewayMsg)
}

func (r *BaseClient) GetBizError(bizCode, bizMsg string) error {
	return fmt.Errorf(r.GetBizTrxMsg(bizCode, bizMsg))
}

func (r *BaseClient) GetBizTrxMsg(bizCode, bizMsg string) string {
	return fmt.Sprintf("bizCode:%s bizMsg:%s", bizCode, bizMsg)
}
