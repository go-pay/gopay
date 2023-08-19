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

// 获取当前汇率
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-GetExchange
func (c *Client) GetExchangeRate(ctx context.Context) (rsp *ExchangeRateRsp, err error) {
	url := fmt.Sprintf(getExchangeRate, c.PartnerCode)
	bs, err := c.doGet(ctx, url, "")
	if err != nil {
		return nil, err
	}
	rsp = new(ExchangeRateRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 获取加密密钥
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-GetEncrypt
func (c *Client) GetEncrypt(ctx context.Context) (rsp *EncryptRsp, err error) {
	url := fmt.Sprintf(getEncrypt, c.PartnerCode)
	bs, err := c.doGet(ctx, url, "")
	if err != nil {
		return nil, err
	}
	rsp = new(EncryptRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 关闭订单
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-CloseOrder
func (c *Client) CloseOrder(ctx context.Context, orderId string) (rsp *ErrorCode, err error) {
	url := fmt.Sprintf(closeOrder, c.PartnerCode, orderId)
	bs, err := c.doPut(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	rsp = new(ErrorCode)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 查询订单状态
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-OrderStatus
func (c *Client) OrderStatus(ctx context.Context, orderId string) (rsp *OrdersRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	url := fmt.Sprintf(getOrderStatus, c.PartnerCode, orderId)
	bs, err := c.doGet(ctx, url, "")
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

// 申请退款
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-RefundOrder
func (c *Client) ApplyRefund(ctx context.Context, orderId string, refundId string, bm gopay.BodyMap) (rsp *RefundRsp, err error) {
	rsp = new(RefundRsp)
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if refundId == gopay.NULL {
		return nil, fmt.Errorf("refundId is empty")
	}
	if err = bm.CheckEmptyError("fee"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(applyRefund, c.PartnerCode, orderId, refundId)
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

// 查询退款状态
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-RefundQuery
func (c *Client) RefundQuery(ctx context.Context, orderId string, refundId string) (rsp *RefundRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if refundId == gopay.NULL {
		return nil, fmt.Errorf("refundId is empty")
	}
	url := fmt.Sprintf(getRefundStatus, c.PartnerCode, orderId, refundId)
	bs, err := c.doGet(ctx, url, "")
	if err != nil {
		return nil, err
	}
	rsp = new(RefundRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 查看订单
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-ListOrder
func (c *Client) OrderList(ctx context.Context, date, status string, page, limit int) (rsp *OrderListRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("date", date).Set("status", status).Set("page", page).Set("limit", limit)
	url := fmt.Sprintf(queryOrderList, c.PartnerCode)
	bs, err := c.doGet(ctx, url, bm.EncodeURLParams())
	if err != nil {
		return nil, err
	}
	rsp = new(OrderListRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 查看账单流水
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-ListTransaction
func (c *Client) TransactionList(ctx context.Context, date string) (rsp *TransactionListRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("date", date)
	url := fmt.Sprintf(queryTransactionList, c.PartnerCode)
	bs, err := c.doGet(ctx, url, bm.EncodeURLParams())
	if err != nil {
		return nil, err
	}
	rsp = new(TransactionListRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 查看清算详情
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-SettleLog
func (c *Client) Settlements(ctx context.Context, date string) (rsp *SettlementsRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("date", date)
	url := fmt.Sprintf(querySettlements, c.PartnerCode)
	bs, err := c.doGet(ctx, url, bm.EncodeURLParams())
	if err != nil {
		return nil, err
	}
	rsp = new(SettlementsRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 查询可用钱包
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-ConsultPayment
func (c *Client) ConsultPayment(ctx context.Context, bm gopay.BodyMap) (rsp *ConsultPaymentRsp, err error) {
	if err = bm.CheckEmptyError("currency", "amount", "terminal_type"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(queryConsultPayment, c.PartnerCode)
	bs, err := c.doPost(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(ConsultPaymentRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 获取优惠券信息
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-GetCoupon
func (c *Client) GetCoupon(ctx context.Context, couponId string) (rsp *GetCouponRsp, err error) {
	url := fmt.Sprintf(getCoupon, c.PartnerCode, couponId)
	bs, err := c.doGet(ctx, url, "")
	if err != nil {
		return nil, err
	}
	rsp = new(GetCouponRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 付款通知
// 文档：https://payjp.lakala.com/docs/cn/#api-CommonApi-PayNotice
func ParseNotify(req *http.Request) (notifyReq *NotifyRequest, err error) {
	var buf bytes.Buffer
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
