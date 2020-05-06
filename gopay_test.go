package gopay

import (
	"fmt"
	"testing"
)

func TestBodyMap_CheckParamsNull(t *testing.T) {
	bm := make(BodyMap)
	bm.Set("name", "jerry")
	bm.Set("age", 2)
	bm.Set("phone", "")
	bm.Set("pi", 3.1415926)
	bm.Set("pi2", 3.1415926)
	err := bm.CheckEmptyError("name", "age", "phone")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
