package chinaunionpay

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

var clnt = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			// DualStack: true,
		}).DialContext,
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		// ExpectContinueTimeout: 1 * time.Second,
		// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout: 40 * time.Second, // 整个http请求发起到等待应答的超时时间
}

// 调用者使用的时候最好设置一下自己的http参数
func SetDefaultHttpClient(client *http.Client) {
	clnt = client
}

func GetDefaultHTTPClient() *http.Client {
	return clnt
}

// (WARNING 带超时)
func httpReq(req *http.Request, headers map[string]string) (resp *http.Response, err error) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err = clnt.Do(req)
	if err != nil {
		return nil, err
	}

	return
}

// (WARNING 带超时)
func HttpGet(url string, headers map[string]string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return httpReq(req, headers)
}

// (WARNING 带超时)
func HttpPost(url string, headers map[string]string, data []byte) (resp *http.Response, err error) {
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	return httpReq(req, headers)
}

// 完整的HTTP GET (WARNING 带超时)
func HttpGetBody(url string, headers map[string]string) ([]byte, error) {
	resp, err := HttpGet(url, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// 完整的HTTP POST (WARNING 带超时)
func HttpPostBody(url string, headers map[string]string, data []byte) ([]byte, error) {
	resp, err := HttpPost(url, headers, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// 使用golang默认的http设置
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// 使用golang默认的http设置
// http上传文件
func PostFile(fullpath string, targetUrl string) ([]byte, error) {
	// 打开文件句柄操作
	fh, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	filename := filepath.Base(fullpath)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// 注意：post时文件太大，会超时，因为使用的clnt默认超时时间
	resp, err := clnt.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func UrlAddParam(targetUrl string, params map[string]string) string {
	urlParsed, err := url.Parse(targetUrl)
	if err != nil {
		return targetUrl
	}

	u := urlParsed.Query()

	for paramKey, paramVal := range params {
		u.Set(paramKey, paramVal)
	}
	urlParsed.RawQuery = u.Encode()

	return urlParsed.String()
}
