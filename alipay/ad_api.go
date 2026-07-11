package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.data.dataservice.ad.conversion.upload(转化数据回传)
// 文档地址：https://opendocs.alipay.com/open/3940a105_alipay.data.dataservice.ad.conversion.upload
func (a *Client) DataServiceAdConversionUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataServiceAdConversionUploadResponse, err error) {
	err = bm.CheckEmptyError("biz_token", "conversion_data_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.ad.conversion.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(DataServiceAdConversionUploadResponse)
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

// alipay.data.dataservice.ad.reportdata.query(广告投放数据通用查询)
// 文档地址：https://opendocs.alipay.com/open/c089ee8d_alipay.data.dataservice.ad.reportdata.query
func (a *Client) DataServiceAdReportdataQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataServiceAdReportdataQueryResponse, err error) {
	err = bm.CheckEmptyError("biz_token", "alipay_pid", "query_type", "ad_level", "start_date", "end_date", "principal_tag")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.ad.reportdata.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataServiceAdReportdataQueryResponse)
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

// alipay.data.dataservice.ad.promotepage.batchquery(自建推广页列表批量查询)
// 文档地址：https://opendocs.alipay.com/open/e060c7d1_alipay.data.dataservice.ad.promotepage.batchquery
func (a *Client) DataServiceAdPromotepageBatchquery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataServiceAdPromotepageBatchqueryResponse, err error) {
	err = bm.CheckEmptyError("biz_token", "principal_tag", "page_no", "page_size")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.ad.promotepage.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(DataServiceAdPromotepageBatchqueryResponse)
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

// alipay.data.dataservice.ad.promotepage.download(自建推广页留资数据查询)
// 文档地址：https://opendocs.alipay.com/open/1df3222a_alipay.data.dataservice.ad.promotepage.download
func (a *Client) DataServiceAdPromotepageDownload(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataServiceAdPromotepageDownloadResponse, err error) {
	err = bm.CheckEmptyError("start_date", "end_date", "page_no", "page_size", "biz_token", "principal_tag", "promote_page_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.ad.promotepage.download"); err != nil {
		return nil, err
	}
	aliRsp = new(DataServiceAdPromotepageDownloadResponse)
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

// alipay.data.dataservice.xlight.task.query(任务广告完成状态查询)
// 文档地址：https://opendocs.alipay.com/open/7275fba1_alipay.data.dataservice.xlight.task.query
func (a *Client) DataServiceXlightTaskQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataServiceXlightTaskQueryResponse, err error) {
	err = bm.CheckEmptyError("biz_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.xlight.task.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataServiceXlightTaskQueryResponse)
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
