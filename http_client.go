package gopay

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

const (
	POST           = "POST"
	GET            = "GET"
	TypeJSON       = "json"
	TypeXML        = "xml"
	TypeUrlencoded = "urlencoded"
	TypeForm       = "form"
	TypeFormData   = "form-data"
	TypeHTML       = "html"
	TypeText       = "text"
	TypeMultipart  = "multipart"
)

var Types = map[string]string{
	TypeJSON:       "application/json",
	TypeXML:        "application/xml",
	TypeForm:       "application/x-www-form-urlencoded",
	TypeFormData:   "application/x-www-form-urlencoded",
	TypeUrlencoded: "application/x-www-form-urlencoded",
	TypeHTML:       "text/html",
	TypeText:       "text/plain",
	TypeMultipart:  "multipart/form-data",
}

type Client struct {
	HttpClient    *http.Client
	Transport     *http.Transport
	Url           string
	Method        string
	ForceType     string
	FormString    string
	ContentType   string
	UnmarshalType string
	Types         map[string]string
	JsonByte      []byte
	Errors        []error
	mu            sync.RWMutex
}

// NewHttpClient , default tls.Config{InsecureSkipVerify: true}
func NewHttpClient() (client *Client) {
	c := new(Client)
	c.HttpClient = new(http.Client)
	c.Transport = &http.Transport{}
	c.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c.ForceType = TypeUrlencoded
	c.UnmarshalType = TypeJSON
	c.Errors = make([]error, 0)
	return c
}

func (c *Client) SetTLSConfig(tlsCfg *tls.Config) (client *Client) {
	c.mu.Lock()
	c.Transport.TLSClientConfig = tlsCfg
	c.mu.Unlock()
	return c
}

func (c *Client) Post(url string) (client *Client) {
	c.mu.Lock()
	c.Method = POST
	c.Url = url
	c.mu.Unlock()
	return c
}

func (c *Client) Type(typeStr string) (client *Client) {
	if _, ok := Types[typeStr]; ok {
		c.ForceType = typeStr
	} else {
		c.Errors = append(c.Errors, errors.New("Type func: incorrect type \""+typeStr+"\""))
	}
	return c
}

func (c *Client) Get(url string) (client *Client) {
	c.mu.Lock()
	c.Method = GET
	c.Url = url
	c.mu.Unlock()
	return c
}

func (c *Client) SendStruct(v interface{}) (client *Client) {
	bs, err := json.Marshal(v)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	c.JsonByte = bs
	return c
}

func (c *Client) SendString(v string) (client *Client) {
	c.mu.Lock()
	c.FormString = v
	c.mu.Unlock()
	return c
}

func (c *Client) EndStruct(v interface{}) (res *http.Response, errs []error) {
	res, bs, errs := c.EndBytes()
	if errs != nil && len(errs) > 0 {
		c.Errors = append(c.Errors, errs...)
		return nil, c.Errors
	}
	switch c.UnmarshalType {
	case TypeJSON:
		err := json.Unmarshal(bs, &v)
		if err != nil {
			c.Errors = append(c.Errors, fmt.Errorf("json.Unmarshal：%s", err.Error()))
			return nil, c.Errors
		}
		return res, nil
	case TypeXML:
		err := xml.Unmarshal(bs, &v)
		if err != nil {
			c.Errors = append(c.Errors, fmt.Errorf("xml.Unmarshal：%s", err.Error()))
			return nil, c.Errors
		}
		return res, nil
	default:
		c.Errors = append(c.Errors, errors.New("UnmarshalType Type Wrong"))
		return nil, c.Errors
	}
}
func (c *Client) EndBytes() (res *http.Response, bs []byte, errs []error) {
	if len(c.Errors) > 0 {
		return nil, nil, c.Errors
	}
	var reader *strings.Reader
	switch c.Method {
	case GET:
		reader = strings.NewReader(null)
	case POST:
		switch c.ForceType {
		case TypeJSON:
			if c.JsonByte != nil {
				reader = strings.NewReader(string(c.JsonByte))
			}
			c.ContentType = Types[TypeJSON]
		case TypeForm, TypeFormData, TypeUrlencoded:
			reader = strings.NewReader(c.FormString)
			c.ContentType = Types[TypeForm]
		case TypeXML:
			reader = strings.NewReader(c.FormString)
			c.ContentType = Types[TypeXML]
			c.UnmarshalType = TypeXML
		default:
			c.Errors = append(c.Errors, errors.New("Request type Error "))
		}
	default:
		reader = strings.NewReader(null)
	}
	req, err := http.NewRequest(c.Method, c.Url, reader)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return nil, nil, c.Errors
	}
	req.Header.Set("Content-Type", c.ContentType)
	c.HttpClient.Transport = c.Transport
	res, err = c.HttpClient.Do(req)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return nil, nil, c.Errors
	}
	defer res.Body.Close()
	bs, err = ioutil.ReadAll(res.Body)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return nil, nil, c.Errors
	}
	return res, bs, nil
}