package wechat

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/errgroup"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xtime"
)

// 获取微信平台证书公钥（获取后自行保存使用，如需定期刷新功能，自行实现）
//	注意事项
//	如果自行实现验证平台签名逻辑的话，需要注意以下事项:
//	  - 程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
//	  - 定期调用该接口，间隔时间小于12小时
//	  - 加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）
//	文档说明：https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
func GetPlatformCerts(ctx context.Context, mchid, apiV3Key, serialNo, privateKey string) (certs *PlatformCertRsp, err error) {
	var (
		eg = new(errgroup.Group)
		mu sync.Mutex
		jb = ""
	)
	// Prepare
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	ts := util.Int642String(timestamp)
	_str := MethodGet + "\n" + v3GetCerts + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	// Sign
	h := sha256.New()
	h.Write([]byte(_str))
	result, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return nil, fmt.Errorf("rsa.SignPKCS1v15(),err:%+v", err)
	}
	sign := base64.StdEncoding.EncodeToString(result)
	// Authorization
	authorization := Authorization + ` mchid="` + mchid + `",nonce_str="` + nonceStr + `",timestamp="` + ts + `",serial_no="` + serialNo + `",signature="` + sign + `"`
	// Request
	var url = v3BaseUrlCh + v3GetCerts
	httpClient := xhttp.NewClient()
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.GetRandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, serialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err := httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	certs = &PlatformCertRsp{Code: Success}
	if res.StatusCode != http.StatusOK {
		certs.Code = res.StatusCode
		certs.Error = string(bs)
		return certs, nil
	}
	// Parse
	certRsp := new(PlatformCert)
	if err = json.Unmarshal(bs, certRsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%+v", string(bs), err)
	}
	for _, v := range certRsp.Data {
		cert := v
		if cert.EncryptCertificate != nil {
			ec := cert.EncryptCertificate
			eg.Go(func(ctx context.Context) error {
				cipherBytes, _ := base64.StdEncoding.DecodeString(ec.Ciphertext)
				pubKeyBytes, err := aes.GCMDecrypt(cipherBytes, []byte(ec.Nonce), []byte(ec.AssociatedData), []byte(apiV3Key))
				if err != nil {
					return fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
				}
				pci := &PlatformCertItem{
					EffectiveTime: cert.EffectiveTime,
					ExpireTime:    cert.ExpireTime,
					PublicKey:     string(pubKeyBytes),
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
	return certs, nil
}

// Deprecated
// 推荐直接使用 GetAndSelectNewestCert() 方法
// 获取微信平台证书公钥（获取后自行保存使用，如需定期刷新功能，自行实现）
//	注意事项
//	如果自行实现验证平台签名逻辑的话，需要注意以下事项:
//	  - 程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
//	  - 定期调用该接口，间隔时间小于12小时
//	  - 加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）
//	文档说明：https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
func (c *ClientV3) GetPlatformCerts() (certs *PlatformCertRsp, err error) {
	var (
		eg = new(errgroup.Group)
		mu sync.Mutex
	)

	authorization, err := c.authorization(MethodGet, v3GetCerts, nil)
	if err != nil {
		return nil, err
	}

	res, _, bs, err := c.doProdGet(c.ctx, v3GetCerts, authorization)
	if err != nil {
		return nil, err
	}
	certs = &PlatformCertRsp{Code: Success}
	if res.StatusCode != http.StatusOK {
		certs.Code = res.StatusCode
		certs.Error = string(bs)
		return certs, nil
	}
	certRsp := new(PlatformCert)
	if err = json.Unmarshal(bs, certRsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%+v", string(bs), err)
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
	return certs, nil
}

// Deprecated
// 设置 微信支付平台证书 和 证书序列号
//	注意：请预先通过 client.GetPlatformCerts() 获取 微信平台公钥证书 和 证书序列号
//	部分接口请求参数中敏感信息加密，使用此 微信支付平台公钥 和 证书序列号
func (c *ClientV3) SetPlatformCert(wxPublicKeyContent []byte, wxSerialNo string) (client *ClientV3) {
	pubKey, err := xpem.DecodePublicKey(wxPublicKeyContent)
	if err != nil {
		xlog.Errorf("SetPlatformCert(%s),err:%+v", wxPublicKeyContent, err)
	}
	if pubKey != nil {
		c.wxPublicKey = pubKey
	}
	c.WxSerialNo = wxSerialNo
	return c
}

// Deprecated
// 解密加密的证书
func (c *ClientV3) DecryptCerts(ciphertext, nonce, additional string) (wxCerts string, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), c.ApiV3Key)
	if err != nil {
		return "", fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	return string(decrypt), nil
}

// 获取 微信平台证书（readonly）
func (c *ClientV3) WxPublicKey() (wxPublicKey *rsa.PublicKey) {
	wxPublicKey = c.wxPublicKey
	return
}

// 获取并选择最新的有效证书
func (c *ClientV3) GetAndSelectNewestCert() (cert, serialNo string, err error) {
	certs, err := c.GetPlatformCerts()
	if err != nil {
		return gopay.NULL, gopay.NULL, err
	}
	if certs.Code == Success && len(certs.Certs) > 0 {
		// only one
		if len(certs.Certs) == 1 {
			formatExpire := xtime.FormatDateTime(certs.Certs[0].ExpireTime)
			expireTime, err := time.ParseInLocation(xtime.TimeLayout, formatExpire, time.Local)
			if err != nil {
				return gopay.NULL, gopay.NULL, fmt.Errorf("time.ParseInLocation(%s, %s),err:%w", xtime.TimeLayout, formatExpire, err)
			}
			if time.Now().Unix() >= expireTime.Unix() {
				// 过期了
				return gopay.NULL, gopay.NULL, fmt.Errorf("wechat platform API cert expired, expired time: %s", formatExpire)
			}
			return certs.Certs[0].PublicKey, certs.Certs[0].SerialNo, nil
		}
		// more one
		var (
			effectiveTs []int
			certMap     = make(map[int]*PlatformCertItem)
		)
		for _, v := range certs.Certs {
			formatEffective := xtime.FormatDateTime(v.EffectiveTime)
			effectiveTime, err := time.ParseInLocation(xtime.TimeLayout, formatEffective, time.Local)
			if err != nil {
				return gopay.NULL, gopay.NULL, fmt.Errorf("time.ParseInLocation(%s, %s),err:%w", xtime.TimeLayout, formatEffective, err)
			}
			eu := int(effectiveTime.Unix())
			effectiveTs = append(effectiveTs, eu)
			certMap[eu] = v
		}
		sort.Ints(effectiveTs)
		// newest cert
		newestCert := certMap[effectiveTs[len(effectiveTs)-1]]
		formatExpire := xtime.FormatDateTime(newestCert.ExpireTime)
		expireTime, err := time.ParseInLocation(xtime.TimeLayout, formatExpire, time.Local)
		if err != nil {
			return gopay.NULL, gopay.NULL, fmt.Errorf("time.ParseInLocation(%s, %s),err:%w", xtime.TimeLayout, formatExpire, err)
		}
		if time.Now().Unix() >= expireTime.Unix() {
			// 过期了
			return gopay.NULL, gopay.NULL, fmt.Errorf("wechat platform API cert expired, expired time: %s", formatExpire)
		}
		return newestCert.PublicKey, newestCert.SerialNo, nil
	}
	// failed
	return gopay.NULL, gopay.NULL, fmt.Errorf("GetPlatformCerts() failed or certs is nil: %+v", certs)
}

func (c *ClientV3) autoCheckCertProc() {
	for {
		time.Sleep(time.Hour * 12)
		wxPk, wxSerialNo, err := c.GetAndSelectNewestCert()
		if err != nil {
			xlog.Errorf("c.GetAndSelectNewestCert()，err:%+v", err)
			err = nil
			continue
		}
		// decode cert
		pubKey, err := xpem.DecodePublicKey([]byte(wxPk))
		if err != nil {
			xlog.Errorf("xpem.DecodePublicKey(%s)，err:%+v", wxPk, err)
			err = nil
			continue
		}
		c.wxPublicKey = pubKey
		c.WxSerialNo = wxSerialNo
	}
}
