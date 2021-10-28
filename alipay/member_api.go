package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// alipay.user.info.share(支付宝会员授权信息查询接口)
//	body：此接口无需body参数
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.info.share
func (a *Client) UserInfoShare(ctx context.Context, authToken string) (aliRsp *UserInfoShareResponse, err error) {
	if authToken == "" {
		return nil, errors.New("auth_token can not be null")
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, nil, "alipay.user.info.share", authToken); err != nil {
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
func (a *Client) UserCertifyOpenInit(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCertifyOpenInitResponse, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code", "identity_param", "merchant_config")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.certify.open.initialize"); err != nil {
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
func (a *Client) UserCertifyOpenCertify(ctx context.Context, bm gopay.BodyMap) (certifyUrl string, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return util.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.certify.open.certify"); err != nil {
		return util.NULL, err
	}
	certifyUrl = string(bs)
	return certifyUrl, nil
}

// alipay.user.certify.open.query(身份认证记录查询)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.query
func (a *Client) UserCertifyOpenQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCertifyOpenQueryResponse, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.certify.open.query"); err != nil {
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
func (a *Client) UserAgreementPageSign(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementPageSignRsp, err error) {
	err = bm.CheckEmptyError("personal_product_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.agreement.page.sign"); err != nil {
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
func (a *Client) UserAgreementPageUnSign(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementPageUnSignRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.agreement.unsign"); err != nil {
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
func (a *Client) UserAgreementQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.agreement.query"); err != nil {
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
func (a *Client) UserAgreementExecutionplanModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementExecutionplanModifyRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "deduct_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.agreement.executionplan.modify"); err != nil {
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
func (a *Client) UserAgreementTransfer(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAgreementTransferRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "target_product_code", "period_rule_params")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.agreement.transfer"); err != nil {
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
func (a *Client) UserTwostageCommonUse(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserTwostageCommonUseRsp, err error) {
	err = bm.CheckEmptyError("dynamic_id", "sence_no", "pay_pid")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.twostage.common.use"); err != nil {
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
func (a *Client) UserAuthZhimaorgIdentityApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAuthZhimaorgIdentityApplyRsp, err error) {
	err = bm.CheckEmptyError("cert_type", "cert_no", "name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.auth.zhimaorg.identity.apply"); err != nil {
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
func (a *Client) UserCharityRecordexistQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCharityRecordexistQueryRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.charity.recordexist.query"); err != nil {
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
func (a *Client) UserAlipaypointSend(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAlipaypointSendRsp, err error) {
	err = bm.CheckEmptyError("budget_code", "partner_biz_no", "point_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.alipaypoint.send"); err != nil {
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
func (a *Client) MemberDataIsvCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MemberDataIsvCreateRsp, err error) {
	err = bm.CheckEmptyError("member_card_id", "member_source", "member_status", "gmt_merber_card_create", "parter_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.member.data.isv.create"); err != nil {
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
func (a *Client) UserFamilyArchiveQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserFamilyArchiveQueryRsp, err error) {
	err = bm.CheckEmptyError("archive_token")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.family.archive.query"); err != nil {
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

// alipay.user.family.archive.initialize(初始化家人信息档案(选人授权)组件)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.family.archive.initialize
func (a *Client) UserFamilyArchiveInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserFamilyArchiveInitializeRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "template_id", "redirect_uri")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.family.archive.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(UserFamilyArchiveInitializeRsp)
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

// alipay.user.certdoc.certverify.preconsult(实名证件信息比对验证预咨询)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certdoc.certverify.preconsult
func (a *Client) UserCertdocCertverifyPreconsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCertdocCertverifyPreconsultRsp, err error) {
	err = bm.CheckEmptyError("user_name", "cert_type", "cert_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.certdoc.certverify.preconsult"); err != nil {
		return nil, err
	}
	aliRsp = new(UserCertdocCertverifyPreconsultRsp)
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

// alipay.user.certdoc.certverify.consult(实名证件信息比对验证咨询)
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certdoc.certverify.consult
func (a *Client) UserCertdocCertverifyConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCertdocCertverifyConsultRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.certdoc.certverify.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(UserCertdocCertverifyConsultRsp)
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

// alipay.user.family.share.zmgo.initialize(初始化家庭芝麻GO共享组件)
//	文档地址：https://opendocs.alipay.com/apis/01n4yx
func (a *Client) UserFamilyShareZmgoInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserFamilyShareZmgoInitializeRsp, err error) {
	err = bm.CheckEmptyError("user_id", "scene_id", "template_id", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.family.share.zmgo.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(UserFamilyShareZmgoInitializeRsp)
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

// alipay.user.dtbank.qrcodedata.query(数字分行银行码明细数据查询)
//	文档地址：https://opendocs.alipay.com/apis/01ozks
func (a *Client) UserDtbankQrcodedataQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserDtbankQrcodedataQueryRsp, err error) {
	err = bm.CheckEmptyError("data_date", "qrcode_id", "qrcode_out_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.dtbank.qrcodedata.query"); err != nil {
		return nil, err
	}
	aliRsp = new(UserDtbankQrcodedataQueryRsp)
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

// alipay.user.alipaypoint.budgetlib.query(查询集分宝预算库详情)
//	文档地址：https://opendocs.alipay.com/apis/01zrby
func (a *Client) UserAlipaypointBudgetlibQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAlipaypointBudgetlibQueryRsp, err error) {
	err = bm.CheckEmptyError("budget_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.user.alipaypoint.budgetlib.query"); err != nil {
		return nil, err
	}
	aliRsp = new(UserAlipaypointBudgetlibQueryRsp)
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
