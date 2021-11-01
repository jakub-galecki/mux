package mux

import (
	"context"
	"net/http"
)

func NewMux() *Router {
	return &Router{}
}

func (_router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_route := _router.match(req.Method, req.URL)
	if _route == nil {
		// TODO handle error
	}
	_, params := parseURL(req.URL)
	ctx := context.WithValue(req.Context(), "params", params)
	req = req.WithContext(ctx)
	_route.handle(w, req)
}
