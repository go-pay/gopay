package apple

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-pay/gopay/pkg/jwt"
)

// ExtractClaims 解析jws格式数据
// signedPayload：jws格式数据
// tran：指针类型的结构体，用于接收解析后的数据
func ExtractClaims(signedPayload string, tran jwt.Claims) (err error) {
	valueOf := reflect.ValueOf(tran)
	if valueOf.Kind() != reflect.Ptr {
		return errors.New("tran must be ptr struct")
	}

	_, err = jwt.ParseWithClaims(signedPayload, tran, func(token *jwt.Token) (any, error) {
		return x5cCertVerify(signedPayload)
	})
	if err != nil {
		return err
	}
	return nil
}

type header struct {
	Alg string   `json:"alg"`
	X5c []string `json:"x5c"`
}

// 解析 x5c root intermedia user证书
// Per doc: https://datatracker.ietf.org/doc/html/rfc7515#section-4.1.6
func (h *header) certParse() (*x509.Certificate, *x509.Certificate, *x509.Certificate, error) {
	if len(h.X5c) != 3 {
		return nil, nil, nil, errors.New("invalid x5c format")
	}
	var certDecode = func(certStr string) (*x509.Certificate, error) {
		certBytes, err := base64.StdEncoding.DecodeString(certStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %w", err)
		}
		cert, err := x509.ParseCertificate(certBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %w", err)
		}
		return cert, nil
	}
	root, err := certDecode(h.X5c[2])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode root certificate: %w", err)
	}
	interCert, err := certDecode(h.X5c[1])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode intermedia certificate: %w", err)
	}
	userCert, err := certDecode(h.X5c[0])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode user certificate: %w", err)
	}
	return root, interCert, userCert, nil
}

func (h *header) x5cCertVerify() (*ecdsa.PublicKey, error) {
	_, i, u, err := h.certParse()
	if err != nil {
		return nil, err
	}
	var iPool = x509.NewCertPool()
	iPool.AddCert(i)
	opts := x509.VerifyOptions{
		Roots:         rootCertPool,
		Intermediates: iPool,
	}
	_, err = u.Verify(opts)
	if err != nil {
		return nil, err
	}
	switch pk := u.PublicKey.(type) {
	case *ecdsa.PublicKey:
		return pk, nil
	default:
		return nil, errors.New("appstore public key must be of type ecdsa.PublicKey")
	}
}

// 苹果根ca证书
var rootCertPool *x509.CertPool

func init() {
	rootCertPool = x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM([]byte(rootPEM))
}

// 验证x5c证书链是否合法
func x5cCertVerify(tokenStr string) (*ecdsa.PublicKey, error) {
	tokenArr := strings.Split(tokenStr, ".")
	if len(tokenArr) != 3 {
		return nil, errors.New("invalid jwt format")
	}
	headerByte, err := base64.RawStdEncoding.DecodeString(tokenArr[0])
	if err != nil {
		return nil, err
	}
	h := &header{}
	err = json.Unmarshal(headerByte, h)
	if err != nil {
		return nil, fmt.Errorf("invalid jwt.header: %w", err)
	}
	return h.x5cCertVerify()
}
