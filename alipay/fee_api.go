package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.open.fee.adjust.apply(特殊费率申请)
// 文档地址：https://opendocs.alipay.com/open/c50c780c_alipay.open.fee.adjust.apply
func (a *Client) OpenFeeAdjustApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenFeeAdjustApplyResponse, err error) {
	err = bm.CheckEmptyError("account", "product_code", "application_fee")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.FileUploadRequest(ctx, bm, "alipay.open.fee.adjust.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenFeeAdjustApplyResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
