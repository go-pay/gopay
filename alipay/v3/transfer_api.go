package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// PayeeInfo 收款方信息
type PayeeInfo struct {
	Identity     string `json:"identity,omitempty"`      // 必选
	IdentityType string `json:"identity_type,omitempty"` // 必选
	CertNo       string `json:"cert_no,omitempty"`       // 可选
	CertType     string `json:"cert_type,omitempty"`     // 可选
	Name         string `json:"name,omitempty"`          // 可选
}

type TransferRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	OutBizNo       string `json:"out_biz_no"`        // 商户订单号
	OrderId        string `json:"order_id"`          // 支付宝转账订单号
	PayFundOrderId string `json:"pay_fund_order_id"` // 支付宝支付资金流水号
	TransDate      string `json:"trans_date"`        // 订单支付时间
	Status         string `json:"status"`
}

// Transfer 单笔转账接口
// StatusCode = 200 is success
func (a *ClientV3) Transfer(ctx context.Context, bm gopay.BodyMap) (aliRsp *TransferRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene", "payee_info", "order_title")
	if err != nil {
		return nil, err
	}

	authorization, err := a.authorization(MethodPost, v3TransUniTransfer, bm)
	if err != nil {
		return nil, err
	}

	res, bs, err := a.doPost(ctx, bm, v3TransUniTransfer, authorization)
	if err != nil {
		return nil, err
	}

	aliRsp = &TransferRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}

	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
