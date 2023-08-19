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
	tokenStr := signedPayload
	rootCertStr, err := extractHeaderByIndex(tokenStr, 2)
	if err != nil {
		return err
	}
	intermediaCertStr, err := extractHeaderByIndex(tokenStr, 1)
	if err != nil {
		return err
	}
	if err = verifyCert(rootCertStr, intermediaCertStr); err != nil {
		return err
	}
	_, err = jwt.ParseWithClaims(tokenStr, tran, func(token *jwt.Token) (any, error) {
		return extractPublicKeyFromToken(tokenStr)
	})
	if err != nil {
		return err
	}
	return nil
}

// Per doc: https://datatracker.ietf.org/doc/html/rfc7515#section-4.1.6
func extractPublicKeyFromToken(tokenStr string) (*ecdsa.PublicKey, error) {
	certStr, err := extractHeaderByIndex(tokenStr, 0)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(certStr)
	if err != nil {
		return nil, err
	}
	switch pk := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		return pk, nil
	default:
		return nil, errors.New("appstore public key must be of type ecdsa.PublicKey")
	}
}

func extractHeaderByIndex(tokenStr string, index int) ([]byte, error) {
	if index > 2 {
		return nil, errors.New("invalid index")
	}
	tokenArr := strings.Split(tokenStr, ".")
	headerByte, err := base64.RawStdEncoding.DecodeString(tokenArr[0])
	if err != nil {
		return nil, err
	}
	type Header struct {
		Alg string   `json:"alg"`
		X5c []string `json:"x5c"`
	}
	header := &Header{}
	err = json.Unmarshal(headerByte, header)
	if err != nil {
		return nil, err
	}
	if len(header.X5c) < index {
		return nil, fmt.Errorf("index[%d] > header.x5c slice len(%d)", index, len(header.X5c))
	}
	certByte, err := base64.StdEncoding.DecodeString(header.X5c[index])
	if err != nil {
		return nil, err
	}
	return certByte, nil
}

func verifyCert(certByte, intermediaCertStr []byte) error {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		return errors.New("failed to parse root certificate")
	}
	interCert, err := x509.ParseCertificate(intermediaCertStr)
	if err != nil {
		return errors.New("failed to parse intermedia certificate")
	}
	intermedia := x509.NewCertPool()
	intermedia.AddCert(interCert)
	cert, err := x509.ParseCertificate(certByte)
	if err != nil {
		return err
	}
	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermedia,
	}
	_, err = cert.Verify(opts)
	return err
}
