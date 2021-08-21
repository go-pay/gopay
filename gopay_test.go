package gopay

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"testing"
	"time"

	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

func TestBodyMap_CheckParamsNull(t *testing.T) {
	bm := make(BodyMap)
	bm.Set("name", "jerry")
	bm.Set("age", 2)
	bm.Set("phone", "")
	bm.Set("pi", 3.1415926)

	err := bm.CheckEmptyError("name", "age", "phone")
	if err != nil {
		xlog.Errorf("bm.CheckEmptyError():error:%+v", err)
		return
	}
}

func TestNewClient(t *testing.T) {
	client := xhttp.NewClient()
	res, _, errs := client.Get("http://www.baidu.com").SetTimeout(30 * time.Second).EndBytes()
	if len(errs) > 0 {
		xlog.Error(errs[0])
		return
	}
	xlog.Info("bs:", res.StatusCode)
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

func TestPkcs(t *testing.T) {
	var PriKey = `-----BEGIN PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDV523KVXZaaZI3
WxQiaz0J8o8nxAYsxBjrfcaKPnLo+r5uFME7GPV+4UHEZWG6ZogJ87yBt8L4IV8q
/2n0MPKV5qNtS0htG7G0Mtyw7lPmdXUXsA0ionbL2mzz0kgJ1S6FJwyZRRZNJ08Q
/GQE3TWqErbxL/2ITuzTeHrdTNL0i9oNxtB92EWFZ0gSL677zEiz41EVog24SyOd
TFqxjGFd9DR0CeRNU/oQPplFnM9YFseRuhEZ/jLATEvubH/U1qGqTlW0UHvIn14j
NqRxyAjDI/HfXl3Bo7Fx0QCJkVkqb+5ou8KFRchbcixRU0khbrxTy7dDJj60vSmr
PySqqZLFAgMBAAECggEBAKHPN9ZfX/B0/A6z7z86MCpeOryyJJmondFGi/H326Uy
SOus959k+hDJBZ8zsgH3neEpZ+gYwnxBgmRcYiI/BMMwfWAoGtmuoXbXIusU3pLv
N2x72PPiQktjKBgpciU+BrrjFzy6bmxe2AjZZC/pxrapAYrh6sA6NBykfwz5GHu0
DQmjHYqSlghDDljCzVR3Gcs/KicCMw6eQ0JlWDqtDEDoENlBkfn4spHwocV7HtSq
0bnUrQqqMtpZjbMJzZxJc39qkyNNDosuNy5GXYLQE7lr9RuRqLxEfg6KfGUS5bAZ
eJ5pizql7+c0viUtiXG17PYp8QR4c5G+54RlQd1pPuECgYEA9UBi5rFJzK0/n4aO
lsrp6BvUOSherp57SNYvpsRuBPU0odyH2/McLNphisKTxfSm0/hADaTmnzAnOUVg
cduc/5/5tVaaqyLL3SemxJhwqVsL3tE/KAN7HUBhhQrqD+H8r39TAoIkyfjCOHzS
74rygZ35x0kXNMavXQFB0RE2fEcCgYEA30dWaLddGmTvUXwhyTWcsiDfrsKbw8+n
MhAlSCXE42v9Uo3ULqD3/rpUQlMhoqaZb3cSyOyQwJvv0tp/g3hM7Q4usLxkdysc
KA9HmmZ4Q2P2838cqvNr/Dz1UAnfdDryMEnbiv1MUKYqFFTVX6oR0iH544JgDFCG
YLQA2M+3GpMCgYEAh+ax51v+pSirxN5vTSgMDc69/x5buS+g6W+m4CahQKYQEFGA
B2XkCwbIXngMIvm7KGK8O9NQ6I1qbtX+55jmmtAvM0lWU9boWRiL1Q0UAQSuwz34
XVfwdPkkEPFHWp3DxAwuF4m+kR0DowGocYzxbNn5e3EJJvmiW0tDCXMcWikCgYEA
tyNxWcUFBdBCh+i0YbCqzWSvdE3Fq8/YSPT7T3lDTHLYPu18W57Gq1Y0JI7BaQMT
mVzmuI1pkcKV7LIxoyl6l3ppi6eLFD/1AVq/FYL1I/mLpl/dqM6vBR8O686dTV3I
Jxl9jTyEayZQH4sR1TzPDze1GwpmM9Oc1RbwFuYRPycCgYEAzYaRKh6EQ+s37HDv
e/ZGMs3PI+CoA/x6lx4Owa7amRsWRKys45NV6gcC8pkbN4IeFaYXVHmJ1Yaef3xn
0VxHAzWI4BF+1pUwXzS2rAMBZR/VKS0XA856NauAC3mKHipoOWVVs+uFP3VMUQ79
hSImAa7UBzss6b6ie7AYxXtZBjY=
-----END PRIVATE KEY-----`

	var (
		pk *rsa.PrivateKey
		ok bool
	)
	block, _ := pem.Decode([]byte(PriKey))
	if block == nil {
		return
	}
	xlog.Debugf("%+v", block.Type)
	//pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	//if err != nil {
	//	xlog.Errorf("1.err:%+v",err)
	//	pk, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	//	if err != nil {
	//		xlog.Errorf("2.err:%+v",err)
	//		return
	//	}
	//}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		xlog.Errorf("1.err:%+v", err)
		pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			xlog.Errorf("2.err:%+v", err)
			return
		}
		pk, ok = pk8.(*rsa.PrivateKey)
		if !ok {
			xlog.Warn("3.err:")
			return
		}
	}
	xlog.Debugf("%+v", pk.PublicKey)
	//xlog.Debugf("%+v", pub)
	//publicKey, ok := pub.PublicKey.(*rsa.PublicKey)
	//if !ok {
	//	return
	//}
	//key := x509.MarshalPKCS1PublicKey(publicKey)
}
