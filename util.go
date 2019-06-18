package gopay

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
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
		return null
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

func convert2String(value interface{}) (valueStr string) {
	switch v := value.(type) {
	case int:
		valueStr = Int2String(v)
	case int64:
		valueStr = Int642String(v)
	case float64:
		valueStr = Float64ToString(v)
	case float32:
		valueStr = Float32ToString(v)
	case string:
		valueStr = v
	default:
		valueStr = null
	}
	return
}

//解析时间
func ParseDateTime(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation(TimeLayout, timeStr, time.Local)
	return
}

//格式化Datetime
func FormatDateTime(timeStr string) (formatTime string) {
	//2019-01-04T15:40:00Z
	//2019-01-18 20:51:30+08:00
	if timeStr == null {
		return null
	}
	replace := strings.Replace(timeStr, "T", " ", 1)
	formatTime = replace[:19]
	return
}

//格式化
func FormatDate(dateStr string) (formatDate string) {
	//2020-12-30T00:00:00+08:00
	if dateStr == null {
		return null
	}
	split := strings.Split(dateStr, "T")
	formatDate = split[0]
	return
}

//字符串转Float
func String2Float(floatStr string) (floatNum float64) {
	floatNum, _ = strconv.ParseFloat(floatStr, 64)
	return
}

//Float64转字符串
//    floatNum：float64数字
//    prec：精度位数（不传则默认float数字精度）
func Float64ToString(floatNum float64, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(floatNum, 'f', prec[0], 64)
		return
	}
	floatStr = strconv.FormatFloat(floatNum, 'f', -1, 64)
	return
}

//Float32转字符串
//    floatNum：float32数字
//    prec：精度位数（不传则默认float数字精度）
func Float32ToString(floatNum float32, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(float64(floatNum), 'f', prec[0], 32)
		return
	}
	floatStr = strconv.FormatFloat(float64(floatNum), 'f', -1, 32)
	return
}

//字符串转Int
func String2Int(intStr string) (intNum int) {
	intNum, _ = strconv.Atoi(intStr)
	return
}

//字符串转Int64
func String2Int64(intStr string) (int64Num int64) {
	intNum, _ := strconv.Atoi(intStr)
	int64Num = int64(intNum)
	return
}

//Int转字符串
func Int2String(intNum int) (intStr string) {
	intStr = strconv.Itoa(intNum)
	return
}

//Int64转字符串
func Int642String(intNum int64) (int64Str string) {
	//10, 代表10进制
	int64Str = strconv.FormatInt(intNum, 10)
	return
}

//解密填充模式（去除补全码） PKCS7UnPadding
//解密时，需要在最后面去掉加密时添加的填充byte
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])   //找到Byte数组最后的填充byte
	return plainText[:(length - unpadding)] //只截取返回有效数字内的byte数组
}

//格式化秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")

	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen

	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

//格式化秘钥
func FormatPublickKey(publickKey string) (pKey string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")

	rawLen := 64
	keyLen := len(publickKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen

	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publickKey[start:])
		} else {
			buffer.WriteString(publickKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}
