package moenet

import (
	"net/url"
	"time"
)

type requestlog struct {
	Requrl *url.URL
	T      time.Time
}

func NewReqlog(method, urlraw string) (this requestlog) {
	u, _ := url.Parse(urlraw)
	this = requestlog{
		Requrl: u,
		T:      time.Now(),
	}
	return
}
