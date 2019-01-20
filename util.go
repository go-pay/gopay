package gopay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/parnurzeal/gorequest"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type BodyMap map[string]interface{}

//设置参数
func (bm BodyMap) Set(key string, value interface{}) {
	bm[key] = value
}

//获取参数
func (bm BodyMap) Get(key string) string {
	if bm == nil {
		return ""
	}
	v := bm[key]
	value, ok := v.(int)
	if ok {
		value := strconv.Itoa(value)
		return value
	}
	return v.(string)
}

//删除参数
func (bm BodyMap) Remove(key string) {
	delete(bm, key)
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

//获取根据Key排序后的请求参数字符串
func sortSignParams(secretKey string, body BodyMap) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(k)
		buffer.WriteString("=")
		value, ok := body[k].(int)
		if ok {
			value := strconv.Itoa(value)
			buffer.WriteString(value)
		} else {
			buffer.WriteString(body[k].(string))
		}
		buffer.WriteString("&")
	}
	buffer.WriteString("key=")
	buffer.WriteString(secretKey)
	return buffer.String()
}

//从微信提供的接口获取：SandboxSignkey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	reqXml := generateXml(reqs)
	//fmt.Println("req:::", reqXml)
	_, byteList, errorList := gorequest.New().
		Post(wxURL_SanBox_GetSignKey).
		Type("xml").
		SendString(reqXml).EndBytes()
	if len(errorList) > 0 {
		return "", errorList[0]
	}
	keyResponse := new(getSignKeyResponse)
	err = xml.Unmarshal(byteList, keyResponse)
	if err != nil {
		return "", err
	}
	if keyResponse.ReturnCode == "FAIL" {
		return "", errors.New(keyResponse.Retmsg)
	}
	return keyResponse.SandboxSignkey, nil
}
