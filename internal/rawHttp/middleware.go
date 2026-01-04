package rawHttp

func (r *Router) UseMiddleware(m Middleware) {
	r.middleware = append(r.middleware, m)
}
