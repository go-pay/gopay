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
