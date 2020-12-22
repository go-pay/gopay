package wecaht

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
)

func (c *ClientV3) V3TransactionJsapi(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := gotil.GetRandomString(32)
	authorization, err := c.Authorization(MethodPost, v3ApiJsapi, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	bs, err := c.doProdPost(bm, v3ApiJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = new(PrepayRsp)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
