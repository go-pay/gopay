package allinpay

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// verifySign 验证响应签名
func (c *Client) verifySign(bs []byte) (err error) {
	bm := gopay.BodyMap{}
	if err = json.Unmarshal(bs, &bm); err != nil {
		return err
	}
	sign := bm.Get("sign")
	bm.Remove("sign")
	signData := bm.EncodeAliPaySignParams()
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	hashs := crypto.SHA1
	c.mu.Lock()
	defer func() {
		c.sha1Hash.Reset()
		c.mu.Unlock()
	}()
	c.sha1Hash.Write([]byte(signData))
	if err = rsa.VerifyPKCS1v15(c.publicKey, hashs, c.sha1Hash.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}
