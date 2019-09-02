package gopay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestMd5(t *testing.T) {
	st := "appid=wxdaa2ab9ef87b5497&nonceStr=9k20rM66parD2U49&package=prepay_id=wx29164301554772fbc70d1d793335446010&signType=MD5&timeStamp=1548751382&key=GFDS8j98rewnmgl45wHTt980jg543wmg"
	hash := md5.New()
	hash.Write([]byte(st))
	sum := hash.Sum(nil)
	upper := strings.ToUpper(hex.EncodeToString(sum))
	fmt.Println(" ssad  ", upper)
}

type Student struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Sign  string `json:"sign,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func TestBodyMap_MarshalXML(t *testing.T) {

	student := new(Student)
	student.Name = "Jerry"
	student.Age = 28
	student.Phone = "18017448610"

	marshal, err := json.Marshal(student)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("marshal:", string(marshal))

	maps := make(map[string]interface{})

	err = json.Unmarshal(marshal, &maps)
	if err != nil {
		fmt.Println("err2:", err)
	}
	fmt.Println("maps:", maps)

	//maps := make(BodyMap)
	//maps.Set("name", "jerry")
	//maps.Set("age", 28)
	//maps.Set("phone", "13212345678")
	//
	//bytes, err := xml.Marshal(&maps)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//fmt.Println("ssss:", string(bytes))

}

func TestVerifyWeChatResponseSign(t *testing.T) {
	student := new(Student)
	student.Name = "Jerry"
	student.Age = 1
	student.Sign = "544E55ED43B787B945FF0BF8344A4D69"
	student.Phone = "18017448610"

	maps := make(BodyMap)
	maps["name"] = "Jerry"
	maps["age"] = 1
	maps["sign"] = "544E55ED43B787B945FF0BF8344A4D69"
	maps["phone"] = "18017448610"

	ok, err := VerifyWeChatSign("ABCDEFG", "MD5", student)
	if err != nil {
		fmt.Println("errrrr:", err)
		return
	}
	fmt.Println("ok:", ok)
}

func TestDecryptWeChatOpenDataToStruct(t *testing.T) {
	//data := "HSIMnigFLkuKzFDVtHM2t2s423ZasY0DXst5Cma4Oih8Tke+HRnHX9G+PRey1SLg5ajCPtlguVBRDqOB+DWpHo3Emvza889koLTkV3M9X4tURa1UwQ4eKZ8A1WoGB4hktr3jFyRt2nMccwQLbpL20V1zR/uwvjYHeB4TDh9rsDxMm3WY5YemSW54b+mbp9BD1rTgakWoGaOwrlEaQUIwbg=="
	//iv := "PYLEcr/lIA/NhK7N6yhDPg=="
	//session := "jSsUV8GCyJJf5Qnz3noNKA=="

	//data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	//iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	//session := "lyY4HPQbaOYzZdG+JcYK9w=="
	//phone := new(WeChatUserPhone)
	////解密开放数据
	////    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	////    iv:加密算法的初始向量
	////    sessionKey:会话密钥
	////    beanPtr:需要解析到的结构体指针
	//err := DecryptWeChatOpenDataToStruct(data, iv, session, phone)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("PhoneNumber:", phone.PhoneNumber)
	//fmt.Println("PurePhoneNumber:", phone.PurePhoneNumber)
	//fmt.Println("CountryCode:", phone.CountryCode)
	//fmt.Println("Watermark:", phone.Watermark)
}
