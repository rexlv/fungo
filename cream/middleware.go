package cream

type Middleware interface {
	Func() MiddlewareFunc
}

//type MiddlewareFunc func(HandlerFunc) HandlerFunc

//type BaseMiddleware struct { }
//
//func (mw BaseMiddleware) Func() MiddlewareFunc {
//	return func(next HandlerFunc) HandlerFunc {
//		return func (c Context) error {
//			return next(c)
//		}
//	}
//}

type MiddlewareFunc func(Context, HandlerFunc) error

func Recovery() MiddlewareFunc {
	return func(c Context, next HandlerFunc) error {
		return next(c)
	}
}
