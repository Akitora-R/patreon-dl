package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type RequestMethod int

const (
	GET RequestMethod = iota
	POST
)

var ProxyUrl string

func init() {
	//_, _, err := NewGetRequest("https://www.google.com").SetTimeout(2000).Exec()
	//if err != nil {
	//	log.Println("上梯")
	//	useProxy = true
	//	return
	//}
	//log.Println("直连")
}

type Request struct {
	Url               string
	Method            RequestMethod
	QueryParam        map[string]string
	ProxyUrl          string
	Header            map[string][]string
	Body              []byte
	BodyForm          map[string]string
	BodyFileForm      map[string][]byte
	TimeoutMillSecond int
}

func NewGetRequest(url string) *Request {
	return &Request{
		Method:            GET,
		Url:               url,
		TimeoutMillSecond: -1,
	}
}

func NewPostRequest(url string) *Request {
	return &Request{
		Method:            POST,
		Url:               url,
		TimeoutMillSecond: -1,
	}
}

func (r *Request) SetQueryParam(queryParam map[string]string) *Request {
	r.QueryParam = queryParam
	return r
}

func (r *Request) SetHeader(header map[string][]string) *Request {
	r.Header = header
	return r
}

func (r *Request) SetBodyFileForm(fileForm map[string][]byte) *Request {
	r.BodyFileForm = fileForm
	return r
}

func (r *Request) SetBodyForm(form map[string]string) *Request {
	r.BodyForm = form
	return r
}

func (r *Request) SetBodyJson(body []byte) *Request {
	r.Body = body
	return r
}

func (r *Request) SetTimeout(millSec int) *Request {
	r.TimeoutMillSecond = millSec
	return r
}

func (r *Request) Exec() ([]byte, *http.Response, error) {
	switch r.Method {
	case GET:
		return r.execGet()
	case POST:
		if r.Body != nil {
			return r.execPostJson()
		} else if r.BodyForm != nil || r.BodyFileForm != nil {
			return r.execPostForm()
		}
	}
	return nil, nil, errors.New("不支持的方法")
}

func (r *Request) execGet() ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return nil, nil, err
	}
	query := req.URL.Query()
	for k, v := range r.QueryParam {
		query.Set(k, v)
	}
	req.URL.RawQuery = query.Encode()
	if r.Header != nil {
		req.Header = r.Header
	}
	client, err := r.getClient()
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return respBody, resp, nil
}

func (r *Request) ExecStream(handlerFunc func(respBytes []byte) error) error {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()
	for k, v := range r.QueryParam {
		query.Set(k, v)
	}
	req.URL.RawQuery = query.Encode()
	if r.Header != nil {
		req.Header = r.Header
	}
	client, err := r.getClient()
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(resp.Body)
	defer resp.Body.Close()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}
		go func() {
			err := handlerFunc(line)
			if err != nil {
				log.Println("读取http stream 发生错误:", err)
			}
			defer func() {
				if err := recover(); err != nil {
					log.Println("读取http stream 发生panic:", err)
				}
			}()
		}()
	}
}

func (r *Request) execPostForm() ([]byte, *http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	if r.BodyFileForm != nil {
		for fieldName, b := range r.BodyFileForm {
			formFileWriter, err := bodyWriter.CreateFormFile(fieldName, "file")
			if err != nil {
				fmt.Println("boterr writing to buffer")
				return nil, nil, err
			}
			if _, err = formFileWriter.Write(b); err != nil {
				return nil, nil, err
			}
		}
	}

	if r.BodyForm != nil {
		for f, v := range r.BodyForm {
			field, err := bodyWriter.CreateFormField(f)
			if err != nil {
				fmt.Println("boterr writing to buffer")
				return nil, nil, err
			}
			if _, err = field.Write([]byte(v)); err != nil {
				return nil, nil, err
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()
	req, err := http.NewRequest("POST", r.Url, bodyBuf)
	if err != nil {
		return nil, nil, err
	}
	if r.Header != nil {
		req.Header = r.Header
	}
	req.Header.Set("Content-Type", contentType)
	client, err := r.getClient()
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return b, resp, nil
}

func (r *Request) execPostJson() ([]byte, *http.Response, error) {
	req, _ := http.NewRequest("POST", r.Url, bytes.NewBuffer(r.Body))
	if r.Header != nil {
		req.Header = r.Header
	}
	req.Header.Set("Content-Type", "application/json")
	client, err := r.getClient()
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return body, resp, nil
}

func (r *Request) SetOptionalProxy() *Request {
	if ProxyUrl != "" {
		r.ProxyUrl = ProxyUrl
	}
	return r
}

func (r *Request) getClient() (*http.Client, error) {
	client := &http.Client{}
	if r.TimeoutMillSecond >= 0 {
		client.Timeout = time.Duration(r.TimeoutMillSecond) * time.Millisecond
	}
	if r.ProxyUrl != "" {
		if proxyUrl, err := url.Parse(r.ProxyUrl); err != nil {
			return nil, err
		} else {
			client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		}
	}
	return client, nil
}
