package apple

import (
	"time"

	"github.com/go-pay/gopay/pkg/jwt"
)

type CustomClaims struct {
	jwt.Claims
	Iss string `json:"iss"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
	Aud string `json:"aud"`
	Bid string `json:"bid"`
}

func (c *Client) generatingToken() (string, error) {
	claims := CustomClaims{
		Iss: c.iss,
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(5 * time.Minute).Unix(),
		Aud: "appstoreconnect-v1",
		Bid: c.bid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header = map[string]any{
		"alg": "ES256",
		"kid": c.kid,
		"typ": "JWT",
	}

	accessToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
