package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 转化数据回传 alipay.data.dataservice.ad.conversion.upload
// StatusCode = 200 is success
func (a *ClientV3) AdConversionUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *AdConversionUploadRsp, err error) {
	err = bm.CheckEmptyError("biz_token", "conversion_data_list")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3DataDataserviceAdConversionUpload, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceAdConversionUpload, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &AdConversionUploadRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 广告投放数据通用查询 alipay.data.dataservice.ad.reportdata.query
// StatusCode = 200 is success
func (a *ClientV3) AdReportdataQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AdReportdataQueryRsp, err error) {
	err = bm.CheckEmptyError("biz_token", "alipay_pid", "query_type", "ad_level", "start_date", "end_date", "principal_tag")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3DataDataserviceAdReportdataQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceAdReportdataQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &AdReportdataQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 自建推广页列表批量查询 alipay.data.dataservice.ad.promotepage.batchquery
// StatusCode = 200 is success
func (a *ClientV3) AdPromotepageBatchquery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AdPromotepageBatchqueryRsp, err error) {
	err = bm.CheckEmptyError("biz_token", "principal_tag", "page_no", "page_size")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceAdPromotepageBatchquery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &AdPromotepageBatchqueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 自建推广页留资数据查询 alipay.data.dataservice.ad.promotepage.download
// StatusCode = 200 is success
func (a *ClientV3) AdPromotepageDownload(ctx context.Context, bm gopay.BodyMap) (aliRsp *AdPromotepageDownloadRsp, err error) {
	err = bm.CheckEmptyError("start_date", "end_date", "page_no", "page_size", "biz_token", "principal_tag", "promote_page_id")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceAdPromotepageDownload + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &AdPromotepageDownloadRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 任务广告完成状态查询接口 alipay.data.dataservice.xlight.task.query
// StatusCode = 200 is success
func (a *ClientV3) XlightTaskQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *XlightTaskQueryRsp, err error) {
	err = bm.CheckEmptyError("biz_id")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceXlightTaskQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceXlightTaskQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &XlightTaskQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 消费明细查询接口 alipay.data.dataservice.ad.consumehistory.query
// StatusCode = 200 is success
func (a *ClientV3) AdConsumehistoryQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AdConsumehistoryQueryRsp, err error) {
	err = bm.CheckEmptyError("biz_token", "alipay_pid", "start_date", "end_date", "group_condition", "biz_scene", "current")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceAdConsumehistoryQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceAdConsumehistoryQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &AdConsumehistoryQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 商品落地页信息创建或更新 alipay.data.dataservice.product.landinginfo.createormodify
// StatusCode = 200 is success
func (a *ClientV3) ProductLandinginfoCreateOrModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *ProductLandinginfoCreateOrModifyRsp, err error) {
	err = bm.CheckEmptyError("item_id", "out_item_id", "landing")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceProductLandinginfoCreateOrModify + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceProductLandinginfoCreateOrModify, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &ProductLandinginfoCreateOrModifyRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 商品落地页信息查询 alipay.data.dataservice.product.landinginfo.query
// StatusCode = 200 is success
func (a *ClientV3) ProductLandinginfoQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ProductLandinginfoQueryRsp, err error) {
	err = bm.CheckEmptyError("item_id", "out_item_id")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceProductLandinginfoQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceProductLandinginfoQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &ProductLandinginfoQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 广告代理商投放数据查询 alipay.data.dataservice.ad.agentreportdata.query
// StatusCode = 200 is success
func (a *ClientV3) AdAgentreportdataQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AdAgentreportdataQueryRsp, err error) {
	err = bm.CheckEmptyError("biz_token", "alipay_pid", "principal_tag", "query_type", "start_date", "end_date")
	if err != nil {
		return nil, err
	}
	uri := v3DataDataserviceAdAgentreportdataQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3DataDataserviceAdAgentreportdataQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &AdAgentreportdataQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
