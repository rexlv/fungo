package cream

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

type Context interface {
	Status(int) error
	Request() Request
	Response() Response

	QueryInt(string, int)
	QueryString(string, string)
	QueryUint(string, uint)
	QueryDuration(string, time.Duration)
	QueryBool(string, bool)

	Get(string interface{})
	Set(string, interface{})

	Cookie(string) (string, error)
	SetCookie(*http.Cookie)
	Cookies() []*http.Cookie

	Bind(i interface{}) error

	JSON(int, interface{}) error
	XML(int, interface{}) error
	Redirect(int, string) error
}

type DebugContext struct {
	BaseContext
}

type BaseContext struct {
	context.Context
	req Request
	rep Response
}

func (ctx BaseContext) Request() Request {
	return ctx.req
}

func (ctx BaseContext) Response() Response {
	return ctx.rep
}

func (ctx BaseContext) Query(key string) (val string, err error) {
	return ctx.req.URL().Get(key)
}

func (ctx BaseContext) QueryInt(key string, expect int) (ret int) {
	ret = expect
	if val, err := ctx.Query(key); err == nil {
		if v, e := strconv.ParseInt(val, 10, 64); e != nil {
			ret = int(v)
		}
	}
	return
}

func (ctx BaseContext) QueryUint(key string, expect uint) (ret uint) {
	ret = expect
	if val, err := ctx.Query(key); err == nil {
		if v, e := strconv.ParseUint(val, 10, 64); e != nil {
			ret = uint(v)
		}
	}
	return
}

func (ctx BaseContext) QueryString(key string, expect string) (ret string) {
	ret = expect
	if val, err := ctx.Query(key); err == nil {
		ret = val
	}

	return
}

func (ctx BaseContext) QueryBool(key string, expect bool) (ret bool) {
	ret = expect
	if val, err := ctx.Query(key); err == nil {
		if v, e := strconv.ParseBool(val); e == nil {
			ret = v
		}
	}
	return
}

func (ctx BaseContext) Cookie(name string) (*http.Cookie, error) {
	return ctx.req.Cookie(name)
}

func (ctx BaseContext) SetCookie(cookie *http.Cookie) {
	ctx.req.SetCookie(cookie)
}

func (ctx BaseContext) Cookies() []*http.Cookie {
	return ctx.req.Cookies()
}

func (ctx BaseContext) Bind(i interface{}) error {
	return nil
}

func (ctx BaseContext) JSON(code int, obj interface{}) error {
	ctx.rep.SetStatus(code)
	ctx.rep.SetData(obj)
	return nil
}

func (ctx BaseContext) XML(int, interface{}) error { return nil }
func (ctx BaseContext) Redirect(int, string) error { return nil }
