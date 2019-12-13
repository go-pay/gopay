package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/iGoogle-ink/gopay/v2"
)

type Client struct {
	AppId              string
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
func (a *Client) TradeFastPayRefundQuery(bm gopay.BodyMap) (aliRsp *TradeFastpayRefundQueryResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.fastpay.refund.query"); err != nil {
		return
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
	return
}

// alipay.trade.order.settle(统一收单交易结算接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (a *Client) TradeOrderSettle(bm gopay.BodyMap) (aliRsp *TradeOrderSettleResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_request_no") == gopay.NULL || bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_request_no or trade_no are not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.order.settle"); err != nil {
		return
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
	return
}

// alipay.trade.create(统一收单交易创建接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.create
func (a *Client) TradeCreate(bm gopay.BodyMap) (aliRsp *TradeCreateResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("buyer_id") == gopay.NULL {
		return nil, errors.New("out_trade_no and buyer_id are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.create"); err != nil {
		return
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
	return
}

// alipay.trade.close(统一收单交易关闭接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.close
func (a *Client) TradeClose(bm gopay.BodyMap) (aliRsp *TradeCloseResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.close"); err != nil {
		return
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
	return
}

// alipay.trade.cancel(统一收单交易撤销接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.cancel
func (a *Client) TradeCancel(bm gopay.BodyMap) (aliRsp *TradeCancelResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.cancel"); err != nil {
		return
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
	return
}

// alipay.trade.refund(统一收单交易退款接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.refund
func (a *Client) TradeRefund(bm gopay.BodyMap) (aliRsp *TradeRefundResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
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
	return
}

// alipay.trade.page.refund(统一收单退款页面接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.page.refund
func (a *Client) TradePageRefund(bm gopay.BodyMap) (aliRsp *TradePageRefundResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.page.refund"); err != nil {
		return
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
	return
}

// alipay.trade.precreate(统一收单线下交易预创建)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.precreate
func (a *Client) TradePrecreate(bm gopay.BodyMap) (aliRsp *TradePrecreateResponse, err error) {
	var bs []byte
	if bm.Get("out_trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.precreate"); err != nil {
		return
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
	return
}

// alipay.trade.pay(统一收单交易支付接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.pay
func (a *Client) TradePay(bm gopay.BodyMap) (aliRsp *TradePayResponse, err error) {
	var bs []byte
	if bm.Get("out_trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.pay"); err != nil {
		return
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
	return
}

// alipay.trade.query(统一收单线下交易查询)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.query
func (a *Client) TradeQuery(bm gopay.BodyMap) (aliRsp *TradeQueryResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("out_trade_no") == gopay.NULL && bm.Get("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.query"); err != nil {
		return
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
	return
}

// alipay.trade.app.pay(app支付接口2.0)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (a *Client) TradeAppPay(bm gopay.BodyMap) (payParam string, err error) {
	var bs []byte
	if bm.Get("out_trade_no") == gopay.NULL {
		return gopay.NULL, errors.New("out_trade_no is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.trade.app.pay"); err != nil {
		return gopay.NULL, err
	}
	payParam = string(bs)
	return
}

// alipay.trade.wap.pay(手机网站支付接口2.0)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.wap.pay
func (a *Client) TradeWapPay(bm gopay.BodyMap) (payUrl string, err error) {
	var bs []byte
	if bm.Get("out_trade_no") == gopay.NULL {
		return gopay.NULL, errors.New("out_trade_no is not allowed to be null")
	}
	bm.Set("product_code", "QUICK_WAP_WAY")
	if bs, err = a.doAliPay(bm, "alipay.trade.wap.pay"); err != nil {
		return gopay.NULL, err
	}
	payUrl = string(bs)
	return
}

// alipay.trade.page.pay(统一收单下单并支付页面接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (a *Client) TradePagePay(bm gopay.BodyMap) (payUrl string, err error) {
	var bs []byte
	if bm.Get("out_trade_no") == gopay.NULL {
		return gopay.NULL, errors.New("out_trade_no is not allowed to be null")
	}
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	if bs, err = a.doAliPay(bm, "alipay.trade.page.pay"); err != nil {
		return gopay.NULL, err
	}
	payUrl = string(bs)
	return
}

// alipay.fund.trans.toaccount.transfer(单笔转账到支付宝账户接口)
//    文档地址：https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
func (a *Client) FundTransToaccountTransfer(bm gopay.BodyMap) (aliRsp *FundTransToaccountTransferResponse, err error) {
	var bs []byte
	if bm.Get("out_biz_no") == gopay.NULL {
		return nil, errors.New("out_biz_no is not allowed to be null")
	}
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
	return
}

// alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync
func (a *Client) TradeOrderinfoSync(body gopay.BodyMap) {

}

// alipay.system.oauth.token(换取授权访问令牌)
//    文档地址：https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (a *Client) SystemOauthToken(bm gopay.BodyMap) (aliRsp *SystemOauthTokenResponse, err error) {
	var bs []byte
	if bm.Get("grant_type") == gopay.NULL {
		return nil, errors.New("grant_type is not allowed to be null")
	}
	if bm.Get("code") == gopay.NULL && bm.Get("refresh_token") == gopay.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	if bs, err = systemOauthToken(a.AppId, a.PrivateKey, bm, "alipay.system.oauth.token", a.IsProd); err != nil {
		return
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
	return
}

// alipay.user.info.share(支付宝会员授权信息查询接口)
//    body：此接口无需body参数
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.info.share
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
	return
}

// alipay.open.auth.token.app(换取应用授权令牌)
//    文档地址：https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
func (a *Client) OpenAuthTokenApp(bm gopay.BodyMap) (aliRsp *OpenAuthTokenAppResponse, err error) {
	var bs []byte
	if bm.Get("grant_type") == gopay.NULL {
		return nil, errors.New("grant_type is not allowed to be null")
	}
	if bm.Get("code") == gopay.NULL && bm.Get("refresh_token") == gopay.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	if bs, err = a.doAliPay(bm, "alipay.open.auth.token.app"); err != nil {
		return
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
	return
}

// zhima.credit.score.get(芝麻分)
//    文档地址：https://docs.open.alipay.com/api_8/zhima.credit.score.get
func (a *Client) ZhimaCreditScoreGet(bm gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("product_code") == gopay.NULL {
		bm.Set("product_code", "w1010100100000000001")
	}
	if bm.Get("transaction_id") == gopay.NULL {
		return nil, errors.New("transaction_id is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "zhima.credit.score.get"); err != nil {
		return
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
	return
}

// alipay.user.certify.open.initialize(身份认证初始化服务)
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.certify.open.initialize
func (a *Client) UserCertifyOpenInit(bm gopay.BodyMap) (aliRsp *UserCertifyOpenInitResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("biz_code") == gopay.NULL {
		return nil, errors.New("biz_code is not allowed to be null")
	}
	if bm.Get("outer_order_no") == gopay.NULL {
		return nil, errors.New("outer_order_no is not allowed to be null")
	}
	if bm.Get("identity_param") == gopay.NULL {
		return nil, errors.New("identity_param is not allowed to be null")
	}
	if bm.Get("merchant_config") == gopay.NULL {
		return nil, errors.New("merchant_config is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.initialize"); err != nil {
		return
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
	return
}

// alipay.user.certify.open.certify(身份认证开始认证)
//    API文档地址：https://docs.open.alipay.com/api_2/alipay.user.certify.open.certify
//    产品文档地址：https://docs.open.alipay.com/20181012100420932508/quickstart
func (a *Client) UserCertifyOpenCertify(bm gopay.BodyMap) (certifyUrl string, err error) {
	var (
		bs []byte
	)
	if bm.Get("certify_id") == gopay.NULL {
		return gopay.NULL, errors.New("certify_id is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.certify"); err != nil {
		return gopay.NULL, err
	}
	certifyUrl = string(bs)
	return
}

// alipay.user.certify.open.query(身份认证记录查询)
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.certify.open.query
func (a *Client) UserCertifyOpenQuery(bm gopay.BodyMap) (aliRsp *UserCertifyOpenQueryResponse, err error) {
	var (
		bs []byte
	)
	if bm.Get("certify_id") == gopay.NULL {
		return nil, errors.New("certify_id is not allowed to be null")
	}
	if bs, err = a.doAliPay(bm, "alipay.user.certify.open.query"); err != nil {
		return
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
	return
}

// 向支付宝发送请求
func (a *Client) doAliPay(bm gopay.BodyMap, method string) (bs []byte, err error) {
	var (
		bodyStr, url string
		bodyBs       []byte
	)
	if bm != nil {
		if bodyBs, err = json.Marshal(bm); err != nil {
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
	if a.LocationName != gopay.NULL && a.location != nil {
		a.mu.RLock()
		pubBody.Set("timestamp", time.Now().In(a.location).Format(gopay.TimeLayout))
		a.mu.RUnlock()
	} else {
		pubBody.Set("timestamp", time.Now().Format(gopay.TimeLayout))
	}
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
	sign, err := getRsaSign(pubBody, pubBody.Get("sign_type"), FormatPrivateKey(a.PrivateKey))
	if err != nil {
		return
	}
	pubBody.Set("sign", sign)
	param := FormatURLParam(pubBody)

	switch method {
	case "alipay.trade.app.pay":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		if !a.IsProd {
			return []byte(zfbSandboxBaseUrl + "?" + param), nil
		}
		return []byte(zfbBaseUrl + "?" + param), nil
	default:
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
