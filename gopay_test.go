package gopay

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/iGoogle-ink/gotil/xhttp"
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
	client := xhttp.NewHttpClient()
	res, _, errs := client.Get("http://www.baidu.com").SetTimeout(30 * time.Second).EndBytes()
	if len(errs) > 0 {
		fmt.Println("err:", errs[0])
		return
	}
	fmt.Println("bs:", res.StatusCode)
}

func TestBodyMap_UnmarshalXML(t *testing.T) {
	xmlData := `<xml>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[TeqClE3i0mvn3DrK]]></nonce_str>
   <out_refund_no_0><![CDATA[1415701182]]></out_refund_no_0>
   <out_trade_no><![CDATA[1415757673]]></out_trade_no>
   <refund_count>1</refund_count>
   <refund_fee_0>1</refund_fee_0>
   <refund_id_0><![CDATA[2008450740201411110000174436]]></refund_id_0>
   <refund_status_0><![CDATA[PROCESSING]]></refund_status_0>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <sign><![CDATA[1F2841558E233C33ABA71A961D27561C]]></sign>
   <transaction_id><![CDATA[1008450740201411110005820873]]></transaction_id>
</xml>`

	mm := make(BodyMap)
	err := xml.Unmarshal([]byte(xmlData), &mm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for k, v := range mm {
		fmt.Printf("%s:%s\n", k, v)
	}
}
