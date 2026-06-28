package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// 扣费服务（预约扣费）—— 小程序预签约
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012525209
// 必填字段：appid, openid, plan_id (integer), out_contract_code, contract_display_account, contract_notify_url
// 选填字段：out_user_code, deduct_schedule{estimated_deduct_date, estimated_deduct_amount{total, currency}, description}
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductPreSignMiniProgram(ctx context.Context, bm gopay.BodyMap) (wxRsp *PapayScheduledPreSignMiniProgramRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScheduledDeductPreSignMiniProgram, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScheduledDeductPreSignMiniProgram, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledPreSignMiniProgramRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledPreSignMiniProgram)}
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

// 扣费服务（预约扣费）—— APP 预签约
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012524934
// 必填字段：appid, plan_id (integer), out_contract_code, contract_display_account, contract_notify_url
// 选填字段：out_user_code, deduct_schedule{...}
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductPreSignApp(ctx context.Context, bm gopay.BodyMap) (wxRsp *PapayScheduledPreSignAppRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScheduledDeductPreSignApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScheduledDeductPreSignApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledPreSignAppRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledPreSignApp)}
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

// 扣费服务（预约扣费）—— H5 预签约
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012489208
// 必填字段：appid, plan_id (integer), out_contract_code, contract_display_account, contract_notify_url
// 选填字段：out_user_code, deduct_schedule{...}
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductPreSignH5(ctx context.Context, bm gopay.BodyMap) (wxRsp *PapayScheduledPreSignH5Rsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScheduledDeductPreSignH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScheduledDeductPreSignH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledPreSignH5Rsp{Code: Success, SignInfo: si, Response: new(PapayScheduledPreSignH5)}
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

// 扣费服务（预约扣费）—— JSAPI 预签约
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012525133
// 必填字段：appid, openid, plan_id (integer), out_contract_code, contract_display_account, contract_notify_url
// 选填字段：out_user_code, deduct_schedule{...}
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductPreSignJsapi(ctx context.Context, bm gopay.BodyMap) (wxRsp *PapayScheduledPreSignJsapiRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScheduledDeductPreSignJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScheduledDeductPreSignJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledPreSignJsapiRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledPreSignJsapi)}
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

// 扣费服务（预约扣费）—— 通过商户协议号查询签约
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012489245
// 路径参数：plan_id（委托代扣模板 ID）、out_contract_code（商户协议号）
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductContractQuery(ctx context.Context, planId, outContractCode string) (wxRsp *PapayScheduledContractRsp, err error) {
	uri := fmt.Sprintf(v3ScheduledDeductContractQuery, planId, outContractCode)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledContractRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledContract)}
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

// 扣费服务（预约扣费）—— 通过商户协议号解约
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012489295
// 路径参数：plan_id、out_contract_code
// 必填字段：contract_termination_remark（解约备注）
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductContractTerminate(ctx context.Context, planId, outContractCode string, bm gopay.BodyMap) (wxRsp *PapayScheduledContractRsp, err error) {
	uri := fmt.Sprintf(v3ScheduledDeductContractTerminate, planId, outContractCode)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledContractRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledContract)}
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

// 扣费服务（预约扣费）—— 创建预约扣费
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012467036
// 路径参数：contract_id（委托代扣协议 ID）
// 必填字段：appid, schedule_amount{total, currency}
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductSchedule(ctx context.Context, contractId string, bm gopay.BodyMap) (wxRsp *PapayScheduledScheduleRsp, err error) {
	uri := fmt.Sprintf(v3ScheduledDeductSchedule, contractId)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledScheduleRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledSchedule)}
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

// 扣费服务（预约扣费）—— 查询预约扣费结果
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012466997
// 路径参数：contract_id
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductScheduleQuery(ctx context.Context, contractId string) (wxRsp *PapayScheduledScheduleRsp, err error) {
	uri := fmt.Sprintf(v3ScheduledDeductScheduleQuery, contractId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledScheduleRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledSchedule)}
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

// 扣费服务（预约扣费）—— 受理扣款
// 微信文档：https://pay.weixin.qq.com/doc/v3/merchant/4012467087
// 必填字段：appid, out_trade_no, description, transaction_notify_url, contract_id, amount{total, currency}
// 选填字段：goods_tag, attach
// Code = 0 is success
func (c *ClientV3) V3ScheduledDeductApply(ctx context.Context, bm gopay.BodyMap) (wxRsp *PapayScheduledApplyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3ScheduledDeductApply, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ScheduledDeductApply, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PapayScheduledApplyRsp{Code: Success, SignInfo: si, Response: new(PapayScheduledApply)}
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
