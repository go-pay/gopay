package apple

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

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
	pkey, ok := parsedKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("ErrNotECPrivateKey")
	}
	return pkey, nil
}
