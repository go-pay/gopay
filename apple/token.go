package apple

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/go-pay/gopay/pkg/jwt"
	"time"
)

type SignConfig struct {
	IssuerID        string // 从AppStore Connect 获得的 Issuer ID
	BundleID        string // app包名
	AppleKeyID      string // 内购密钥id
	ApplePrivateKey string // 内购密钥.p8文件内容
}

func generatingToken(ctx context.Context, signConfig *SignConfig) (string, error) {
	type CustomClaims struct {
		jwt.Claims
		Iss string `json:"iss"`
		Iat int64  `json:"iat"`
		Exp int64  `json:"exp"`
		Aud string `json:"aud"`
		Bid string `json:"bid"`
	}
	claims := CustomClaims{
		Iss: signConfig.IssuerID,
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(5 * time.Minute).Unix(),
		Aud: "appstoreconnect-v1",
		Bid: signConfig.BundleID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header = map[string]interface{}{
		"alg": "ES256",
		"kid": signConfig.AppleKeyID,
		"typ": "JWT",
	}

	privateKey, err := ParseECPrivateKeyFromPEM([]byte(signConfig.ApplePrivateKey))
	if err != nil {
		return "", err
	}

	accessToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// ParseECPrivateKeyFromPEM parses a PEM encoded Elliptic Curve Private Key Structure
func ParseECPrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("ErrKeyMustBePEMEncoded")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParseECPrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, errors.New("ErrNotECPrivateKey")
	}

	return pkey, nil
}
