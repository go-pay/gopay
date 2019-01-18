package gopay

import (
	"bytes"
	"math/rand"
	"time"
)

type requestBody map[string]string

//设置参数
func (w requestBody) Set(key string, value string) {
	w[key] = value
}

//获取参数
func (w requestBody) Get(key string) string {
	if w == nil {
		return ""
	}
	ws := w[key]
	return ws
}

//删除参数
func (w requestBody) Remove(key string) {
	delete(w, key)
}

//获取随机字符串
//    length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func generateXml(w requestBody) (reqXml string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")

	for k, v := range w {
		buffer.WriteString("<")
		buffer.WriteString(k)
		buffer.WriteString("><![CDATA[")
		buffer.WriteString(v)
		buffer.WriteString("]]></")
		buffer.WriteString(k)
		buffer.WriteString(">")
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
