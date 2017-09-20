package middleware

import (
	"fmt"

	"github.com/rexlv/fungo/cream"
)

var DefaultRecovery Recovery

type Recovery struct {
}

func (mw Recovery) Func() cream.MiddlewareFunc {
	return func(next cream.HandlerFunc) cream.HandlerFunc {
		return func(c cream.Context) error {
			defer func() {
				fmt.Println("recovery")
			}()
			return next(c)
		}
	}
}
