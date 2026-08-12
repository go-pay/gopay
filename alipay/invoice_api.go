package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.commerce.ec.industryinvoice.invoiceapply.create(支付开票开票申请创建)
// 支持蓝票（BLUE）和红票（RED）场景。
// product_id 和 product_code 二选一传入。
// 文档地址: https://opendocs.alipay.com/mini/031f9074_alipay.commerce.ec.industryinvoice.invoiceapply.create?pathHash=a4e5340c
func (a *Client) IndustryInvoiceApplyCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *IndustryInvoiceApplyCreateRsp, err error) {
	err = bm.CheckEmptyError("outer_apply_id", "invoice_type", "invoice_kind", "tax_no", "trade_list", "buyer_name", "show_seller_bank_info")
	if err != nil {
		return nil, err
	}
	if (bm.GetString("product_id") == gopay.NULL) == (bm.GetString("product_code") == gopay.NULL) {
		return nil, fmt.Errorf("product_id and product_code must choose one")
	}
	if bm.GetString("invoice_type") == IndustryInvoiceTypeRed {
		if bm.GetString("invoice_red_reason") == gopay.NULL {
			return nil, fmt.Errorf("invoice_red_reason is required when invoice_type is RED")
		}
		if bm.GetString("related_blue_invoice_no") == gopay.NULL {
			return nil, fmt.Errorf("related_blue_invoice_no is required when invoice_type is RED")
		}
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.ec.industryinvoice.invoiceapply.create"); err != nil {
		return nil, err
	}
	aliRsp = new(IndustryInvoiceApplyCreateRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.commerce.ec.industryinvoice.invoiceapply.query(支付开票开票申请查询)
// 根据 invoice_apply_id 或 outer_apply_id 查询发票申请状态及正式发票号码等信息。
// 文档地址: https://opendocs.alipay.com/mini/45168e10_alipay.commerce.ec.industryinvoice.invoiceapply.query?scene=common&pathHash=9e5a0b67
func (a *Client) IndustryInvoiceApplyQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *IndustryInvoiceApplyQueryRsp, err error) {
	err = bm.CheckEmptyError("tax_no", "product_id")
	if err != nil {
		return nil, err
	}
	if (bm.GetString("invoice_apply_id") == gopay.NULL) == (bm.GetString("outer_apply_id") == gopay.NULL) {
		return nil, fmt.Errorf("invoice_apply_id and outer_apply_id must choose one")
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.ec.industryinvoice.invoiceapply.query"); err != nil {
		return nil, err
	}
	aliRsp = new(IndustryInvoiceApplyQueryRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.commerce.ec.industryinvoice.company.query(企业信息查询)
// 根据企业税号查询企业已开通的发票产品信息。
// 文档地址: https://opendocs.alipay.com/mini/a52ef6fc_alipay.commerce.ec.industryinvoice.company.query?scene=common&pathHash=83faa4b3
func (a *Client) IndustryInvoiceCompanyQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *IndustryInvoiceCompanyQueryRsp, err error) {
	err = bm.CheckEmptyError("tax_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.ec.industryinvoice.company.query"); err != nil {
		return nil, err
	}
	aliRsp = new(IndustryInvoiceCompanyQueryRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.commerce.ec.industryinvoice.invoiceapply.retry(支付开票开票申请重试)
// 用于支付开票开票申请的异常重试。
// 文档地址: https://opendocs.alipay.com/mini/99125d7c_alipay.commerce.ec.industryinvoice.invoiceapply.retry?scene=common&pathHash=d5ce1397
func (a *Client) IndustryInvoiceApplyRetry(ctx context.Context, bm gopay.BodyMap) (aliRsp *IndustryInvoiceApplyRetryRsp, err error) {
	err = bm.CheckEmptyError("invoice_apply_id", "product_id", "tax_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.commerce.ec.industryinvoice.invoiceapply.retry"); err != nil {
		return nil, err
	}
	aliRsp = new(IndustryInvoiceApplyRetryRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
