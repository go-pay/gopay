package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 特约商户余额提现API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter6_1.shtml
func (c *ClientV3) V3Withdraw(ctx context.Context, bm gopay.BodyMap) (*WithdrawRsp, error) {
	authorization, err := c.authorization(MethodPost, v3Withdraw, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3Withdraw, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &WithdrawRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Withdraw)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询特约商户提现状态API
//	注意：withdrawId 和 outRequestNo 二选一
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter6_2.shtml
func (c *ClientV3) V3WithdrawStatus(ctx context.Context, withdrawId, outRequestNo string, bm gopay.BodyMap) (*WithdrawStatusRsp, error) {
	if withdrawId == gopay.NULL && outRequestNo == gopay.NULL {
		return nil, fmt.Errorf("withdrawId[%s] and outRequestNo[%s] empty at the same time", withdrawId, outRequestNo)
	}
	var uri string
	if withdrawId != gopay.NULL {
		uri = fmt.Sprintf(v3WithdrawStatusById, withdrawId) + "?" + bm.EncodeURLParams()
	} else {
		uri = fmt.Sprintf(v3WithdrawStatusByNo, outRequestNo) + "?" + bm.EncodeURLParams()
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &WithdrawStatusRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(WithdrawStatus)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 按日下载提现异常文件API
//	注意：如 bill_date 为空，默认查前一天的
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter6_3.shtml
func (c *ClientV3) V3WithdrawDownloadErrBill(ctx context.Context, bm gopay.BodyMap) (wxRsp *BillRsp, err error) {
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
		bm.Remove("bill_type")
	}
	uri := fmt.Sprintf(v3WithdrawDownloadErrBill, "NO_SUCC") + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &BillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TradeBill)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
