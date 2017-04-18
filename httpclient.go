package moenet

import (
	"net/http"

	"bytes"
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
	mc.client.Jar = *(mc.Browser.cookies)
	return mc
}
func CopyClient(bs BrowserState) *MoeClient {
	mc := &MoeClient{
		client:  http.Client{},
		Browser: bs,
	}
	mc.client.Jar = *(mc.Browser.cookies)
	return mc
}
func (this *MoeClient) Do(Method string, urlraw string, bs []byte) (resp *http.Response, e error) {
	req, e := http.NewRequest(Method, urlraw, bytes.NewReader(bs))
	if e != nil {
		return nil, e
	}
	this.Browser.AddReqlog(NewReqlog(Method, urlraw))
	this.Browser.cookies = NewMemoryCookieStorage(this.client.Jar.Cookies(nil))
	return this.client.Do(req)
}
