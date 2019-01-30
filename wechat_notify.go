//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/28 20:16
//==================================
package gopay

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

func ParseNotifyResult(req *http.Request) (notifyRsp *WeChatNotifyRequest, err error) {
	notifyRsp = new(WeChatNotifyRequest)
	defer req.Body.Close()
	err = xml.NewDecoder(req.Body).Decode(notifyRsp)
	if err != nil {
		return nil, err
	}
	return
}

type WeChatNotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

func (this *WeChatNotifyResponse) ToXmlString() (xmlStr string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(this.ReturnCode)
	buffer.WriteString("]]></return_code>")

	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(this.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}
