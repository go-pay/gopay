package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 生成下一个发票号码（Generate invoice number）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_generate-next-invoice-number
func (c *Client) InvoiceNumberGenerate(ctx context.Context, invoiceNumber string) (inRsp *InvoiceNumberGenerateRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("invoice_number", invoiceNumber)
	res, bs, err := c.doPayPalPost(ctx, bm, generateInvoiceNumber)
	if err != nil {
		return nil, err
	}
	inRsp = &InvoiceNumberGenerateRsp{Code: Success}
	inRsp.Response = new(InvoiceNumber)
	if err = json.Unmarshal(bs, inRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		inRsp.Code = res.StatusCode
		inRsp.Error = string(bs)
		inRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, inRsp.ErrorResponse)
	}
	return inRsp, nil
}

// 发票列表（List invoices）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_list
func (c *Client) InvoiceList(ctx context.Context, bm gopay.BodyMap) (ppRsp *InvoiceListRsp, err error) {
	uri := invoiceList + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceListRsp{Code: Success}
	ppRsp.Response = new(Invoice)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
