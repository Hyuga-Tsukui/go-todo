package middleware

import (
	"net/http"
)

type handler struct {
	path string
	fn   http.HandlerFunc
}

func NewHandler(path string, fn http.HandlerFunc) handler {
	return handler{
		path: path,
		fn:   fn,
	}
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Router struct {
	middlewares []Middleware
	routes      map[string]http.HandlerFunc
}

func NewRouter(middlewares ...Middleware) *Router {
	return &Router{
		middlewares: middlewares,
		routes:      make(map[string]http.HandlerFunc),
	}
}

func (r *Router) RegistrationHandler(handlers ...handler) {
	for _, h := range handlers {
		fn := h.fn
		for i := len(r.middlewares) - 1; i >= 0; i-- {
			fn = r.middlewares[i](fn)
		}
		r.routes[h.path] = fn
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := r.routes[req.URL.Path]; ok {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}
