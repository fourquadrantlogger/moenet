package moenet

import "net/http"

var MoeClient = &http.Client{
	Jar: NewMemoryCookieStorage(),
}
