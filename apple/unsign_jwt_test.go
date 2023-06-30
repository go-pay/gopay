package apple

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/go-pay/gopay/pkg/xlog"
)

func TestExample(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	xlog.Warnf("%+v", privateKey)

	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)
	fmt.Printf("signature2: %s\n", hex.EncodeToString(sig))

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}
