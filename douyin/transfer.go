package douyin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// Transfer 商户转账到抖音零钱
// bm 关键字段：appid、out_bill_no、transfer_scene_id、openid、transfer_amount、transfer_remark、notify_url
// 可选：user_name（≥2000.00 元时必填，需先调用 (c *Client).EncryptText 加密）、user_recv_perception、transfer_scene_report_infos
// 商户号通过 Authorization 头传递（无 body mchid 字段）
func (c *Client) Transfer(ctx context.Context, bm gopay.BodyMap) (dyRsp *TransferRsp, err error) {
	authorization, err := c.authorization(MethodPost, transfer, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, transfer, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &TransferRsp{Code: Success, SignInfo: si, Response: new(TransferBill)}
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

// TransferQueryByOutBillNo 使用商户订单号查询转账单
func (c *Client) TransferQueryByOutBillNo(ctx context.Context, outBillNo string) (dyRsp *TransferRsp, err error) {
	if outBillNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	uri := fmt.Sprintf(transferQueryByOutBillNo, outBillNo)
	return c.transferQuery(ctx, uri)
}

// TransferQueryByTransferBillNo 使用抖音转账单号查询转账单
func (c *Client) TransferQueryByTransferBillNo(ctx context.Context, transferBillNo string) (dyRsp *TransferRsp, err error) {
	if transferBillNo == gopay.NULL {
		return nil, gopay.MissParamErr
	}
	uri := fmt.Sprintf(transferQueryByTransferBillNo, transferBillNo)
	return c.transferQuery(ctx, uri)
}

func (c *Client) transferQuery(ctx context.Context, uri string) (dyRsp *TransferRsp, err error) {
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &TransferRsp{Code: Success, SignInfo: si, Response: new(TransferBill)}
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
