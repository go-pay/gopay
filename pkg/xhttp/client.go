package xhttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

type Client struct {
	HttpClient       *http.Client
	Transport        *http.Transport
	Header           http.Header
	Timeout          time.Duration
	Host             string
	bodySize         int // body size limit(MB), default is 10MB
	url              string
	method           string
	requestType      RequestType
	FormString       string
	ContentType      string
	unmarshalType    string
	multipartBodyMap map[string]any
	jsonByte         []byte
	err              error
}

// NewClient , default tls.Config{InsecureSkipVerify: true}
func NewClient() (client *Client) {
	client = &Client{
		HttpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				DisableKeepAlives: true,
				Proxy:             http.ProxyFromEnvironment,
			},
		},
		Transport:     nil,
		Header:        make(http.Header),
		bodySize:      10, // default is 10MB
		requestType:   TypeJSON,
		unmarshalType: string(TypeJSON),
	}
	return client
}

func (c *Client) SetTransport(transport *http.Transport) (client *Client) {
	c.Transport = transport
	return c
}

func (c *Client) SetTLSConfig(tlsCfg *tls.Config) (client *Client) {
	c.Transport = &http.Transport{TLSClientConfig: tlsCfg, DisableKeepAlives: true, Proxy: http.ProxyFromEnvironment}
	return c
}

func (c *Client) SetTimeout(timeout time.Duration) (client *Client) {
	c.Timeout = timeout
	return c
}

func (c *Client) SetHost(host string) (client *Client) {
	c.Host = host
	return c
}

// set body size (MB), default is 10MB
func (c *Client) SetBodySize(sizeMB int) (client *Client) {
	c.bodySize = sizeMB
	return c
}

func (c *Client) Type(typeStr RequestType) (client *Client) {
	if _, ok := types[typeStr]; ok {
		c.requestType = typeStr
	}
	return c
}

func (c *Client) Get(url string) (client *Client) {
	c.method = GET
	c.url = url
	return c
}

func (c *Client) Post(url string) (client *Client) {
	c.method = POST
	c.url = url
	return c
}

func (c *Client) Put(url string) (client *Client) {
	c.method = PUT
	c.url = url
	return c
}

func (c *Client) Delete(url string) (client *Client) {
	c.method = DELETE
	c.url = url
	return c
}

func (c *Client) Patch(url string) (client *Client) {
	c.method = PATCH
	c.url = url
	return c
}

func (c *Client) SendStruct(v any) (client *Client) {
	if v == nil {
		return c
	}
	bs, err := json.Marshal(v)
	if err != nil {
		c.err = fmt.Errorf("[%w]: %v, value: %v", gopay.MarshalErr, err, v)
		return c
	}
	switch c.requestType {
	case TypeJSON:
		c.jsonByte = bs
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		body := make(map[string]any)
		if err = json.Unmarshal(bs, &body); err != nil {
			c.err = fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
			return c
		}
		c.FormString = FormatURLParam(body)
	}
	return c
}

func (c *Client) SendBodyMap(bm map[string]any) (client *Client) {
	if bm == nil {
		return c
	}
	switch c.requestType {
	case TypeJSON:
		bs, err := json.Marshal(bm)
		if err != nil {
			c.err = fmt.Errorf("[%w]: %v, value: %v", gopay.MarshalErr, err, bm)
			return c
		}
		c.jsonByte = bs
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		c.FormString = FormatURLParam(bm)
	}
	return c
}

func (c *Client) SendMultipartBodyMap(bm map[string]any) (client *Client) {
	if bm == nil {
		return c
	}
	switch c.requestType {
	case TypeJSON:
		bs, err := json.Marshal(bm)
		if err != nil {
			c.err = fmt.Errorf("[%w]: %v, value: %v", gopay.MarshalErr, err, bm)
			return c
		}
		c.jsonByte = bs
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		c.FormString = FormatURLParam(bm)
	case TypeMultipartFormData:
		c.multipartBodyMap = bm
	}
	return c
}

// encodeStr: url.Values.Encode() or jsonBody
func (c *Client) SendString(encodeStr string) (client *Client) {
	switch c.requestType {
	case TypeJSON:
		c.jsonByte = []byte(encodeStr)
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		c.FormString = encodeStr
	}
	return c
}

func (c *Client) EndStruct(ctx context.Context, v any) (res *http.Response, err error) {
	res, bs, err := c.EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return res, fmt.Errorf("StatusCode(%d) != 200", res.StatusCode)
	}

	switch c.unmarshalType {
	case string(TypeJSON):
		err = json.Unmarshal(bs, &v)
		if err != nil {
			return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
		}
		return res, nil
	case string(TypeXML):
		err = xml.Unmarshal(bs, &v)
		if err != nil {
			return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
		}
		return res, nil
	default:
		return nil, errors.New("unmarshalType Type Wrong")
	}
}

func (c *Client) EndBytes(ctx context.Context) (res *http.Response, bs []byte, err error) {
	if c.err != nil {
		return nil, nil, c.err
	}
	var (
		body io.Reader
		bw   *multipart.Writer
	)
	// multipart-form-data
	if c.requestType == TypeMultipartFormData {
		body = &bytes.Buffer{}
		bw = multipart.NewWriter(body.(io.Writer))
	}

	reqFunc := func() (err error) {
		switch c.method {
		case GET:
			switch c.requestType {
			case TypeJSON:
				c.ContentType = types[TypeJSON]
			case TypeForm, TypeFormData, TypeUrlencoded:
				c.ContentType = types[TypeForm]
			case TypeMultipartFormData:
				c.ContentType = bw.FormDataContentType()
			case TypeXML:
				c.ContentType = types[TypeXML]
				c.unmarshalType = string(TypeXML)
			default:
				return errors.New("Request type Error ")
			}
		case POST, PUT, DELETE, PATCH:
			switch c.requestType {
			case TypeJSON:
				if c.jsonByte != nil {
					body = strings.NewReader(string(c.jsonByte))
				}
				c.ContentType = types[TypeJSON]
			case TypeForm, TypeFormData, TypeUrlencoded:
				body = strings.NewReader(c.FormString)
				c.ContentType = types[TypeForm]
			case TypeMultipartFormData:
				for k, v := range c.multipartBodyMap {
					// file 参数
					if file, ok := v.(*util.File); ok {
						fw, err := bw.CreateFormFile(k, file.Name)
						if err != nil {
							return err
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
				c.ContentType = bw.FormDataContentType()
			case TypeXML:
				body = strings.NewReader(c.FormString)
				c.ContentType = types[TypeXML]
				c.unmarshalType = string(TypeXML)
			default:
				return errors.New("Request type Error ")
			}
		default:
			return errors.New("Only support GET and POST and PUT and DELETE ")
		}

		req, err := http.NewRequestWithContext(ctx, c.method, c.url, body)
		if err != nil {
			return err
		}
		req.Header = c.Header
		req.Header.Set("Content-Type", c.ContentType)
		if c.Transport != nil {
			c.HttpClient.Transport = c.Transport
		}
		if c.Host != "" {
			req.Host = c.Host
		}
		if c.Timeout > 0 {
			c.HttpClient.Timeout = c.Timeout
		}
		res, err = c.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		bs, err = ioutil.ReadAll(io.LimitReader(res.Body, int64(c.bodySize<<20))) // default 10MB change the size you want
		if err != nil {
			return err
		}
		return nil
	}

	if err = reqFunc(); err != nil {
		return nil, nil, err
	}
	return res, bs, nil
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
			v = convertToString(body[k])
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

func convertToString(v any) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
