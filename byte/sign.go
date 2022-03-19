package byte

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// 生成签名
func (c *Client) signature(timestamp, nonceStr, msg string) string {
	var signSlice []string
	signSlice = append(signSlice, c.Token)
	signSlice = append(signSlice, timestamp)
	signSlice = append(signSlice, nonceStr)
	signSlice = append(signSlice, msg)
	sort.Strings(signSlice)
	h := sha1.New()
	h.Write([]byte(strings.Join(signSlice, "")))
	return hex.EncodeToString(h.Sum(nil))
}
