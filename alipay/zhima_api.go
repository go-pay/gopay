package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// Deprecated
// zhima.credit.score.get(查询芝麻用户的芝麻分)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.score.get
func (a *Client) ZhimaCreditScoreGet(bm gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	if bm.GetString("product_code") == util.NULL {
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
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
