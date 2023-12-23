package gopay

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"testing"

	"github.com/go-pay/xlog"
)

func TestBodyMap_CheckParamsNull(t *testing.T) {
	bm := make(BodyMap)
	bm.Set("name", "jerry")
	bm.Set("age", 2)
	bm.Set("phone", "123")
	bm.Set("pi", 3.1415926)

	err := bm.CheckEmptyError("name", "age", "phone")
	if err != nil {
		xlog.Errorf("bm.CheckEmptyError(),err: %+v", err)
		return
	}
	h := sha1.New()
	h.Write([]byte("golang"))
	bs := h.Sum(nil)
	_signature := hex.EncodeToString(bs)
	xlog.Info(_signature) // 771e417b9dcae54aead2f3cbbbff340787bc462f
	// 771e417b9dcae54aead2f3cbbbff340787bc462f
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
		xlog.Errorf("xml.Unmarshal(%s),error:%+v", xmlData, err)
		return
	}
	for k, v := range mm {
		xlog.Infof("%s:%s\n", k, v)
	}
}

type QueryRedRecordResponse struct {
	HbType string  `xml:"hb_type,omitempty" json:"hb_type,omitempty"`
	Hblist *hbList `xml:"hblist,omitempty" json:"hblist,omitempty"`
}
type hbList struct {
	HbinfoList []*hbinfo `xml:"hbinfo,omitempty" json:"hbinfo,omitempty"`
}

type hbinfo struct {
	Openid  string `xml:"openid,omitempty" json:"openid,omitempty"`
	Amount  string `xml:"amount,omitempty" json:"amount,omitempty"`
	RcvTime string `xml:"rcv_time,omitempty" json:"rcv_time,omitempty"`
}

func TestBodyMap_UnmarshalXML2(t *testing.T) {
	xmlData := `
<xml>
<hb_type><![CDATA[NORMAL]]></hb_type>
<hblist>
	<hbinfo>
		<openid><![CDATA[111]]></openid>
		<amount>222</amount>
		<rcv_time><![CDATA[333]]></rcv_time>
	</hbinfo>
	<hbinfo>
		<openid><![CDATA[444]]></openid>
		<amount>555</amount>
		<rcv_time><![CDATA[666]]></rcv_time>
	</hbinfo>
</hblist>
</xml>`

	mm := new(QueryRedRecordResponse)
	err := xml.Unmarshal([]byte(xmlData), &mm)
	if err != nil {
		xlog.Errorf("xml.Unmarshal(%s),error:%+v", xmlData, err)
		return
	}
	xlog.Debugf("%+v", mm)
	for _, v := range mm.Hblist.HbinfoList {
		xlog.Debugf("%+v", v)
	}
}
