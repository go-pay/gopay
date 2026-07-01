package douyin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// OrderQueryByTransactionId 使用抖音支付订单号查询订单
// 查询参数：transactionId（路径参数）、mchid（query 参数，若不传则使用 c.Mchid）
func (c *Client) OrderQueryByTransactionId(ctx context.Context, transactionId string, mchid ...string) (dyRsp *OrderQueryRsp, err error) {
	if transactionId == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	m := c.Mchid
	if len(mchid) > 0 && mchid[0] != gopay.NULL {
		m = mchid[0]
	}
	uri := fmt.Sprintf(queryByTransactionId, transactionId) + "?mchid=" + m
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &OrderQueryRsp{Code: Success, SignInfo: si, Response: new(OrderQuery)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}

// OrderQueryByOutTradeNo 使用商户订单号查询订单
func (c *Client) OrderQueryByOutTradeNo(ctx context.Context, outTradeNo string, mchid ...string) (dyRsp *OrderQueryRsp, err error) {
	if outTradeNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	m := c.Mchid
	if len(mchid) > 0 && mchid[0] != gopay.NULL {
		m = mchid[0]
	}
	uri := fmt.Sprintf(queryByOutTradeNo, outTradeNo) + "?mchid=" + m
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &OrderQueryRsp{Code: Success, SignInfo: si, Response: new(OrderQuery)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}

// CloseOrder 关闭订单
// outTradeNo 商户订单号（路径参数），Body 参数 mchid 若不传自动填入 c.Mchid
// 关单成功时 HTTP 200 且 Body 为空
func (c *Client) CloseOrder(ctx context.Context, outTradeNo string, bm gopay.BodyMap) (dyRsp *EmptyRsp, err error) {
	if outTradeNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	if bm == nil {
		bm = make(gopay.BodyMap)
	}
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	uri := fmt.Sprintf(closeOrder, outTradeNo)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	return dyRsp, c.verifySyncSign(si)
}
