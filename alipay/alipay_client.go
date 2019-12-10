package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/iGoogle-ink/gopay"
)

type Client struct {
	AppId              string
	PrivateKey         string
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
	mu                 sync.RWMutex
}

// 初始化支付宝客户端
//    注意：如果使用支付宝公钥证书验签，请设置 支付宝根证书SN（client.SetAlipayRootCertSN()）、应用公钥证书SN（client.SetAppCertSN()）
//    appId：应用ID
//    PrivateKey：应用私钥
//    IsProd：是否是正式环境
func NewClient(appId, privateKey string, isProd bool) (client *Client) {
	return &Client{
		AppId:      appId,
		PrivateKey: privateKey,
		IsProd:     isProd,
	}
}

// alipay.trade.fastpay.refund.query(统一收单交易退款查询)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (a *Client) TradeFastPayRefundQuery(body gopay.BodyMap) (aliRsp *AliPayTradeFastpayRefundQueryResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.fastpay.refund.query"); err != nil {
		return
	}
	aliRsp = new(AliPayTradeFastpayRefundQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.order.settle(统一收单交易结算接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (a *Client) TradeOrderSettle(body gopay.BodyMap) (aliRsp *AliPayTradeOrderSettleResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_request_no") == gopay.NULL || body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_request_no or trade_no are not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.order.settle"); err != nil {
		return
	}
	aliRsp = new(AliPayTradeOrderSettleResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.create(统一收单交易创建接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.create
func (a *Client) TradeCreate(body gopay.BodyMap) (aliRsp *AliPayTradeCreateResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("buyer_id") == gopay.NULL {
		return nil, errors.New("out_trade_no and buyer_id are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.create"); err != nil {
		return
	}
	aliRsp = new(AliPayTradeCreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.close(统一收单交易关闭接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.close
func (a *Client) AliPayTradeClose(body gopay.BodyMap) (aliRsp *AliPayTradeCloseResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.close"); err != nil {
		return
	}
	aliRsp = new(AliPayTradeCloseResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.cancel(统一收单交易撤销接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.cancel
func (a *Client) AliPayTradeCancel(body gopay.BodyMap) (aliRsp *AliPayTradeCancelResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.cancel"); err != nil {
		return
	}
	aliRsp = new(AliPayTradeCancelResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.refund(统一收单交易退款接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.refund
func (a *Client) AliPayTradeRefund(body gopay.BodyMap) (aliRsp *AliPayTradeRefundResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(AliPayTradeRefundResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.refund(统一收单退款页面接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.page.refund
func (a *Client) AliPayTradePageRefund(body gopay.BodyMap) (aliRsp *AliPayTradePageRefundResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "	alipay.trade.page.refund"); err != nil {
		return
	}
	aliRsp = new(AliPayTradePageRefundResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.precreate(统一收单线下交易预创建)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.precreate
func (a *Client) AliPayTradePrecreate(body gopay.BodyMap) (aliRsp *AliPayTradePrecreateResponse, err error) {
	var bs []byte
	if body.Get("out_trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.precreate"); err != nil {
		return
	}
	aliRsp = new(AliPayTradePrecreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.pay(统一收单交易支付接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.pay
func (a *Client) AliPayTradePay(body gopay.BodyMap) (aliRsp *AliPayTradePayResponse, err error) {
	var bs []byte
	if body.Get("out_trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.pay"); err != nil {
		return
	}
	aliRsp = new(AliPayTradePayResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.query(统一收单线下交易查询)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.query
func (a *Client) AliPayTradeQuery(body gopay.BodyMap) (aliRsp *AliPayTradeQueryResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("out_trade_no") == gopay.NULL && body.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.query"); err != nil {
		return
	}
	aliRsp = new(AliPayTradeQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.app.pay(app支付接口2.0)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (a *Client) AliPayTradeAppPay(body gopay.BodyMap) (payParam string, err error) {
	var bs []byte
	if body.Get("out_trade_no") == gopay.NULL {
		return gopay.NULL, errors.New("out_trade_no is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.trade.app.pay"); err != nil {
		return gopay.NULL, err
	}
	payParam = string(bs)
	return
}

// alipay.trade.wap.pay(手机网站支付接口2.0)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.wap.pay
func (a *Client) AliPayTradeWapPay(body gopay.BodyMap) (payUrl string, err error) {
	var bs []byte
	if body.Get("out_trade_no") == gopay.NULL {
		return gopay.NULL, errors.New("out_trade_no is not allowed to be NULL")
	}
	body.Set("product_code", "QUICK_WAP_WAY")
	if bs, err = a.doAliPay(body, "alipay.trade.wap.pay"); err != nil {
		return gopay.NULL, err
	}
	payUrl = string(bs)
	return
}

// alipay.trade.page.pay(统一收单下单并支付页面接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (a *Client) AliPayTradePagePay(body gopay.BodyMap) (payUrl string, err error) {
	var bs []byte
	if body.Get("out_trade_no") == gopay.NULL {
		return gopay.NULL, errors.New("out_trade_no is not allowed to be NULL")
	}
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	if bs, err = a.doAliPay(body, "alipay.trade.page.pay"); err != nil {
		return gopay.NULL, err
	}
	payUrl = string(bs)
	return
}

// alipay.fund.trans.toaccount.transfer(单笔转账到支付宝账户接口)
//    文档地址：https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
func (a *Client) AliPayFundTransToaccountTransfer(body gopay.BodyMap) (aliRsp *AliPayFundTransToaccountTransferResponse, err error) {
	var bs []byte
	if body.Get("out_biz_no") == gopay.NULL {
		return nil, errors.New("out_biz_no is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.fund.trans.toaccount.transfer"); err != nil {
		return
	}
	aliRsp = new(AliPayFundTransToaccountTransferResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync
func (a *Client) AliPayTradeOrderinfoSync(body gopay.BodyMap) {

}

// alipay.system.oauth.token(换取授权访问令牌)
//    文档地址：https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (a *Client) SystemOauthToken(body gopay.BodyMap) (aliRsp *AliPaySystemOauthTokenResponse, err error) {
	var bs []byte
	if body.Get("grant_type") == gopay.NULL {
		return nil, errors.New("grant_type is not allowed to be NULL")
	}
	if body.Get("code") == gopay.NULL && body.Get("refresh_token") == gopay.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be NULL at the same time")
	}
	if bs, err = systemOauthToken(a.AppId, a.PrivateKey, body, "alipay.system.oauth.token", a.IsProd); err != nil {
		return
	}
	aliRsp = new(AliPaySystemOauthTokenResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.AccessToken == gopay.NULL {
		info := aliRsp.ErrorResponse
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.user.info.share(支付宝会员授权信息查询接口)
//    body：此接口无需body参数
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.info.share
func (a *Client) AliPayUserInfoShare() (aliRsp *AliPayUserInfoShareResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(nil, "alipay.user.info.share"); err != nil {
		return nil, err
	}
	aliRsp = new(AliPayUserInfoShareResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.open.auth.token.app(换取应用授权令牌)
//    文档地址：https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
func (a *Client) AliPayOpenAuthTokenApp(body gopay.BodyMap) (aliRsp *AliPayOpenAuthTokenAppResponse, err error) {
	var bs []byte
	if body.Get("grant_type") == gopay.NULL {
		return nil, errors.New("grant_type is not allowed to be NULL")
	}
	if body.Get("code") == gopay.NULL && body.Get("refresh_token") == gopay.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be NULL at the same time")
	}
	if bs, err = a.doAliPay(body, "alipay.open.auth.token.app"); err != nil {
		return
	}
	aliRsp = new(AliPayOpenAuthTokenAppResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// zhima.credit.score.get(芝麻分)
//    文档地址：https://docs.open.alipay.com/api_8/zhima.credit.score.get
func (a *Client) ZhimaCreditScoreGet(body gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("product_code") == gopay.NULL {
		body.Set("product_code", "w1010100100000000001")
	}
	if body.Get("transaction_id") == gopay.NULL {
		return nil, errors.New("transaction_id is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "zhima.credit.score.get"); err != nil {
		return
	}
	aliRsp = new(ZhimaCreditScoreGetResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.user.certify.open.initialize(身份认证初始化服务)
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.certify.open.initialize
func (a *Client) AliPayUserCertifyOpenInit(body gopay.BodyMap) (aliRsp *AliPayUserCertifyOpenInitResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("biz_code") == gopay.NULL {
		return nil, errors.New("biz_code is not allowed to be NULL")
	}
	if body.Get("outer_order_no") == gopay.NULL {
		return nil, errors.New("outer_order_no is not allowed to be NULL")
	}
	if body.Get("identity_param") == gopay.NULL {
		return nil, errors.New("identity_param is not allowed to be NULL")
	}
	if body.Get("merchant_config") == gopay.NULL {
		return nil, errors.New("merchant_config is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.user.certify.open.initialize"); err != nil {
		return
	}
	aliRsp = new(AliPayUserCertifyOpenInitResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// alipay.user.certify.open.certify(身份认证开始认证)
//    API文档地址：https://docs.open.alipay.com/api_2/alipay.user.certify.open.certify
//    产品文档地址：https://docs.open.alipay.com/20181012100420932508/quickstart
func (a *Client) AliPayUserCertifyOpenCertify(body gopay.BodyMap) (certifyUrl string, err error) {
	var (
		bs []byte
	)
	if body.Get("certify_id") == gopay.NULL {
		return gopay.NULL, errors.New("certify_id is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.user.certify.open.certify"); err != nil {
		return gopay.NULL, err
	}
	certifyUrl = string(bs)
	return
}

// alipay.user.certify.open.query(身份认证记录查询)
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.certify.open.query
func (a *Client) AliPayUserCertifyOpenQuery(body gopay.BodyMap) (aliRsp *AliPayUserCertifyOpenQueryResponse, err error) {
	var (
		bs []byte
	)
	if body.Get("certify_id") == gopay.NULL {
		return nil, errors.New("certify_id is not allowed to be NULL")
	}
	if bs, err = a.doAliPay(body, "alipay.user.certify.open.query"); err != nil {
		return
	}
	aliRsp = new(AliPayUserCertifyOpenQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return
}

// 向支付宝发送请求
func (a *Client) doAliPay(body gopay.BodyMap, method string) (bs []byte, err error) {
	var (
		bodyStr, sign, url string
		bodyBs             []byte
	)
	if body != nil {
		if bodyBs, err = json.Marshal(body); err != nil {
			return nil, fmt.Errorf("json.Marshal：%s", err.Error())
		}
		bodyStr = string(bodyBs)
	}
	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", a.AppId)
	pubBody.Set("method", method)
	pubBody.Set("format", "JSON")
	if a.AppCertSN != gopay.NULL {
		a.mu.RLock()
		pubBody.Set("app_cert_sn", a.AppCertSN)
		a.mu.RUnlock()
	}
	if a.AliPayRootCertSN != gopay.NULL {
		a.mu.RLock()
		pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
		a.mu.RUnlock()
	}
	if a.ReturnUrl != gopay.NULL {
		a.mu.RLock()
		pubBody.Set("return_url", a.ReturnUrl)
		a.mu.RUnlock()
	}
	if a.Charset == gopay.NULL {
		pubBody.Set("charset", "utf-8")
	} else {
		a.mu.RLock()
		pubBody.Set("charset", a.Charset)
		a.mu.RUnlock()
	}
	if a.SignType == gopay.NULL {
		pubBody.Set("sign_type", "RSA2")
	} else {
		a.mu.RLock()
		pubBody.Set("sign_type", a.SignType)
		a.mu.RUnlock()
	}
	pubBody.Set("timestamp", time.Now().Format(gopay.TimeLayout))
	pubBody.Set("version", "1.0")
	if a.NotifyUrl != gopay.NULL {
		a.mu.RLock()
		pubBody.Set("notify_url", a.NotifyUrl)
		a.mu.RUnlock()
	}
	if a.AppAuthToken != gopay.NULL {
		a.mu.RLock()
		pubBody.Set("app_auth_token", a.AppAuthToken)
		a.mu.RUnlock()
	}
	if a.AuthToken != gopay.NULL {
		a.mu.RLock()
		pubBody.Set("auth_token", a.AuthToken)
		a.mu.RUnlock()
	}
	if bodyStr != gopay.NULL {
		pubBody.Set("biz_content", bodyStr)
	}
	if sign, err = getRsaSign(pubBody, pubBody.Get("sign_type"), FormatPrivateKey(a.PrivateKey)); err != nil {
		return
	}
	pubBody.Set("sign", sign)
	param := FormatURLParam(pubBody)
	if method == "alipay.trade.app.pay" {
		return []byte(param), nil
	}
	if method == "alipay.user.certify.open.certify" {
		if !a.IsProd {
			return []byte(zfbSandboxBaseUrl + "?" + param), nil
		} else {
			return []byte(zfbBaseUrl + "?" + param), nil
		}
	}
	if method == "alipay.trade.page.pay" {
		if !a.IsProd {
			return []byte(zfbSandboxBaseUrl + "?" + param), nil
		} else {
			return []byte(zfbBaseUrl + "?" + param), nil
		}
	}
	httpClient := gopay.NewHttpClient()
	if !a.IsProd {
		url = zfbSandboxBaseUrlUtf8
	} else {
		url = zfbBaseUrlUtf8
	}
	res, bs, errs := httpClient.Type(gopay.TypeForm).Post(url).SendString(param).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if method == "alipay.trade.wap.pay" {
		if res.Request.URL.String() == zfbSandboxBaseUrl || res.Request.URL.String() == zfbBaseUrl {
			return nil, errors.New("alipay.trade.wap.pay error,please check the parameters")
		}
		return []byte(res.Request.URL.String()), nil
	}
	return bs, nil
}

func getSignData(bs []byte) (signData string) {
	str := string(bs)
	indexStart := strings.Index(str, `":`)
	indexEnd := strings.Index(str, `,"sign"`)
	signData = str[indexStart+2 : indexEnd]
	return
}