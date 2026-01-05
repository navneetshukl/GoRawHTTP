package rawHttp

func (r *Router) UseMiddleware(m ...Handler) {
	r.routes = append(r.routes, Route{handler: m})
}
