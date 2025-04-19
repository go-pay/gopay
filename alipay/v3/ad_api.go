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
