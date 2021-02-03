package wechat

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"errors"
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
	Amount          *Amount            `json:"amount"`
	Appid           string             `json:"appid"`
	Attach          string             `json:"attach"`
	BankType        string             `json:"bank_type"`
	Mchid           string             `json:"mchid"`
	OutTradeNo      string             `json:"out_trade_no"`
	Payer           *Payer             `json:"payer"`
	SuccessTime     string             `json:"success_time"`
	TradeState      string             `json:"trade_state"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	TradeType       string             `json:"trade_type"`
	TransactionId   string             `json:"transaction_id"`
	PromotionDetail []*PromotionDetail `json:"promotion_detail"`
	SceneInfo       *SceneInfo         `json:"scene_info"`
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

// V3ParseNotify 解析微信回调请求的参数到 V3NotifyReq 结构体
func V3ParseNotify(req *http.Request) (notifyReq *V3NotifyReq, err error) {
	bs, err := ioutil.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		return nil, errors.Errorf("read request body error:%+v", err)
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
		return nil, errors.Errorf("json.Unmarshal(%s,%#v)：%+v", string(bs), notifyReq, err)
	}
	return notifyReq, nil
}

// VerifySign 异步通知验签
func (v *V3NotifyReq) VerifySign(wxPkContent string) (err error) {
	if v.SignInfo != nil {
		return V3VerifySign(v.SignInfo.HeaderTimestamp, v.SignInfo.HeaderNonce, v.SignInfo.SignBody, v.SignInfo.HeaderSignature, wxPkContent)
	}
	return errors.New("verify notify sign, bug SignInfo is nil")
}

// DecryptCipherText 解密回调中的加密订单信息
func (v *V3NotifyReq) DecryptCipherText(apiV3Key string) (result *V3DecryptResult, err error) {
	if v.Resource != nil {
		result, err = V3DecryptNotifyCipherText(v.Resource.Ciphertext, v.Resource.Nonce, v.Resource.AssociatedData, apiV3Key)
		if err != nil {
			bytes, _ := json.Marshal(v)
			return nil, errors.Errorf("V3NotifyReq(%s) decrypt cipher text error(%+v)", string(bytes), err)
		}
		return result, nil
	}
	return nil, errors.New("notify data Resource is nil")
}

// Deprecated
// V3ParseNotifyToBodyMap 解析微信回调请求的参数到 gopay.BodyMap
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
		return nil, errors.Errorf("json.Unmarshal(%s)：%+v", string(bs), err)
	}
	return bm, nil
}

// V3DecryptNotifyCipherText 解密回调中的加密订单信息
func V3DecryptNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, errors.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, errors.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}
