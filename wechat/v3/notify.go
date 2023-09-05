package wechat

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

type Resource struct {
	OriginalType   string `json:"original_type,omitempty"`
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

type V3DecryptPartnerResult struct {
	SpAppid         string             `json:"sp_appid"`
	SpMchid         string             `json:"sp_mchid"`
	SubAppid        string             `json:"sub_appid"`
	SubMchid        string             `json:"sub_mchid"`
	OutTradeNo      string             `json:"out_trade_no"`
	TransactionId   string             `json:"transaction_id"`
	TradeType       string             `json:"trade_type"`
	TradeState      string             `json:"trade_state"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	BankType        string             `json:"bank_type"`
	Attach          string             `json:"attach"`
	SuccessTime     string             `json:"success_time"`
	Payer           *PartnerPayer      `json:"payer"`
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

type V3DecryptPartnerRefundResult struct {
	SpMchid             string        `json:"sp_mchid"`
	SubMchid            string        `json:"sub_mchid"`
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

type V3DecryptBusifavorResult struct {
	EventType    string               `json:"event_type"`    // 事件类型
	CouponCode   string               `json:"coupon_code"`   // 券code
	StockId      string               `json:"stock_id"`      // 批次号
	SendTime     string               `json:"send_time"`     // 发放时间
	Openid       string               `json:"openid"`        // 用户标识
	Unionid      string               `json:"unionid"`       // 用户统一标识
	SendChannel  string               `json:"send_channel"`  // 发放渠道
	SendMerchant string               `json:"send_merchant"` // 发券商户号
	AttachInfo   *BusifavorAttachInfo `json:"attach_info"`   // 发券附加信息
}

type BusifavorAttachInfo struct {
	TransactionId   string `json:"transaction_id"`     // 交易订单编号
	ActCode         string `json:"act_code"`           // 支付有礼活动编号/营销馆活动ID
	HallCode        string `json:"hall_code"`          // 营销馆ID
	HallBelongMchID int    `json:"hall_belong_mch_id"` // 营销馆所属商户号
	CardID          string `json:"card_id"`            // 会员卡ID
	Code            string `json:"code"`               // 会员卡code
	ActivityID      string `json:"activity_id"`        // 会员活动ID
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

// =====================================================================================================================

// 解析微信回调请求的参数到 V3NotifyReq 结构体
func V3ParseNotify(req *http.Request) (notifyReq *V3NotifyReq, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(5<<20))) // default 5MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read request body error:%w", err)
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
		return nil, fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), notifyReq, err)
	}
	return notifyReq, nil
}

// Deprecated
// 推荐使用 VerifySignByPK()
func (v *V3NotifyReq) VerifySign(wxPkContent string) (err error) {
	if v.SignInfo != nil {
		return V3VerifySign(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, wxPkContent)
	}
	return errors.New("verify notify sign, bug SignInfo is nil")
}

// 异步通知验签
// wxPublicKey：微信平台证书公钥内容，通过 client.WxPublicKeyMap() 获取，然后根据 signInfo.HeaderSerial 获取相应的公钥
// 推荐使用 VerifySignByPKMap()
func (v *V3NotifyReq) VerifySignByPK(wxPublicKey *rsa.PublicKey) (err error) {
	if v.SignInfo != nil {
		return V3VerifySignByPK(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, wxPublicKey)
	}
	return errors.New("verify notify sign, bug SignInfo is nil")
}

// 异步通知验签
// wxPublicKey：微信平台证书公钥内容，通过 client.WxPublicKeyMap() 获取
func (v *V3NotifyReq) VerifySignByPKMap(wxPublicKeyMap map[string]*rsa.PublicKey) (err error) {
	if v.SignInfo != nil && wxPublicKeyMap != nil {
		return V3VerifySignByPK(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, wxPublicKeyMap[v.SignInfo.HeaderSerial])
	}
	return errors.New("verify notify sign, bug SignInfo or wxPublicKeyMap is nil")
}

// 解密 普通支付 回调中的加密信息
func (v *V3NotifyReq) DecryptCipherText(apiV3Key string) (result *V3DecryptResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 服务商支付 回调中的加密信息
func (v *V3NotifyReq) DecryptPartnerCipherText(apiV3Key string) (result *V3DecryptPartnerResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptPartnerNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 服务商退款 回调中的加密信息
func (v *V3NotifyReq) DecryptPartnerRefundCipherText(apiV3Key string) (result *V3DecryptPartnerRefundResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptPartnerRefundNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密商家券回调中的加密信息
func (v *V3NotifyReq) DecryptBusifavorCipherText(apiV3Key string) (result *V3DecryptBusifavorResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptBusifavorNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text error(%w)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// Deprecated
// 暂时不推荐此方法，请使用 wechat.V3ParseNotify()
// 解析微信回调请求的参数到 gopay.BodyMap
func V3ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	bm = make(gopay.BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return bm, nil
}
