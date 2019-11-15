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

func TestDecryptRefundNotifyReqInfo(t *testing.T) {
	key := "ziR0QKsTUfMOuochC9RfCdmfHECorQAP"
	data := "YYwp8C48th0wnQzTqeI+41pflB26v+smFj9z6h9RPBgxTyZyxc+4YNEz7QEgZNWj/6rIb2MfyWMZmCc41CfjKSssoSZPXxOhUayb6KvNSZ1p6frOX1PDWzhyruXK7ouNND+gDsG4yZ0XXzsL4/pYNwLLba/71QrnkJ/BHcByk4EXnglju5DLup9pJQSnTxjomI9Rxu57m9jg5lLQFxMWXyeASZJNvof0ulnHlWJswS4OxKOkmW7VEyKyLGV6npoOm03Qsx2wkRxLsSa9gPpg4hdaReeUqh1FMbm7aWjyrVYT/MEZWg98p4GomEIYvz34XfDncTezX4bf/ZiSLXt79aE1/YTZrYfymXeCrGjlbe0rg/T2ezJHAC870u2vsVbY1/KcE2A443N+DEnAziXlBQ1AeWq3Rqk/O6/TMM0lomzgctAOiAMg+bh5+Gu1ubA9O3E+vehULydD5qx2o6i3+qA9ORbH415NyRrQdeFq5vmCiRikp5xYptWiGZA0tkoaLKMPQ4ndE5gWHqiBbGPfULZWokI+QjjhhBmwgbd6J0VqpRorwOuzC/BHdkP72DCdNcm7IDUpggnzBIy0+seWIkcHEryKjge3YDHpJeQCqrAH0CgxXHDt1xtbQbST1VqFyuhPhUjDXMXrknrGPN/oE1t0rLRq+78cI+k8xe5E6seeUXQsEe8r3358mpcDYSmXWSXVZxK6er9EF98APqHwcndyEJD2YyCh/mMVhERuX+7kjlRXSiNUWa/Cv/XAKFQuvUYA5ea2eYWtPRHa4DpyuF1SNsaqVKfgqKXZrJHfAgslVpSVqUpX4zkKszHF4kwMZO3M7J1P94Mxa7Tm9mTOJePOoHPXeEB+m9rX6pSfoi3mJDQ5inJ+Vc4gOkg/Wd/lqiy6TTyP/dHDN6/v+AuJx5AXBo/2NDD3fWhHjkqEKIuARr2ClZt9ZRQO4HkXdZo7CN06sGCHk48Tg8PmxnxKcMZm7Aoquv5yMIM2gWSWIRJhwJ8cUpafIHc+GesDlbF6Zbt+/KXkafJAQq2RklEN+WvZ/zFz113EPgWPjp16TwBoziq96MMekvWKY/vdhjol8VFtGH9F61Oy1Xwf6DJtPw=="
	refundNotify, err := DecryptRefundNotifyReqInfo(data, key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("refundNotify:", *refundNotify)
}

func TestBodyMap_Set_Get(t *testing.T) {
	bm := make(BodyMap)
	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.gopay.ink"
	h5Info["wap_name"] = "H5测试支付"
	sceneInfo["h5_info"] = h5Info
	bm.Set("scene_info", sceneInfo)
	fmt.Println("Get1:", bm.Get("scene_info"))

	bm.Set("student", &Student{
		Name:  "Jerry",
		Age:   26,
		Sign:  "123465",
		Phone: "123654987",
	})

	fmt.Println("Get2:", bm.Get("student"))

}
