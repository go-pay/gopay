package gopay

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"

	"github.com/go-pay/gopay/pkg/util"
)

type BodyMap map[string]any

type xmlMapMarshal struct {
	XMLName xml.Name
	Value   any `xml:",cdata"`
}

type xmlMapUnmarshal struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

// 设置参数
func (bm BodyMap) Set(key string, value any) BodyMap {
	bm[key] = value
	return bm
}

func (bm BodyMap) SetBodyMap(key string, value func(b BodyMap)) BodyMap {
	_bm := make(BodyMap)
	value(_bm)
	bm[key] = _bm
	return bm
}

// 设置 FormFile
func (bm BodyMap) SetFormFile(key string, file *util.File) BodyMap {
	bm[key] = file
	return bm
}

// 获取参数，同 GetString()
func (bm BodyMap) Get(key string) string {
	return bm.GetString(key)
}

// 获取参数转换string
func (bm BodyMap) GetString(key string) string {
	if bm == nil {
		return NULL
	}
	value, ok := bm[key]
	if !ok {
		return NULL
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

// 获取原始参数
func (bm BodyMap) GetInterface(key string) any {
	if bm == nil {
		return nil
	}
	return bm[key]
}

// 删除参数
func (bm BodyMap) Remove(key string) {
	delete(bm, key)
}

// 置空BodyMap
func (bm BodyMap) Reset() {
	for k := range bm {
		delete(bm, k)
	}
}

func (bm BodyMap) JsonBody() (jb string) {
	bs, err := json.Marshal(bm)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

// Unmarshal to struct or slice point
func (bm BodyMap) Unmarshal(ptr any) (err error) {
	bs, err := json.Marshal(bm)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}

func (bm BodyMap) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if len(bm) == 0 {
		return nil
	}
	start.Name = xml.Name{Space: NULL, Local: "xml"}
	if err = e.EncodeToken(start); err != nil {
		return
	}
	for k := range bm {
		if v := bm.GetString(k); v != NULL {
			_ = e.Encode(xmlMapMarshal{XMLName: xml.Name{Local: k}, Value: v})
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
	if bm == nil {
		return NULL
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm.GetString(k); v != NULL {
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
	if bm == nil {
		return NULL
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm.GetString(k); v != NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return NULL
	}
	return buf.String()[:buf.Len()-1]
}

// ("bar=baz&foo=quux") sorted by key.
func (bm BodyMap) EncodeURLParams() string {
	if bm == nil {
		return NULL
	}
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range bm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v := bm.GetString(k); v != NULL {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return NULL
	}
	return buf.String()[:buf.Len()-1]
}

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.GetString(k); v == NULL {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return fmt.Errorf("[%w], %v", MissParamErr, strings.Join(emptyKeys, ", "))
	}
	return nil
}

func (bm BodyMap) CheckNotAllEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.GetString(k); v == NULL {
			emptyKeys = append(emptyKeys, k)
		}
	}
	// if all key is empty, return error
	if len(emptyKeys) == len(keys) {
		return fmt.Errorf("[%w], %v", MissParamErr, strings.Join(emptyKeys, ", "))
	}
	return nil
}

func convertToString(v any) (str string) {
	if v == nil {
		return NULL
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return NULL
	}
	str = string(bs)
	return
}
