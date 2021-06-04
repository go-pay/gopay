package xhttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/iGoogle-ink/gopay/pkg/util"
)

type Client struct {
	// A Client is an HTTP client.
	HttpClient *http.Client

	// Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	Transport *http.Transport

	// Header contains the request header fields either received
	// by the server or to be sent by the client.
	//
	// If a server received a request with header lines,
	//
	//	Host: example.com
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	fOO: Bar
	//	foo: two
	//
	// then
	//
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
	//
	// For incoming requests, the Host header is promoted to the
	// Request.Host field and removed from the Header map.
	//
	// HTTP defines that header names are case-insensitive. The
	// request parser implements this by using CanonicalHeaderKey,
	// making the first character and any characters following a
	// hyphen uppercase and the rest lowercase.
	//
	// For client requests, certain headers such as Content-Length
	// and Connection are automatically written when needed and
	// values in Header may be ignored. See the documentation
	// for the Request.Write method.
	Header http.Header

	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	//
	// The Client cancels requests to the underlying Transport
	// as if the Request's Context ended.
	//
	// For compatibility, the Client will also use the deprecated
	// CancelRequest method on Transport if found. New
	// RoundTripper implementations should use the Request's Context
	// for cancellation instead of implementing CancelRequest.
	Timeout time.Duration

	// request URL
	url string

	// For server requests, Host specifies the host on which the
	// URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
	// is either the value of the "Host" header or the host name
	// given in the URL itself. For HTTP/2, it is the value of the
	// ":authority" pseudo-header field.
	// It may be of the form "host:port". For international domain
	// names, Host may be in Punycode or Unicode form. Use
	// golang.org/x/net/idna to convert it to either format if
	// needed.
	// To prevent DNS rebinding attacks, server Handlers should
	// validate that the Host header has a value for which the
	// Handler considers itself authoritative. The included
	// ServeMux supports patterns registered to particular host
	// names and thus protects its registered Handlers.
	//
	// For client requests, Host optionally overrides the Host
	// header to send. If empty, the Request.Write method uses
	// the value of URL.Host. Host may contain an international
	// domain name.
	Host string

	// method request method, now only support GET and POST
	method string

	// requestType
	requestType RequestType

	FormString string

	// ContentType, now only support json, form, form-data, urlencoded, xml
	ContentType string

	// unmarshalType json or xml
	unmarshalType string

	// types            map[string]string
	multipartBodyMap map[string]interface{}

	jsonByte []byte

	Errors []error

	//mu sync.RWMutex
}

// NewClient , default tls.Config{InsecureSkipVerify: true}
func NewClient() (client *Client) {
	client = &Client{
		HttpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				DisableKeepAlives: true,
			},
		},
		Transport:     nil,
		Header:        make(http.Header),
		requestType:   TypeUrlencoded,
		unmarshalType: string(TypeJSON),
		Errors:        make([]error, 0),
	}
	return client
}

func (c *Client) SetTLSConfig(tlsCfg *tls.Config) (client *Client) {
	c.Transport = &http.Transport{TLSClientConfig: tlsCfg, DisableKeepAlives: true}
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

func (c *Client) Type(typeStr RequestType) (client *Client) {
	if _, ok := types[typeStr]; ok {
		c.requestType = typeStr
	} else {
		c.Errors = append(c.Errors, errors.New("Type func: incorrect type \""+string(typeStr)+"\""))
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

func (c *Client) SendStruct(v interface{}) (client *Client) {
	if v == nil {
		return c
	}
	bs, err := json.Marshal(v)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	switch c.requestType {
	case TypeJSON:
		c.jsonByte = bs
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		c.FormString = string(bs)
	}
	return c
}

// 参数可直接传 gopay.BodyMap
func (c *Client) SendBodyMap(bm map[string]interface{}) (client *Client) {
	if bm == nil {
		return c
	}
	bs, err := json.Marshal(bm)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	switch c.requestType {
	case TypeJSON:
		c.jsonByte = bs
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		c.FormString = string(bs)
	}
	return c
}

// 参数可直接传 gopay.BodyMap
func (c *Client) SendMultipartBodyMap(bm map[string]interface{}) (client *Client) {
	if bm == nil {
		return c
	}
	bs, err := json.Marshal(bm)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	switch c.requestType {
	case TypeJSON:
		c.jsonByte = bs
	case TypeXML, TypeUrlencoded, TypeForm, TypeFormData:
		c.FormString = string(bs)
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

func (c *Client) EndStruct(v interface{}) (res *http.Response, errs []error) {
	res, bs, errs := c.EndBytes()
	if errs != nil && len(errs) > 0 {
		c.Errors = append(c.Errors, errs...)
		return nil, c.Errors
	}
	if res.StatusCode != 200 {
		c.Errors = append(c.Errors, errors.New(string(bs)))
		return res, c.Errors
	}

	switch c.unmarshalType {
	case string(TypeJSON):
		err := json.Unmarshal(bs, &v)
		if err != nil {
			c.Errors = append(c.Errors, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err))
			return nil, c.Errors
		}
		return res, nil
	case string(TypeXML):
		err := xml.Unmarshal(bs, &v)
		if err != nil {
			c.Errors = append(c.Errors, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err))
			return nil, c.Errors
		}
		return res, nil
	default:
		c.Errors = append(c.Errors, errors.New("unmarshalType Type Wrong"))
		return nil, c.Errors
	}
}

func (c *Client) EndBytes() (res *http.Response, bs []byte, errs []error) {
	if len(c.Errors) > 0 {
		return nil, nil, c.Errors
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

	req, err := func() (*http.Request, error) {
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
				return nil, errors.New("Request type Error ")
			}
		case POST, PUT, DELETE:
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
							return nil, err
						}
						fw.Write(file.Content)
						continue
					}
					// text 参数
					vs, ok2 := v.(string)
					if ok2 {
						bw.WriteField(k, vs)
					} else if ss := util.ConvertToString(v); ss != "" {
						bw.WriteField(k, ss)
					}
				}
				bw.Close()
				c.ContentType = bw.FormDataContentType()
			case TypeXML:
				body = strings.NewReader(c.FormString)
				c.ContentType = types[TypeXML]
				c.unmarshalType = string(TypeXML)
			default:
				return nil, errors.New("Request type Error ")
			}
		default:
			return nil, errors.New("Only support GET and POST and PUT and DELETE ")
		}

		req, err := http.NewRequest(c.method, c.url, body)
		if err != nil {
			return nil, err
		}
		req.Header = c.Header
		req.Header.Set("Content-Type", c.ContentType)
		if c.Transport != nil {
			c.HttpClient.Transport = c.Transport
		}
		return req, nil
	}()
	if err != nil {
		c.Errors = append(c.Errors, err)
		return nil, nil, c.Errors
	}

	if c.Timeout != time.Duration(0) {
		c.HttpClient.Timeout = c.Timeout
	}
	if c.Host != "" {
		req.Host = c.Host
	}
	res, err = c.HttpClient.Do(req)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return nil, nil, c.Errors
	}
	defer res.Body.Close()
	bs, err = ioutil.ReadAll(io.LimitReader(res.Body, int64(5<<20))) // default 5MB change the size you want
	if err != nil {
		c.Errors = append(c.Errors, err)
		return nil, nil, c.Errors
	}
	return res, bs, nil
}
