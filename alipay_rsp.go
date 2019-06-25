//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/16 0:30
//==================================
package gopay

type AliPayNotifyRequest struct {
	NotifyTime        string                  `json:"notify_time"`
	NotifyType        string                  `json:"notify_type"`
	NotifyId          string                  `json:"notify_id"`
	AppId             string                  `json:"app_id"`
	Charset           string                  `json:"charset"`
	Version           string                  `json:"version"`
	SignType          string                  `json:"sign_type"`
	Sign              string                  `json:"sign"`
	AuthAppId         string                  `json:"auth_app_id"`
	TradeNo           string                  `json:"trade_no"`
	OutTradeNo        string                  `json:"out_trade_no"`
	OutBizNo          string                  `json:"out_biz_no"`
	BuyerId           string                  `json:"buyer_id"`
	BuyerLogonId      string                  `json:"buyer_logon_id"`
	SellerId          string                  `json:"seller_id"`
	SellerEmail       string                  `json:"seller_email"`
	TradeStatus       string                  `json:"trade_status"`
	TotalAmount       string                  `json:"total_amount"`
	ReceiptAmount     string                  `json:"receipt_amount"`
	InvoiceAmount     string                  `json:"invoice_amount"`
	BuyerPayAmount    string                  `json:"buyer_pay_amount"`
	PointAmount       string                  `json:"point_amount"`
	RefundFee         string                  `json:"refund_fee"`
	Subject           string                  `json:"subject"`
	Body              string                  `json:"body"`
	GmtCreate         string                  `json:"gmt_create"`
	GmtPayment        string                  `json:"gmt_payment"`
	GmtRefund         string                  `json:"gmt_refund"`
	GmtClose          string                  `json:"gmt_close"`
	FundBillList      []FundBillListInfo      `json:"fund_bill_list"`
	PassbackParams    string                  `json:"passback_params"`
	VoucherDetailList []VoucherDetailListInfo `json:"voucher_detail_list"`
}

type FundBillListInfo struct {
	FundChannel string `json:"fundChannel"`
	Amount      string `json:"amount"`
}

type VoucherDetailListInfo struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Amount             string `json:"amount"`
	MerchantContribute string `json:"merchant_contribute"`
	OtherContribute    string `json:"other_contribute"`
	Memo               string `json:"memo"`
}

type AliPayTradePayAppResponse struct {
	AlipayTradeAppPayResponse AliPayAppInfo `json:"alipay_trade_app_pay_response"`
	Sign                      string        `json:"sign"`
}

type AliPayAppInfo struct {
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	SubCode         string `json:"sub_code"`
	SubMsg          string `json:"sub_msg"`
	OutTradeNo      string `json:"out_trade_no"`
	TradeNo         string `json:"trade_no"`
	TotalAmount     string `json:"total_amount"`
	SellerId        string `json:"seller_id"`
	MerchantOrderNo string `json:"merchant_order_no"`
}
