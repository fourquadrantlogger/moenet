package moenet

import (
	"encoding/json"
	"net/url"
)

type BrowserState struct {
	cookies *MemoryCookieStorage
	history []*requestlog
}

func (this *BrowserState) AddReqlog(req requestlog) {
	this.history = append(req)
}
func (this *BrowserState) GetCookies() *MemoryCookieStorage {
	return this.cookies
}
func (this *BrowserState) NowUrl() *url.URL {
	if len(this.history) == 0 {
		return nil
	}
	return this.history[len(this.history)-1].requrl
}
func (this *BrowserState) String() string {
	var js map[string]interface{}
	js["cookies"] = this.cookies
	js["history"] = this.history

	bs, _ := json.Marshal(js)
	return string(bs)
}
func ParseBrowserState(s string) (this *BrowserState) {
	json.Unmarshal([]byte(s), this)
	return
}
