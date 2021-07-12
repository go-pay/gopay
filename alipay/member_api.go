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
