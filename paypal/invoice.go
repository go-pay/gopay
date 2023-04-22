package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
func (c *Client) InvoiceList(ctx context.Context, query gopay.BodyMap) (ppRsp *InvoiceListRsp, err error) {
	uri := invoiceList + "?" + query.EncodeURLParams()
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
func (c *Client) InvoiceCreate(ctx context.Context, body gopay.BodyMap) (ppRsp *InvoiceCreateRsp, err error) {
	res, bs, err := c.doPayPalPost(ctx, body, createDraftInvoice)
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

// 更新发票（Fully update invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_update
func (c *Client) InvoiceUpdate(ctx context.Context, invoiceId string, sendToInvoicer, sendToRecipient bool, body gopay.BodyMap) (ppRsp *InvoiceUpdateRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(fullyUpdateInvoice, invoiceId)
	if sendToInvoicer {
		url = url + "?send_to_invoicer=true"
	}
	if sendToRecipient {
		if sendToInvoicer {
			url = url + "&send_to_recipient=true"
		} else {
			url = url + "?send_to_recipient=true"
		}
	}
	res, bs, err := c.doPayPalPut(ctx, body, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceUpdateRsp{Code: Success}
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

// 获取发票详情（Show invoice details）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_get
func (c *Client) InvoiceDetail(ctx context.Context, invoiceId string) (ppRsp *InvoiceDetailRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(showInvoiceDetail, invoiceId)
	res, bs, err := c.doPayPalGet(ctx, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceDetailRsp{Code: Success}
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

// 取消已发送发票（Cancel sent invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_cancel
func (c *Client) InvoiceCancel(ctx context.Context, invoiceId string, body gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(cancelSentInvoice, invoiceId)
	res, bs, err := c.doPayPalPost(ctx, body, url)
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

// 生成发票二维码（Generate QR code）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_generate-qr-code
func (c *Client) InvoiceGenerateQRCode(ctx context.Context, invoiceId string, body gopay.BodyMap) (ppRsp *InvoiceGenerateQRCodeRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(generateInvoiceQRCode, invoiceId)
	res, bs, err := c.doPayPalPost(ctx, body, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceGenerateQRCodeRsp{Code: Success}
	ppRsp.Response = &QRCodeBase64{Base64Image: string(bs)}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 发票付款记录（Record payment for invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_payments
func (c *Client) InvoicePaymentRecord(ctx context.Context, invoiceId string, body gopay.BodyMap) (ppRsp *InvoicePaymentRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(recordPaymentForInvoice, invoiceId)
	res, bs, err := c.doPayPalPost(ctx, body, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoicePaymentRsp{Code: Success}
	ppRsp.Response = new(InvoicePayment)
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

// 发票付款删除（Delete external payment）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_payments-delete
func (c *Client) InvoicePaymentDelete(ctx context.Context, invoiceId, transactionId string) (ppRsp *EmptyRsp, err error) {
	if invoiceId == gopay.NULL || transactionId == gopay.NULL {
		return nil, errors.New("invoice_id or transaction_id is empty")
	}
	url := fmt.Sprintf(deleteExternalPayment, invoiceId, transactionId)
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

// 发票退款记录（Record refund for invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_refunds
func (c *Client) InvoiceRefundRecord(ctx context.Context, invoiceId string, body gopay.BodyMap) (ppRsp *InvoiceRefundRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(recordRefundForInvoice, invoiceId)
	res, bs, err := c.doPayPalPost(ctx, body, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceRefundRsp{Code: Success}
	ppRsp.Response = new(InvoiceRefund)
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

// 发票退款删除（Delete external refund）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_refunds-delete
func (c *Client) InvoiceRefundDelete(ctx context.Context, invoiceId, transactionId string) (ppRsp *EmptyRsp, err error) {
	if invoiceId == gopay.NULL || transactionId == gopay.NULL {
		return nil, errors.New("invoice_id or transaction_id is empty")
	}
	url := fmt.Sprintf(deleteExternalRefund, invoiceId, transactionId)
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

// 发送发票提醒（Send invoice reminder）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_remind
func (c *Client) InvoiceSendRemind(ctx context.Context, invoiceId string, body gopay.BodyMap) (ppRsp *EmptyRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(sendInvoiceReminder, invoiceId)
	res, bs, err := c.doPayPalPost(ctx, body, url)
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

// 发送发票（Send invoice）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_send
func (c *Client) InvoiceSend(ctx context.Context, invoiceId string, body gopay.BodyMap) (ppRsp *InvoiceSendRsp, err error) {
	if invoiceId == gopay.NULL {
		return nil, errors.New("invoice_id is empty")
	}
	url := fmt.Sprintf(sendInvoice, invoiceId)
	res, bs, err := c.doPayPalPost(ctx, body, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceSendRsp{Code: Success}
	ppRsp.Response = new(InvoiceSend)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusAccepted {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 发票搜索（Search for invoices）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#search-invoices_search-invoices
func (c *Client) InvoiceSearch(ctx context.Context, page, pageSize int, totalRequired bool, body gopay.BodyMap) (ppRsp *InvoiceSearchRsp, err error) {
	uri := searchInvoice
	if page != 0 && pageSize != 0 {
		uri += uri + "?page=" + strconv.Itoa(page) + "&page_size=" + strconv.Itoa(pageSize)
	}
	if totalRequired {
		if page != 0 && pageSize != 0 {
			uri += uri + "&total_required=true"
		} else {
			uri += uri + "?total_required=true"
		}
	}
	res, bs, err := c.doPayPalPost(ctx, body, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceSearchRsp{Code: Success}
	ppRsp.Response = new(InvoiceSearch)
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

// 发票模板列表（List templates）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_list
func (c *Client) InvoiceTemplateList(ctx context.Context, query gopay.BodyMap) (ppRsp *InvoiceTemplateListRsp, err error) {
	uri := invoiceTemplateList + "?" + query.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceTemplateListRsp{Code: Success}
	ppRsp.Response = new(InvoiceTemplate)
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

// 创建发票模板（Create template）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_create
func (c *Client) InvoiceTemplateCreate(ctx context.Context, body gopay.BodyMap) (ppRsp *InvoiceTemplateCreateRsp, err error) {
	res, bs, err := c.doPayPalPost(ctx, body, createInvoiceTemplate)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceTemplateCreateRsp{Code: Success}
	ppRsp.Response = new(Template)
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

// 删除发票模板（Delete template）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_delete
func (c *Client) InvoiceTemplateDelete(ctx context.Context, templateId string) (ppRsp *EmptyRsp, err error) {
	if templateId == gopay.NULL {
		return nil, errors.New("template_id is empty")
	}
	url := fmt.Sprintf(deleteInvoiceTemplate, templateId)
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

// 更新发票模板（Fully update template）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_update
func (c *Client) InvoiceTemplateUpdate(ctx context.Context, templateId string, body gopay.BodyMap) (ppRsp *InvoiceTemplateUpdateRsp, err error) {
	if templateId == gopay.NULL {
		return nil, errors.New("template_id is empty")
	}
	url := fmt.Sprintf(fullyUpdateInvoiceTemplate, templateId)
	res, bs, err := c.doPayPalPut(ctx, body, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &InvoiceTemplateUpdateRsp{Code: Success}
	ppRsp.Response = new(Template)
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
