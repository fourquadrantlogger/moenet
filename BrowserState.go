package moenet

import "net/url"

type BrowserState struct {
	cookies *MemoryCookieStorage
	history []*requestlog
}

func (this *BrowserState)AddReqlog(req requestlog){
	this.history=append(req)
}
func (this *BrowserState)GetCookies()(*MemoryCookieStorage){
	return this.cookies
}
func (this *BrowserState)NowUrl()(*url.URL){
	if(len(this.history)==0){
		return nil
	}
	return this.history[len(this.history)-1].requrl
}
