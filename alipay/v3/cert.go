package alipay

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"os"
	"strings"

	"github.com/go-pay/gopay"
)

// 允许进行 sn 提取的证书签名算法
var allowSignatureAlgorithm = map[string]bool{
	"MD2-RSA":       true,
	"MD5-RSA":       true,
	"SHA1-RSA":      true,
	"SHA256-RSA":    true,
	"SHA384-RSA":    true,
	"SHA512-RSA":    true,
	"SHA256-RSAPSS": true,
	"SHA384-RSAPSS": true,
	"SHA512-RSAPSS": true,
}

// certPathOrData x509证书文件路径(appPublicCert.crt、alipayPublicCert.crt) 或证书 buffer
// 返回 sn：证书序列号(app_cert_sn、alipay_cert_sn)
// 返回 err：error 信息
func getCertSN(certPathOrData any) (sn string, err error) {
	var certData []byte
	switch pathOrData := certPathOrData.(type) {
	case string:
		certData, err = os.ReadFile(pathOrData)
		if err != nil {
			return gopay.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return gopay.NULL, errors.New("certPathOrData 证书类型断言错误")
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return gopay.NULL, err
		}
		name := cert.Issuer.String()
		serialNumber := cert.SerialNumber.String()
		h := md5.New()
		h.Write([]byte(name))
		h.Write([]byte(serialNumber))
		sn = hex.EncodeToString(h.Sum(nil))
	}
	if sn == gopay.NULL {
		return gopay.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// rootCertPathOrData x509证书文件路径(alipayRootCert.crt) 或文件 buffer
// 返回 sn：证书序列号(alipay_root_cert_sn)
// 返回 err：error 信息
func getRootCertSN(rootCertPathOrData any) (sn string, err error) {
	var (
		certData []byte
		certEnd  = `-----END CERTIFICATE-----`
	)
	switch pathOrData := rootCertPathOrData.(type) {
	case string:
		certData, err = os.ReadFile(pathOrData)
		if err != nil {
			return gopay.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return gopay.NULL, errors.New("rootCertPathOrData 断言异常")
	}

	pems := strings.Split(string(certData), certEnd)
	for _, c := range pems {
		if block, _ := pem.Decode([]byte(c + certEnd)); block != nil {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				continue
			}
			if !allowSignatureAlgorithm[cert.SignatureAlgorithm.String()] {
				continue
			}
			name := cert.Issuer.String()
			serialNumber := cert.SerialNumber.String()
			h := md5.New()
			h.Write([]byte(name))
			h.Write([]byte(serialNumber))
			if sn == gopay.NULL {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_" + hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == gopay.NULL {
		return gopay.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}
