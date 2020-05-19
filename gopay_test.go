package gopay

import (
	"fmt"
	"testing"
	"time"
)

func TestBodyMap_CheckParamsNull(t *testing.T) {
	bm := make(BodyMap)
	bm.Set("name", "jerry")
	bm.Set("age", 2)
	bm.Set("phone", "")
	bm.Set("pi", 3.1415926)

	err := bm.CheckEmptyError("name", "age", "phone")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

func TestNewHttpClient(t *testing.T) {
	client := NewHttpClient()
	res, _, errs := client.Get("http://www.baidu.com").SetTimeout(30 * time.Second).EndBytes()
	if len(errs) > 0 {
		fmt.Println("err:", errs[0])
		return
	}
	fmt.Println("bs:", res.StatusCode)
}
