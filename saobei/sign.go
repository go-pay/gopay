package saobei

import (
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// getRsaSign 获取签名字符串
func (c *Client) getRsaSign(bm gopay.BodyMap) (sign string) {
	signParams := bm.EncodeAliPaySignParams()
	c.mu.Lock()
	defer func() {
		c.md5Hash.Reset()
		c.mu.Unlock()
	}()
	c.md5Hash.Write([]byte(signParams + "&access_token=" + c.accessToken))
	return fmt.Sprintf("%x", c.md5Hash.Sum(nil))
}

// verifySign 验证响应签名
func (c *Client) verifySign(bs []byte) (err error) {
	bm := gopay.BodyMap{}
	if err = json.Unmarshal(bs, &bm); err != nil {
		return err
	}
	sign := bm.Get("key_sign")
	bm.Remove("key_sign")
	s := c.getRsaSign(bm)
	if s != sign {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, "验签失败")
	}
	return nil
}
