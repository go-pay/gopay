package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/go-pay/gopay/pkg/util"
)

type Request struct {
	client           *Client
	Header           http.Header
	formString       string
	jsonByte         []byte
	url              string
	method           string
	requestType      string
	responseType     string
	multipartBodyMap map[string]any
	err              error
}

func (r *Request) Get(url string) *Request {
	r.method = GET
	r.url = url
	return r
}

func (r *Request) Post(url string) *Request {
	r.method = POST
	r.url = url
	return r
}

func (r *Request) Put(url string) *Request {
	r.method = PUT
	r.url = url
	return r
}

func (r *Request) Delete(url string) *Request {
	r.method = DELETE
	r.url = url
	return r
}

func (r *Request) Patch(url string) *Request {
	r.method = PATCH
	r.url = url
	return r
}

// =====================================================================================================================

func (r *Request) SendStruct(v any) (c *Request) {
	if v == nil {
		return r
	}
	bs, err := json.Marshal(v)
	if err != nil {
		r.err = fmt.Errorf("json.Marshal(%+v)：%w", v, err)
		return r
	}
	switch r.requestType {
	case TypeJSON:
		r.jsonByte = bs
	case TypeXML, TypeFormData:
		body := make(map[string]any)
		if err = json.Unmarshal(bs, &body); err != nil {
			r.err = fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), body, err)
			return r
		}
		r.formString = FormatURLParam(body)
	}
	return r
}

func (r *Request) SendBodyMap(bm map[string]any) (client *Request) {
	if bm == nil {
		return r
	}
	switch r.requestType {
	case TypeJSON:
		bs, err := json.Marshal(bm)
		if err != nil {
			r.err = fmt.Errorf("json.Marshal(%+v)：%w", bm, err)
			return r
		}
		r.jsonByte = bs
	case TypeXML, TypeFormData:
		r.formString = FormatURLParam(bm)
	}
	return r
}

func (r *Request) SendMultipartBodyMap(bm map[string]any) (client *Request) {
	if bm == nil {
		return r
	}
	switch r.requestType {
	case TypeJSON:
		bs, err := json.Marshal(bm)
		if err != nil {
			r.err = fmt.Errorf("json.Marshal(%+v)：%w", bm, err)
			return r
		}
		r.jsonByte = bs
	case TypeXML, TypeFormData:
		r.formString = FormatURLParam(bm)
	case TypeMultipartFormData:
		r.multipartBodyMap = bm
	}
	return r
}

// encodeStr: url.Values.Encode() or jsonBody
func (r *Request) SendString(encodeStr string) (client *Request) {
	switch r.requestType {
	case TypeJSON:
		r.jsonByte = []byte(encodeStr)
	case TypeXML, TypeFormData:
		r.formString = encodeStr
	}
	return r
}

// =====================================================================================================================

func (r *Request) EndBytes(ctx context.Context) (res *http.Response, bs []byte, err error) {
	if r.err != nil {
		return nil, nil, r.err
	}
	var (
		body io.Reader
		bw   *multipart.Writer
	)
	// multipart-form-data
	if r.requestType == TypeMultipartFormData {
		body = &bytes.Buffer{}
		bw = multipart.NewWriter(body.(io.Writer))
	}

	switch r.method {
	case GET:
		// do nothing
	case POST, PUT, DELETE, PATCH:
		switch r.requestType {
		case TypeJSON:
			if r.jsonByte != nil {
				body = strings.NewReader(string(r.jsonByte))
			}
		case TypeFormData:
			if r.formString != "" {
				body = strings.NewReader(r.formString)
			}
		case TypeMultipartFormData:
			for k, v := range r.multipartBodyMap {
				// file 参数
				if file, ok := v.(*util.File); ok {
					fw, e := bw.CreateFormFile(k, file.Name)
					if e != nil {
						return nil, nil, e
					}
					_, _ = fw.Write(file.Content)
					continue
				}
				// text 参数
				vs, ok2 := v.(string)
				if ok2 {
					_ = bw.WriteField(k, vs)
				} else if ss := util.ConvertToString(v); ss != "" {
					_ = bw.WriteField(k, ss)
				}
			}
			_ = bw.Close()
			r.Header.Set("Content-Type", bw.FormDataContentType())
		case TypeXML:
			if r.formString != "" {
				body = strings.NewReader(r.formString)
			}
		default:
			return nil, nil, errors.New("Request type Error ")
		}
	default:
		return nil, nil, errors.New("Only support GET and POST and PUT and DELETE ")
	}

	// request
	req, err := http.NewRequestWithContext(ctx, r.method, r.url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header = r.Header
	res, err = r.client.HttpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	bs, err = io.ReadAll(io.LimitReader(res.Body, int64(r.client.bodySize<<20))) // default 10MB change the size you want
	if err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}

func (r *Request) EndStruct(ctx context.Context, v any) (res *http.Response, err error) {
	res, bs, err := r.EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return res, fmt.Errorf("StatusCode(%d) != 200", res.StatusCode)
	}

	switch r.responseType {
	case ResTypeJSON:
		err = json.Unmarshal(bs, &v)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), v, err)
		}
		return res, nil
	case ResTypeXML:
		err = xml.Unmarshal(bs, &v)
		if err != nil {
			return nil, fmt.Errorf("xml.Unmarshal(%s, %+v)：%w", string(bs), v, err)
		}
		return res, nil
	default:
		return nil, errors.New("responseType Type Wrong")
	}
}

func FormatURLParam(body map[string]any) (urlParam string) {
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range body {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v, ok := body[k].(string)
		if !ok {
			v = util.ConvertToString(body[k])
		}
		if v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}
