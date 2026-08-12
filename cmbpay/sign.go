package cmbpay

import (
	"sort"
	"strings"
)

// buildSignString 依据招行签名规则拼接待签名字符串（详见接口文档 2.4.1.1 / 2.4.1.2）：
//
//  1. 获取所有参数，剔除 sign 字段；
//  2. 按参数名（key）的 ASCII 码递增排序（字母升序）；
//  3. 组合成 "参数=参数值" 并用 & 连接。
//
// 说明：空字符串值的字段仍参与拼接（招行样例中值为空的字段照常参与），
// 调用方在装配 params 时应只放入实际参与签名的字段。
func buildSignString(params map[string]string) string {
	// 先剔除 sign，再排序，避免 sign 排序靠前时产生前导 '&'。
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	for i, k := range keys {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(params[k])
	}
	return b.String()
}
