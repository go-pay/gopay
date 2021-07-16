package alipay

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// alipay.user.info.share(支付宝会员授权信息查询接口)
//	body：此接口无需body参数
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.info.share
func (a *Client) UserInfoShare(authToken string) (aliRsp *UserInfoShareResponse, err error) {
	if authToken == "" {
		return nil, errors.New("auth_token can not be null")
	}
	var bs []byte
	if bs, err = a.doAliPay(nil, "alipay.user.info.share", authToken); err != nil {
		return nil, err
	}
	aliRsp = new(UserInfoShareResponse)
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

// alipay.user.certify.open.initialize(身份认证初始化服务)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.initialize
func (a *Client) UserCertifyOpenInit(bm gopay.BodyMap) (aliRsp *UserCertifyOpenInitResponse, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code", "identity_param", "merchant_config")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(UserCertifyOpenInitResponse)
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

// alipay.user.certify.open.certify(身份认证开始认证)
//	API文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.certify
//	产品文档地址：https://opendocs.alipay.com/open/20181012100420932508/quickstart
func (a *Client) UserCertifyOpenCertify(bm gopay.BodyMap) (certifyUrl string, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return util.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.certify"); err != nil {
		return util.NULL, err
	}
	certifyUrl = string(bs)
	return certifyUrl, nil
}

// alipay.user.certify.open.query(身份认证记录查询)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.query
func (a *Client) UserCertifyOpenQuery(bm gopay.BodyMap) (aliRsp *UserCertifyOpenQueryResponse, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.query"); err != nil {
		return nil, err
	}
	aliRsp = new(UserCertifyOpenQueryResponse)
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

// alipay.user.agreement.page.sign(支付宝个人协议页面签约接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.agreement.page.sign
func (a *Client) UserAgreementPageSign(bm gopay.BodyMap) (aliRsp *UserAgreementPageSignRsp, err error) {
	err = bm.CheckEmptyError("personal_product_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.agreement.page.sign"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAgreementPageSignRsp)
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

// alipay.user.agreement.unsign(支付宝个人代扣协议解约接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.agreement.page.unsign
func (a *Client) UserAgreementPageUnSign(bm gopay.BodyMap) (aliRsp *UserAgreementPageUnSignRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.agreement.unsign"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAgreementPageUnSignRsp)
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

// alipay.user.agreement.query(支付宝个人代扣协议查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.agreement.query
func (a *Client) UserAgreementQuery(bm gopay.BodyMap) (aliRsp *UserAgreementQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.agreement.query"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAgreementQueryRsp)
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

// alipay.user.agreement.executionplan.modify(周期性扣款协议执行计划修改接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.agreement.executionplan.modify
func (a *Client) UserAgreementExecutionplanModify(bm gopay.BodyMap) (aliRsp *UserAgreementExecutionplanModifyRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "deduct_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.agreement.executionplan.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAgreementExecutionplanModifyRsp)
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

// alipay.user.agreement.transfer(协议由普通通用代扣协议产品转移到周期扣协议产品)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.agreement.transfer
func (a *Client) UserAgreementTransfer(bm gopay.BodyMap) (aliRsp *UserAgreementTransferRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "target_product_code", "period_rule_params")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.agreement.transfer"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAgreementTransferRsp)
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

// alipay.user.twostage.common.use(通用当面付二阶段接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.twostage.common.use
func (a *Client) UserTwostageCommonUse(bm gopay.BodyMap) (aliRsp *UserTwostageCommonUseRsp, err error) {
	err = bm.CheckEmptyError("dynamic_id", "sence_no", "pay_pid")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.twostage.common.use"); err != nil {
		return nil, err
	}
	aliRsp = new(UserTwostageCommonUseRsp)
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

// alipay.user.auth.zhimaorg.identity.apply(芝麻企业征信基于身份的协议授权)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.auth.zhimaorg.identity.apply
func (a *Client) UserAuthZhimaorgIdentityApply(bm gopay.BodyMap) (aliRsp *UserAuthZhimaorgIdentityApplyRsp, err error) {
	err = bm.CheckEmptyError("cert_type", "cert_no", "name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.auth.zhimaorg.identity.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAuthZhimaorgIdentityApplyRsp)
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

// alipay.user.charity.recordexist.query(查询是否在支付宝公益捐赠的接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.charity.recordexist.query
func (a *Client) UserCharityRecordexistQuery(bm gopay.BodyMap) (aliRsp *UserCharityRecordexistQueryRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.charity.recordexist.query"); err != nil {
		return nil, err
	}
	aliRsp = new(UserCharityRecordexistQueryRsp)
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

// alipay.user.alipaypoint.send(集分宝发放接口)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.alipaypoint.send
func (a *Client) UserAlipaypointSend(bm gopay.BodyMap) (aliRsp *UserAlipaypointSendRsp, err error) {
	err = bm.CheckEmptyError("budget_code", "partner_biz_no", "point_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.alipaypoint.send"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAlipaypointSendRsp)
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

// koubei.member.data.isv.create(isv 会员CRM数据回流)
//	文档地址：https://opendocs.alipay.com/apis/api_2/koubei.member.data.isv.create
func (a *Client) MemberDataIsvCreate(bm gopay.BodyMap) (aliRsp *MemberDataIsvCreateRsp, err error) {
	err = bm.CheckEmptyError("member_card_id", "member_source", "member_status", "gmt_merber_card_create", "parter_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "koubei.member.data.isv.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MemberDataIsvCreateRsp)
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

// alipay.user.family.archive.query(查询家人信息档案(选人授权)组件已选的家人档案信息)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.family.archive.query
func (a *Client) UserFamilyArchiveQuery(bm gopay.BodyMap) (aliRsp *UserFamilyArchiveQueryRsp, err error) {
	err = bm.CheckEmptyError("archive_token")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.family.archive.query"); err != nil {
		return nil, err
	}
	aliRsp = new(UserFamilyArchiveQueryRsp)
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
