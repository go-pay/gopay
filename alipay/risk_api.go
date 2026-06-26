package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.security.risk.complaint.process.finish(处理消费者投诉)
// 文档地址：https://opendocs.alipay.com/open/da75e1ec_alipay.security.risk.complaint.process.finish
func (a *Client) SecurityRiskComplaintProcessFinish(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskComplaintProcessFinishResponse, err error) {
	err = bm.CheckEmptyError("id_list", "process_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.complaint.process.finish"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskComplaintProcessFinishResponse)
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

// alipay.security.risk.complaint.file.upload(投诉处理附件图片上传)
// 文档地址：https://opendocs.alipay.com/open/20ea7441_alipay.security.risk.complaint.file.upload
func (a *Client) SecurityRiskComplaintFileUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskComplaintFileUploadResponse, err error) {
	err = bm.CheckEmptyError("file_content")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.FileUploadRequest(ctx, bm, "alipay.security.risk.complaint.file.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskComplaintFileUploadResponse)
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

// alipay.security.risk.complaint.info.query(查询消费者投诉详情)
// 文档地址：https://opendocs.alipay.com/open/271499b9_alipay.security.risk.complaint.info.query
func (a *Client) SecurityRiskComplaintInfoQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskComplaintInfoQueryResponse, err error) {
	err = bm.CheckEmptyError("complaint_ids")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.complaint.info.query"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskComplaintInfoQueryResponse)
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

// alipay.security.risk.complaint.info.batchquery(查询消费者投诉列表)
// 文档地址：https://opendocs.alipay.com/open/8ad1ac86_alipay.security.risk.complaint.info.batchquery
func (a *Client) SecurityRiskComplaintInfoBatchquery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskComplaintInfoBatchqueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.complaint.info.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskComplaintInfoBatchqueryResponse)
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

// alipay.security.risk.marketing.awarding.query(营销风险识别发奖)
// 文档地址：https://opendocs.alipay.com/open/f4427923_alipay.security.risk.marketing.awarding.query
func (a *Client) SecurityRiskMarketingAwardingQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskMarketingAwardingQueryResponse, err error) {
	err = bm.CheckEmptyError("role", "risk_type", "service", "merchant_scene", "scene", "business_code", "channel")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.marketing.awarding.query"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskMarketingAwardingQueryResponse)
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

// alipay.security.risk.marketing.purchase.query(营销风险识别抢购)
// 文档地址：https://opendocs.alipay.com/open/91f83d97_alipay.security.risk.marketing.purchase.query
func (a *Client) SecurityRiskMarketingPurchaseQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskMarketingPurchaseQueryResponse, err error) {
	err = bm.CheckEmptyError("role", "risk_type", "service", "merchant_scene", "scene", "business_code", "channel")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.marketing.purchase.query"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskMarketingPurchaseQueryResponse)
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

// alipay.security.risk.industry.scalper.query(行业风险识别黄牛)
// 文档地址：https://opendocs.alipay.com/open/5e344142_alipay.security.risk.industry.scalper.query
func (a *Client) SecurityRiskIndustryScalperQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskIndustryScalperQueryResponse, err error) {
	err = bm.CheckEmptyError("role", "risk_type", "service", "merchant_scene", "scene", "business_code", "channel")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.industry.scalper.query"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskIndustryScalperQueryResponse)
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

// alipay.security.risk.industry.farming.query(行业风险识别刷单)
// 文档地址：https://opendocs.alipay.com/open/e76efd50_alipay.security.risk.industry.farming.query
func (a *Client) SecurityRiskIndustryFarmingQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskIndustryFarmingQueryResponse, err error) {
	err = bm.CheckEmptyError("role", "risk_type", "service", "merchant_scene", "scene", "business_code", "channel")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.industry.farming.query"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskIndustryFarmingQueryResponse)
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

// alipay.security.risk.industry.nsf.query(行业风险识别先享后付违约)
// 文档地址：https://opendocs.alipay.com/open/399e7ee9_alipay.security.risk.industry.nsf.query
func (a *Client) SecurityRiskIndustryNsfQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskIndustryNsfQueryResponse, err error) {
	err = bm.CheckEmptyError("role", "risk_type", "service", "merchant_scene", "scene", "business_code", "channel")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.industry.nsf.query"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskIndustryNsfQueryResponse)
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

// alipay.security.risk.content.sync.detect(内容风险同步识别)
// 文档地址：https://opendocs.alipay.com/open/8513019b_alipay.security.risk.content.sync.detect
func (a *Client) SecurityRiskContentSyncDetect(ctx context.Context, bm gopay.BodyMap) (aliRsp *SecurityRiskContentSyncDetectResponse, err error) {
	err = bm.CheckEmptyError("request_id", "products", "channel", "content_type", "data_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.security.risk.content.sync.detect"); err != nil {
		return nil, err
	}
	aliRsp = new(SecurityRiskContentSyncDetectResponse)
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
