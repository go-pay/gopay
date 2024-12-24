package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 支付宝个人代扣协议查询接口 alipay.user.agreement.query
// StatusCode = 200 is success
func (a *ClientV3) UserAgreementQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementQueryRsp, err error) {
	if bm.GetString("alipay_user_id") == gopay.NULL && bm.GetString("alipay_open_id") == gopay.NULL {
		return nil, errors.New("alipay_user_id and alipay_open_id are not allowed to be null at the same time")
	}
	uri := v3UserAgreementQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserAgreementQueryRsp{StatusCode: res.StatusCode}
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

// 支付宝个人代扣协议解约接口 alipay.user.agreement.unsign
// StatusCode = 200 is success
func (a *ClientV3) UserAgreementPageUnSign(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementPageUnSignRsp, err error) {
	if bm.GetString("alipay_user_id") == gopay.NULL && bm.GetString("alipay_open_id") == gopay.NULL && bm.GetString("alipay_logon_id") == gopay.NULL {
		return nil, errors.New("alipay_user_id and alipay_open_id and alipay_logon_id are not allowed to be null at the same time")
	}
	if bm.GetString("external_agreement_no") == gopay.NULL && bm.GetString("agreement_no") == gopay.NULL {
		return nil, errors.New("external_agreement_no and agreement_no are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3UserAgreementPageUnSign, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3UserAgreementPageUnSign, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserAgreementPageUnSignRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 分账关系绑定 alipay.trade.royalty.relation.bind
// StatusCode = 200 is success
func (a *ClientV3) TradeRelationBind(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationBindRsp, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradeRoyaltyRelationBind, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeRoyaltyRelationBind, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeRelationBindRsp{StatusCode: res.StatusCode}
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

// 分账关系解绑 alipay.trade.royalty.relation.unbind
// StatusCode = 200 is success
func (a *ClientV3) TradeRelationUnbind(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationUnbindRsp, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradeRoyaltyRelationUnbind, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeRoyaltyRelationUnbind, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeRelationUnbindRsp{StatusCode: res.StatusCode}
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

// 分账关系查询 alipay.trade.royalty.relation.batchquery
// StatusCode = 200 is success
func (a *ClientV3) TradeRelationBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationBatchQueryRsp, err error) {
	authorization, err := a.authorization(MethodPost, v3TradeRoyaltyRelationBatchQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeRoyaltyRelationBatchQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeRelationBatchQueryRsp{StatusCode: res.StatusCode}
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

// 分账比例查询 alipay.trade.royalty.rate.query
// StatusCode = 200 is success
func (a *ClientV3) TradeRoyaltyRateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRoyaltyRateQueryRsp, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	uri := v3TradeRoyaltyRateQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeRoyaltyRateQueryRsp{StatusCode: res.StatusCode}
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

// 统一收单交易结算接口 alipay.trade.order.settle
// StatusCode = 200 is success
func (a *ClientV3) TradeOrderSettle(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderSettleRsp, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "royalty_parameters")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradeOrderSettle, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeOrderSettle, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeOrderSettleRsp{StatusCode: res.StatusCode}
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

// 交易分账查询接口 alipay.trade.order.settle.query
// StatusCode = 200 is success
func (a *ClientV3) TradeOrderSettleQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderSettleQueryRsp, err error) {
	uri := v3TradeOrderSettleQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeOrderSettleQueryRsp{StatusCode: res.StatusCode}
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

// 分账剩余金额查询 alipay.trade.order.onsettle.query
// StatusCode = 200 is success
func (a *ClientV3) TradeOrderOnSettleQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderOnSettleQueryRsp, err error) {
	err = bm.CheckEmptyError("trade_no")
	if err != nil {
		return nil, err
	}
	uri := v3TradeOrderOnSettleQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeOrderOnSettleQueryRsp{StatusCode: res.StatusCode}
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
