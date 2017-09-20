package fungo

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

////////////////////////////////////////
// Service Options
///////////////////////////////////////
type ServiceOption func(*ServiceOptions)
type ServiceOptions struct {
	name string
	host string
	port int
	addr string
}

func (opts ServiceOptions) Addr() *url.URL {
	addr := opts.addr
	if opts.port != 0 {
		addr = fmt.Sprintf("%s:%d", opts.host, opts.port)
	}

	u, e := url.Parse(addr)
	if e != nil {
		panic("invalid addr")
	}

	return u
}

func Host(host string) ServiceOption {
	return func(o *ServiceOptions) {
		o.host = host
	}
}

func Port(port int) ServiceOption {
	return func(o *ServiceOptions) {
		o.port = port
	}
}

func PortStr(port string) ServiceOption {
	if strings.HasPrefix(port, ":") {
		port = port[1:]
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		panic("invalid port")
	}
	return func(o *ServiceOptions) {
		o.port = p
	}
}

func Addr(addr string) ServiceOption {
	return func(o *ServiceOptions) {
		o.addr = addr
	}
}

////////////////////////////////////////
// Work Options
///////////////////////////////////////
type WorkOption func(*WorkOptions)
type WorkOptions struct {
	label    string
	interval time.Duration
	cronSpec string
}

func Label(label string) WorkOption {
	return func(o *WorkOptions) {
		o.label = label
	}
}

func Interval(interval time.Duration) WorkOption {
	return func(o *WorkOptions) {
		o.interval = interval
	}
}

func CronSpec(spec string) WorkOption {
	return func(o *WorkOptions) {
		o.cronSpec = spec
	}
}
