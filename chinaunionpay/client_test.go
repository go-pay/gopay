package chinaunionpay

import (
	"fmt"
	"testing"
	"time"
)

const (
	appkey       = "xxxxxx"
	appid        = "xxxxxxxx"
	merchantCode = "1111111"
	terminalCode = "2222222"
	isProEnv     = false
)

var (
	requestId = NewRandomBase64(10)
	client    = NewClient(requestId, appid, appkey, isProEnv)
)

// {
//    "merchantCode":"00000000",
//    "terminalCode":"00000000",
//    "transactionAmount":"1",
//    "transactionCurrencyCode":"156",
//    "merchantOrderId":"00000000",
//    "payMode":"CODE_SCAN",
//    "payCode":"00000000"
// }
//
// {
//    "errCode":"00",
//    "errInfo":"10000成功响应码",
//    "transactionTime":"091011",
//    "transactionDate":"0101",
//    "settlementDate":"0101",
//    "transactionDateWithYear":"202000101",
//    "settlementDateWithYear":"202000101",
//    "retrievalRefNum":"00000000",
//    "transactionAmount":1,
//    "actualTransactionAmount":1,
//    "amount":1,
//    "orderId":"00000000",
//    "thirdPartyDiscountInstrution":"微信钱包支付0.01元",
//    "thirdPartyDiscountInstruction":"微信钱包支付0.01元",
//    "thirdPartyName":"微信钱包",
//    "userId":"xsssssdsdsdsdsd",
//    "thirdPartyBuyerId":"fffffffffff",
//    "thirdPartyOrderId":"42000000000000000",
//    "thirdPartyPayInformation":"现金:1",
//    "cardAttr":"01",
//    "mchntName":"测试"
// }
func TestClient_PrepayMchScanUser(t *testing.T) {
	req := &MchScanRequest{
		MerchantCode:            merchantCode,
		TerminalCode:            terminalCode,
		TransactionAmount:       1,
		TransactionCurrencyCode: TransactionCurrencyCode,
		MerchantOrderId:         fmt.Sprintf("%d", time.Now().Unix()),
		MerchantRemark:          "",
		PayMode:                 PayMode_CODE_SCAN,
		PayCode:                 "00000000",
	}

	response, err := client.PrepayMchScanUser(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", response)
}

func TestClient_QueryOrder(t *testing.T) {
	req := &QueryRequest{
		MerchantCode:    merchantCode,
		TerminalCode:    terminalCode,
		MerchantOrderId: "00000000",
		OriginalOrderId: "",
	}

	response, err := client.QueryOrder(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", response)
}

// {
//    "errCode":"00",
//    "errInfo":"10000成功响应码",
//    "transactionTime":"160455",
//    "transactionDate":"0723",
//    "settlementDate":"0723",
//    "transactionDateWithYear":"20200101",
//    "settlementDateWithYear":"20200101",
//    "retrievalRefNum":"00000000",
//    "thirdPartyName":"微信钱包",
//    "cardAttr":"01",
//    "refundInvoiceAmount":"1",
//    "transactionAmount":1
// }
func TestClient_Refund(t *testing.T) {
	req := &RefundRequest{
		MerchantCode:      merchantCode,
		TerminalCode:      terminalCode,
		MerchantOrderId:   "159000000",
		OriginalOrderId:   "",
		RefundRequestId:   "159000000_1",
		TransactionAmount: 1,
	}

	response, err := client.Refund(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", response)
}

func TestClient_QueryRefundOrder(t *testing.T) {
	req := &QueryRefundRequest{
		MerchantCode:    merchantCode,
		TerminalCode:    terminalCode,
		MerchantOrderId: "159000002",
		OriginalOrderId: "",
		RefundRequestId: "159000002_1",
	}

	response, err := client.QueryRefund(req)
	if err != nil {
		t.Error("错误========", err)
		return
	}

	t.Logf("%+v", response)
}

