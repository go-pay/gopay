package apple

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-pay/gopay/pkg/jwt"
)

// rootPEM is from `openssl x509 -inform der -in AppleRootCA-G3.cer -out apple_root.pem`
const rootPEM = `
-----BEGIN CERTIFICATE-----
MIICQzCCAcmgAwIBAgIILcX8iNLFS5UwCgYIKoZIzj0EAwMwZzEbMBkGA1UEAwwS
QXBwbGUgUm9vdCBDQSAtIEczMSYwJAYDVQQLDB1BcHBsZSBDZXJ0aWZpY2F0aW9u
IEF1dGhvcml0eTETMBEGA1UECgwKQXBwbGUgSW5jLjELMAkGA1UEBhMCVVMwHhcN
MTQwNDMwMTgxOTA2WhcNMzkwNDMwMTgxOTA2WjBnMRswGQYDVQQDDBJBcHBsZSBS
b290IENBIC0gRzMxJjAkBgNVBAsMHUFwcGxlIENlcnRpZmljYXRpb24gQXV0aG9y
aXR5MRMwEQYDVQQKDApBcHBsZSBJbmMuMQswCQYDVQQGEwJVUzB2MBAGByqGSM49
AgEGBSuBBAAiA2IABJjpLz1AcqTtkyJygRMc3RCV8cWjTnHcFBbZDuWmBSp3ZHtf
TjjTuxxEtX/1H7YyYl3J6YRbTzBPEVoA/VhYDKX1DyxNB0cTddqXl5dvMVztK517
IDvYuVTZXpmkOlEKMaNCMEAwHQYDVR0OBBYEFLuw3qFYM4iapIqZ3r6966/ayySr
MA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMAoGCCqGSM49BAMDA2gA
MGUCMQCD6cHEFl4aXTQY2e3v9GwOAEZLuN+yRhHFD/3meoyhpmvOwgPUnPWTxnS4
at+qIxUCMG1mihDK1A3UT82NQz60imOlM27jbdoXt2QfyFMm+YhidDkLF1vLUagM
6BgD56KyKA==
-----END CERTIFICATE-----
`

// ExtractClaims 解析jws格式数据
func ExtractClaims(signedPayload string, tran jwt.Claims) (interface{}, error) {
	tokenStr := signedPayload
	rootCertStr, err := extractHeaderByIndex(tokenStr, 2)
	if err != nil {
		return nil, err
	}
	intermediaCertStr, err := extractHeaderByIndex(tokenStr, 1)
	if err != nil {
		return nil, err
	}
	if err = verifyCert(rootCertStr, intermediaCertStr); err != nil {
		return nil, err
	}
	_, err = jwt.ParseWithClaims(tokenStr, tran, func(token *jwt.Token) (interface{}, error) {
		return extractPublicKeyFromToken(tokenStr)
	})
	if err != nil {
		return nil, err
	}
	return tran, nil
}

// DecodeSignedPayload 解析SignedPayload数据
func DecodeSignedPayload(signedPayload string) (payload *NotificationV2Payload, err error) {
	if signedPayload == "" {
		return nil, fmt.Errorf("signedPayload is empty")
	}
	payload = &NotificationV2Payload{}
	_, err = ExtractClaims(signedPayload, payload)
	if err != nil {
		return nil, err
	}
	return
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
