package wechat

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/aes"
	"github.com/iGoogle-ink/gopay/pkg/xlog"
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
func (v *V3NotifyReq) VerifySign(wxPkContent string) (err error) {
	if v.SignInfo != nil {
		return V3VerifySign(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, wxPkContent)
	}
	return errors.New("verify notify sign, bug SignInfo is nil")
}

// 解密普通支付回调中的加密订单信息
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

// 解密普通退款回调中的加密订单信息
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

// 解密合单支付回调中的加密订单信息
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

// 解密普通支付回调中的加密订单信息
func V3DecryptNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密普通退款回调中的加密订单信息
func V3DecryptRefundNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptRefundResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptRefundResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密合单支付回调中的加密订单信息
func V3DecryptCombineNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptCombineResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptCombineResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
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
