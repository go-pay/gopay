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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
