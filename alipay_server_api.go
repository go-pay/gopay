//==================================
//  * Name：Jerry
//  * DateTime：2019/6/18 19:24
//  * Desc：
//==================================
package gopay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//解析支付宝支付完成后的Notify信息
func ParseAliPayNotifyResult(req *http.Request) (notifyRsp *AliPayNotifyRequest, err error) {
	notifyRsp = new(AliPayNotifyRequest)
	notifyRsp.NotifyTime = req.FormValue("notify_time")
	notifyRsp.NotifyType = req.FormValue("notify_type")
	notifyRsp.NotifyId = req.FormValue("notify_id")
	notifyRsp.AppId = req.FormValue("app_id")
	notifyRsp.Charset = req.FormValue("charset")
	notifyRsp.Version = req.FormValue("version")
	notifyRsp.SignType = req.FormValue("sign_type")
	notifyRsp.Sign = req.FormValue("sign")
	notifyRsp.AuthAppId = req.FormValue("auth_app_id")
	notifyRsp.TradeNo = req.FormValue("trade_no")
	notifyRsp.OutTradeNo = req.FormValue("out_trade_no")
	notifyRsp.OutBizNo = req.FormValue("out_biz_no")
	notifyRsp.BuyerId = req.FormValue("buyer_id")
	notifyRsp.BuyerLogonId = req.FormValue("buyer_logon_id")
	notifyRsp.SellerId = req.FormValue("seller_id")
	notifyRsp.SellerEmail = req.FormValue("seller_email")
	notifyRsp.TradeStatus = req.FormValue("trade_status")
	notifyRsp.TotalAmount = req.FormValue("total_amount")
	notifyRsp.ReceiptAmount = req.FormValue("receipt_amount")
	notifyRsp.InvoiceAmount = req.FormValue("invoice_amount")
	notifyRsp.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	notifyRsp.PointAmount = req.FormValue("point_amount")
	notifyRsp.RefundFee = req.FormValue("refund_fee")
	notifyRsp.Subject = req.FormValue("subject")
	notifyRsp.Body = req.FormValue("body")
	notifyRsp.GmtCreate = req.FormValue("gmt_create")
	notifyRsp.GmtPayment = req.FormValue("gmt_payment")
	notifyRsp.GmtRefund = req.FormValue("gmt_refund")
	notifyRsp.GmtClose = req.FormValue("gmt_close")
	billList := req.FormValue("fund_bill_list")
	//log.Println("billList:", billList)
	if billList != null {
		bills := make([]FundBillListInfo, 0)
		err = json.Unmarshal([]byte(billList), &bills)
		if err != nil {
			return nil, err
		}
		notifyRsp.FundBillList = bills
	} else {
		notifyRsp.FundBillList = nil
	}
	notifyRsp.PassbackParams = req.FormValue("passback_params")
	detailList := req.FormValue("voucher_detail_list")
	//log.Println("detailList:", detailList)
	if detailList != null {
		details := make([]VoucherDetailListInfo, 0)
		err = json.Unmarshal([]byte(detailList), &details)
		if err != nil {
			return nil, err
		}
		notifyRsp.VoucherDetailList = details
	} else {
		notifyRsp.VoucherDetailList = nil
	}
	return notifyRsp, err
}

//支付通知的签名验证和参数签名后的Sign
//    aliPayPublicKey：支付宝公钥
//    notifyRsp：利用 gopay.ParseAliPayNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数sign：根据参数计算的sign值，非支付宝返回参数中的Sign
func VerifyAliPayResultSign(aliPayPublicKey string, notifyRsp *AliPayNotifyRequest) (ok bool, sign string) {
	body := make(BodyMap)
	body.Set("notify_time", notifyRsp.NotifyTime)
	body.Set("notify_type", notifyRsp.NotifyType)
	body.Set("notify_id", notifyRsp.NotifyId)
	body.Set("app_id", notifyRsp.AppId)
	body.Set("charset", notifyRsp.Charset)
	body.Set("version", notifyRsp.Version)
	body.Set("auth_app_id", notifyRsp.AuthAppId)
	body.Set("trade_no", notifyRsp.TradeNo)
	body.Set("out_trade_no", notifyRsp.OutTradeNo)
	body.Set("out_biz_no", notifyRsp.OutBizNo)
	body.Set("buyer_id", notifyRsp.BuyerId)
	body.Set("buyer_logon_id", notifyRsp.BuyerLogonId)
	body.Set("seller_id", notifyRsp.SellerId)
	body.Set("seller_email", notifyRsp.SellerEmail)
	body.Set("trade_status", notifyRsp.TradeStatus)
	body.Set("total_amount", notifyRsp.TotalAmount)
	body.Set("receipt_amount", notifyRsp.ReceiptAmount)
	body.Set("invoice_amount", notifyRsp.InvoiceAmount)
	body.Set("buyer_pay_amount", notifyRsp.BuyerPayAmount)
	body.Set("point_amount", notifyRsp.PointAmount)
	body.Set("refund_fee", notifyRsp.RefundFee)
	body.Set("subject", notifyRsp.Subject)
	body.Set("body", notifyRsp.Body)
	body.Set("gmt_create", notifyRsp.GmtCreate)
	body.Set("gmt_payment", notifyRsp.GmtPayment)
	body.Set("gmt_refund", notifyRsp.GmtRefund)
	body.Set("gmt_close", notifyRsp.GmtClose)
	body.Set("fund_bill_list", jsonToString(notifyRsp.FundBillList))
	body.Set("passback_params", notifyRsp.PassbackParams)
	body.Set("voucher_detail_list", jsonToString(notifyRsp.VoucherDetailList))

	newBody := make(BodyMap)
	for k, v := range body {
		if v != null {
			newBody.Set(k, v)
		}
	}

	sign, err := getRsaSign(newBody, aliPayPublicKey)
	if err != nil {
		return false, ""
	}
	ok = sign == notifyRsp.Sign
	return
}

func jsonToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	bs, err := json.Marshal(v)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	//log.Println("string:", string(bs))
	return string(bs)
}

//格式化秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")

	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen

	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

//格式化秘钥
func FormatAliPayPublicKey(publickKey string) (pKey string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")

	rawLen := 64
	keyLen := len(publickKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen

	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publickKey[start:])
		} else {
			buffer.WriteString(publickKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}
