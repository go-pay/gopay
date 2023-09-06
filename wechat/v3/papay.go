package wechat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 预扣费通知API
// Code = 0 is success
func (c *ClientV3) V3EntrustPayNotify(ctx context.Context, contractId string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3EntrustPayNotify, contractId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
