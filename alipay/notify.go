package alipay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 解析支付宝支付异步通知的参数到BodyMap
// req：*http.Request
// 返回参数bm：Notify请求的参数
// 返回参数err：错误信息
// 文档：https://opendocs.alipay.com/open/203/105286
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	var form map[string][]string = req.Form
	bm = make(gopay.BodyMap, len(form)+1)
	for k, v := range form {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}

// 通过 url.Values 解析支付宝支付异步通知的参数到BodyMap
// value：url.Values
// 返回参数notifyReq：Notify请求的参数
// 返回参数err：错误信息
// 文档：https://opendocs.alipay.com/open/203/105286
func ParseNotifyByURLValues(value url.Values) (bm gopay.BodyMap, err error) {
	bm = make(gopay.BodyMap, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}

// Deprecated
// 推荐使用 ParseNotifyToBodyMap()，以防阿里云通知参数变动，NotifyRequest 无法解析。
// 解析支付宝支付异步通知的参数到Struct
// req：*http.Request
// 返回参数notifyReq：Notify请求的参数
// 返回参数err：错误信息
// 文档：https://opendocs.alipay.com/open/203/105286
func ParseNotifyResult(req *http.Request) (notifyReq *NotifyRequest, err error) {
	notifyReq = new(NotifyRequest)
	if err = req.ParseForm(); err != nil {
		return
	}
	notifyReq.NotifyTime = req.Form.Get("notify_time")
	notifyReq.NotifyType = req.Form.Get("notify_type")
	notifyReq.NotifyId = req.Form.Get("notify_id")
	notifyReq.AppId = req.Form.Get("app_id")
	notifyReq.Charset = req.Form.Get("charset")
	notifyReq.Version = req.Form.Get("version")
	notifyReq.SignType = req.Form.Get("sign_type")
	notifyReq.Sign = req.Form.Get("sign")
	notifyReq.AuthAppId = req.Form.Get("auth_app_id")
	notifyReq.TradeNo = req.Form.Get("trade_no")
	notifyReq.OutTradeNo = req.Form.Get("out_trade_no")
	notifyReq.OutBizNo = req.Form.Get("out_biz_no")
	notifyReq.BuyerId = req.Form.Get("buyer_id")
	notifyReq.BuyerLogonId = req.Form.Get("buyer_logon_id")
	notifyReq.SellerId = req.Form.Get("seller_id")
	notifyReq.SellerEmail = req.Form.Get("seller_email")
	notifyReq.TradeStatus = req.Form.Get("trade_status")
	notifyReq.TotalAmount = req.Form.Get("total_amount")
	notifyReq.ReceiptAmount = req.Form.Get("receipt_amount")
	notifyReq.InvoiceAmount = req.Form.Get("invoice_amount")
	notifyReq.BuyerPayAmount = req.Form.Get("buyer_pay_amount")
	notifyReq.PointAmount = req.Form.Get("point_amount")
	notifyReq.RefundFee = req.Form.Get("refund_fee")
	notifyReq.Subject = req.Form.Get("subject")
	notifyReq.Body = req.Form.Get("body")
	notifyReq.GmtCreate = req.Form.Get("gmt_create")
	notifyReq.GmtPayment = req.Form.Get("gmt_payment")
	notifyReq.GmtRefund = req.Form.Get("gmt_refund")
	notifyReq.GmtClose = req.Form.Get("gmt_close")
	notifyReq.PassbackParams = req.Form.Get("passback_params")

	billList := req.Form.Get("fund_bill_list")
	if billList != util.NULL {
		bills := make([]*FundBillListInfo, 0)
		if err = json.Unmarshal([]byte(billList), &bills); err != nil {
			return nil, fmt.Errorf(`"fund_bill_list" xml.Unmarshal(%s)：%w`, billList, err)
		}
		notifyReq.FundBillList = bills
	} else {
		notifyReq.FundBillList = nil
	}

	detailList := req.Form.Get("voucher_detail_list")
	if detailList != util.NULL {
		details := make([]*VoucherDetail, 0)
		if err = json.Unmarshal([]byte(detailList), &details); err != nil {
			return nil, fmt.Errorf(`"voucher_detail_list" xml.Unmarshal(%s)：%w`, detailList, err)
		}
		notifyReq.VoucherDetailList = details
	} else {
		notifyReq.VoucherDetailList = nil
	}
	return
}
