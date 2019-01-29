//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/28 20:16
//==================================
package gopay

import (
	"encoding/xml"
	"net/http"
)

func ParseNotifyResult(req *http.Request) (notifyRsp *WeChatNotifyResponse, err error) {
	notifyRsp = new(WeChatNotifyResponse)
	defer req.Body.Close()
	err = xml.NewDecoder(req.Body).Decode(notifyRsp)
	if err != nil {
		return nil, err
	}
	return
}
