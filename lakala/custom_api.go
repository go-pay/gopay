package lakala

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// 创建报关单（非拆单）
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_report_single
func (c *Client) CreateReportSingle(ctx context.Context, partnerReportId string, bm gopay.BodyMap) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	if err = bm.CheckEmptyError("partner_order_id", "customs", "mch_customs_id", "mch_customs_name"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(createReportSingle, c.PartnerCode, partnerReportId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 创建报关单（拆单）
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_report_separate
func (c *Client) CreateReportSeparate(ctx context.Context, partnerReportId, partnerSubReportId string, bm gopay.BodyMap) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	if partnerSubReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerSubReportId is empty")
	}
	if err = bm.CheckEmptyError("partner_order_id", "customs", "mch_customs_id", "mch_customs_name", "order_fee", "product_fee", "transport_fee"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(createReportSeparate, c.PartnerCode, partnerReportId, partnerSubReportId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 报关状态查询
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_query_single
func (c *Client) ReportStatus(ctx context.Context, partnerReportId string) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	url := fmt.Sprintf(queryReportStatus, c.PartnerCode, partnerReportId)
	bs, err := c.doGet(ctx, url, "")
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 报关子单状态查询
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_query_separate
func (c *Client) ReportSubStatus(ctx context.Context, partnerReportId, partnerSubReportId string) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	url := fmt.Sprintf(queryReportSubStatus, c.PartnerCode, partnerReportId, partnerSubReportId)
	bs, err := c.doGet(ctx, url, "")
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 修改报关信息（非拆单）
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_modify_single
func (c *Client) ModifyReportSingle(ctx context.Context, partnerReportId string, bm gopay.BodyMap) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	if err = bm.CheckEmptyError("customs", "mch_customs_id", "mch_customs_name"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(modifyReportSingle, c.PartnerCode, partnerReportId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 修改报关信息（拆单）
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_modify_separate
func (c *Client) ModifyReportSeparate(ctx context.Context, partnerReportId, partnerSubReportId string, bm gopay.BodyMap) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	if partnerSubReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerSubReportId is empty")
	}
	if err = bm.CheckEmptyError("customs", "mch_customs_id", "mch_customs_name", "order_fee", "product_fee", "transport_fee"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(modifyReportSeparate, c.PartnerCode, partnerReportId, partnerSubReportId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 重推报关（非拆单）
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_resend_single
func (c *Client) ResendReportSingle(ctx context.Context, partnerReportId string) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	url := fmt.Sprintf(reSendReportSingle, c.PartnerCode, partnerReportId)
	bs, err := c.doPut(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 报关单子单重推
// 文档：https://payjp.lakala.com/docs/cn/#api-Custom-declare_resend_separate
func (c *Client) ResendReportSeparate(ctx context.Context, partnerReportId, partnerSubReportId string) (rsp *ReportRsp, err error) {
	if partnerReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerReportId is empty")
	}
	if partnerSubReportId == gopay.NULL {
		return nil, fmt.Errorf("partnerSubReportId is empty")
	}
	url := fmt.Sprintf(modifyReportSeparate, c.PartnerCode, partnerReportId, partnerSubReportId)
	bs, err := c.doPut(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	rsp = new(ReportRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}
