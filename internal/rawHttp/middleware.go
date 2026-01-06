package rawHttp

func (r *Router) UseMiddleware(m ...Handler) {
	r.middleware = append(r.middleware, m...)
}
