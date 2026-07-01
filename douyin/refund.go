package douyin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// Refund 申请退款
// transaction_id 与 out_trade_no 二选一传入
func (c *Client) Refund(ctx context.Context, bm gopay.BodyMap) (dyRsp *RefundRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, refund, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, refund, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &RefundRsp{Code: Success, SignInfo: si, Response: new(Refund)}
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

// RefundQuery 查询单笔退款（通过商户退款单号）
// outRefundNo 商户退款单号（路径参数），mchid 可选（默认 c.Mchid），appid 可选
func (c *Client) RefundQuery(ctx context.Context, outRefundNo string, mchid, appid string) (dyRsp *RefundRsp, err error) {
	if outRefundNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	m := c.Mchid
	if mchid != gopay.NULL {
		m = mchid
	}
	uri := fmt.Sprintf(refundQuery, outRefundNo) + "?mchid=" + m
	if appid != gopay.NULL {
		uri += "&appid=" + appid
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &RefundRsp{Code: Success, SignInfo: si, Response: new(Refund)}
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
