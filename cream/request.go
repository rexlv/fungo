package cream

import (
	"net/http"
	"net/url"
)

// Request request
type Request interface {
	Reset(*http.Request)
	URL() url.URL
	Cookie(string) (string, error)
	SetCookie(c *http.Cookie)
	Cookies() []*http.Cookie
}

// BaseRequest base request
type BaseRequest struct {
	*http.Request
	url url.URL
}

// Reset reset BaseRequest
func (req *BaseRequest) Reset(r *http.Request) {
	req.Request = r
}

// URL reset BaseRequest's URL
func (req *BaseRequest) URL() url.URL {
	return req.url
}

func (req *BaseRequest) Cookie(name string) (val string, err error) {
	return
}

func (req *BaseRequest) SetCookie(c *http.Cookie) {
	return
}

func (req *BaseRequest) Cookies() []*http.Cookie {
	return nil
}
