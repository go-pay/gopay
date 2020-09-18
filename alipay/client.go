package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xhttp"
)

type Client struct {
	AppId              string
	PrivateKeyType     PKCSType
	PrivateKey         string
	LocationName       string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	SignType           string
	AppAuthToken       string
	AuthToken          string
	IsProd             bool
	location           *time.Location
	mu                 sync.RWMutex
}

// 初始化支付宝客户端
//	注意：如果使用支付宝公钥证书验签，请设置 支付宝根证书SN（client.SetAlipayRootCertSN()）、应用公钥证书SN（client.SetAppCertSN()）
//	appId：应用ID
//	privateKey：应用私钥，支持PKCS1和PKCS8
//	isProd：是否是正式环境
func NewClient(appId, privateKey string, isProd bool) (client *Client) {
	return &Client{
		AppId:      appId,
		PrivateKey: privateKey,
		IsProd:     isProd,
	}
}

// PostAliPayAPISelf 支付宝接口自行实现方法
//	示例：请参考 client_test.go 的 TestClient_PostAliPayAPISelf() 方法
func (a *Client) PostAliPayAPISelf(bm gopay.BodyMap, method string, aliRsp interface{}) (err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, method); err != nil {
		return err
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return err
	}
	return nil
}

// alipay.trade.fastpay.refund.query(统一收单交易退款查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.fastpay.refund.query
func (a *Client) TradeFastPayRefundQuery(bm gopay.BodyMap) (aliRsp *TradeFastpayRefundQueryResponse, err error) {
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("trade_no") == gotil.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.fastpay.refund.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeFastpayRefundQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.order.settle(统一收单交易结算接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.order.settle
func (a *Client) TradeOrderSettle(bm gopay.BodyMap) (aliRsp *TradeOrderSettleResponse, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "royalty_parameters")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.order.settle"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.create(统一收单交易创建接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.create
func (a *Client) TradeCreate(bm gopay.BodyMap) (aliRsp *TradeCreateResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.create"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.close(统一收单交易关闭接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.close
func (a *Client) TradeClose(bm gopay.BodyMap) (aliRsp *TradeCloseResponse, err error) {
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("trade_no") == gotil.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.close"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCloseResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.cancel(统一收单交易撤销接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.cancel
func (a *Client) TradeCancel(bm gopay.BodyMap) (aliRsp *TradeCancelResponse, err error) {
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("trade_no") == gotil.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCancelResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.refund(统一收单交易退款接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.refund
func (a *Client) TradeRefund(bm gopay.BodyMap) (aliRsp *TradeRefundResponse, err error) {
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("trade_no") == gotil.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("refund_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRefundResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.page.refund(统一收单退款页面接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.refund
func (a *Client) TradePageRefund(bm gopay.BodyMap) (aliRsp *TradePageRefundResponse, err error) {
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("trade_no") == gotil.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("out_request_no", "refund_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.page.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePageRefundResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.precreate(统一收单线下交易预创建)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate
func (a *Client) TradePrecreate(bm gopay.BodyMap) (aliRsp *TradePrecreateResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.precreate"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePrecreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	if aliRsp.NullResponse != nil {
		info := aliRsp.NullResponse
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.pay(统一收单交易支付接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.pay
func (a *Client) TradePay(bm gopay.BodyMap) (aliRsp *TradePayResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "scene", "auth_code", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.pay"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePayResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.query(统一收单线下交易查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.query
func (a *Client) TradeQuery(bm gopay.BodyMap) (aliRsp *TradeQueryResponse, err error) {
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("trade_no") == gotil.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.app.pay(app支付接口2.0)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
func (a *Client) TradeAppPay(bm gopay.BodyMap) (payParam string, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return gotil.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.app.pay"); err != nil {
		return gotil.NULL, err
	}
	payParam = string(bs)
	return payParam, nil
}

// alipay.trade.wap.pay(手机网站支付接口2.0)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay
func (a *Client) TradeWapPay(bm gopay.BodyMap) (payUrl string, err error) {
	bm.Set("product_code", "QUICK_WAP_WAY")
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return gotil.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.wap.pay"); err != nil {
		return gotil.NULL, err
	}
	payUrl = string(bs)
	return payUrl, nil
}

// alipay.trade.page.pay(统一收单下单并支付页面接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
func (a *Client) TradePagePay(bm gopay.BodyMap) (payUrl string, err error) {
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return gotil.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.page.pay"); err != nil {
		return gotil.NULL, err
	}
	payUrl = string(bs)
	return payUrl, nil
}

// alipay.fund.trans.toaccount.transfer(单笔转账到支付宝账户接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.toaccount.transfer
//	注意：此接口官方以升级替换为 alipay.fund.trans.uni.transfer
func (a *Client) FundTransToaccountTransfer(bm gopay.BodyMap) (aliRsp *FundTransToaccountTransferResponse, err error) {
	if bm.Get("out_biz_no") == gotil.NULL {
		return nil, errors.New("out_biz_no is not allowed to be null")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.toaccount.transfer"); err != nil {
		return
	}
	aliRsp = new(FundTransToaccountTransferResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.uni.transfer(单笔转账接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.uni.transfer
func (a *Client) FundTransUniTransfer(bm gopay.BodyMap) (aliRsp *FundTransUniTransferResponse, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "payee_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.uni.transfer"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransUniTransferResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.common.query(转账业务单据查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.common.query
func (a *Client) FundTransCommonQuery(bm gopay.BodyMap) (aliRsp *FundTransCommonQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.common.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransCommonQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.account.query(支付宝资金账户资产查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.account.query
func (a *Client) FundAccountQuery(bm gopay.BodyMap) (aliRsp *FundAccountQueryResponse, err error) {
	err = bm.CheckEmptyError("alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.account.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAccountQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.orderinfo.sync
func (a *Client) TradeOrderinfoSync(body gopay.BodyMap) {

}

// alipay.system.oauth.token(换取授权访问令牌)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token
func (a *Client) SystemOauthToken(bm gopay.BodyMap) (aliRsp *SystemOauthTokenResponse, err error) {
	if bm.Get("code") == gotil.NULL && bm.Get("refresh_token") == gotil.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("grant_type")
	if err != nil {
		return nil, err
	}

	if a.AppCertSN != gotil.NULL {
		a.mu.RLock()
		bm.Set("app_cert_sn", a.AppCertSN)
		a.mu.RUnlock()
	}
	if a.AliPayRootCertSN != gotil.NULL {
		a.mu.RLock()
		bm.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
		a.mu.RUnlock()
	}

	var bs []byte
	if bs, err = systemOauthToken(a.AppId, a.PrivateKeyType, a.PrivateKey, bm, "alipay.system.oauth.token", a.IsProd, a.SignType); err != nil {
		return nil, err
	}
	aliRsp = new(SystemOauthTokenResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.ErrorResponse != nil {
		info := aliRsp.ErrorResponse
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.user.info.share(支付宝会员授权信息查询接口)
//	body：此接口无需body参数
//	文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.info.share
func (a *Client) UserInfoShare() (aliRsp *UserInfoShareResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(nil, "alipay.user.info.share"); err != nil {
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
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.open.auth.token.app(换取应用授权令牌)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.auth.token.app
func (a *Client) OpenAuthTokenApp(bm gopay.BodyMap) (aliRsp *OpenAuthTokenAppResponse, err error) {
	if bm.Get("code") == gotil.NULL && bm.Get("refresh_token") == gotil.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("grant_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.open.auth.token.app"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAuthTokenAppResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// zhima.credit.score.get(芝麻分)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.score.get
func (a *Client) ZhimaCreditScoreGet(bm gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	if bm.Get("product_code") == gotil.NULL {
		bm.Set("product_code", "w1010100100000000001")
	}
	err = bm.CheckEmptyError("transaction_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.score.get"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditScoreGetResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
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
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.user.certify.open.certify(身份认证开始认证)
//	API文档地址：https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.certify
//	产品文档地址：https://opendocs.alipay.com/open/20181012100420932508/quickstart
func (a *Client) UserCertifyOpenCertify(bm gopay.BodyMap) (certifyUrl string, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return gotil.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.certify"); err != nil {
		return gotil.NULL, err
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
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.user.info.auth(用户登陆授权)
//	文档地址：https://opendocs.alipay.com/apis/api_9/alipay.user.info.auth
func (a *Client) UserInfoAuth(bm gopay.BodyMap) (aliRsp *UserInfoAuthResponse, err error) {
	err = bm.CheckEmptyError("scopes", "state")
	if err != nil {
		return nil, err
	}

	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.user.info.auth"); err != nil {
		return nil, err
	}
	if strings.Contains(string(bs), "<head>") {
		return nil, errors.New(string(bs))
	}
	aliRsp = new(UserInfoAuthResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.data.bill.balance.query(支付宝商家账户当前余额查询)
//	文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
func (a *Client) DataBillBalanceQuery(bm gopay.BodyMap) (aliRsp *DataBillBalanceQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.data.bill.balance.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillBalanceQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.data.dataservice.bill.downloadurl.query(查询对账单下载地址)
//	文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.dataservice.bill.downloadurl.query
func (a *Client) DataBillDownloadUrlQuery(bm gopay.BodyMap) (aliRsp *DataBillDownloadUrlQueryResponse, err error) {
	err = bm.CheckEmptyError("bill_type", "bill_date")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.data.dataservice.bill.downloadurl.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillDownloadUrlQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// 向支付宝发送请求
func (a *Client) doAliPay(bm gopay.BodyMap, method string) (bs []byte, err error) {
	var (
		bodyStr, url string
		bodyBs       []byte
	)
	if bm != nil {
		if bodyBs, err = json.Marshal(bm); err != nil {
			return nil, fmt.Errorf("json.Marshal：%w", err)
		}
		bodyStr = string(bodyBs)
	}
	pubBody := make(gopay.BodyMap)
	func() {
		a.mu.RLock()
		defer a.mu.RUnlock()

		pubBody.Set("app_id", a.AppId)
		pubBody.Set("method", method)
		pubBody.Set("format", "JSON")
		if a.AppCertSN != gotil.NULL {
			pubBody.Set("app_cert_sn", a.AppCertSN)
		}
		if a.AliPayRootCertSN != gotil.NULL {
			pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
		}
		if a.ReturnUrl != gotil.NULL {
			pubBody.Set("return_url", a.ReturnUrl)
		}
		pubBody.Set("charset", "utf-8")
		if a.Charset != gotil.NULL {
			pubBody.Set("charset", a.Charset)
		}
		pubBody.Set("sign_type", RSA2)
		if a.SignType != gotil.NULL {
			pubBody.Set("sign_type", a.SignType)
		}
		pubBody.Set("timestamp", time.Now().Format(gotil.TimeLayout))
		if a.LocationName != gotil.NULL && a.location != nil {
			pubBody.Set("timestamp", time.Now().In(a.location).Format(gotil.TimeLayout))
		}
		pubBody.Set("version", "1.0")
		if a.NotifyUrl != gotil.NULL {
			pubBody.Set("notify_url", a.NotifyUrl)
		}
		if a.AppAuthToken != gotil.NULL {
			pubBody.Set("app_auth_token", a.AppAuthToken)
		}
		if a.AuthToken != gotil.NULL {
			pubBody.Set("auth_token", a.AuthToken)
		}
	}()

	if bodyStr != gotil.NULL {
		pubBody.Set("biz_content", bodyStr)
	}
	sign, err := GetRsaSign(pubBody, pubBody.Get("sign_type"), a.PrivateKeyType, a.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %v", err)
	}
	pubBody.Set("sign", sign)
	param := FormatURLParam(pubBody)

	switch method {
	case "alipay.trade.app.pay":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		if !a.IsProd {
			return []byte(sandboxBaseUrl + "?" + param), nil
		}
		return []byte(baseUrl + "?" + param), nil
	default:
		httpClient := xhttp.NewClient()
		if !a.IsProd {
			url = sandboxBaseUrlUtf8
		} else {
			url = baseUrlUtf8
		}
		res, bs, errs := httpClient.Type(xhttp.TypeForm).Post(url).SendString(param).EndBytes()
		if len(errs) > 0 {
			return nil, errs[0]
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
		}
		return bs, nil
	}
}

func getSignData(bs []byte) (signData string) {
	str := string(bs)
	indexStart := strings.Index(str, `":`)
	indexEnd := strings.Index(str, `,"sign"`)
	signData = str[indexStart+2 : indexEnd]
	return
}
