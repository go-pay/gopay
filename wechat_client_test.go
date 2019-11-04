package gopay

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="
	phone := new(WeChatUserPhone)
	//解密开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    iv:加密算法的初始向量
	//    sessionKey:会话密钥
	//    beanPtr:需要解析到的结构体指针
	err := DecryptWeChatOpenDataToStruct(data, iv, session, phone)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("PhoneNumber:", phone.PhoneNumber)
	fmt.Println("PurePhoneNumber:", phone.PurePhoneNumber)
	fmt.Println("CountryCode:", phone.CountryCode)
	fmt.Println("Watermark:", phone.Watermark)

	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

	//微信小程序 用户信息
	userInfo := new(WeChatAppletUserInfo)

	err = DecryptWeChatOpenDataToStruct(encryptedData, iv2, sessionKey, userInfo)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("NickName:", userInfo.NickName)
	fmt.Println("AvatarUrl:", userInfo.AvatarUrl)
	fmt.Println("Country:", userInfo.Country)
	fmt.Println("Province:", userInfo.Province)
	fmt.Println("City:", userInfo.City)
	fmt.Println("Gender:", userInfo.Gender)
	fmt.Println("OpenId:", userInfo.OpenId)
	fmt.Println("UnionId:", userInfo.UnionId)
	fmt.Println("Watermark:", userInfo.Watermark)
}

func TestDecryptWeChatOpenDataToBodyMap(t *testing.T) {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="

	//解密开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    iv:加密算法的初始向量
	//    sessionKey:会话密钥
	bm, err := DecryptWeChatOpenDataToBodyMap(data, iv, session)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("WeChatUserPhone:", bm)
}
