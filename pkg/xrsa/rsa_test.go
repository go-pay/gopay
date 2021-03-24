package xrsa

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/iGoogle-ink/gopay/pkg/xlog"
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

	aliPriKeyPKCS8 = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDj//wXZNW0ZNTXxcB1sng75LwP5iu1A56kOcDJEyWLnY1A9t/h7jVemuwhBhnyqRtGOYsJNcj4/PmDkBlldVdDbzg5jQ5TJLkvbxKURfFPIW6/+Me2+3Va8c60w9cvp6nJ6e+jY4AgQ/t8BVCPJ+1YT3D7HUDTUnlA6Vau5zgzXl7CSnkOzKN+M0XkwbnXBKZxawYIAhnQ8482u9/+li1par/h4roYkCUn6jK1ppkETRP4yligvaOTqtqc5io8rstUszLe3FEgky2mONv4NzZLBXe0h5grC5i7riB/oWJCMfClIQVD4IUx0dhgSjEiC3xbXEHdIYidXp3kdn1//AtbAgMBAAECggEBAIAT2boOg9oNVl9LkBzNv6rtiBNqWh0aRnC8nKaXody3CZV32TZiiQp2/rRiaAsfodmmE4t98kKkYUndldafvwB7ptr1eRmNpJWX0JkIXVD1V/BxQUdyJG+NLG/xH51+kuuBBYQELgJyDgCSP7nbVQHpxm7Z9QOA7epf6bjg3HpFthbQVd8fziHwn/3ZfnX3ZicXDufUEdne4iROQyvT5rt7VWrKOB/KgOtbRzyoQ3n9Q3QPNKFdC4m1d+9nXvPrVG/ypdjMhU8cVMCsoKq0LmR79Hk5DpnPfLEZvmCuBqkOCf1YlfTeT01I8aTpa+/Vz7a6SBs45KYHeh/a2OPmV6kCgYEA/9NoERuRPeO+KwiHzHQ2OyxZwXQg5a/kUQK/iIh4RzJyFtIZ+1T/FjqBvYeSL4CQKMdMgPSEbUxEEL6ldwGUVo0A95wmLI9XlZBXE/JOr8ARxh6+TTKiO6BCY8DRg7w/dW6KAROODQ+HrtfIc7yGO2WsWEB3Lg/QZ76aqyQywU8CgYEA5Ce6U8Kn4vPJRBCDvtd7tD7Ow0ndZS5NHB8p3v6+Ce7jkhCgyjWqUrS4X1Fr/jCLvToqSgPl2E75WpidyynXor30GSZPYWheVCXdp8pZ3PAOXkrnwhhLKtPvFQiLzBakQIO1kZpNK1xxEZCbfYoAVk+MuJlQU0JSnEKukaKoGjUCgYEA52bCazcUWXb3PNWqgDVeGIg1Z9czKo+/ZcemmElZAo2OxTWmasCnO+QQFu/aOz1/KGOdv8dtQbAKoVrRMKQtGVzR0ydbIFNIVfEJ0vEKW87QRPTqDcoPFQARDUoMOay9eetSyUZ25zl08aIdnxKYGIIPMoSdSPzk9fJi1QsZHeECgYB+S5eTAi/ZK8X/zXyASdwKXLIr3FmXxtecd4vACxe7e+exVU8MYESWX7+n0/SJ0umnp+j868dxw9W4Vr8SVWkuhyP2QAS1EQcWYZefaB3ydVgeziwYf0YswQCZjg+EvAvUQpRhsx6AoD1e0cOg7KFfJc3p6cRTMmV0Zl0n9PhE+QKBgGfd+oW3E59XtnEzFuK9vPXlENpnSTuTpztJ/8dp8M4v5VOsKLqUhd/90aNbQV1tPGYhBIwB7l9USrMVo1bP/+A81asdEOKTN/L2MgoN7+fFXSUQ5hc+YaNt3P80fqSWmVqurTbGZeTqnDF3XVt+rcC9yStR2PLCJF8bDoLX5s77"
	aliPubKeyPKCS8 = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4//8F2TVtGTU18XAdbJ4O+S8D+YrtQOepDnAyRMli52NQPbf4e41XprsIQYZ8qkbRjmLCTXI+Pz5g5AZZXVXQ284OY0OUyS5L28SlEXxTyFuv/jHtvt1WvHOtMPXL6epyenvo2OAIEP7fAVQjyftWE9w+x1A01J5QOlWruc4M15ewkp5DsyjfjNF5MG51wSmcWsGCAIZ0POPNrvf/pYtaWq/4eK6GJAlJ+oytaaZBE0T+MpYoL2jk6ranOYqPK7LVLMy3txRIJMtpjjb+Dc2SwV3tIeYKwuYu64gf6FiQjHwpSEFQ+CFMdHYYEoxIgt8W1xB3SGInV6d5HZ9f/wLWwIDAQAB"

	label = "www.igoogle.ink"
)

func TestRsaEncryptAndDecryptData(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("数据：", originData)
	encryptData, err := RsaEncryptData(PKCS8, []byte(originData), FormatAlipayPublicKey(aliPubKeyPKCS8))
	if err != nil {
		xlog.Error("RsaEncryptData:", err)
		return
	}
	origin, err := RsaDecryptData(PKCS8, encryptData, FormatAlipayPrivateKey(aliPriKeyPKCS8))
	if err != nil {
		xlog.Error("RsaDecryptData:", err)
		return
	}
	xlog.Debug("decrypt:", string(origin))
}

func TestRsaEncryptOAEPAndDecryptOAEPData(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("数据：", originData)
	encryptData, err := RsaEncryptOAEPData(sha256.New(), PKCS1, publicPKCS1, []byte(originData), []byte(label))
	if err != nil {
		xlog.Error("RsaEncryptData:", err)
		return
	}
	base64EncryptData := base64.StdEncoding.EncodeToString(encryptData)
	xlog.Debug("base64EncryptData:", base64EncryptData)
	bytes, err := base64.StdEncoding.DecodeString(base64EncryptData)
	if err != nil {
		xlog.Error("base64.StdEncoding.DecodeString:", err)
		return
	}
	origin, err := RsaDecryptOAEPData(sha256.New(), PKCS1, privatePKCS1, bytes, []byte(label))
	if err != nil {
		xlog.Error("RsaDecryptData:", err)
		return
	}
	xlog.Debug("decrypt:", string(origin))
}

func TestRsaEncryptOAEPAndDecryptOAEPDataPKCS8(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("数据：", originData)
	key := FormatAlipayPublicKey(aliPubKeyPKCS8)
	fmt.Println(key)
	encryptData, err := RsaEncryptOAEPData(sha256.New(), PKCS8, key, []byte(originData), []byte(label))
	if err != nil {
		xlog.Error("RsaEncryptData:", err)
		return
	}
	base64EncryptData := base64.StdEncoding.EncodeToString(encryptData)
	xlog.Debug("base64EncryptData:", base64EncryptData)
	bytes, err := base64.StdEncoding.DecodeString(base64EncryptData)
	if err != nil {
		xlog.Error("base64.StdEncoding.DecodeString:", err)
		return
	}
	privateKey := FormatAlipayPrivateKey(aliPriKeyPKCS8)
	fmt.Println(privateKey)
	origin, err := RsaDecryptOAEPData(sha256.New(), PKCS8, privateKey, bytes, []byte(label))
	if err != nil {
		xlog.Error("RsaDecryptData:", err)
		return
	}
	xlog.Debug("decrypt:", string(origin))
}
