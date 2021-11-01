package mux

import (
	"net/http"
	"net/url"
)

type route struct {
	method         string
	pattern        string
	expectedParams []string
	handler        http.HandlerFunc
}

func (_route *route) handle(w http.ResponseWriter, req *http.Request) {
	_route.handler.ServeHTTP(w, req)
}

func parseURL(url *url.URL) (string, url.Values) {
	pattern := url.Path
	return pattern, url.Query()
}
