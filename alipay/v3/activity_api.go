package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建商家券活动 alipay.marketing.activity.ordervoucher.create
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherCreateRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "merchant_access_mode", "activity_base_info", "voucher_send_mode_info",
		"voucher_deduct_info", "voucher_available_scope_info", "voucher_use_rule_info", "voucher_customer_guide_info", "voucher_display_pattern_info")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3MarketingActivityOrderVoucherCreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingActivityOrderVoucherCreate, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherCreateRsp{StatusCode: res.StatusCode}
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

// 同步商家券券码 alipay.marketing.activity.ordervoucher.codedeposit
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherCodeDeposit(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherCodeDepositRsp, err error) {
	err = bm.CheckEmptyError("voucher_codes", "out_biz_no", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityOrderVoucherCodeDeposit, activityId)
	authorization, err := a.authorization(MethodPost, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, url, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherCodeDepositRsp{StatusCode: res.StatusCode}
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

// 修改商家券活动基本信息 alipay.marketing.activity.ordervoucher.modify
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherModify(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherModifyRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "merchant_access_mode", "activity_base_info")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityOrderVoucherModify, activityId)
	authorization, err := a.authorization(MethodPatch, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPatch(ctx, bm, url, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherModifyRsp{StatusCode: res.StatusCode}
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

// 停止商家券活动 alipay.marketing.activity.ordervoucher.stop
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherStop(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherStopRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityOrderVoucherStop, activityId)
	authorization, err := a.authorization(MethodPatch, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPatch(ctx, bm, url, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherStopRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 修改商家券活动发券数量上限 alipay.marketing.activity.ordervoucher.append
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherAppend(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherAppendRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "merchant_access_mode", "voucher_quantity")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityOrderVoucherAppend, activityId)
	authorization, err := a.authorization(MethodPatch, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPatch(ctx, bm, url, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherAppendRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 同步券核销状态 alipay.marketing.activity.ordervoucher.use
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherUse(ctx context.Context, activityId, voucherCode string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherUseRsp, err error) {
	err = bm.CheckEmptyError("biz_dt", "trade_channel", "total_fee", "out_biz_no")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityOrderVoucherUse, activityId, voucherCode)
	authorization, err := a.authorization(MethodPost, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, url, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherUseRsp{StatusCode: res.StatusCode}
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

// 取消券核销状态 alipay.marketing.activity.ordervoucher.refund
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherRefund(ctx context.Context, activityId, voucherCode string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherRefundRsp, err error) {
	err = bm.CheckEmptyError("biz_dt", "total_fee", "out_biz_no")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityOrderVoucherRefund, activityId, voucherCode)
	authorization, err := a.authorization(MethodPost, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, url, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherRefundRsp{StatusCode: res.StatusCode}
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

// 活动领取咨询接口 alipay.marketing.activity.consult
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityConsultRsp, err error) {
	err = bm.CheckEmptyError("consult_activity_info_list", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	if bm.GetString("user_id") == gopay.NULL && bm.GetString("open_id") == gopay.NULL {
		return nil, errors.New("user_id and open_id are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3MarketingActivityConsult, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingActivityConsult, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityConsultRsp{StatusCode: res.StatusCode}
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

// 查询商家券活动 alipay.marketing.activity.ordervoucher.query
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3MarketingActivityOrderVoucherQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherQueryRsp{StatusCode: res.StatusCode}
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

// 查询活动详情 alipay.marketing.activity.query
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityQuery(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryRsp, err error) {
	err = bm.CheckEmptyError("merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := fmt.Sprintf(v3MarketingActivityQuery, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityQueryRsp{StatusCode: res.StatusCode}
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

// 统计商家券券码数量 alipay.marketing.activity.ordervoucher.codecount
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityOrderVoucherCodeCount(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityOrderVoucherCodeCountRsp, err error) {
	err = bm.CheckEmptyError("merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := fmt.Sprintf(v3MarketingActivityOrderVoucherCodeCount, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityOrderVoucherCodeCountRsp{StatusCode: res.StatusCode}
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

// 条件查询活动列表 alipay.marketing.activity.batchquery
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("page_num", "page_size", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3MarketingActivityBatchQuery, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingActivityBatchQuery, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityBatchQueryRsp{StatusCode: res.StatusCode}
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

// 条件查询用户券 alipay.marketing.activity.user.batchqueryvoucher
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityQueryUserBatchQueryVoucher(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryUserBatchQueryVoucherRsp, err error) {
	err = bm.CheckEmptyError("auth_token", "page_num", "page_size", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	if bm.GetString("user_id") == gopay.NULL && bm.GetString("open_id") == gopay.NULL {
		return nil, errors.New("user_id and open_id are not allowed to be null at the same time")
	}
	if bm.GetString("belong_merchant_id") == gopay.NULL && bm.GetString("sender_merchant_id") == gopay.NULL {
		return nil, errors.New("belong_merchant_id and sender_merchant_id are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3MarketingActivityQueryUserBatchQueryVoucher + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityQueryUserBatchQueryVoucherRsp{StatusCode: res.StatusCode}
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

// 查询用户券详情 alipay.marketing.activity.user.queryvoucher
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityQueryUserQueryVoucher(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryUserQueryVoucherRsp, err error) {
	if bm.GetString("user_id") == gopay.NULL && bm.GetString("open_id") == gopay.NULL {
		return nil, errors.New("user_id and open_id are not allowed to be null at the same time")
	}
	if bm.GetString("voucher_id") == gopay.NULL && bm.GetString("voucher_code") == gopay.NULL {
		return nil, errors.New("voucher_id and voucher_code are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3MarketingActivityQueryUserQueryVoucher + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityQueryUserQueryVoucherRsp{StatusCode: res.StatusCode}
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

// 查询活动可用小程序 alipay.marketing.activity.app.batchquery
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityQueryAppBatchQuery(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryAppBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("page_num", "page_size", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := fmt.Sprintf(v3MarketingActivityQueryAppBatchQuery, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityQueryAppBatchQueryRsp{StatusCode: res.StatusCode}
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

// 查询活动可用门店 alipay.marketing.activity.shop.batchquery
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityQueryShopBatchQuery(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryShopBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("page_num", "page_size")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := fmt.Sprintf(v3MarketingActivityQueryShopBatchQuery, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityQueryShopBatchQueryRsp{StatusCode: res.StatusCode}
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

// 查询活动适用商品 alipay.marketing.activity.goods.batchquery
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityQueryGoodsBatchQuery(ctx context.Context, activityId string, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryGoodsBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("page_num", "page_size")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := fmt.Sprintf(v3MarketingActivityQueryGoodsBatchQuery, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityQueryGoodsBatchQueryRsp{StatusCode: res.StatusCode}
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
