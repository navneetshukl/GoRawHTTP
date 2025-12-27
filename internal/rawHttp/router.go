package rawHttp

type Handler func(*Context)

type Route struct {
	method  string
	path    string
	handler Handler
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{routes: make([]Route, 0)}
}

func (r *Router) Handle(method, path string, h Handler) {
	r.routes = append(r.routes, Route{
		method:  method,
		path:    path,
		handler: h,
	})
}

// GET Method
func (r *Router) GET(path string, h Handler) {
	r.Handle("GET", path, h)
}

// POST Method
func (r *Router) POST(path string, h Handler) {
	r.Handle("POST", path, h)
}

// PUT Method
func (r *Router) PUT(path string, h Handler) {
	r.Handle("PUT", path, h)
}

// PATCH Method
func (r *Router) PATCH(path string, h Handler) {
	r.Handle("PATCH", path, h)
}

// DELETE Method
func (r *Router) DELETE(path string, h Handler) {
	r.Handle("DELETE", path, h)
}

