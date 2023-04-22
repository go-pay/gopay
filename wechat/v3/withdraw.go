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

// 特约商户余额提现、二级商户预约提现
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_14.shtml
func (c *ClientV3) V3Withdraw(ctx context.Context, bm gopay.BodyMap) (*WithdrawRsp, error) {
	if err := bm.CheckEmptyError("sub_mchid", "out_request_no", "amount"); err != nil {
		return nil, err
	}
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
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询特约商户提现状态、二级商户查询预约提现状态
// 注意：withdrawId 和 outRequestNo 二选一
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_15.shtml
func (c *ClientV3) V3WithdrawStatus(ctx context.Context, withdrawId, outRequestNo string, bm gopay.BodyMap) (*WithdrawStatusRsp, error) {
	if withdrawId == gopay.NULL && outRequestNo == gopay.NULL {
		return nil, fmt.Errorf("[%w]: withdrawId[%s] and outRequestNo[%s] empty at the same time", gopay.MissParamErr, withdrawId, outRequestNo)
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
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 电商平台预约提现
// Code = 0 is success
// 电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_2.shtml
func (c *ClientV3) V3EcommerceWithdraw(ctx context.Context, bm gopay.BodyMap) (*EcommerceWithdrawRsp, error) {
	if err := bm.CheckEmptyError("out_request_no", "amount", "account_type"); err != nil {
		return nil, err
	}
	authorization, err := c.authorization(MethodPost, v3EcommerceWithdraw, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceWithdraw, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &EcommerceWithdrawRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceWithdraw)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 电商平台查询预约提现状态
// 注意：withdrawId 和 outRequestNo 二选一
// Code = 0 is success
// 电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_6.shtml
func (c *ClientV3) V3EcommerceWithdrawStatus(ctx context.Context, withdrawId, outRequestNo string) (*EcommerceWithdrawStatusRsp, error) {
	if withdrawId == gopay.NULL && outRequestNo == gopay.NULL {
		return nil, fmt.Errorf("[%w]: withdrawId[%s] and outRequestNo[%s] empty at the same time", gopay.MissParamErr, withdrawId, outRequestNo)
	}
	var uri string
	if withdrawId != gopay.NULL {
		uri = fmt.Sprintf(v3EcommerceWithdrawStatusById, withdrawId)
	} else {
		uri = fmt.Sprintf(v3EcommerceWithdrawStatusByNo, outRequestNo)
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceWithdrawStatusRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceWithdrawStatus)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 按日下载提现异常文件
// 注意：如 bill_date 为空，默认查前一天的
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_16.shtml
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
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
