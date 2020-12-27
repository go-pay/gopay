package wechat

import (
	"encoding/json"

	"github.com/iGoogle-ink/gotil/aes"
	"github.com/pkg/errors"
)

type V3NotifyRsp struct {
	Id           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     *Resource `json:"resource"`
}

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type V3DecryptResult struct {
	Amount         *Amount `json:"amount"`
	Appid          string  `json:"appid"`
	Attach         string  `json:"attach"`
	BankType       string  `json:"bank_type"`
	Mchid          string  `json:"mchid"`
	OutTradeNo     string  `json:"out_trade_no"`
	Payer          *Payer  `json:"payer"`
	SuccessTime    string  `json:"success_time"`
	TradeState     string  `json:"trade_state"`
	TradeStateDesc string  `json:"trade_state_desc"`
	TradeType      string  `json:"trade_type"`
	TransactionID  string  `json:"transaction_id"`
}

func V3DecryptNotifyCipher(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptResult, err error) {
	decrypt, err := aes.GCMDecrypt([]byte(ciphertext), []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, errors.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, errors.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}
