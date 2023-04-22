package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.commerce.transport.nfccard.send(NFC用户卡信息同步)
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.commerce.transport.nfccard.send
func (a *Client) CommerceTransportNfccardSend(ctx context.Context, bm gopay.BodyMap) (aliRsp *CommerceTransportNfccardSendRsp, err error) {
	err = bm.CheckEmptyError("issue_org_no", "card_no", "card_status")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.transport.nfccard.send"); err != nil {
		return nil, err
	}
	aliRsp = new(CommerceTransportNfccardSendRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.commerce.air.callcenter.trade.apply(航司电话订票待申请接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.commerce.air.callcenter.trade.apply
func (a *Client) CommerceAirCallcenterTradeApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *CommerceAirCallcenterTradeApplyRsp, err error) {
	err = bm.CheckEmptyError("scene_code", "op_code", "channel", "target_id", "target_id_type", "trade_apply_params")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.air.callcenter.trade.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(CommerceAirCallcenterTradeApplyRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.commerce.operation.gamemarketing.benefit.apply(申请权益发放)
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.commerce.operation.gamemarketing.benefit.apply
func (a *Client) CommerceBenefitApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *CommerceBenefitApplyRsp, err error) {
	err = bm.CheckEmptyError("activity_code", "trade_no", "user_account", "platform")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.operation.gamemarketing.benefit.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(CommerceBenefitApplyRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.commerce.operation.gamemarketing.benefit.verify(权益核销)
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.commerce.operation.gamemarketing.benefit.verify
func (a *Client) CommerceBenefitVerify(ctx context.Context, bm gopay.BodyMap) (aliRsp *CommerceBenefitVerifyRsp, err error) {
	err = bm.CheckEmptyError("activity_code", "voucher_code", "user_account", "trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.operation.gamemarketing.benefit.verify"); err != nil {
		return nil, err
	}
	aliRsp = new(CommerceBenefitVerifyRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
