package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 人脸核身初始化 datadigital.fincloud.generalsaas.face.verification.initialize
// StatusCode = 200 is success
func (a *ClientV3) FaceVerificationInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceVerificationInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code", "identity_type", "cert_type", "cert_name", "cert_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FaceVerificationInitialize, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FaceVerificationInitialize, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceVerificationInitializeRsp{StatusCode: res.StatusCode}
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

// 人脸核身结果查询 datadigital.fincloud.generalsaas.face.verification.query
// StatusCode = 200 is success
func (a *ClientV3) FaceVerificationQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceVerificationQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	uri := v3FaceVerificationQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceVerificationQueryRsp{StatusCode: res.StatusCode}
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

// 跳转支付宝人脸核身初始化 datadigital.fincloud.generalsaas.face.certify.initialize
// StatusCode = 200 is success
func (a *ClientV3) FaceCertifyInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCertifyInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "identity_param", "merchant_config")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FaceCertifyInitialize, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FaceCertifyInitialize, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceCertifyInitializeRsp{StatusCode: res.StatusCode}
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

// 跳转支付宝人脸核身开始认证 datadigital.fincloud.generalsaas.face.certify.verify
// StatusCode = 200 is success
func (a *ClientV3) FaceCertifyVerify(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCertifyVerifyRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FaceCertifyVerify, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FaceCertifyVerify, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceCertifyVerifyRsp{StatusCode: res.StatusCode}
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

// 跳转支付宝人脸核身查询记录 datadigital.fincloud.generalsaas.face.certify.query
// StatusCode = 200 is success
func (a *ClientV3) FaceCertifyQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCertifyQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	uri := v3FaceCertifyQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceCertifyQueryRsp{StatusCode: res.StatusCode}
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

// 纯服务端人脸核身 datadigital.fincloud.generalsaas.face.source.certify
// StatusCode = 200 is success
func (a *ClientV3) FaceSourceCertify(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceSourceCertifyRsp, err error) {
	err = bm.CheckEmptyError("outer_biz_no", "cert_type", "cert_no", "cert_name")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FaceSourceCertify, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FaceSourceCertify, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceSourceCertifyRsp{StatusCode: res.StatusCode}
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

// 活体检测初始化 datadigital.fincloud.generalsaas.face.check.initialize
// StatusCode = 200 is success
func (a *ClientV3) FaceCheckInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCheckInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3FaceCheckInitialize, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3FaceCheckInitialize, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceCheckInitializeRsp{StatusCode: res.StatusCode}
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

// 活体检测结果查询 datadigital.fincloud.generalsaas.face.check.query
// StatusCode = 200 is success
func (a *ClientV3) FaceCheckQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FaceCheckQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	uri := v3FaceCheckQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &FaceCheckQueryRsp{StatusCode: res.StatusCode}
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

// 身份证二要素核验 datadigital.fincloud.generalsaas.twometa.check
// StatusCode = 200 is success
func (a *ClientV3) IDCardTwoMetaCheck(ctx context.Context, bm gopay.BodyMap) (aliRsp *IDCardTwoMetaCheckRsp, err error) {
	err = bm.CheckEmptyError("outer_biz_no", "cert_name", "cert_no", "cert_type")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3IDCardTwoMetaCheck, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3IDCardTwoMetaCheck, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &IDCardTwoMetaCheckRsp{StatusCode: res.StatusCode}
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

// 银行卡核验 datadigital.fincloud.generalsaas.bankcard.check
// StatusCode = 200 is success
func (a *ClientV3) BankCardCheck(ctx context.Context, bm gopay.BodyMap) (aliRsp *BankCardCheckRsp, err error) {
	err = bm.CheckEmptyError("outer_biz_no", "product_type", "cert_name", "bankcard_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3BankCardCheck, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3BankCardCheck, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &BankCardCheckRsp{StatusCode: res.StatusCode}
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

// 手机号三要素核验简版 datadigital.fincloud.generalsaas.mobilethreemeta.simple.check
// StatusCode = 200 is success
func (a *ClientV3) MobileThreeMetaSimpleCheck(ctx context.Context, bm gopay.BodyMap) (aliRsp *MobileThreeMetaSimpleCheckRsp, err error) {
	err = bm.CheckEmptyError("outer_biz_no", "cert_name", "cert_no", "phone")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MobileThreeMetaSimpleCheck, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MobileThreeMetaSimpleCheck, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MobileThreeMetaSimpleCheckRsp{StatusCode: res.StatusCode}
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

// 手机号三要素核验详版 datadigital.fincloud.generalsaas.mobilethreemeta.detail.check
// StatusCode = 200 is success
func (a *ClientV3) MobileThreeMetaDetailCheck(ctx context.Context, bm gopay.BodyMap) (aliRsp *MobileThreeMetaDetailCheckRsp, err error) {
	err = bm.CheckEmptyError("outer_biz_no", "cert_name", "cert_no", "phone")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MobileThreeMetaDetailCheck, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MobileThreeMetaDetailCheck, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MobileThreeMetaDetailCheckRsp{StatusCode: res.StatusCode}
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

// 服务端OCR datadigital.fincloud.generalsaas.ocr.server.detect
// StatusCode = 200 is success
func (a *ClientV3) OcrServerDetect(ctx context.Context, bm gopay.BodyMap) (aliRsp *OcrServerDetectRsp, err error) {
	err = bm.CheckEmptyError("ocr_type", "outer_order_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3OcrServerDetect, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OcrServerDetect, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &OcrServerDetectRsp{StatusCode: res.StatusCode}
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

// App端OCR初始化 datadigital.fincloud.generalsaas.ocr.mobile.initialize
// StatusCode = 200 is success
func (a *ClientV3) OcrMobileInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *OcrMobileInitializeRsp, err error) {
	err = bm.CheckEmptyError("biz_code", "outer_order_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3OcrMobileInitialize, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OcrMobileInitialize, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &OcrMobileInitializeRsp{StatusCode: res.StatusCode}
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
