package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

func TestAdReportdataQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 	err = bm.CheckEmptyError("biz_token", "alipay_pid", "query_type", "ad_level", "start_date", "end_date", "principal_tag")
	bm.Set("biz_token", "e09d869511189c24ce13fe3xxxxxx").
		Set("alipay_pid", "123456789010").
		Set("query_type", "SUM").
		Set("ad_level", "PLAN").
		Set("start_date", "20210801").
		Set("end_date", "20210802").
		Set("principal_tag", "shybo")
	// 创建订单
	aliRsp, err := client.AdReportdataQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AdReportdataQuery(), err:%v", err)
		return
	}
	xlog.Debugf("aliRsp:%s", js.MarshalString(aliRsp))

	if aliRsp.StatusCode != Success {
		xlog.Errorf("aliRsp.StatusCode:%d", aliRsp.StatusCode)
		return
	}
	xlog.Debug("aliRsp.CrowdNo:", aliRsp.DataList)
}

func TestAdPromotepageBatchquery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// err = bm.CheckEmptyError("biz_token", "principal_tag", "page_no", "page_size")
	bm.Set("biz_token", "e09d869511189c24ce13fe3xxxxxx").
		Set("principal_tag", "shybo").
		Set("page_no", 1).
		Set("page_size", 10)
	// 创建订单
	aliRsp, err := client.AdPromotepageBatchquery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AdPromotepageBatchquery(), err:%v", err)
		return
	}
	xlog.Debugf("aliRsp:%s", js.MarshalString(aliRsp))

	if aliRsp.StatusCode != Success {
		xlog.Errorf("aliRsp.StatusCode:%d", aliRsp.StatusCode)
		return
	}
	xlog.Debug("aliRsp.Total:", aliRsp.Total)
	xlog.Debug("aliRsp.List:", aliRsp.List)
}
