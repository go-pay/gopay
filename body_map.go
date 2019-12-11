package gopay

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"sort"
	"strings"
)

type BodyMap map[string]interface{}

// 设置参数
func (bm BodyMap) Set(key string, value interface{}) {
	bm[key] = value
}

// 获取参数
func (bm BodyMap) Get(key string) string {
	if bm == nil {
		return null
	}
	var (
		value interface{}
		ok    bool
		v     string
	)
	if value, ok = bm[key]; !ok {
		return null
	}
	if v, ok = value.(string); ok {
		return v
	}
	return convertToString(value)
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return null
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return null
	}
	str = string(bs)
	return
}

// 删除参数
func (bm BodyMap) Remove(key string) {
	delete(bm, key)
}

type xmlMapMarshal struct {
	XMLName xml.Name
	Value   interface{} `xml:",cdata"`
}

type xmlMapUnmarshal struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

func (bm BodyMap) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if len(bm) == 0 {
		return nil
	}
	start.Name = xml.Name{null, "xml"}
	if err = e.EncodeToken(start); err != nil {
		return
	}
	for k := range bm {
		if v := bm.Get(k); v != null {
			e.Encode(xmlMapMarshal{XMLName: xml.Name{Local: k}, Value: v})
		}
	}
	return e.EncodeToken(start.End())
}

func (bm *BodyMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for {
		var e xmlMapUnmarshal
		err = d.Decode(&e)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		bm.Set(e.XMLName.Local, e.Value)
	}
}

// ("bar=baz&foo=quux") sorted by key.
func (bm BodyMap) EncodeWeChatSignParams(apiKey string) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm.Get(k); v != null {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	buf.WriteString("key")
	buf.WriteByte('=')
	buf.WriteString(apiKey)
	return buf.String()
}

// ("bar=baz&foo=quux") sorted by key.
func (bm BodyMap) EncodeAliPaySignParams() string {
	var (
		buf     strings.Builder
		keyList []string
	)
	keyList = make([]string, 0, len(bm))
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm.Get(k); v != null {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	return buf.String()[:buf.Len()-1]
}
