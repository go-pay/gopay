package douyin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// ProfitRequest 请求分账（异步受理）
// bm 关键字段：appid、transaction_id、out_order_no、receivers（数组）、unfreeze_unsplit、notify_url
// mchid 不设置时自动填 c.Mchid
// receivers[].name 若有值必须先经 (c *Client).EncryptText 加密（RSA-PKCS1v15）
func (c *Client) ProfitRequest(ctx context.Context, bm gopay.BodyMap) (dyRsp *ProfitRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, profitRequest, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, profitRequest, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitRsp{Code: Success, SignInfo: si, Response: new(Profit)}
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

// ProfitQuery 查询分账结果
// outOrderNo 商户分账单号（路径参数）
// bm 关键 query 参数：mchid（必需，未设置时自动填 c.Mchid）、transaction_id 或 order_id（二选一）
func (c *Client) ProfitQuery(ctx context.Context, outOrderNo string, bm gopay.BodyMap) (dyRsp *ProfitRsp, err error) {
	if outOrderNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	if bm == nil {
		bm = make(gopay.BodyMap)
	}
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	uri := fmt.Sprintf(profitQuery, outOrderNo) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitRsp{Code: Success, SignInfo: si, Response: new(Profit)}
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

// ProfitRollback 请求分账回退
// bm 关键字段：mchid、order_id、out_order_no、out_return_no、return_mchid、amount、description
// mchid 不设置时自动填 c.Mchid
func (c *Client) ProfitRollback(ctx context.Context, bm gopay.BodyMap) (dyRsp *ProfitReturnRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, profitRollback, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, profitRollback, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitReturnRsp{Code: Success, SignInfo: si, Response: new(ProfitReturn)}
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

// ProfitRollbackQuery 查询分账回退结果
// outReturnNo 商户回退单号（路径参数）
// bm 关键 query 参数：mchid（必需，未设置时自动填 c.Mchid）、out_order_no
func (c *Client) ProfitRollbackQuery(ctx context.Context, outReturnNo string, bm gopay.BodyMap) (dyRsp *ProfitReturnRsp, err error) {
	if outReturnNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	if bm == nil {
		bm = make(gopay.BodyMap)
	}
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	uri := fmt.Sprintf(profitRollbackQuery, outReturnNo) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitReturnRsp{Code: Success, SignInfo: si, Response: new(ProfitReturn)}
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

// ProfitComplete 完结分账（unfreeze_unsplit=false 下单后需要主动完结）
// bm 关键字段：mchid、transaction_id、out_order_no、description、notify_url
func (c *Client) ProfitComplete(ctx context.Context, bm gopay.BodyMap) (dyRsp *ProfitRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, profitComplete, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, profitComplete, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitRsp{Code: Success, SignInfo: si, Response: new(Profit)}
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

// ProfitBalanceQuery 查询订单剩余待分账金额
// transactionId 抖音支付订单号（路径参数）
// mchid 可选，未传时自动使用 c.Mchid
func (c *Client) ProfitBalanceQuery(ctx context.Context, transactionId, mchid string) (dyRsp *ProfitBalanceRsp, err error) {
	if transactionId == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	m := c.Mchid
	if mchid != gopay.NULL {
		m = mchid
	}
	uri := fmt.Sprintf(profitBalanceQuery, transactionId) + "?mchid=" + m
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitBalanceRsp{Code: Success, SignInfo: si, Response: new(ProfitBalance)}
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

// ProfitReceiverAdd 添加分账接收方
// bm 关键字段：mchid、appid、type、account、name（需 EncryptText 加密）、relation_type
func (c *Client) ProfitReceiverAdd(ctx context.Context, bm gopay.BodyMap) (dyRsp *ProfitReceiverRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, receiverAdd, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, receiverAdd, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitReceiverRsp{Code: Success, SignInfo: si, Response: new(ProfitReceiverInfo)}
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

// ProfitReceiverDelete 删除分账接收方
// bm 关键字段：mchid、appid、type、account
func (c *Client) ProfitReceiverDelete(ctx context.Context, bm gopay.BodyMap) (dyRsp *ProfitReceiverRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, receiverDelete, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, receiverDelete, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &ProfitReceiverRsp{Code: Success, SignInfo: si, Response: new(ProfitReceiverInfo)}
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
