package wechat

import (
	"testing"

	"github.com/go-pay/xlog"
)

var (
	publicPKCS1 = `
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA4//8F2TVtGTU18XAdbJ4O+S8D+YrtQOepDnAyRMli52NQPbf4e41
XprsIQYZ8qkbRjmLCTXI+Pz5g5AZZXVXQ284OY0OUyS5L28SlEXxTyFuv/jHtvt1
WvHOtMPXL6epyenvo2OAIEP7fAVQjyftWE9w+x1A01J5QOlWruc4M15ewkp5Dsyj
fjNF5MG51wSmcWsGCAIZ0POPNrvf/pYtaWq/4eK6GJAlJ+oytaaZBE0T+MpYoL2j
k6ranOYqPK7LVLMy3txRIJMtpjjb+Dc2SwV3tIeYKwuYu64gf6FiQjHwpSEFQ+CF
MdHYYEoxIgt8W1xB3SGInV6d5HZ9f/wLWwIDAQAB
-----END RSA PUBLIC KEY-----`

	privatePKCS1 = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4//8F2TVtGTU18XAdbJ4O+S8D+YrtQOepDnAyRMli52NQPbf
4e41XprsIQYZ8qkbRjmLCTXI+Pz5g5AZZXVXQ284OY0OUyS5L28SlEXxTyFuv/jH
tvt1WvHOtMPXL6epyenvo2OAIEP7fAVQjyftWE9w+x1A01J5QOlWruc4M15ewkp5
DsyjfjNF5MG51wSmcWsGCAIZ0POPNrvf/pYtaWq/4eK6GJAlJ+oytaaZBE0T+MpY
oL2jk6ranOYqPK7LVLMy3txRIJMtpjjb+Dc2SwV3tIeYKwuYu64gf6FiQjHwpSEF
Q+CFMdHYYEoxIgt8W1xB3SGInV6d5HZ9f/wLWwIDAQABAoIBAQCAE9m6DoPaDVZf
S5Aczb+q7YgTalodGkZwvJyml6HctwmVd9k2YokKdv60YmgLH6HZphOLffJCpGFJ
3ZXWn78Ae6ba9XkZjaSVl9CZCF1Q9VfwcUFHciRvjSxv8R+dfpLrgQWEBC4Ccg4A
kj+521UB6cZu2fUDgO3qX+m44Nx6RbYW0FXfH84h8J/92X5192YnFw7n1BHZ3uIk
TkMr0+a7e1VqyjgfyoDrW0c8qEN5/UN0DzShXQuJtXfvZ17z61Rv8qXYzIVPHFTA
rKCqtC5ke/R5OQ6Zz3yxGb5grgapDgn9WJX03k9NSPGk6Wvv1c+2ukgbOOSmB3of
2tjj5lepAoGBAP/TaBEbkT3jvisIh8x0NjssWcF0IOWv5FECv4iIeEcychbSGftU
/xY6gb2Hki+AkCjHTID0hG1MRBC+pXcBlFaNAPecJiyPV5WQVxPyTq/AEcYevk0y
ojugQmPA0YO8P3VuigETjg0Ph67XyHO8hjtlrFhAdy4P0Ge+mqskMsFPAoGBAOQn
ulPCp+LzyUQQg77Xe7Q+zsNJ3WUuTRwfKd7+vgnu45IQoMo1qlK0uF9Ra/4wi706
KkoD5dhO+VqYncsp16K99BkmT2FoXlQl3afKWdzwDl5K58IYSyrT7xUIi8wWpECD
tZGaTStccRGQm32KAFZPjLiZUFNCUpxCrpGiqBo1AoGBAOdmwms3FFl29zzVqoA1
XhiINWfXMyqPv2XHpphJWQKNjsU1pmrApzvkEBbv2js9fyhjnb/HbUGwCqFa0TCk
LRlc0dMnWyBTSFXxCdLxClvO0ET06g3KDxUAEQ1KDDmsvXnrUslGduc5dPGiHZ8S
mBiCDzKEnUj85PXyYtULGR3hAoGAfkuXkwIv2SvF/818gEncClyyK9xZl8bXnHeL
wAsXu3vnsVVPDGBEll+/p9P0idLpp6fo/OvHccPVuFa/ElVpLocj9kAEtREHFmGX
n2gd8nVYHs4sGH9GLMEAmY4PhLwL1EKUYbMegKA9XtHDoOyhXyXN6enEUzJldGZd
J/T4RPkCgYBn3fqFtxOfV7ZxMxbivbz15RDaZ0k7k6c7Sf/HafDOL+VTrCi6lIXf
/dGjW0FdbTxmIQSMAe5fVEqzFaNWz//gPNWrHRDikzfy9jIKDe/nxV0lEOYXPmGj
bdz/NH6klplarq02xmXk6pwxd11bfq3AvckrUdjywiRfGw6C1+bO+w==
-----END RSA PRIVATE KEY-----`
)

func TestV3EncryptTextAndV3DecryptText(t *testing.T) {
	text := "I love GoPay"
	cipherText, err := V3EncryptText(text, []byte(publicPKCS1))
	if err != nil {
		xlog.Errorf("V3EncryptText.Err:", err)
		return
	}
	xlog.Debugf("encrypt text: %s", cipherText)

	originText, err := V3DecryptText(cipherText, []byte(privatePKCS1))
	if err != nil {
		xlog.Errorf("V3DecryptText.Err:", err)
		return
	}
	xlog.Debugf("decrypt text: %s", originText)
}
