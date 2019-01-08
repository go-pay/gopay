package go_pay

import (
	"fmt"
	"testing"
)

func TestGetSign(t *testing.T) {
	necessary := new(WechatParamsNecessary)
	necessary.Appid = "zbcdefg"
	necessary.MchId = "1234455"
	necessary.NonceStr = "dyUNIkNS29hvDUC1CmoF0alSdfCQGg9I"
	necessary.Body = "测试充值"
	necessary.OutTradeNo = "GYsadfjksdhgflkhfgnlsdkf"
	necessary.TotalFee = 15
	necessary.SpbillCreateIp = "127.0.0.1"
	necessary.NotifyUrl = "http://www.igoogle.ink"
	necessary.TradeType = "APP"

	params := new(WechatParams)
	params.DeviceInfo = "WEB"
	params.NecessaryParams = *necessary

	s := getSign("asdfdsagsfdg", params)

	fmt.Println("sign:", s)
}
