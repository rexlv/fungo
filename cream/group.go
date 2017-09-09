package cream

// Group route group
type Group struct {
	s   *Server
	mws []Middleware
}

// GET method
func (g *Group) GET(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("GET", path, handler, append(g.mws, mws...)...)
}

// POST method
func (g *Group) POST(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("POST", path, handler, append(g.mws, mws...)...)
}

// PUT method
func (g *Group) PUT(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("PUT", path, handler, append(g.mws, mws...)...)
}

// PATCH method
func (g *Group) PATCH(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("PATCH", path, handler, append(g.mws, mws...)...)
}

// HEAD method
func (g *Group) HEAD(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("HEAD", path, handler, append(g.mws, mws...)...)
}

// DELETE method
func (g *Group) DELETE(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("DELETE", path, handler, append(g.mws, mws...)...)
}

// OPTIONS method
func (g *Group) OPTIONS(path string, handler HandlerFunc, mws ...Middleware) {
	g.add("OPTIONS", path, handler, append(g.mws, mws...)...)
}

func (g *Group) add(method, path string, handler HandlerFunc, mws ...Middleware) {
	g.s.add(method, path, handler, append(g.mws, mws...)...)
}
