package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 生成下一个发票号码（Generate invoice number）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_generate-next-invoice-number
func (c *Client) InvoiceNumberGenerate(ctx context.Context, invoiceNumber string) (ppRsp *InvoiceNumberGenerateRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("invoice_number", invoiceNumber)
	res, bs, err := c.doPayPalPost(ctx, bm, generateInvoiceNumber)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceNumberGenerateRsp{Code: Success}
	ppRsp.Response = new(InvoiceNumber)
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
	ppRsp.Response = new(InvoiceList)
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

// 创建虚拟发票（Create draft invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_create
func (c *Client) InvoiceCreate(ctx context.Context, bm gopay.BodyMap) (ppRsp *InvoiceCreateRsp, err error) {
	res, bs, err := c.doPayPalPost(ctx, bm, createDraftInvoice)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceCreateRsp{Code: Success}
	ppRsp.Response = new(Invoice)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 删除发票（Delete invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_delete
func (c *Client) InvoiceDelete(ctx context.Context, invoiceId string) (ppRsp *EmptyRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(deleteInvoice, invoiceId)
	res, bs, err := c.doPayPalDelete(ctx, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &EmptyRsp{Code: Success}
	if res.StatusCode != http.StatusNoContent {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
