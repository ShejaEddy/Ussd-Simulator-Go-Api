package routes

import (
	"net/http"
)

type Routes []Route

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}
