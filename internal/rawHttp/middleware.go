package rawHttp

import "log"

func (r *Router) UseMiddleware(f func()) {
	log.Println("Triggered use middleware")
	r.middleware = append(r.middleware, Middleware{
		middleware: f,
	})
}
