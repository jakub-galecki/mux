package mux

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Router struct {
	routes []*route
}

func (_router *Router) registerRoute(method string, pattern string,
	expectedParams []string, handler http.HandlerFunc) {
	newRoute := route{method, pattern, expectedParams, handler}
	_router.routes = append(_router.routes, &newRoute)
}

func (_router *Router) match(method string, url *url.URL) *route {
	pattern, expectedParams := parseURL(url)
	for _, _route := range _router.routes {
		if strings.EqualFold(method, _route.method) && strings.EqualFold(pattern, _route.pattern) &&
			reflect.DeepEqual(reflect.ValueOf(expectedParams).MapKeys(), _route.expectedParams) {
			return _route
		}
	}
	return nil
}
