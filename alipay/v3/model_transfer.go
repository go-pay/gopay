package alipay

type FundTransUniTransferRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutBizNo       string `json:"out_biz_no"`        // 商户订单号
	OrderId        string `json:"order_id"`          // 支付宝转账订单号
	PayFundOrderId string `json:"pay_fund_order_id"` // 支付宝支付资金流水号
	TransDate      string `json:"trans_date"`        // 订单支付时间
	Status         string `json:"status"`
}
