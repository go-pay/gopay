package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 资金授权操作查询接口 alipay.fund.auth.operation.detail.query
// StatusCode = 200 is success
func (a *ClientV3) FundAuthOperationDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOperationDetailQueryRsp, err error) {
	authorization, err := a.authorization(MethodPost, v3FundAuthOperationDetailQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundAuthOperationDetailQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundAuthOperationDetailQueryRsp{StatusCode: res.StatusCode}
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

// 资金授权冻结接口 alipay.fund.auth.order.freeze
// StatusCode = 200 is success
func (a *ClientV3) FundAuthOrderFreeze(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOrderFreezeRsp, err error) {
	err = bm.CheckEmptyError("auth_code", "auth_code_type", "out_order_no", "out_request_no", "order_title", "product_code", "amount")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FundAuthOrderFreeze, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundAuthOrderFreeze, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundAuthOrderFreezeRsp{StatusCode: res.StatusCode}
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

// 资金授权解冻接口 alipay.fund.auth.order.unfreeze
// StatusCode = 200 is success
func (a *ClientV3) FundAuthOrderUnfreeze(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOrderUnfreezeRsp, err error) {
	err = bm.CheckEmptyError("auth_no", "out_request_no", "amount", "remark")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FundAuthOrderUnfreeze, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundAuthOrderUnfreeze, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundAuthOrderUnfreezeRsp{StatusCode: res.StatusCode}
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

// 资金授权发码接口 alipay.fund.auth.order.voucher.create
// StatusCode = 200 is success
func (a *ClientV3) FundAuthOrderVoucherCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOrderVoucherCreateRsp, err error) {
	err = bm.CheckEmptyError("out_order_no", "out_request_no", "order_title", "amount", "product_code")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FundAuthOrderVoucherCreate, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundAuthOrderVoucherCreate, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundAuthOrderVoucherCreateRsp{StatusCode: res.StatusCode}
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
