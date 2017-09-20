package cream

import (
	"net"
	"net/http"
	"strings"
	"sync"
)

// Server http server
type Server struct {
	router   *Router
	listener net.Listener

	mws []Middleware

	pool *sync.Pool
}

// New returns new http server instance
func New() *Server {
	return &Server{
		router: newRouter(),
		mws:    make([]Middleware, 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return nil
			},
		},
	}
}

// Serve implements fungo#Service interface
func (s *Server) Serve(lis net.Listener) error {
	return nil
}

// Stop implements fungo#Service interface
func (s *Server) Stop() {
	s.listener.Close()
}

// GracefulStop implements fungo#Service interface
func (s *Server) GracefulStop() {
}

// Addr addr
func (s *Server) Addr() net.Addr {
	return s.listener.Addr()
}

// Use use
func (s *Server) Use(mws ...Middleware) {
	s.mws = append(s.mws, mws...)
}

func (s *Server) Handle(path string, handler Handler, mws ...Middleware) {
	for mpath, fn := range handler.Route() {
		fields := strings.Split(mpath, " ")
		if len(fields) != 2 {
			panic("invalid method_path")
		}
		method := strings.TrimSpace(fields[0])
		path := strings.TrimSpace(fields[1])
		s.add(method, path, fn, mws...)
	}
}

func (s *Server) GET(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("GET", path, handler, append(s.mws, mws...)...)
}
func (s *Server) POST(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("POST", path, handler, append(s.mws, mws...)...)
}
func (s *Server) PUT(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("PUT", path, handler, append(s.mws, mws...)...)
}
func (s *Server) PATCH(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("PATCH", path, handler, append(s.mws, mws...)...)
}
func (s *Server) HEAD(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("HEAD", path, handler, append(s.mws, mws...)...)
}
func (s *Server) DELETE(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("DELETE", path, handler, append(s.mws, mws...)...)
}
func (s *Server) OPTIONS(path string, handler HandlerFunc, mws ...Middleware) {
	s.add("OPTIONS", path, handler, append(s.mws, mws...)...)
}

func (s *Server) add(method, path string, handler HandlerFunc, mws ...Middleware) {
	h := handler
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i].Func()(h)
	}
	s.router.add(method, path, h)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// c := s.pool.Get().(Context)
	c := &BaseContext{
		req: &BaseRequest{},
		rep: &BaseResponse{},
	}

	c.Request().Reset(r)
	c.Response().Reset(w)

	method := r.Method
	path := r.URL.RawPath
	if path == "" {
		path = r.URL.Path
	}

	if err := func(c Context) error {
		node, _ := s.router.find(method, path)
		handler := node.h[method].(HandlerFunc)
		return handler(c)
	}(c); err != nil {
		panic(err)
	}

	s.pool.Put(c)
}
