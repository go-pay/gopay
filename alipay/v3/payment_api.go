package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xtime"
)

// 统一收单交易支付接口 alipay.trade.pay
// StatusCode = 200 is success
func (a *ClientV3) TradePay(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePayRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject", "auth_code", "scene")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradePay, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradePay, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradePayRsp{StatusCode: res.StatusCode}
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

// 统一收单交易查询 alipay.trade.query
// StatusCode = 200 is success
func (a *ClientV3) TradeQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeQueryRsp, err error) {
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeQuery, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeQuery, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeQueryRsp{StatusCode: res.StatusCode}
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

// 统一收单交易退款接口 alipay.trade.refund
// StatusCode = 200 is success
func (a *ClientV3) TradeRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRefundRsp, err error) {
	err = bm.CheckEmptyError("refund_amount")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeRefund, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeRefund, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeRefundRsp{StatusCode: res.StatusCode}
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

// 统一收单交易退款查询 alipay.trade.fastpay.refund.query
// StatusCode = 200 is success
func (a *ClientV3) TradeFastPayRefundQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeFastPayRefundQueryRsp, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeFastPayRefundQuery, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeFastPayRefundQuery, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeFastPayRefundQueryRsp{StatusCode: res.StatusCode}
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

// 统一收单交易撤销接口 alipay.trade.cancel
// StatusCode = 200 is success
func (a *ClientV3) TradeCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCancelRsp, err error) {
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeCancel, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeCancel, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeCancelRsp{StatusCode: res.StatusCode}
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

// 统一收单交易关闭接口 alipay.trade.close
// StatusCode = 200 is success
func (a *ClientV3) TradeClose(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCloseRsp, err error) {
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeClose, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeClose, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeCloseRsp{StatusCode: res.StatusCode}
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

// 查询对账单下载地址 alipay.data.dataservice.bill.downloadurl.query
// StatusCode = 200 is success
func (a *ClientV3) DataBillDownloadUrlQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillDownloadUrlQueryRsp, err error) {
	err = bm.CheckEmptyError("bill_type", "bill_date")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3DataBillDownloadUrlQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &DataBillDownloadUrlQueryRsp{StatusCode: res.StatusCode}
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

// 统一收单线下交易预创建 alipay.trade.precreate
// StatusCode = 200 is success
func (a *ClientV3) TradePrecreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePrecreateRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradePrecreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradePrecreate, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradePrecreateRsp{StatusCode: res.StatusCode}
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

// alipay.trade.app.pay(app支付接口2.0)
// 文档地址：https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay
func (a *ClientV3) TradeAppPay(ctx context.Context, bm gopay.BodyMap) (orderStr string, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return gopay.NULL, err
	}

	aat := bm.Get(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	bizContent := bm.JsonBody()
	if aat != gopay.NULL {
		bm.Set(HeaderAppAuthToken, aat)
	}

	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", a.AppId).
		Set("method", "alipay.trade.app.pay").
		Set("format", "JSON").
		Set("charset", "utf-8").
		Set("sign_type", "RSA2").
		Set("version", "1.0").
		Set("timestamp", time.Now().Format(xtime.TimeLayout))

	// 前置参数校验赋值
	if a.AppCertSN != gopay.NULL {
		pubBody.Set("app_cert_sn", a.AppCertSN)
	}
	if a.AliPayRootCertSN != gopay.NULL {
		pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}
	// default use app_auth_token
	if a.AppAuthToken != gopay.NULL {
		pubBody.Set("app_auth_token", a.AppAuthToken)
	}
	if bm != nil {
		// version
		if version := bm.GetString("version"); version != gopay.NULL {
			pubBody.Set("version", version)
		}
		if returnUrl := bm.GetString("return_url"); returnUrl != gopay.NULL {
			pubBody.Set("return_url", returnUrl)
		}
		if notifyUrl := bm.GetString("notify_url"); notifyUrl != gopay.NULL {
			pubBody.Set("notify_url", notifyUrl)
		}
		// if user set app_auth_token in body_map, use this
		if aat := bm.Get(HeaderAppAuthToken); aat != gopay.NULL {
			pubBody.Set("app_auth_token", aat)
		}
	}
	if bizContent != gopay.NULL {
		if a.aesKey == gopay.NULL {
			pubBody.Set("biz_content", bizContent)
		} else {
			// AES Encrypt biz_content
			encryptBizContent, err := a.encryptBizContent(bizContent)
			if err != nil {
				return "", fmt.Errorf("EncryptBizContent Error: %w", err)
			}
			if a.DebugSwitch == gopay.DebugOn {
				a.logger.Debugf("Alipay_Origin_BizContent: %s", bizContent)
				a.logger.Debugf("Alipay_Encrypt_BizContent: %s", encryptBizContent)
			}
			pubBody.Set("biz_content", encryptBizContent)
		}
	}
	// sign
	sign, err := a.rsaSign(pubBody.EncodeAliPaySignParams())
	if err != nil {
		return "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	pubBody.Set("sign", sign)
	orderStr = pubBody.EncodeURLParams()
	return orderStr, nil
}

// alipay.trade.page.pay(统一收单下单并支付页面接口)
// 文档地址：https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay
func (a *ClientV3) TradePagePay(ctx context.Context, bm gopay.BodyMap) (payUrl string, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return gopay.NULL, err
	}
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")

	aat := bm.Get(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	bizContent := bm.JsonBody()
	if aat != gopay.NULL {
		bm.Set(HeaderAppAuthToken, aat)
	}

	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", a.AppId).
		Set("method", "alipay.trade.app.pay").
		Set("format", "JSON").
		Set("charset", "utf-8").
		Set("sign_type", "RSA2").
		Set("version", "1.0").
		Set("timestamp", time.Now().Format(xtime.TimeLayout))

	// 前置参数校验赋值
	if a.AppCertSN != gopay.NULL {
		pubBody.Set("app_cert_sn", a.AppCertSN)
	}
	if a.AliPayRootCertSN != gopay.NULL {
		pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}
	// default use app_auth_token
	if a.AppAuthToken != gopay.NULL {
		pubBody.Set("app_auth_token", a.AppAuthToken)
	}
	if bm != nil {
		// version
		if version := bm.GetString("version"); version != gopay.NULL {
			pubBody.Set("version", version)
		}
		if returnUrl := bm.GetString("return_url"); returnUrl != gopay.NULL {
			pubBody.Set("return_url", returnUrl)
		}
		if notifyUrl := bm.GetString("notify_url"); notifyUrl != gopay.NULL {
			pubBody.Set("notify_url", notifyUrl)
		}
		// if user set app_auth_token in body_map, use this
		if aat := bm.Get(HeaderAppAuthToken); aat != gopay.NULL {
			pubBody.Set("app_auth_token", aat)
		}
	}
	if bizContent != gopay.NULL {
		if a.aesKey == gopay.NULL {
			pubBody.Set("biz_content", bizContent)
		} else {
			// AES Encrypt biz_content
			encryptBizContent, err := a.encryptBizContent(bizContent)
			if err != nil {
				return "", fmt.Errorf("EncryptBizContent Error: %w", err)
			}
			if a.DebugSwitch == gopay.DebugOn {
				a.logger.Debugf("Alipay_Origin_BizContent: %s", bizContent)
				a.logger.Debugf("Alipay_Encrypt_BizContent: %s", encryptBizContent)
			}
			pubBody.Set("biz_content", encryptBizContent)
		}
	}
	// sign
	sign, err := a.rsaSign(pubBody.EncodeAliPaySignParams())
	if err != nil {
		return "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	pubBody.Set("sign", sign)
	param := pubBody.EncodeURLParams()
	if !a.IsProd {
		return "https://openapi-sandbox.dl.alipaydev.com/gateway.do?" + param, nil
	}
	return "https://openapi.alipay.com/gateway.do?" + param, nil
}

// 统一收单交易创建接口 alipay.trade.create
// StatusCode = 200 is success
func (a *ClientV3) TradeCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCreateRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject", "product_code", "op_app_id")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeCreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeCreate, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeCreateRsp{StatusCode: res.StatusCode}
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

// 支付宝订单信息同步接口 alipay.trade.orderinfo.sync
// StatusCode = 200 is success
func (a *ClientV3) TradeOrderInfoSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderInfoSyncRsp, err error) {
	err = bm.CheckEmptyError("trade_no", "out_request_no", "biz_type")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3TradeOrderInfoSync, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeOrderInfoSync, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeOrderInfoSyncRsp{StatusCode: res.StatusCode}
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

// 刷脸支付初始化 zoloz.authentication.smilepay.initialize
// StatusCode = 200 is success
func (a *ClientV3) ZolozAuthenticationSmilepayInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZolozAuthenticationSmilepayInitializeRsp, err error) {
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZolozAuthenticationSmilepayInitialize, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZolozAuthenticationSmilepayInitialize, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZolozAuthenticationSmilepayInitializeRsp{StatusCode: res.StatusCode}
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

// 查询刷脸结果信息接口 zoloz.authentication.customer.ftoken.query
// StatusCode = 200 is success
func (a *ClientV3) ZolozAuthenticationCustomerFtokenQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZolozAuthenticationCustomerFtokenQueryRsp, err error) {
	err = bm.CheckEmptyError("ftoken", "biz_type")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZolozAuthenticationCustomerFtokenQuery, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZolozAuthenticationCustomerFtokenQuery, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZolozAuthenticationCustomerFtokenQueryRsp{StatusCode: res.StatusCode}
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
