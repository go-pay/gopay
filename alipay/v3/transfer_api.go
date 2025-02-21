package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 支付宝资金账户资产查询接口 alipay.fund.account.query
// StatusCode = 200 is success
func (a *ClientV3) FundAccountQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAccountQueryRsp, err error) {
	err = bm.CheckEmptyError("account_type")
	if err != nil {
		return nil, err
	}
	if bm.GetString("alipay_user_id") == gopay.NULL && bm.GetString("alipay_open_id") == gopay.NULL {
		return nil, errors.New("alipay_user_id and alipay_open_id are not allowed to be null at the same time")
	}
	uri := v3FundAccountQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundAccountQueryRsp{StatusCode: res.StatusCode}
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

// 转账额度查询接口 alipay.fund.quota.query
// StatusCode = 200 is success
func (a *ClientV3) FundQuotaQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundQuotaQueryRsp, err error) {
	err = bm.CheckEmptyError("product_code", "biz_scene")
	if err != nil {
		return nil, err
	}
	uri := v3FundQuotaQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundQuotaQueryRsp{StatusCode: res.StatusCode}
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

// 单笔转账接口 alipay.fund.trans.uni.transfer
// StatusCode = 200 is success
func (a *ClientV3) FundTransUniTransfer(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransUniTransferRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene", "payee_info", "order_title")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FundTransUniTransfer, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundTransUniTransfer, authorization)
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

// 申请电子回单(incubating) alipay.data.bill.ereceipt.apply
// StatusCode = 200 is success
func (a *ClientV3) DataBillEreceiptApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillEreceiptApplyRsp, err error) {
	err = bm.CheckEmptyError("type", "key")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3DataBillEreceiptApply, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataBillEreceiptApply, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &DataBillEreceiptApplyRsp{StatusCode: res.StatusCode}
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

// 查询电子回单状态(incubating) alipay.data.bill.ereceipt.query
// StatusCode = 200 is success
func (a *ClientV3) DataBillEreceiptQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillEreceiptQueryRsp, err error) {
	err = bm.CheckEmptyError("file_id")
	if err != nil {
		return nil, err
	}
	uri := v3DataBillEreceiptQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &DataBillEreceiptQueryRsp{StatusCode: res.StatusCode}
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

// 转账业务单据查询接口 alipay.fund.trans.common.query
// StatusCode = 200 is success
func (a *ClientV3) FundTransCommonQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransCommonQueryRsp, err error) {
	uri := v3FundTransCommonQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundTransCommonQueryRsp{StatusCode: res.StatusCode}
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

// 多步转账创建并支付 alipay.fund.trans.multistep.transfer
// StatusCode = 200 is success
func (a *ClientV3) FundTransMultistepTransfer(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransMultistepTransferRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "product_code", "biz_scene", "total_amount", "total_count", "order_details")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FundTransMultistepTransfer, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundTransMultistepTransfer, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundTransMultistepTransferRsp{StatusCode: res.StatusCode}
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

// 多步转账查询接口 alipay.fund.trans.multistep.query
// StatusCode = 200 is success
func (a *ClientV3) FundTransMultistepQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransMultistepQueryRsp, err error) {
	err = bm.CheckEmptyError("product_code", "biz_scene")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FundTransMultistepQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FundTransMultistepQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FundTransMultistepQueryRsp{StatusCode: res.StatusCode}
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
