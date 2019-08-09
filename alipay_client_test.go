package gopay

import (
	"encoding/json"
	"fmt"
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
