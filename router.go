package resttk

import (
	"net/http"
)

type Router struct {
	routes []*Route
}

func (r *Router) AddRoute(path string, controller func() ControllerInterface) *Route {
	route := NewRoute(path, controller)
	r.routes = append(r.routes, route)
	return route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//Find the proper handler by parsing the req.Path and looking up the proper route
	route := r.findRoute(req)
	if route != nil {
		// If not a websocket apply any filters available
		if req.Header.Get("Upgrade") != "websocket" {
			for _, filter := range route.filters {
				if req.Method != "OPTIONS" {
					if !filter.Apply(w, req) {
						return
					}
				}
			}
		}

		// After filters have been run create the controller instance and serve it
		handler := route.handler()
		handler.Init(w, req, handler, route.path)
		handler.ServeHTTP(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (r *Router) findRoute(request *http.Request) *Route {
	for _, route := range r.routes {
		if route.regex.MatchString(request.URL.Path) {
			return route
		}
	}
	return nil
}

func NewRouter() *Router {
	return &Router{}
}
