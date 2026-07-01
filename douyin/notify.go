package douyin

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

// Resource 回调 resource 字段（AES-256-GCM 加密的资源）
type Resource struct {
	Algorithm      string `json:"algorithm"`
	OriginalType   string `json:"original_type,omitempty"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data,omitempty"`
	Nonce          string `json:"nonce"`
	Mchid          string `json:"mchid,omitempty"` // 支付通知的 resource 内含商户号，退款通知无
}

// NotifyReq 抖音支付回调通用请求体
// event_type 覆盖：TRANSACTION.SUCCESS / REFUND.SUCCESS / ASYNC_SPLIT.FINISH / SPLIT.SUCCESS / TRANSFER.SUCCESS
type NotifyReq struct {
	Id           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary,omitempty"`
	ResourceType string    `json:"resource_type"`
	Resource     *Resource `json:"resource"`
	SignInfo     *SignInfo `json:"-"`
}

// NotifyRsp 回调应答体（接收失败时按此格式返回）
type NotifyRsp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ParseNotify 解析抖音支付回调请求为 NotifyReq 结构体（推荐使用）
// 内部读取 Body 并绑定验签所需的 SignInfo（Douyinpay-Timestamp / Nonce / Signature / Serial）
func ParseNotify(req *http.Request) (notifyReq *NotifyReq, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(5<<20))) // 默认 5MB
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read request body error: %w", err)
	}
	si := &SignInfo{
		HeaderTimestamp: req.Header.Get(HeaderTimestamp),
		HeaderNonce:     req.Header.Get(HeaderNonce),
		HeaderSignature: req.Header.Get(HeaderSignature),
		HeaderSerial:    req.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	notifyReq = &NotifyReq{SignInfo: si}
	if err = js.UnmarshalBytes(bs, notifyReq); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	return notifyReq, nil
}

// ParseNotifyToBodyMap 解析抖音支付回调请求到 gopay.BodyMap（备用，一般推荐 ParseNotify）
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(5<<20)))
	defer req.Body.Close()
	if err != nil {
		xlog.Error("ParseNotifyToBodyMap, io.ReadAll, err:", err)
		return nil, err
	}
	bm = make(gopay.BodyMap)
	if err = js.UnmarshalBytes(bs, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	return bm, nil
}

// VerifySignByPK 回调验签（外部传入平台公钥）
// 平台公钥应与 v.SignInfo.HeaderSerial 匹配
func (v *NotifyReq) VerifySignByPK(publicKey *rsa.PublicKey) error {
	if v.SignInfo == nil {
		return errors.New("verify notify sign, but SignInfo is nil")
	}
	return VerifySignByPK(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, publicKey)
}

// VerifySignByPKMap 回调验签（传入平台公钥 map，按 Serial 匹配）
// 推荐搭配 Client.PlatformCertMap() 使用
func (v *NotifyReq) VerifySignByPKMap(publicKeyMap map[string]*rsa.PublicKey) error {
	if v.SignInfo == nil {
		return errors.New("verify notify sign, but SignInfo is nil")
	}
	pk, ok := publicKeyMap[v.SignInfo.HeaderSerial]
	if !ok {
		return fmt.Errorf("verify notify sign, but public key of serial(%s) not found", v.SignInfo.HeaderSerial)
	}
	return VerifySignByPK(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, pk)
}

// DecryptCipherTextToStruct 解密 resource 密文到任意结构体指针
// apiKey：商户在【产品中心】->【密钥管理】->【接口加密密钥】中设置的 32 字节 AES 密钥
func (v *NotifyReq) DecryptCipherTextToStruct(apiKey string, objPtr any) error {
	if v.Resource == nil {
		return errors.New("notify data Resource is nil")
	}
	if err := DecryptNotifyCipherTextToStruct(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiKey, objPtr); err != nil {
		return fmt.Errorf("NotifyReq(%s) decrypt cipher text err: %w", js.MarshalString(v), err)
	}
	return nil
}

// DecryptPayCipherText 支付成功通知（event_type=TRANSACTION.SUCCESS）密文解密
func (v *NotifyReq) DecryptPayCipherText(apiKey string) (result *DecryptPayResult, err error) {
	result = &DecryptPayResult{}
	if err = v.DecryptCipherTextToStruct(apiKey, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptRefundCipherText 退款结果通知（event_type=REFUND.*）密文解密
func (v *NotifyReq) DecryptRefundCipherText(apiKey string) (result *DecryptRefundResult, err error) {
	result = &DecryptRefundResult{}
	if err = v.DecryptCipherTextToStruct(apiKey, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptProfitResultCipherText 分账结果通知（event_type=ASYNC_SPLIT.FINISH）密文解密
func (v *NotifyReq) DecryptProfitResultCipherText(apiKey string) (result *DecryptProfitResult, err error) {
	result = &DecryptProfitResult{}
	if err = v.DecryptCipherTextToStruct(apiKey, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptProfitDynamicCipherText 分账动账通知（event_type=SPLIT.SUCCESS）密文解密
func (v *NotifyReq) DecryptProfitDynamicCipherText(apiKey string) (result *DecryptProfitDynamic, err error) {
	result = &DecryptProfitDynamic{}
	if err = v.DecryptCipherTextToStruct(apiKey, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptTransferCipherText 转账结果通知（event_type=TRANSFER.SUCCESS）密文解密
func (v *NotifyReq) DecryptTransferCipherText(apiKey string) (result *DecryptTransferResult, err error) {
	result = &DecryptTransferResult{}
	if err = v.DecryptCipherTextToStruct(apiKey, result); err != nil {
		return nil, err
	}
	return result, nil
}
