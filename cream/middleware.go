package cream

type MiddlewareFunc func(HandlerFunc) HandlerFunc
type Middleware interface {
	Func() MiddlewareFunc
}

type BaseMiddleware struct {
}

func (mw BaseMiddleware) Func() MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return next(c)
		}
	}
}
