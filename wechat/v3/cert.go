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
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/errgroup"
	"github.com/go-pay/gopay/pkg/retry"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xtime"
)

// 获取平台RSA证书列表（获取后自行保存使用，如需定期刷新功能，自行实现）
// 注意事项
// 如果自行实现验证平台签名逻辑的话，需要注意以下事项:
// - 程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
// - 定期调用该接口，间隔时间小于12小时
// - 加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）
// 文档说明：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/wechatpay5_1.shtml
func GetPlatformCerts(ctx context.Context, mchid, apiV3Key, serialNo, privateKey string, certType ...CertType) (certs *PlatformCertRsp, err error) {
	var (
		eg  = new(errgroup.Group)
		mu  sync.Mutex
		jb  = ""
		uri = v3GetCerts
	)
	if len(certType) > 1 {
		return nil, fmt.Errorf("certType must be one of `RSA` or `SM2` or `ALL`")
	}
	if len(certType) == 1 {
		uri += "?algorithm_type=" + string(certType[0])
	}
	// Prepare
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().Unix()
	nonceStr := util.RandomString(32)
	ts := util.Int642String(timestamp)
	_str := MethodGet + "\n" + uri + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	// Sign
	h := sha256.New()
	h.Write([]byte(_str))
	result, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return nil, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	sign := base64.StdEncoding.EncodeToString(result)
	// Authorization
	authorization := Authorization + ` mchid="` + mchid + `",nonce_str="` + nonceStr + `",timestamp="` + ts + `",serial_no="` + serialNo + `",signature="` + sign + `"`
	// Request
	var url = v3BaseUrlCh + uri
	hc := xhttp.NewClient().Req()
	hc.Header.Add(HeaderAuthorization, authorization)
	hc.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	hc.Header.Add(HeaderSerial, serialNo)
	hc.Header.Add("Accept", "application/json")
	res, bs, err := hc.Get(url).EndBytes(ctx)
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

// 获取平台RSA证书列表
func GetPlatformRSACerts(ctx context.Context, mchid, apiV3Key, serialNo, privateKey string) (certs *PlatformCertRsp, err error) {
	return GetPlatformCerts(ctx, mchid, apiV3Key, serialNo, privateKey, CertTypeRSA)
}

// 获取国密平台证书
func GetPlatformSM2Certs(ctx context.Context, mchid, apiV3Key, serialNo, privateKey string) (certs *PlatformCertRsp, err error) {
	return GetPlatformCerts(ctx, mchid, apiV3Key, serialNo, privateKey, CertTypeSM2)
}

// 设置 微信支付平台证书 和 证书序列号
// 注意1：如已开启自动验签功能 client.AutoVerifySign()，无需再调用此方法设置
// 注意2：请预先通过 wechat.GetPlatformCerts() 获取 微信平台公钥证书 和 证书序列号
// 部分接口请求参数中敏感信息加密，使用此 微信支付平台公钥 和 证书序列号
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

// 获取最新的 微信平台证书
func (c *ClientV3) WxPublicKey() (wxPublicKey *rsa.PublicKey) {
	return c.wxPublicKey
}

// 获取 微信平台证书 Map（readonly）
// wxPublicKeyMap: key:SerialNo, value:WxPublicKey
func (c *ClientV3) WxPublicKeyMap() (wxPublicKeyMap map[string]*rsa.PublicKey) {
	wxPublicKeyMap = make(map[string]*rsa.PublicKey, len(c.SnCertMap))
	for k, v := range c.SnCertMap {
		wxPublicKeyMap[k] = v
	}
	return wxPublicKeyMap
}

// 获取证书Map集并选择最新的有效证书序列号（默认RSA证书）
// 文档说明：https://pay.weixin.qq.com/docs/merchant/apis/platform-certificate/api-v3-get-certificates/get.html
func (c *ClientV3) GetAndSelectNewestCert(certType ...CertType) (serialNo string, snCertMap map[string]string, err error) {
	certs, err := c.getPlatformCerts(certType...)
	if err != nil {
		return gopay.NULL, nil, err
	}
	if certs.Code == Success && len(certs.Certs) > 0 {
		snCertMap = make(map[string]string)
		// only one
		if len(certs.Certs) == 1 {
			formatExpire := xtime.FormatDateTime(certs.Certs[0].ExpireTime)
			expireTime, err := time.ParseInLocation(xtime.TimeLayout, formatExpire, time.Local)
			if err != nil {
				return gopay.NULL, nil, fmt.Errorf("time.ParseInLocation(%s, %s),err:%w", xtime.TimeLayout, formatExpire, err)
			}
			if time.Since(expireTime) > 0 {
				return gopay.NULL, nil, fmt.Errorf("wechat platform API cert expired, expired time: %s", formatExpire)
			}
			serialNo = certs.Certs[0].SerialNo
			snCertMap[serialNo] = certs.Certs[0].PublicKey
			return serialNo, snCertMap, nil
		}
		// more one
		var (
			effectiveTs []int
			certMap     = make(map[int]*PlatformCertItem)
		)
		for _, v := range certs.Certs {
			formatEffective := xtime.FormatDateTime(v.EffectiveTime)
			formatExpire := xtime.FormatDateTime(v.ExpireTime)
			effectiveTime, err := time.ParseInLocation(xtime.TimeLayout, formatEffective, time.Local)
			if err != nil {
				return gopay.NULL, nil, fmt.Errorf("time.ParseInLocation(%s, %s),err:%w", xtime.TimeLayout, formatEffective, err)
			}
			expireTime, err := time.ParseInLocation(xtime.TimeLayout, formatExpire, time.Local)
			if err != nil {
				return gopay.NULL, nil, fmt.Errorf("time.ParseInLocation(%s, %s),err:%w", xtime.TimeLayout, formatExpire, err)
			}
			if time.Since(expireTime) > 0 {
				// expired
				continue
			}
			eu := int(effectiveTime.Unix())
			effectiveTs = append(effectiveTs, eu)
			certMap[eu] = v
			snCertMap[v.SerialNo] = v.PublicKey
		}
		sort.Ints(effectiveTs)
		// newest cert
		newestCert := certMap[effectiveTs[len(effectiveTs)-1]]
		return newestCert.SerialNo, snCertMap, nil
	}
	// failed
	return gopay.NULL, nil, fmt.Errorf("GetAndSelectNewestCert() failed or certs is empty: %+v", certs)
}

// 获取证书Map集并选择最新的有效RSA证书序列号
func (c *ClientV3) GetAndSelectNewestCertRSA() (serialNo string, snCertMap map[string]string, err error) {
	return c.GetAndSelectNewestCert(CertTypeRSA)
}

// 获取证书Map集并选择最新的有效SM2证书序列号
// 文档：https://pay.weixin.qq.com/docs/merchant/development/shangmi/key-and-certificate.html
func (c *ClientV3) GetAndSelectNewestCertSM2() (serialNo string, snCertMap map[string]string, err error) {
	return c.GetAndSelectNewestCert(CertTypeSM2)
}

// 获取证书Map集并选择最新的有效RSA+SM2证书序列号
func (c *ClientV3) GetAndSelectNewestCertALL() (serialNo string, snCertMap map[string]string, err error) {
	return c.GetAndSelectNewestCert(CertTypeALL)
}

// 推荐直接使用 client.GetAndSelectNewestCert() 方法
// 获取微信平台证书公钥（获取后自行保存使用，如需定期刷新功能，自行实现）
// 注意事项
// 如果自行实现验证平台签名逻辑的话，需要注意以下事项:
//   - 程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
//   - 定期调用该接口，间隔时间小于12小时
//   - 加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）
//
// 文档说明：https://pay.weixin.qq.com/docs/merchant/apis/platform-certificate/api-v3-get-certificates/get.html
func (c *ClientV3) getPlatformCerts(certType ...CertType) (certs *PlatformCertRsp, err error) {
	var (
		eg  = new(errgroup.Group)
		mu  sync.Mutex
		uri = v3GetCerts
	)
	if len(certType) > 1 {
		return nil, fmt.Errorf("certType must be one of `RSA` or `SM2` or `ALL`")
	}
	if len(certType) == 1 {
		uri += "?algorithm_type=" + string(certType[0])
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, _, bs, err := c.doProdGet(c.ctx, uri, authorization)
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
				pubKey, err := c.decryptCerts(ec.Ciphertext, ec.Nonce, ec.AssociatedData)
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

// 解密加密的证书
func (c *ClientV3) decryptCerts(ciphertext, nonce, additional string) (wxCerts string, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), c.ApiV3Key)
	if err != nil {
		return "", fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	return string(decrypt), nil
}

func (c *ClientV3) autoCheckCertProc() {
	xlog.Info("auto refresh wechat platform public key")
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			xlog.Errorf("autoCheckCertProc: panic recovered: %s\n%s", r, buf)
			// 重启
			c.autoCheckCertProc()
		}
	}()
	for {
		time.Sleep(time.Hour * 12)
		err := retry.Retry(func() error {
			serialNo, snCertMap, err := c.GetAndSelectNewestCert()
			if err != nil {
				return err
			}
			snPkMap := make(map[string]*rsa.PublicKey)
			for sn, cert := range snCertMap {
				pubKey, err := xpem.DecodePublicKey([]byte(cert))
				if err != nil {
					return err
				}
				snPkMap[sn] = pubKey
			}
			c.rwMu.Lock()
			c.SnCertMap = snPkMap
			c.WxSerialNo = serialNo
			c.wxPublicKey = snPkMap[serialNo]
			c.rwMu.Unlock()
			return nil
		}, 3, time.Second)
		if err != nil {
			xlog.Errorf("c.GetAndSelectNewestCert()，err:%+v", err)
			continue
		}
	}
}
