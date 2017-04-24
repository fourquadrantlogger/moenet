package moenet

import (
	"encoding/json"

	"strings"
	//"fmt"
)

type BrowserState struct {
	Cookies *MemoryCookieStorage `json:"cookies"`
	History []*requestlog        `json:"history"`
	Obj     map[string]string    `json:"obj"`
}

func NewBroserState() *BrowserState {
	this := new(BrowserState)
	this.Cookies = new(MemoryCookieStorage)
	this.History = make([]*requestlog, 0)
	this.Obj = make(map[string]string)
	return this
}
func (this *BrowserState) AddReqlog(req requestlog) {
	this.History = append(this.History, &req)

}
func (this *BrowserState) GetCookies() *MemoryCookieStorage {
	return this.Cookies
}

func (this *BrowserState) LastReq() *requestlog {
	if len(this.History) == 0 {
		return nil
	}

	re := this.History[len(this.History)-1]
	return re
}
func (this *BrowserState) LastHost() string {
	if len(this.History) == 0 {
		return ""
	}

	host := this.History[len(this.History)-1].Requrl.Host
	if strings.Contains(host, ":") {
		host = host[:strings.Index(host, ":")]
	}
	return host
}
func (this *BrowserState) Bytes() []byte {

	bs, _ := json.Marshal(this)
	return bs
}
func ParseBrowserState(s []byte) (this *BrowserState) {
	this = new(BrowserState)
	json.Unmarshal(s, this)
	return
}
