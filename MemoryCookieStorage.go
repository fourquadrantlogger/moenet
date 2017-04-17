package moenet

import (
	"net/http"
	"net/url"
	"log"

)
var speaker log.Logger
type MemoryCookieStorage struct {
	cookiedb map[string]*http.Cookie
}

func (this *MemoryCookieStorage)SetCookies(u *url.URL, cookies []*http.Cookie){
	for _,c:=range cookies{
		v,has:=this.cookiedb[c.Name]
		if(has){
			speaker.Println(c.Name,"由",v.Value,"更新为",c.Value)
		}else {
			speaker.Println("新增",c.Value,"=",v.Value)
		}
		this.cookiedb[c.Name]=c
	}

}

func (this *MemoryCookieStorage)Cookies(u *url.URL) []*http.Cookie{
	result:=make([]*http.Cookie,0)
	for _,c:=range this.cookiedb{
		if(len(u.Host)>=len(c.Domain)){
			//全局域名
			global:=u.Host[len(u.Host)-len(c.Domain):]
			if(global==c.Domain){
				result=append(c)
			}
		}
	}
	return result
}