package v3

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
)

func TestClientRSA2(t *testing.T) {
	block, _ := pem.Decode([]byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY=\n-----END RSA PRIVATE KEY-----"))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatal(err)
		os.Exit(1)
	}
	c := NewClient(key, "", "")
	if str, _ := c.rsa2("test"); str != "cX2ta9egtLPv9SRaWNNpNgRVNAEgfQG5KLNox9dtE7K2FnbPC4ajCF+LXeo+KyVG+d6cpJy+qZzYk1I0cvtN5mS/zyEfw4VnESjaOi6SZS3sCsEak6oDAtgGp3O63/W+JJkAzaV7MGbSgi9MuhZhjzlnTrTtkjJXZJ2j4DVY5iPYFf6keCU+VYMtsfbjPLRDA90et1K5nN7rbRwycmgGlMW+ijExQLmKgSoTCU4/0Pp1g9pajuN9ugGBfhRvdIb2Q5JI4tJ0GBzejv9aNqVqztJKzHUK7B6CqKkvN+O7ZWDmiAK7MG7qFjo64IahO5fCUuMwqp0hDNy7K8ioM8kROw==" {
		t.Fatal("test rsa2 fail")
	}
	str, err := c.authorization("GET", "/test", "1606205766", "HJGKBT(&", "")
	if str != `WECHATPAY2-SHA256-RSA2048 mchid="",nonce_str="HJGKBT(&",timestamp="1606205766",serial_no="",signature="soHh+jRyPNq9NAujXm8T4E+6uzwx6usyLhVwAaZs1JqDe0+dLMlWipVJW5qm1qf7ioVHMjBOo+gtx/7Pba7YfHTWsqsV+afJRd1sBIESerJGe04Zw1ywVsAuHg8T2w46zPya+Ir7+M2i3649u54bIGcYm0jZ15uauXWWm0mCmfaLufI/duI76CJ9C9oggQO+sWPcVmEbmuT/X7ZEffM5+PGwem3/Spds0B6L00VoMIFKUZrzJxj9qpX/kMkxGz17vmiNqDAD1V4VcI6R1kz5vLwBC4mitUYjzGbUEtx8LKOsdNSJI5FWUJFIJEUa066EaUCVgdRnpXfCZv/vjlDC7A=="` {
		t.Fatal("get authorization header error")
	}
}
