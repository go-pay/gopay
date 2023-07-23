package lakala

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-pay/gopay"
)

// 文档：https://payjp.lakala.com/docs/cn/#api-Channel_Web_Gateway-WebGateway
/*
CommonApi - 获取当前汇率
1.2.0
获取当前各渠道JPY兑CNY汇率值(1JPY=?CNY)，该汇率仅做参考，以实际成交汇率为准
*/
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/channel_exchange_rate
func (c *Client) GetChannelExchangeRate(ctx context.Context) (rsp *QRCodeRsp, err error) {
	url := fmt.Sprintf(rateUrl, c.PartnerCode)
	bs, err := c.doGet(ctx, url)
	if err != nil {
		return nil, err
	}
	rsp = new(QRCodeRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 文档：https://payjp.lakala.com/docs/cn/#api-Channel_Web_Gateway-WebGateway
/*
CommonApi - 获取加密密钥
1.2.0
获取系统随机密钥，确保数据加密传输
*/
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/encrypt
func (c *Client) GetEncrypt(ctx context.Context) (rsp *QRCodeRsp, err error) {
	url := fmt.Sprintf(encryptUrl, c.PartnerCode)
	bs, err := c.doGet(ctx, url)
	if err != nil {
		return nil, err
	}
	rsp = new(QRCodeRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 文档：https://payjp.lakala.com/docs/cn/#api-SDKPayment-SDK
/*
CommonApi - 查询订单状态
1.2.0
查询范围包括所有订单
*/
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}
func (c *Client) GetOrders(ctx context.Context, orderId string) (rsp *OrdersRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	url := fmt.Sprintf(getOrdersUrl, c.PartnerCode, orderId)
	bs, err := c.doGet(ctx, url)
	if err != nil {
		return nil, err
	}
	rsp = new(OrdersRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 文档：https://payjp.lakala.com/docs/cn/#api-SDKPayment-SDK
/*
CommonApi - 付款通知
1.2.0
若订单创建时提供了notify_url，系统会在用户支付成功后向这个地址主动发送支付成功状态推送，请求方式为POST 与服务器API不同，推送校验参数会包含在json内，商户系统应该验证校验参数，确定来源正确后再进行后续操作。 商户系统收到请求完成订单确认处理后应确保接口的HTTP状态码为200，否则视为商户未接收成功，推送动作最多会重试3次，商户系统应当能够处理收到的重复请求。
建议在收到推送后主动调用查询接口确认支付状态，避免其他第三方伪造推送!!!
*/
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}
func ParseNotifyResult(req *http.Request) (notifyReq *NotifyRequest, err error) {
	var (
		buf bytes.Buffer
	)
	if _, err = buf.ReadFrom(req.Body); err != nil {
		return nil, err
	}
	if err = req.Body.Close(); err != nil {
		return nil, err
	}
	body := buf.Bytes()
	req.Body = io.NopCloser(bytes.NewReader(body))

	notifyReq = new(NotifyRequest)
	err = json.Unmarshal(body, notifyReq)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(body))
	}
	return
}

// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-RefundOrder
/*
CommonApi - 申请退款
1.2.0
一笔支付订单可以分多次退款，退款总金额不得超过实际支付金额，退款币种与支付订单一致
*/
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}/refunds/{refund_id}

func (c *Client) CreateOrdersRefunds(ctx context.Context, orderId string, refundId string, bm gopay.BodyMap) (rsp *OrdersRefundsRsp, err error) {
	rsp = new(OrdersRefundsRsp)
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if refundId == gopay.NULL {
		return nil, fmt.Errorf("refundId is empty")
	}
	if err = bm.CheckEmptyError("fee"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(ordersRefundsUrl, c.PartnerCode, orderId, refundId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}
