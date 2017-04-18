package moenet

import (
	"fmt"
	"net/http"
	"net/url"
)

type MemoryCookieStorage struct {
	cookiedb map[string]*http.Cookie
}

func NewMemoryCookieStorage(cookies []*http.Cookie) (this *MemoryCookieStorage) {
	this = new(MemoryCookieStorage)
	this.cookiedb = make(map[string]*http.Cookie)
	for _, c := range cookies {
		this.cookiedb[c.Name] = c
	}
	return
}
func (this *MemoryCookieStorage) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if this.cookiedb == nil {
		this.cookiedb = make(map[string]*http.Cookie)
	}
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

func (this *MemoryCookieStorage) Cookies(u *url.URL) []*http.Cookie {
	result := make([]*http.Cookie, 0)

	for _, c := range this.cookiedb {
		if u == nil {
			result = append(result, c)
			continue
		}

		if len(u.Host) >= len(c.Domain) {
			//全局域名
			global := u.Host[len(u.Host)-len(c.Domain):]
			if global == c.Domain {
				result = append(result, c)
			}
		}
	}
	fmt.Println("当前所有", this.cookiedb, u, "使用cookie", result)
	return result
}
func (this *MemoryCookieStorage) Cookie(key string) *http.Cookie {

	for _, c := range this.cookiedb {
		if key == c.Name {
			return c
		}
	}

	return nil
}
