package moenet

import (
	"fmt"
	"net/http"
	"net/url"
)

type MemoryCookieStorage struct {
	cookiedb map[string]*http.Cookie
}

func NewMemoryCookieStorage() (this *MemoryCookieStorage) {
	this = new(MemoryCookieStorage)
	this.cookiedb = make(map[string]*http.Cookie)
	return
}
func (this MemoryCookieStorage) SetCookies(u *url.URL, cookies []*http.Cookie) {
	for _, c := range cookies {
		v, has := this.cookiedb[c.Name]
		if has {
			fmt.Println(c.Name, "由", v.Value, "更新为", c.Value)
		} else {
			fmt.Println("新增", c.Name, "=", c.Value)
		}
		this.cookiedb[c.Name] = c
	}

}

func (this MemoryCookieStorage) Cookies(u *url.URL) []*http.Cookie {
	result := make([]*http.Cookie, 0)
	for _, c := range this.cookiedb {
		if len(u.Host) >= len(c.Domain) {
			//全局域名
			global := u.Host[len(u.Host)-len(c.Domain):]
			if global == c.Domain {
				result = append(result, c)
			}
		}
	}
	return result
}
