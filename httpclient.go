package moenet

import (
	"bytes"
	"net/http"
	"net/url"
)

type MoeClient struct {
	client  http.Client
	Browser BrowserState
}

func NewClient() *MoeClient {
	mc := &MoeClient{
		client:  http.Client{},
		Browser: *NewBroserState(),
	}
	mc.client.Jar = (mc.Browser.cookies)
	mc.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		mc.Browser.AddReqlog(NewReqlog(req.Method, req.URL.RequestURI()))
		return nil
	}
	return mc
}
func CopyClient(bs BrowserState) *MoeClient {
	mc := &MoeClient{
		client:  http.Client{},
		Browser: bs,
	}
	mc.client.Jar = (mc.Browser.cookies)
	mc.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		mc.Browser.AddReqlog(NewReqlog(req.Method, req.URL.RequestURI()))
		return nil
	}
	return mc
}

type MoeReq struct {
	Method, urlraw string
	forms, headers map[string]string
	Body           []byte
}

func MakeRequest(method string) (req MoeReq) {
	mr := MoeReq{
		Method:  method,
		headers: make(map[string]string),
		forms:   make(map[string]string),
	}
	return mr
}
func (req MoeReq) DefaultSetting() MoeReq {
	req.Header("upgrade-insecure-requests", "1")
	req.Header("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/56.0.2924.76 Chrome/56.0.2924.76 Safari/537.36")
	req.Header("content-type", "application/x-www-form-urlencoded")
	req.Header("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header("accept-language", "en,zh-CN;q=0.8,zh;q=0.6")
	return req
}
func (req MoeReq) Url(url string) MoeReq {
	req.urlraw = url
	return req
}

func (req MoeReq) Form(key, value string) MoeReq {
	req.forms[key] = value
	return req
}
func (req MoeReq) Header(key, value string) MoeReq {
	req.headers[key] = value
	return req
}
func (req MoeReq) Referer(value string) MoeReq {
	req.headers["Referer"] = value
	return req
}
func (req MoeReq) ContentType(value string) MoeReq {
	req.headers["Content-Type"] = value
	return req
}
func (req MoeReq) Origin(value string) MoeReq {
	req.headers["origin"] = value
	return req
}

func (this *MoeClient) Do(req MoeReq) (resp *http.Response, e error) {
	if len(req.forms) > 0 {
		vs := url.Values{}
		for k, v := range req.forms {
			vs.Add(k, v) //
		}

		if req.Body != nil {
			panic("exist body,form it?")
		}
		req.Body = []byte(vs.Encode())
	}
	var httpreq *http.Request
	if req.Method == "GET" {
		httpreq, e = http.NewRequest(req.Method, req.urlraw, nil)
		if e != nil {
			return nil, e
		}
	} else {
		httpreq, e = http.NewRequest(req.Method, req.urlraw, bytes.NewReader(req.Body))
		if e != nil {
			return nil, e
		}
	}

	for k, v := range req.headers {
		httpreq.Header.Add(k, v)
	}

	this.Browser.AddReqlog(NewReqlog(req.Method, req.urlraw))

	resp, e = this.client.Do(httpreq)
	this.Browser.cookies = NewMemoryCookieStorage(this.client.Jar.Cookies(nil))
	return
}
