package gopay

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"sort"
	"strings"
	"sync"
)

type BodyMap map[string]interface{}

var mu sync.RWMutex

// 设置参数
func (bm BodyMap) Set(key string, value interface{}) {
	mu.Lock()
	bm[key] = value
	mu.Unlock()
}

// 获取参数
func (bm BodyMap) Get(key string) string {
	if bm == nil {
		return NULL
	}
	mu.RLock()
	defer mu.RUnlock()
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

// GetArrayMap 获取获取的key对应的值，并尝试转换为 []BodyMap
// 在请求单次分账接口输入参数有用到 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
func (bm BodyMap) GetArrayBodyMap(key string) (array []BodyMap, err error) {
	if bm == nil {
		return nil, errors.New("param is nil")
	}
	mu.RLock()
	defer mu.RUnlock()
	value, ok := bm[key]
	if !ok {
		return nil, errors.New("param is not exist")
	}
	array, ok = value.([]BodyMap)
	if !ok {
		return nil, errors.New("can not transfer " + key + " to []BodyMap")
	}
	return
}

// GetBodyMap  获取获取的key对应的值，并尝试转换为 BodyMap
// 在添加分账接收方会用到 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4
func (bm BodyMap) GetBodyMap(key string) (b BodyMap, err error) {
	if bm == nil {
		return nil, errors.New("param is nil")
	}
	mu.RLock()
	defer mu.RUnlock()
	value, ok := bm[key]
	if !ok {
		return nil, errors.New("param is not exist")
	}
	b, ok = value.(BodyMap)
	if !ok {
		return nil, errors.New("can not transfer " + key + " to BodyMap")
	}
	return
}

func convertToString(v interface{}) (str string) {
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

// 删除参数
func (bm BodyMap) Remove(key string) {
	mu.Lock()
	delete(bm, key)
	mu.Unlock()
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
	start.Name = xml.Name{NULL, "xml"}
	if err = e.EncodeToken(start); err != nil {
		return
	}
	for k := range bm {
		if v := bm.Get(k); v != NULL {
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
	mu.RLock()
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	mu.RUnlock()
	for _, k := range keyList {
		if v := bm.Get(k); v != NULL {
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
	mu.RLock()
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	mu.RUnlock()
	for _, k := range keyList {
		if v := bm.Get(k); v != NULL {
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

func (bm BodyMap) EncodeGetParams() string {
	var (
		buf strings.Builder
	)
	for k, _ := range bm {
		if v := bm.Get(k); v != NULL {
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

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.Get(k); v == NULL {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return errors.New(strings.Join(emptyKeys, ", ") + " : cannot be empty")
	}
	return nil
}
