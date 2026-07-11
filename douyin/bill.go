package douyin

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// ApplyTradeBill 申请交易账单
// 必填：bill_date（yyyy-MM-dd，3 个月内）
// 未设置的字段会自动补：mchid=c.Mchid、bill_type=TRADE、tar_type=GZIP
func (c *Client) ApplyTradeBill(ctx context.Context, bm gopay.BodyMap) (dyRsp *BillRsp, err error) {
	if bm == nil {
		bm = make(gopay.BodyMap)
	}
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	if bm.GetString("bill_type") == gopay.NULL {
		bm.Set("bill_type", "TRADE")
	}
	if bm.GetString("tar_type") == gopay.NULL {
		bm.Set("tar_type", "GZIP")
	}
	return c.applyBill(ctx, applyTradeBill, bm)
}

// ApplyFundBill 申请资金账单
// 必填：bill_date
// 未设置的字段会自动补：mchid=c.Mchid、tar_type=GZIP
// 可选：account_type（BaseAccount / OperationAccount，缺省 BaseAccount）
func (c *Client) ApplyFundBill(ctx context.Context, bm gopay.BodyMap) (dyRsp *BillRsp, err error) {
	if bm == nil {
		bm = make(gopay.BodyMap)
	}
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	if bm.GetString("tar_type") == gopay.NULL {
		bm.Set("tar_type", "GZIP")
	}
	return c.applyBill(ctx, applyFundBill, bm)
}

// ApplyProfitBill 申请分账账单
// 必填：bill_date
// 未设置的字段会自动补：mchid=c.Mchid、tar_type=GZIP
func (c *Client) ApplyProfitBill(ctx context.Context, bm gopay.BodyMap) (dyRsp *BillRsp, err error) {
	if bm == nil {
		bm = make(gopay.BodyMap)
	}
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	if bm.GetString("tar_type") == gopay.NULL {
		bm.Set("tar_type", "GZIP")
	}
	return c.applyBill(ctx, applyProfitBill, bm)
}

func (c *Client) applyBill(ctx context.Context, path string, bm gopay.BodyMap) (dyRsp *BillRsp, err error) {
	uri := path + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &BillRsp{Code: Success, SignInfo: si, Response: new(Bill)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}

// DownloadBillFile 下载账单文件（原始字节，未解压）
// downloadUrl 为 Apply*Bill 返回的 download_url，5min 内有效
// 抖音下载域名（download.douyinpay.com）与 API 域名（api.douyinpay.com）不同，故走完整 URL
// 若返回结果为 GZIP 压缩包，请配合 UngzipBill；校验完整性请配合 VerifyBillHash
func (c *Client) DownloadBillFile(ctx context.Context, downloadUrl string) (fileBytes []byte, err error) {
	if downloadUrl == gopay.NULL {
		return nil, errors.New("invalid download url")
	}
	u, err := url.Parse(downloadUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid download url: %v", err)
	}
	if u.Host == "" {
		return nil, errors.New("invalid download url: empty host")
	}
	pathAndQuery := u.RequestURI()
	if c.proxyHost != "" {
		downloadUrl = c.proxyHost + pathAndQuery
	}
	authorization, err := c.authorization(MethodGet, pathAndQuery, nil)
	if err != nil {
		return nil, err
	}
	req := c.hc.Req()
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestId())
	req.Header.Add("Accept", "*/*")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Douyin_Url: %s", downloadUrl)
		c.logger.Debugf("Douyin_Req_Headers: %#v", req.Header)
	}
	res, bs, err := req.Get(downloadUrl).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Douyin_Rsp_Status: %d", res.StatusCode)
		c.logger.Debugf("Douyin_Rsp_Headers: %#v", res.Header)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download bill file fail, http=%d, body=%s", res.StatusCode, string(bs))
	}
	return bs, nil
}

// UngzipBill 解压 GZIP 格式的账单文件
func UngzipBill(gzipData []byte) ([]byte, error) {
	if len(gzipData) == 0 {
		return nil, errors.New("empty gzip data")
	}
	zr, err := gzip.NewReader(bytes.NewReader(gzipData))
	if err != nil {
		return nil, fmt.Errorf("gzip new reader: %v", err)
	}
	defer zr.Close()
	return io.ReadAll(zr)
}

// VerifyBillHash 使用账单申请返回的哈希类型与哈希值校验原始账单摘要
// data 为原始账单字节（已解压），hashType 目前仅支持 SHA1
func VerifyBillHash(data []byte, hashType, hashValue string) error {
	switch strings.ToUpper(hashType) {
	case "SHA1":
		sum := sha1.Sum(data)
		got := hex.EncodeToString(sum[:])
		if !strings.EqualFold(got, hashValue) {
			return fmt.Errorf("bill hash mismatch: expect=%s, got=%s", hashValue, got)
		}
		return nil
	default:
		return fmt.Errorf("unsupported hash type: %s", hashType)
	}
}
