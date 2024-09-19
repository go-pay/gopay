package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.merchant.qipan.crowd.create(上传创建人群)
// 文档地址：https://opendocs.alipay.com/open/e93d9a54_alipay.merchant.qipan.crowd.create
func (a *Client) MerchantQipanCrowdCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MerchantQipanCrowdCreateRsp, err error) {
	err = bm.CheckEmptyError("crowd_name", "external_crowd_code", "user_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.crowd.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MerchantQipanCrowdCreateRsp)
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

// alipay.merchant.qipan.crowduser.add(人群中追加用户)
// 文档地址：https://opendocs.alipay.com/open/04330914_alipay.merchant.qipan.crowduser.add
func (a *Client) MerchantQipanCrowdUserAdd(ctx context.Context, bm gopay.BodyMap) (aliRsp *MerchantQipanCrowdUserAddRsp, err error) {
	err = bm.CheckEmptyError("crowd_code", "user_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.crowduser.add"); err != nil {
		return nil, err
	}
	aliRsp = new(MerchantQipanCrowdUserAddRsp)
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

// alipay.merchant.qipan.crowduser.delete(人群中删除用户)
// 文档地址：https://opendocs.alipay.com/open/80646633_alipay.merchant.qipan.crowduser.delete
func (a *Client) MerchantQipanCrowdUserDelete(ctx context.Context, bm gopay.BodyMap) (aliRsp *MerchantQipanCrowdUserDeleteRsp, err error) {
	err = bm.CheckEmptyError("crowd_code", "user_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.crowduser.delete"); err != nil {
		return nil, err
	}
	aliRsp = new(MerchantQipanCrowdUserDeleteRsp)
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

// alipay.marketing.qipan.tagbase.batchquery(棋盘人群圈选标签基本信息查询)
// 文档地址：https://opendocs.alipay.com/open/ce6aee00_alipay.marketing.qipan.tagbase.batchquery
func (a *Client) MarketingQipanTagBaseBatchQuery(ctx context.Context) (aliRsp *MarketingQipanTagBaseBatchQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, nil, "alipay.marketing.qipan.tagbase.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanTagBaseBatchQueryRsp)
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

// alipay.marketing.qipan.tag.query(棋盘标签圈选值查询)
// 文档地址：https://opendocs.alipay.com/open/e322cb35_alipay.marketing.qipan.tag.query
func (a *Client) MarketingQipanTagQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanTagQueryRsp, err error) {
	err = bm.CheckEmptyError("tag_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.qipan.tag.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanTagQueryRsp)
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

// alipay.marketing.qipan.crowdoperation.create(棋盘人群创建)
// 文档地址：https://opendocs.alipay.com/open/09c10677_alipay.marketing.qipan.crowdoperation.create
func (a *Client) MarketingQipanCrowdOperationCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanCrowdOperationCreateRsp, err error) {
	err = bm.CheckEmptyError("crowd_name", "operation_pool_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.qipan.crowdoperation.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdOperationCreateRsp)
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

// alipay.marketing.qipan.crowdtag.query(查询圈选标签列表)
// 文档地址：https://opendocs.alipay.com/open/8e411cec_alipay.marketing.qipan.crowdtag.query
func (a *Client) MarketingQipanCrowdTagQuery(ctx context.Context) (aliRsp *MarketingQipanCrowdTagQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, nil, "alipay.marketing.qipan.crowdtag.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdTagQueryRsp)
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

// alipay.marketing.qipan.crowdwithtag.create(标签圈选创建人群)
// 文档地址：https://opendocs.alipay.com/open/cbf0efa4_alipay.marketing.qipan.crowdwithtag.create
func (a *Client) MarketingQipanCrowdWithTagCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanCrowdWithTagCreateRsp, err error) {
	err = bm.CheckEmptyError("crowd_name", "apply_channel_list", "select_tag_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.qipan.crowdwithtag.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdWithTagCreateRsp)
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

// alipay.marketing.qipan.crowdwithtag.query(标签圈选预估人群规模)
// 文档地址：https://opendocs.alipay.com/open/adb5dd04_alipay.marketing.qipan.crowdwithtag.query
func (a *Client) MarketingQipanCrowdWithTagQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanCrowdWithTagQueryRsp, err error) {
	err = bm.CheckEmptyError("crowd_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.qipan.crowdwithtag.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdWithTagQueryRsp)
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

// alipay.merchant.qipan.crowd.batchquery(查询人群列表)
// 文档地址：https://opendocs.alipay.com/open/b7f3caec_alipay.merchant.qipan.crowd.batchquery
func (a *Client) MarketingQipanCrowdBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanCrowdBatchQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.crowd.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdBatchQueryRsp)
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

// alipay.merchant.qipan.crowd.query(查询人群详情)
// 文档地址：https://opendocs.alipay.com/open/be384367_alipay.merchant.qipan.crowd.query
func (a *Client) MarketingQipanCrowdQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanCrowdQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.crowd.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdQueryRsp)
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

// alipay.merchant.qipan.crowd.modify(修改人群)
// 文档地址：https://opendocs.alipay.com/open/6c8ebb31_alipay.merchant.qipan.crowd.modify
func (a *Client) MarketingQipanCrowdModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanCrowdModifyRsp, err error) {
	err = bm.CheckEmptyError("crowd_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.crowd.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanCrowdModifyRsp)
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

// alipay.merchant.qipan.board.query(看板分析)
// 文档地址：https://opendocs.alipay.com/open/aa5066aa_alipay.merchant.qipan.board.query
func (a *Client) MarketingQipanBoardQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanBoardQueryRsp, err error) {
	err = bm.CheckEmptyError("scene_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.board.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanBoardQueryRsp)
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

// alipay.merchant.qipan.insight.query(画像分析)
// 文档地址：https://opendocs.alipay.com/open/46a75f4e_alipay.merchant.qipan.insight.query
func (a *Client) MarketingQipanInsightQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanInsightQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.insight.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanInsightQueryRsp)
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

// alipay.merchant.qipan.behavior.query(行为分析)
// 文档地址：https://opendocs.alipay.com/open/56735ac5_alipay.merchant.qipan.behavior.query
func (a *Client) MarketingQipanBehaviorQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanBehaviorQueryRsp, err error) {
	err = bm.CheckEmptyError("scene_code", "request_params")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.behavior.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanBehaviorQueryRsp)
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

// alipay.merchant.qipan.trend.query(趋势分析)
// 文档地址：https://opendocs.alipay.com/open/8ee7795a_alipay.merchant.qipan.trend.query
func (a *Client) MarketingQipanTrendQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanTrendQueryRsp, err error) {
	err = bm.CheckEmptyError("index_key", "request_params")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.trend.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanTrendQueryRsp)
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

// alipay.merchant.qipan.insightcity.query(常住省市查询)
// 文档地址：https://opendocs.alipay.com/open/f7d99821_alipay.merchant.qipan.insightcity.query
func (a *Client) MarketingQipanInsightCityQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingQipanInsightCityQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.qipan.insightcity.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingQipanInsightCityQueryRsp)
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
