package xrsa

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"

	"github.com/iGoogle-ink/gopay/pkg/xlog"
)

var (
	pubKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw1WT4WTewZ0F6BsGr96b
iuorTwQrn9+o5XOIk/borLojBpKwW/UprulunqoxlENm34IQ8cMNJSdtJNDhBz7v
PI1U1SkSe4wEjYd/Jo6K86hzAhhNO7h5P+8sPMJRzyzk1Pa9hBLLfX8tNSVj68h2
tJkXjHuadbV0EqNddDpkOpaStd+9bT02YSRRTqFBk2hIIb7eQLm7A6QFe43cwsvx
cBXwyTml15TkHNOtPnZqlX8M+Waam4AeXql8eA7polrWFCKLxaxIRtFP8Abckbpi
WAaywgXPf6b2eV62gShSsHq6owjm+HRKiatf/LXcG3NIaHw8sECw1p8lMWwoBo7C
lQIDAQAB
-----END PUBLIC KEY-----`

	priKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAw1WT4WTewZ0F6BsGr96biuorTwQrn9+o5XOIk/borLojBpKw
W/UprulunqoxlENm34IQ8cMNJSdtJNDhBz7vPI1U1SkSe4wEjYd/Jo6K86hzAhhN
O7h5P+8sPMJRzyzk1Pa9hBLLfX8tNSVj68h2tJkXjHuadbV0EqNddDpkOpaStd+9
bT02YSRRTqFBk2hIIb7eQLm7A6QFe43cwsvxcBXwyTml15TkHNOtPnZqlX8M+Waa
m4AeXql8eA7polrWFCKLxaxIRtFP8AbckbpiWAaywgXPf6b2eV62gShSsHq6owjm
+HRKiatf/LXcG3NIaHw8sECw1p8lMWwoBo7ClQIDAQABAoIBAGmb1jVRnSIe7Hee
TRI/D+eeTcYN4iww95b+zQP6xbfFd3RxUMqXFW9NJBLCv8WCX5dOMi0UIQJKe7cg
+9k4DI63dvs8lpKXlEqUzIkjHjs4gt3KF8/HID8R59/9y89wXVRLBxHmO4Dhhqaq
TYjIPG3OK643kb48WcJN8xQJEfGeTo8EaH+sbmxRoAGKDvghCaWWnDpOdrTQu378
yhxcGRTKnoBZ9ZwVorTn+PSGLHk6lP2npaNLUavYsH0KxOjk5H6Ivkr9uxUL8k3D
jDStxoXajIl27BzzUNrRaVrSkstJ5hxV8+pt0x9XfSfj7IDvkw1i+At39KUl/1Ge
TvDORm0CgYEA808J7AmPOACiSrZFgnjDScYxyfQ2o1+SKZvY0EX/0XT61EypAm/Q
fRcloHFgUpMX35AFwkLxQvbr3WNwQ/FxL66fa9MknatGRjJ5Btpwtm9tPU+8s6zV
ERBILV3lU2QKlyMUEc4eoY0YfeutDVBXsmmIaRYyXutszdD0fySlEycCgYEAzYXs
mf7VwGkpulXJpA1853qjU3sznkoWb+24slde8GJO/5grA1oLmvj5PZZZeScMHs7y
2AVvks7nChN58Bacow44HfoWRzzljJIyzvk3FJQU/pxyYWDqjIWEHyz1QpPwGnAn
sbXDPlL8haK3fvHckCxEUsV4ariYk9off1TLYeMCgYAOhBqtcGyRBMip+HHxNM9B
6Ycy61UPLjaEMP4gfuyiTH5iiChVMuKXc/gUuG7svkXgWrFdQ60tN5oe3T3nb9I2
7A4q2rAGkB9jNMOvuvyahc9Ypxc1CQy1Nel4e2+hKAjRt5AM9+Uv8kER0ivC7ZYt
2MoAqwhaTWVyahMMOt12ZQKBgFSL01lGTAhrOo8820ZxhgBzotykwgdL0w8Ya/Jx
nsykpHrgzlECqeOGqIF9aDE+ru29lIjpM4zhMIokmPT7WFR7dPpf9uo3UvcQ7XXX
b6E3iat0EFme1N4ZztBEXmCVGyDFIo8ohrkMynTvCy1dsa5dh33FXlQIvDuWb36i
E0ihAoGAJ6NtUi67b+X1qOHoZsbX7QJlG8rlzbU6ICTu3Uwwo3JIzjEX3P8KW+TG
2TNqD6LL4AgC3CV0n4V5RTiW9v+5PgQxZed3drDBYee8yOy10zoxSouMZk3M+Ht5
iRp6Zaz/mXn3pegEJN+jz2JCquuFs9fakucLTOXqIuLK+rshi3Q=
-----END RSA PRIVATE KEY-----`

	aliPriKeyPKCS8 = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDj//wXZNW0ZNTXxcB1sng75LwP5iu1A56kOcDJEyWLnY1A9t/h7jVemuwhBhnyqRtGOYsJNcj4/PmDkBlldVdDbzg5jQ5TJLkvbxKURfFPIW6/+Me2+3Va8c60w9cvp6nJ6e+jY4AgQ/t8BVCPJ+1YT3D7HUDTUnlA6Vau5zgzXl7CSnkOzKN+M0XkwbnXBKZxawYIAhnQ8482u9/+li1par/h4roYkCUn6jK1ppkETRP4yligvaOTqtqc5io8rstUszLe3FEgky2mONv4NzZLBXe0h5grC5i7riB/oWJCMfClIQVD4IUx0dhgSjEiC3xbXEHdIYidXp3kdn1//AtbAgMBAAECggEBAIAT2boOg9oNVl9LkBzNv6rtiBNqWh0aRnC8nKaXody3CZV32TZiiQp2/rRiaAsfodmmE4t98kKkYUndldafvwB7ptr1eRmNpJWX0JkIXVD1V/BxQUdyJG+NLG/xH51+kuuBBYQELgJyDgCSP7nbVQHpxm7Z9QOA7epf6bjg3HpFthbQVd8fziHwn/3ZfnX3ZicXDufUEdne4iROQyvT5rt7VWrKOB/KgOtbRzyoQ3n9Q3QPNKFdC4m1d+9nXvPrVG/ypdjMhU8cVMCsoKq0LmR79Hk5DpnPfLEZvmCuBqkOCf1YlfTeT01I8aTpa+/Vz7a6SBs45KYHeh/a2OPmV6kCgYEA/9NoERuRPeO+KwiHzHQ2OyxZwXQg5a/kUQK/iIh4RzJyFtIZ+1T/FjqBvYeSL4CQKMdMgPSEbUxEEL6ldwGUVo0A95wmLI9XlZBXE/JOr8ARxh6+TTKiO6BCY8DRg7w/dW6KAROODQ+HrtfIc7yGO2WsWEB3Lg/QZ76aqyQywU8CgYEA5Ce6U8Kn4vPJRBCDvtd7tD7Ow0ndZS5NHB8p3v6+Ce7jkhCgyjWqUrS4X1Fr/jCLvToqSgPl2E75WpidyynXor30GSZPYWheVCXdp8pZ3PAOXkrnwhhLKtPvFQiLzBakQIO1kZpNK1xxEZCbfYoAVk+MuJlQU0JSnEKukaKoGjUCgYEA52bCazcUWXb3PNWqgDVeGIg1Z9czKo+/ZcemmElZAo2OxTWmasCnO+QQFu/aOz1/KGOdv8dtQbAKoVrRMKQtGVzR0ydbIFNIVfEJ0vEKW87QRPTqDcoPFQARDUoMOay9eetSyUZ25zl08aIdnxKYGIIPMoSdSPzk9fJi1QsZHeECgYB+S5eTAi/ZK8X/zXyASdwKXLIr3FmXxtecd4vACxe7e+exVU8MYESWX7+n0/SJ0umnp+j868dxw9W4Vr8SVWkuhyP2QAS1EQcWYZefaB3ydVgeziwYf0YswQCZjg+EvAvUQpRhsx6AoD1e0cOg7KFfJc3p6cRTMmV0Zl0n9PhE+QKBgGfd+oW3E59XtnEzFuK9vPXlENpnSTuTpztJ/8dp8M4v5VOsKLqUhd/90aNbQV1tPGYhBIwB7l9USrMVo1bP/+A81asdEOKTN/L2MgoN7+fFXSUQ5hc+YaNt3P80fqSWmVqurTbGZeTqnDF3XVt+rcC9yStR2PLCJF8bDoLX5s77"
	aliPubKeyPKCS8 = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4//8F2TVtGTU18XAdbJ4O+S8D+YrtQOepDnAyRMli52NQPbf4e41XprsIQYZ8qkbRjmLCTXI+Pz5g5AZZXVXQ284OY0OUyS5L28SlEXxTyFuv/jHtvt1WvHOtMPXL6epyenvo2OAIEP7fAVQjyftWE9w+x1A01J5QOlWruc4M15ewkp5DsyjfjNF5MG51wSmcWsGCAIZ0POPNrvf/pYtaWq/4eK6GJAlJ+oytaaZBE0T+MpYoL2jk6ranOYqPK7LVLMy3txRIJMtpjjb+Dc2SwV3tIeYKwuYu64gf6FiQjHwpSEFQ+CFMdHYYEoxIgt8W1xB3SGInV6d5HZ9f/wLWwIDAQAB"

	label = "www.igoogle.ink"
)

func TestRsaEncryptAndDecryptData(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("数据：", originData)
	encryptData, err := RsaEncryptDataV2(PKCS8, []byte(originData), FormatAlipayPublicKey(aliPubKeyPKCS8))
	if err != nil {
		xlog.Error("RsaEncryptData:", err)
		return
	}
	origin, err := RsaDecryptDataV2(PKCS8, encryptData, FormatAlipayPrivateKey(aliPriKeyPKCS8))
	if err != nil {
		xlog.Error("RsaDecryptData:", err)
		return
	}
	xlog.Debug("decrypt:", string(origin))
}

func TestRsaEncryptOAEPAndDecryptOAEPData(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("数据：", originData)
	encryptData, err := RsaEncryptOAEPData(sha256.New(), PKCS1, pubKey, []byte(originData), []byte(label))
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
	origin, err := RsaDecryptOAEPData(sha256.New(), PKCS1, priKey, bytes, []byte(label))
	if err != nil {
		xlog.Error("RsaDecryptData:", err)
		return
	}
	xlog.Debug("decrypt:", string(origin))
}

func TestRsaEncryptOAEPAndDecryptOAEPDataPKCS8(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("数据：", originData)
	encryptData, err := RsaEncryptOAEPData(sha256.New(), PKCS8, FormatAlipayPublicKey(aliPubKeyPKCS8), []byte(originData), []byte(label))
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
	origin, err := RsaDecryptOAEPData(sha256.New(), PKCS8, FormatAlipayPrivateKey(aliPriKeyPKCS8), bytes, []byte(label))
	if err != nil {
		xlog.Error("RsaDecryptData:", err)
		return
	}
	xlog.Debug("decrypt:", string(origin))
}
