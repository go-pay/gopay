package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// datadigital.fincloud.generalsaas.face.verification.initialize(人脸核身初始化)
// 文档地址：https://opendocs.alipay.com/open/04jg6r
func (a *Client) FaceVerificationInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceVerificationInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code", "identity_type", "cert_type", "cert_name", "cert_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.verification.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceVerificationInitializeRsp)
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

// datadigital.fincloud.generalsaas.face.verification.query(人脸核身结果查询)
// 文档地址：https://opendocs.alipay.com/open/04jg6s
func (a *Client) FaceVerificationQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceVerificationQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.verification.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceVerificationQueryRsp)
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

// datadigital.fincloud.generalsaas.face.certify.initialize(H5人脸核身初始化)
// 文档地址：https://opendocs.alipay.com/open/02zloa
func (a *Client) FaceCertifyInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCertifyInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "identity_param", "merchant_config")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.certify.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceCertifyInitializeRsp)
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

// datadigital.fincloud.generalsaas.face.certify.verify(H5人脸核身开始认证)
// 文档地址：https://opendocs.alipay.com/open/02zlob
func (a *Client) FaceCertifyVerify(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCertifyVerifyRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.certify.verify"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceCertifyVerifyRsp)
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

// datadigital.fincloud.generalsaas.face.certify.query(H5人脸核身查询记录)
// 文档地址：https://opendocs.alipay.com/open/02zloc
func (a *Client) FaceCertifyQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCertifyQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.certify.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceCertifyQueryRsp)
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

// datadigital.fincloud.generalsaas.face.source.certify(纯服务端人脸核身)
// 文档地址：https://opendocs.alipay.com/open/04pxq6
func (a *Client) FaceSourceCertify(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceSourceCertifyRsp, err error) {
	err = bm.CheckEmptyError("outer_biz_no", "cert_type", "cert_no", "cert_name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.source.certify"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceSourceCertifyRsp)
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

// datadigital.fincloud.generalsaas.face.check.initialize(活体检测初始化)
// 文档地址：https://opendocs.alipay.com/open/03nisu
func (a *Client) FaceCheckInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCheckInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.check.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceCheckInitializeRsp)
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

// datadigital.fincloud.generalsaas.face.check.query(活体检测结果查询)
// 文档地址：https://opendocs.alipay.com/open/03nisv
func (a *Client) FaceCheckQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCheckQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.face.check.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FaceCheckQueryRsp)
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
