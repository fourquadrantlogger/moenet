package moenet

import (
	"net/url"
	"time"
)

type requestlog struct {
	requrl *url.URL
	t time.Time
}
