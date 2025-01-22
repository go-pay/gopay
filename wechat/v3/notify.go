package wechat

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
	"io"
	"net/http"
)

type Resource struct {
	Algorithm      string `json:"algorithm"`
	OriginalType   string `json:"original_type,omitempty"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type V3DecryptPayResult struct {
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

// 服务商支付通知 解密结果
type V3DecryptPartnerPayResult struct {
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

// 服务商子商户处置记录回调通知
type V3DecryptViolationResult struct {
	SubMchid          string `json:"sub_mchid"`
	CompanyName       string `json:"company_name"`
	RecordId          string `json:"record_id"`
	PunishPlan        string `json:"punish_plan"`
	PunishTime        string `json:"punish_time"`
	PunishDescription string `json:"punish_description"`
	RiskType          string `json:"risk_type"`
	RiskDescription   string `json:"risk_description"`
}

// 退款通知 解密结果
type V3DecryptRefundResult struct {
	Mchid               string              `json:"mchid"`
	OutTradeNo          string              `json:"out_trade_no"`
	TransactionId       string              `json:"transaction_id"`
	OutRefundNo         string              `json:"out_refund_no"`
	RefundId            string              `json:"refund_id"`
	RefundStatus        string              `json:"refund_status"`
	SuccessTime         string              `json:"success_time"`
	UserReceivedAccount string              `json:"user_received_account"`
	Amount              *RefundNotifyAmount `json:"amount"`
}

// 服务商退款通知 解密结果
type V3DecryptPartnerRefundResult struct {
	SpMchid             string              `json:"sp_mchid"`
	SubMchid            string              `json:"sub_mchid"`
	OutTradeNo          string              `json:"out_trade_no"`
	TransactionId       string              `json:"transaction_id"`
	OutRefundNo         string              `json:"out_refund_no"`
	RefundId            string              `json:"refund_id"`
	RefundStatus        string              `json:"refund_status"`
	SuccessTime         string              `json:"success_time"`
	UserReceivedAccount string              `json:"user_received_account"`
	Amount              *RefundNotifyAmount `json:"amount"`
}

// 合单支付通知 解密结果
type V3DecryptCombineResult struct {
	CombineAppid      string       `json:"combine_appid"`
	CombineMchid      string       `json:"combine_mchid"`
	CombineOutTradeNo string       `json:"combine_out_trade_no"`
	SceneInfo         *SceneInfo   `json:"scene_info"`
	SubOrders         []*SubOrders `json:"sub_orders"`         // 最多支持子单条数：50
	CombinePayerInfo  *Payer       `json:"combine_payer_info"` // 支付者信息
}

// 支付分 确认订单回调通知 解密结果
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

// 支付分 开启/解除授权服通知 解密结果
type V3DecryptScorePermissionResult struct {
	Appid             string `json:"appid"`
	Mchid             string `json:"mchid"`
	OutOrderNo        string `json:"out_order_no"`
	ServiceId         string `json:"service_id"`
	Openid            string `json:"openid"`
	UserServiceStatus string `json:"user_service_status"`
	OpenorcloseTime   string `json:"openorclose_time"`
	AuthorizationCode string `json:"authorization_code"`
}

// 分账动账通知 解密结果
type V3DecryptProfitShareResult struct {
	Mchid         string    `json:"mchid"`
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

// 领券事件通知 解密结果
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

// 停车入场状态变更通知 解密结果
type V3DecryptParkingInResult struct {
	SpMchid                 string `json:"sp_mchid"`       // 调用接口提交的商户号
	ParkingId               string `json:"parking_id"`     // 停车入场id
	OutParkingNo            string `json:"out_parking_no"` // 商户入场id
	PlateNumber             string `json:"plate_number"`   // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor              string `json:"plate_color"`    // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	StartTime               string `json:"start_time"`     // 入场时间
	ParkingName             string `json:"parking_name"`   // 所在停车位车场的名称
	FreeDuration            int    `json:"free_duration"`  // 停车场的免费停车时长，单位为秒
	ParkingState            string `json:"parking_state"`  // 本次入场车牌的服务状态
	BlockedStateDescription string `json:"blocked_state_description,omitempty"`
	StateUpdateTime         string `json:"state_update_time"`
}

// 停车支付结果通知 解密结果
type V3DecryptParkingPayResult struct {
	Appid                 string             `json:"appid"`
	SpMchid               string             `json:"sp_mchid"` // 调用接口提交的商户号
	OutTradeNo            string             `json:"out_trade_no"`
	TransactionId         string             `json:"transaction_id"`
	Description           string             `json:"description"`
	CreateTime            string             `json:"create_time"`
	TradeState            string             `json:"trade_state"`
	TradeStateDescription string             `json:"trade_state_description"`
	SuccessTime           string             `json:"success_time"`
	BankType              string             `json:"bank_type"`
	Attach                string             `json:"attach"`
	UserRepaid            string             `json:"user_repaid"`
	TradeScene            string             `json:"trade_scene"`
	ParkingInfo           *ParkingInfo       `json:"parking_info"`
	Payer                 *Payer             `json:"payer"`
	Amount                *Amount            `json:"amount"`
	PromotionDetail       []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

// 代金券核销事件通知 解密结果
type V3DecryptCouponResult struct {
	StockCreatorMchid       string                   `json:"stock_creator_mchid"`       // 创建批次的商户号
	StockId                 string                   `json:"stock_id"`                  // 微信为每个代金券批次分配的唯一Id
	CouponId                string                   `json:"coupon_id"`                 // 微信为代金券唯一分配的id
	SingleitemDiscountOff   *SingleitemDiscountOff   `json:"singleitem_discount_off"`   // 单品优惠特定信息
	DiscountTo              *DiscountTo              `json:"discount_to,omitempty"`     // 减至优惠限定字段，仅减至优惠场景有返回。
	CouponName              string                   `json:"coupon_name"`               // 代金券名称
	Status                  string                   `json:"status"`                    // 代金券状态
	Description             string                   `json:"description"`               // 使用说明
	CreateTime              string                   `json:"create_time"`               // 领券时间，遵循rfc3339标准格式
	CouponType              string                   `json:"coupon_type"`               // 券类型
	NoCash                  bool                     `json:"no_cash"`                   // 是否无资金流
	AvailableBeginTime      string                   `json:"available_begin_time"`      // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime        string                   `json:"available_end_time"`        // 可用结束时间，遵循rfc3339标准格式
	Singleitem              bool                     `json:"singleitem"`                // 是否单品优惠
	NormalCouponInformation *NormalCouponInformation `json:"normal_coupon_information"` // 普通满减券面额、门槛信息
	ConsumeInformation      *ConsumeInformation      `json:"consume_information"`       // 已实扣代金券信息
}

type SingleitemDiscountOff struct {
	SinglePriceMax int `json:"single_price_max"`
}

type DiscountTo struct {
	CutToPrice int `json:"cut_to_price"`
	MaxPrice   int `json:"max_price"`
}

// 用户发票抬头填写完成通知 解密结果
type V3DecryptInvoiceTitleResult struct {
	Mchid         string `json:"mchid"`
	FapiaoApplyId string `json:"fapiao_apply_id"`
	ApplyTime     string `json:"apply_time"`
}

// 发票卡券作废/发票开具成功/发票冲红成功/发票插入用户卡包成功通知 解密结果
type V3DecryptInvoiceResult struct {
	Mchid             string               `json:"mchid"`
	FapiaoApplyId     string               `json:"fapiao_apply_id"`
	FapiaoInformation []*FapiaoInformation `json:"fapiao_information"`
}

type FapiaoInformation struct {
	FapiaoId     string `json:"fapiao_id"`
	FapiaoStatus string `json:"fapiao_status"`
	CardStatus   string `json:"card_status"`
}

// 投诉通知 解密结果
type V3DecryptComplaintResult struct {
	ComplaintId string `json:"complaint_id"`
	ActionType  string `json:"action_type"`
}

// 商家转账批次回调通知 解密结果
type V3DecryptTransferBatchResult struct {
	Mchid         string `json:"mchid,omitempty"`
	OutBatchNo    string `json:"out_batch_no"`
	BatchId       string `json:"batch_id"`
	BatchStatus   string `json:"batch_status"`
	TotalNum      int    `json:"total_num"`
	TotalAmount   int    `json:"total_amount"`
	SuccessAmount int    `json:"success_amount"`
	SuccessNum    int    `json:"success_num"`
	FailAmount    int    `json:"fail_amount"`
	FailNum       int    `json:"fail_num"`
	UpdateTime    string `json:"update_time"`
	CloseReason   string `json:"close_reason,omitempty"`
}

// 商家转账新版本回调通知 解密结果
type V3DecryptTransferBillsResult struct {
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	State          string `json:"state"`
	MchId          string `json:"mch_id"`
	TransferAmount int    `json:"transfer_amount"`
	Openid         string `json:"openid"`
	FailReason     string `json:"fail_reason"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

// =====================================================================================================================

type V3NotifyReq struct {
	Id           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	ResourceType string    `json:"resource_type"`
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
	if err = js.UnmarshalBytes(bs, notifyReq); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s, %+v): %w", string(bs), notifyReq, err)
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

// 解密 统一数据 到指针结构体对象
func (v *V3NotifyReq) DecryptCipherTextToStruct(apiV3Key string, objPtr any) (err error) {
	if v.Resource != nil {
		err = V3DecryptNotifyCipherTextToStruct(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key, objPtr)
		if err != nil {
			return fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return nil
	}
	return errors.New("notify data Resource is nil")
}

// 解密 普通支付 回调中的加密信息
func (v *V3NotifyReq) DecryptPayCipherText(apiV3Key string) (result *V3DecryptPayResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptPayNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 服务商支付 回调中的加密信息
func (v *V3NotifyReq) DecryptPartnerPayCipherText(apiV3Key string) (result *V3DecryptPartnerPayResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptPartnerPayNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
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
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 支付分确认订单 回调中的加密信息
func (v *V3NotifyReq) DecryptScoreCipherText(apiV3Key string) (result *V3DecryptScoreResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptScoreNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 支付分开启/解除授权服务 回调中的加密信息
func (v *V3NotifyReq) DecryptScorePermissionCipherText(apiV3Key string) (result *V3DecryptScorePermissionResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptScorePermissionNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 分账动账 回调中的加密信息
func (v *V3NotifyReq) DecryptProfitShareCipherText(apiV3Key string) (result *V3DecryptProfitShareResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptProfitShareNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 领券事件 回调中的加密信息
func (v *V3NotifyReq) DecryptBusifavorCipherText(apiV3Key string) (result *V3DecryptBusifavorResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptBusifavorNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 停车入场状态变更 回调中的加密信息
func (v *V3NotifyReq) DecryptParkingInCipherText(apiV3Key string) (result *V3DecryptParkingInResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptParkingInNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 订单支付结果 回调中的加密信息
func (v *V3NotifyReq) DecryptParkingPayCipherText(apiV3Key string) (result *V3DecryptParkingPayResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptParkingPayNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 代金券核销事件 回调中的加密信息
func (v *V3NotifyReq) DecryptCouponUseCipherText(apiV3Key string) (result *V3DecryptCouponResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptCouponUseNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 用户发票抬头填写完成 回调中的加密信息
func (v *V3NotifyReq) DecryptInvoiceTitleCipherText(apiV3Key string) (result *V3DecryptInvoiceTitleResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptInvoiceTitleNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 发票卡券作废/发票开具成功/发票冲红成功/发票插入用户卡包成功 回调中的加密信息
func (v *V3NotifyReq) DecryptInvoiceCipherText(apiV3Key string) (result *V3DecryptInvoiceResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptInvoiceNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 服务商子商户处置记录 回调中的加密信息
func (v *V3NotifyReq) DecryptViolationCipherText(apiV3Key string) (result *V3DecryptViolationResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptViolationNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 商家转账批次回调通知 回调中的加密信息
func (v *V3NotifyReq) DecryptTransferBatchCipherText(apiV3Key string) (result *V3DecryptTransferBatchResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptTransferBatchNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// 解密 新版商家转账通知 回调中的加密信息
func (v *V3NotifyReq) DecryptTransferBillsNotifyCipherText(apiV3Key string) (result *V3DecryptTransferBillsResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptTransferBillsNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			return nil, fmt.Errorf("V3NotifyReq(%s) decrypt cipher text err(%w)", js.MarshalString(v), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// Deprecated
// 暂时不推荐此方法，请使用 wechat.V3ParseNotify()
// 解析微信回调请求的参数到 gopay.BodyMap
func V3ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(5<<20))) // default 5MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		xlog.Error("V3ParseNotifyToBodyMap, io.ReadAll, err:", err)
		return
	}
	bm = make(gopay.BodyMap)
	if err = js.UnmarshalBytes(bs, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	return bm, nil
}
