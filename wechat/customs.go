package wechat

import (
	"encoding/xml"
	"fmt"

	"github.com/go-pay/gopay"
)

// 订单附加信息提交（正式环境）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/external/declarecustom.php?chapter=18_1
func (w *Client) CustomsDeclareOrder(bm gopay.BodyMap) (wxRsp *CustomsDeclareOrderResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "transaction_id", "customs", "mch_customs_no")
	if err != nil {
		return nil, err
	}
	bm.Set("sign_type", SignType_MD5)
	bs, err := w.doProdPost(bm, customsDeclareOrder, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(CustomsDeclareOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 订单附加信息查询（正式环境）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/external/declarecustom.php?chapter=18_2
func (w *Client) CustomsDeclareQuery(bm gopay.BodyMap) (wxRsp *CustomsDeclareQueryResponse, err error) {
	err = bm.CheckEmptyError("customs")
	if err != nil {
		return nil, err
	}
	bm.Set("sign_type", SignType_MD5)
	bs, err := w.doProdPost(bm, customsDeclareQuery, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(CustomsDeclareQueryResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 订单附加信息重推（正式环境）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/external/declarecustom.php?chapter=18_4&index=3
func (w *Client) CustomsReDeclareOrder(bm gopay.BodyMap) (wxRsp *CustomsReDeclareOrderResponse, err error) {
	err = bm.CheckEmptyError("customs", "mch_customs_no")
	if err != nil {
		return nil, err
	}
	bm.Set("sign_type", SignType_MD5)
	bs, err := w.doProdPost(bm, customsReDeclareOrder, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(CustomsReDeclareOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
