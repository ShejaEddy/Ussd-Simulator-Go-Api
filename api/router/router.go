package router

import (
	"github.com/gorilla/mux"
	Routes "github.com/projects/ussd/api/router/routes"
)

type RouterHandler struct {
	Router *mux.Router
}

func NewRouter() *RouterHandler {
	var r RouterHandler
	r.Router = mux.NewRouter().StrictSlash(true)

	routes := Routes.GetRoutes()

	for _, route := range routes {
		r.Router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Path).
			Handler(route.Handler)
	}

	return &r
}
