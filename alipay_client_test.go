package gopay

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type List struct {
	BillList []fundBillListInfo `json:"bill_list"`
}

func TestJsonToString(t *testing.T) {

	list := new(List)
	infos := make([]fundBillListInfo, 0)

	infos = append(infos, fundBillListInfo{Amount: "1.0.0", FundChannel: "iguiyu"})
	infos = append(infos, fundBillListInfo{Amount: "2.0.2", FundChannel: "Jerry"})

	list.BillList = infos

	bs, err := json.Marshal(list)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("string:", string(bs))
}

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestAliPayParams(t *testing.T) {
	bodyMap := make(BodyMap)

	//people := new(People)
	//people.Name = "Jerry"
	//people.Age = 18
	people := make(map[string]interface{})
	people["name"] = "jerry"
	people["age"] = 18
	bodyMap.Set("people", people)

	fmt.Println("result:", bodyMap.Get("people"))
}

func TestVerifyAliPaySign(t *testing.T) {
	signData := `{"code":"10000","msg":"Success","buyer_logon_id":"854***@qq.com","buyer_pay_amount":"0.01","buyer_user_id":"2088102363632794","fund_bill_list":[{"amount":"0.01","fund_channel":"PCREDIT"}],"gmt_payment":"2019-08-29 20:14:05","invoice_amount":"0.01","out_trade_no":"GZ201901301040361012","point_amount":"0.00","receipt_amount":"0.01","total_amount":"0.01","trade_no":"2019082922001432790585537960"}`
	sign := "bk3SzX0CZRI811IJioS2XKQHcgMixUT8mYyGQj+vcOAQas7GIYi6LpykqqSc3m7+yvqoG0TdX/c2JjYnpw/J53JxtC2IC4vsLuIPIgghVo5qafsfSxEJ22w20RZDatI2dYqFVcj8Jp+4aesQ8zMMNw7cX9NLyk7kw3DecYeyQp+zrZMueZPqLh88Z+54G+e6QuSU++0ouqQVd4PkpPqy6YI+8MdMUX4Ve0jOQxMmYH8BC6n5ZsTH/uEaLEtzYVZdSw/xdSQ7K1SH73aEH8XbRYx6rL7RkKksrdvhezX+ThDjQ+fTWjvNFrGcg3fmqXRy2elvoalu+BQmqlkWWjEJYA=="
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"
	pKey := FormatAliPayPublicKey(aliPayPublicKey)
	err := verifyAliPaySign(signData, sign, "RSA2", pKey)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestSubString(t *testing.T) {
	str := `{"alipay_trade_pay_response":{"code":"10000","msg":"Success","buyer_logon_id":"854***@qq.com","buyer_pay_amount":"0.01","buyer_user_id":"2088102363632794","fund_bill_list":[{"amount":"0.01","fund_channel":"PCREDIT"}],"gmt_payment":"2019-08-29 20:22:02","invoice_amount":"0.01","out_trade_no":"GZ201901301040361013","point_amount":"0.00","receipt_amount":"0.01","total_amount":"0.01","trade_no":"2019082922001432790585666965"},"sign":"DSX/wmE0nnuxQrWfJZtq0fNntcx5UYtVV35P2VZpoTC2KlIWr4eGNiXcetbb7AkI/1Tyd0+cNtcGMgB7SYzTB15/wDE0vJ+eT5ucqhNkER1kcuCC0k9OkZzU5w8wCJzOgAy52Wso9KnrwkY86mJWt3dC8DNCCi1rlf1a8bTGIBG/diJaKAgP1lGT3aW8jeGGM98zLabqDUNvck2qkgctGR49kBb0ZYmIzmY0x5goVyKnaCkcC/d1VTIIMz81mJbeqU8UZk6TqEplCC8J+dYEUj04pAO4/lwIg/YZdKj3Pz1136/+uy669Pew88+74J/u/zPsehC44PxcUk9YKmkNyw=="}`

	index := strings.Index(str, `":`)
	fmt.Println("index:", index)
	indexEnd := strings.Index(str, `,"sign"`)
	fmt.Println("indexEnd:", indexEnd)

	fmt.Println("sub:", str[index+2:indexEnd])
}

func TestGetCertSN(t *testing.T) {

	//sn, err := GetCertSN("alipay_cert/alipayCertPublicKey_RSA2.crt")
	//sn, err := GetCertSN("alipay_cert/appCertPublicKey.crt")
	//sn, err := GetCertSN("alipay_cert/alipayRootCert.crt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//fmt.Println("sn:", sn)
}
