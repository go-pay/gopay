package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 单笔转账接口 alipay.fund.trans.uni.transfer
// StatusCode = 200 is success
func (a *ClientV3) FundTransUniTransfer(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransUniTransferRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene", "payee_info", "order_title")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TransUniTransfer, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TransUniTransfer, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundTransUniTransferRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
