package wechat

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/cedarwu/gopay"
)

// 订单附加信息提交（正式环境）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/external/declarecustom.php?chapter=18_1
func (w *Client) CustomsDeclareOrder(bm gopay.BodyMap) (wxRsp *CustomsDeclareOrderResponse, header http.Header, err error) {
	err = bm.CheckEmptyError("out_trade_no", "transaction_id", "customs", "mch_customs_no")
	if err != nil {
		return nil, nil, err
	}
	bm.Set("sign_type", SignType_MD5)
	bs, _, _, header, err := w.doProdPost(context.Background(), bm, customsDeclareOrder, nil)
	if err != nil {
		return nil, header, err
	}
	wxRsp = new(CustomsDeclareOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, header, nil
}

// 订单附加信息查询（正式环境）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/external/declarecustom.php?chapter=18_2
func (w *Client) CustomsDeclareQuery(bm gopay.BodyMap) (wxRsp *CustomsDeclareQueryResponse, header http.Header, err error) {
	err = bm.CheckEmptyError("customs")
	if err != nil {
		return nil, nil, err
	}
	bm.Set("sign_type", SignType_MD5)
	bs, _, _, header, err := w.doProdPost(context.Background(), bm, customsDeclareQuery, nil)
	if err != nil {
		return nil, header, err
	}
	wxRsp = new(CustomsDeclareQueryResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, header, nil
}

// 订单附加信息重推（正式环境）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/external/declarecustom.php?chapter=18_4&index=3
func (w *Client) CustomsReDeclareOrder(bm gopay.BodyMap) (wxRsp *CustomsReDeclareOrderResponse, header http.Header, err error) {
	err = bm.CheckEmptyError("customs", "mch_customs_no")
	if err != nil {
		return nil, nil, err
	}
	bm.Set("sign_type", SignType_MD5)
	bs, _, _, header, err := w.doProdPost(context.Background(), bm, customsReDeclareOrder, nil)
	if err != nil {
		return nil, header, err
	}
	wxRsp = new(CustomsReDeclareOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, header, nil
}
