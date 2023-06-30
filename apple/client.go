package apple

import (
	"crypto/ecdsa"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// Client AppleClient
type Client struct {
	iss        string // Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	bid        string // Your app’s bundle ID (Ex: “com.example.testbundleid2021”)
	kid        string // Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	privateKey *ecdsa.PrivateKey
}

// NewClient 初始化Apple客户端
// iss：issuer ID
// bid：bundle ID
// kid：private key ID
// privateKey：私钥文件读取后的字符串内容
func NewClient(iss, bid, kid, privateKey string) (client *Client, err error) {
	if iss == util.NULL || bid == util.NULL || kid == util.NULL || privateKey == util.NULL {
		return nil, gopay.MissAppleInitParamErr
	}
	ecPrivateKey, err := ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	client = &Client{
		iss:        iss,
		bid:        bid,
		kid:        kid,
		privateKey: ecPrivateKey,
	}
	return client, nil
}
