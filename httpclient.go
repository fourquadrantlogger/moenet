package moenet

import (
	"net/http"

	"bytes"
)

type MoeClient struct {
	http.Client
	Browser BrowserState
}

func NewClient() *MoeClient {
	return &MoeClient{
		Client: http.Client{
			Jar: NewMemoryCookieStorage(),
		},
		Browser: *NewBroserState(),
	}
}
func (this *MoeClient) Do(Method string, urlraw string, bs []byte) (resp *http.Response, e error) {
	req, e := http.NewRequest(Method, urlraw, bytes.NewReader(bs))
	if e != nil {
		return nil, e
	}
	return this.Client.Do(req)
}
