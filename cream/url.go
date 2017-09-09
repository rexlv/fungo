package cream

import "net/url"

type URL struct {
	url.URL
}

func (u URL) Get(key string) (val string, err error) {
	return
}