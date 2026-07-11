package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// 发起转账并完成免确认收款授权（一步式）
// 文档：https://pay.weixin.qq.com/doc/v3/merchant/4014399293.md
// 2025-05 新增。一次接口同时拉起授权 + 发起转账，返回 package_info 由前端调起授权页。
func (c *ClientV3) V3TransferPreTransferWithAuth(ctx context.Context, bm gopay.BodyMap) (*TransferPreTransferWithAuthRsp, error) {
	authorization, err := c.authorization(MethodPost, V3TransferPreTransferWithAuth, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, V3TransferPreTransferWithAuth, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferPreTransferWithAuthRsp{Code: Success, SignInfo: si, Response: &TransferPreTransferWithAuth{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 发起免确认收款授权（单独申请授权，不发起转账）
// 文档：https://pay.weixin.qq.com/doc/v3/merchant/4015901167.md
// 2025-05 新增。返回 package_info 由前端调起授权页让用户确认。
func (c *ClientV3) V3TransferUserConfirmAuth(ctx context.Context, bm gopay.BodyMap) (*TransferUserConfirmAuthRsp, error) {
	authorization, err := c.authorization(MethodPost, V3TransferUserConfirmAuth, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, V3TransferUserConfirmAuth, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferUserConfirmAuthRsp{Code: Success, SignInfo: si, Response: &TransferUserConfirmAuth{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户单号查询授权结果
// 文档：https://pay.weixin.qq.com/doc/v3/merchant/4014399423.md
// query 可选传 is_display_authorization=true 让响应里返回 package_info（2026-06 新增）。
func (c *ClientV3) V3TransferUserConfirmAuthQuery(ctx context.Context, outAuthorizationNo string, bm gopay.BodyMap) (*TransferUserConfirmAuthQryRsp, error) {
	uri := fmt.Sprintf(V3TransferUserConfirmAuthQry, outAuthorizationNo)
	if bm != nil {
		if q := bm.EncodeURLParams(); q != "" {
			uri = uri + "?" + q
		}
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferUserConfirmAuthQryRsp{Code: Success, SignInfo: si, Response: &TransferUserConfirmAuthQry{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 解除免确认收款授权
// 文档：https://pay.weixin.qq.com/doc/v3/merchant/4015653811.md
// 无请求体。
func (c *ClientV3) V3TransferUserConfirmAuthClose(ctx context.Context, outAuthorizationNo string) (*TransferUserConfirmAuthCloseRsp, error) {
	uri := fmt.Sprintf(V3TransferUserConfirmAuthClose, outAuthorizationNo)
	authorization, err := c.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, nil, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferUserConfirmAuthCloseRsp{Code: Success, SignInfo: si, Response: &TransferUserConfirmAuthClose{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}
