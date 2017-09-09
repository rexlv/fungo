package cream

type HandlerFunc func (Context) error

type Handler interface {
	Route() map[string]HandlerFunc
}

type RestHandler struct { }
func (handler RestHandler)	List(c Context) error {return notFoundHandler(c)}
func (handler RestHandler)	Get(c Context) error{return notFoundHandler(c)}
func (handler RestHandler)	Put(c Context) error{return notFoundHandler(c)}
func (handler RestHandler)	Post(c Context) error{return notFoundHandler(c)}
func (handler RestHandler)	Delete(c Context) error{return notFoundHandler(c)}
func (handler RestHandler)	Head(c Context) error{return notFoundHandler(c)}
func (handler RestHandler)	Patch(c Context) error{return notFoundHandler(c)}
func (handler RestHandler) Route() map[string]HandlerFunc {
	return map[string]HandlerFunc{
		"get /": handler.List,
		"get /:id": handler.Get,
		"put /:id": handler.Put,
		"post /": handler.Post,
		"delete /:id": handler.Delete,
		"head /": handler.Head,
		"patch /:id": handler.Patch,
	}
}
