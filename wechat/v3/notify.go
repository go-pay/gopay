package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type V3DecryptResult struct {
	Appid           string             `json:"appid"`
	Mchid           string             `json:"mchid"`
	OutTradeNo      string             `json:"out_trade_no"`
	TransactionId   string             `json:"transaction_id"`
	TradeType       string             `json:"trade_type"`
	TradeState      string             `json:"trade_state"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	BankType        string             `json:"bank_type"`
	Attach          string             `json:"attach"`
	SuccessTime     string             `json:"success_time"`
	Payer           *Payer             `json:"payer"`
	Amount          *Amount            `json:"amount"`
	SceneInfo       *SceneInfo         `json:"scene_info"`
	PromotionDetail []*PromotionDetail `json:"promotion_detail"`
}

type V3DecryptRefundResult struct {
	Mchid               string        `json:"mchid"`
	OutTradeNo          string        `json:"out_trade_no"`
	TransactionId       string        `json:"transaction_id"`
	OutRefundNo         string        `json:"out_refund_no"`
	RefundId            string        `json:"refund_id"`
	RefundStatus        string        `json:"refund_status"`
	SuccessTime         string        `json:"success_time"`
	UserReceivedAccount string        `json:"user_received_account"`
	Amount              *RefundAmount `json:"amount"`
}

type V3DecryptCombineResult struct {
	CombineAppid      string       `json:"combine_appid"`
	CombineMchid      string       `json:"combine_mchid"`
	CombineOutTradeNo string       `json:"combine_out_trade_no"`
	SceneInfo         *SceneInfo   `json:"scene_info"`
	SubOrders         []*SubOrders `json:"sub_orders"`         // 最多支持子单条数：50
	CombinePayerInfo  *Payer       `json:"combine_payer_info"` // 支付者信息
}

type V3DecryptScoreResult struct {
	Appid               string           `json:"appid"`
	Mchid               string           `json:"mchid"`
	OutOrderNo          string           `json:"out_order_no"`
	ServiceId           string           `json:"service_id"`
	Openid              string           `json:"openid"`
	State               string           `json:"state"`
	StateDescription    string           `json:"state_description"`
	TotalAmount         int              `json:"total_amount"`
	ServiceIntroduction string           `json:"service_introduction"`
	PostPayments        []*PostPayments  `json:"post_payments"`
	PostDiscounts       []*PostDiscounts `json:"post_discounts"`
	RiskFund            *RiskFund        `json:"risk_fund"`
	TimeRange           *TimeRange       `json:"time_range"`
	Location            *Location        `json:"location"`
	Attach              string           `json:"attach"`
	NotifyUrl           string           `json:"notify_url"`
	OrderId             string           `json:"order_id"`
	NeedCollection      bool             `json:"need_collection"`
	Collection          *Collection      `json:"collection"`
}

type V3DecryptProfitShareResult struct {
	SpMchid       string    `json:"sp_mchid"`       // 服务商商户号
	SubMchid      string    `json:"sub_mchid"`      // 子商户号
	TransactionId string    `json:"transaction_id"` // 微信订单号
	OrderId       string    `json:"order_id"`       // 微信分账/回退单号
	OutOrderNo    string    `json:"out_order_no"`   // 商户分账/回退单号
	Receiver      *Receiver `json:"receiver"`
	SuccessTime   string    `json:"success_time"` // 成功时间
}

type Receiver struct {
	Type        string `json:"type"`        // 分账接收方类型
	Account     string `json:"account"`     // 分账接收方账号
	Amount      int    `json:"amount"`      // 分账动账金额
	Description string `json:"description"` // 分账/回退描述
}

type V3NotifyReq struct {
	Id           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     *Resource `json:"resource"`
	SignInfo     *SignInfo `json:"-"`
}

type V3NotifyRsp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// 解析微信回调请求的参数到 V3NotifyReq 结构体
func V3ParseNotify(req *http.Request) (notifyReq *V3NotifyReq, err error) {
	bs, err := ioutil.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read request body error:%+v", err)
	}
	si := &SignInfo{
		HeaderTimestamp: req.Header.Get(HeaderTimestamp),
		HeaderNonce:     req.Header.Get(HeaderNonce),
		HeaderSignature: req.Header.Get(HeaderSignature),
		HeaderSerial:    req.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	notifyReq = &V3NotifyReq{SignInfo: si}
	if err = json.Unmarshal(bs, notifyReq); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s,%#v)：%+v", string(bs), notifyReq, err)
	}
	return notifyReq, nil
}

// 异步通知验签
//	wxPubKeyContent 是通过client.GetPlatformCerts()接口向微信获取的微信平台公钥证书内容
func (v *V3NotifyReq) VerifySign(wxPkContent string) (err error) {
	if v.SignInfo != nil {
		return V3VerifySign(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, wxPkContent)
	}
	return errors.New("verify notify sign, bug SignInfo is nil")
}

// 解密 普通支付 回调中的加密信息
func (v *V3NotifyReq) DecryptCipherText(apiV3Key string) (result *V3DecryptResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%+v)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 普通退款 回调中的加密信息
func (v *V3NotifyReq) DecryptRefundCipherText(apiV3Key string) (result *V3DecryptRefundResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptRefundNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%+v)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 合单支付 回调中的加密信息
func (v *V3NotifyReq) DecryptCombineCipherText(apiV3Key string) (result *V3DecryptCombineResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptCombineNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%+v)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 支付分 回调中的加密信息
func (v *V3NotifyReq) DecryptScoreCipherText(apiV3Key string) (result *V3DecryptScoreResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptScoreNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%+v)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密分账动账回调中的加密信息
func (v *V3NotifyReq) DecryptProfitShareCipherText(apiV3Key string) (result *V3DecryptProfitShareResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptProfitShareNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%+v)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// Deprecated
// 解析微信回调请求的参数到 gopay.BodyMap
//	暂时不推荐此方法，请使用 wechat.V3ParseNotify()
func V3ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := ioutil.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	bm = make(gopay.BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%+v", string(bs), err)
	}
	return bm, nil
}
