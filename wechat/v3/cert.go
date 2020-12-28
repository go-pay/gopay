package wechat

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/aes"
	"github.com/iGoogle-ink/gotil/errgroup"
	"github.com/pkg/errors"
)

// GetPlatformCerts 获取微信平台证书
func (c *ClientV3) GetPlatformCerts() (certs *PlatformCertRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = gotil.GetRandomString(32)
		eg       = new(errgroup.Group)
		mu       sync.Mutex
	)

	authorization, err := c.authorization(MethodGet, v3GetCerts, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}

	res, si, bs, err := c.doProdGet(v3GetCerts, authorization)
	if err != nil {
		return nil, err
	}
	certs = &PlatformCertRsp{StatusCode: res.StatusCode, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		certs.Error = string(bs)
		return certs, nil
	}
	certRsp := new(PlatformCert)
	if err = json.Unmarshal(bs, certRsp); err != nil {
		return nil, errors.Errorf("json.Unmarshal(%s)：%+v", string(bs), err)
	}
	for _, v := range certRsp.Data {
		cert := v
		if cert.EncryptCertificate != nil {
			ec := cert.EncryptCertificate
			eg.Go(func(ctx context.Context) error {
				pubKey, err := c.DecryptCerts(ec.Ciphertext, ec.Nonce, ec.AssociatedData)
				if err != nil {
					return err
				}
				pci := &PlatformCertItem{
					EffectiveTime: cert.EffectiveTime,
					ExpireTime:    cert.ExpireTime,
					PublicKey:     pubKey,
					SerialNo:      cert.SerialNo,
				}
				mu.Lock()
				certs.Certs = append(certs.Certs, pci)
				mu.Unlock()
				return nil
			})
		}
	}
	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return certs, c.verifySyncSign(si)
}

// DecryptCerts 解密加密的证书
func (c *ClientV3) DecryptCerts(ciphertext, nonce, additional string) (wxCerts string, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(c.apiV3Key))
	if err != nil {
		return "", errors.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	return string(decrypt), nil
}
