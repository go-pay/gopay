package paygateway

import "bytes"

func bytesReader(bs []byte) *bytes.Reader {
	return bytes.NewReader(bs)
}
