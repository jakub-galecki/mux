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

func (_router *Router) RegisterRoute(method string, pattern string,
	expectedParams []string, handler http.HandlerFunc) {
	newRoute := &route{method, pattern, expectedParams, handler}
	_router.routes = append(_router.routes, newRoute)
}

func (_router *Router) match(method string, url *url.URL) (*route, *Error) {
	pattern, expectedParams := parseURL(url)
	checkParams := func(arr1 []reflect.Value, arr2 []string) bool {
		if len(arr1) == 0 && len(arr2) == 0 {
			return true
		}
		if len(arr1) != len(arr2) {
			return false
		}
		exists := make(map[string]bool)
		for _, v := range arr1 {
			exists[v.Interface().(string)] = true
		}
		for _, v := range arr2 {
			if !exists[v] {
				return false
			}
		}
		return true
	}
	for _, _route := range _router.routes {
		if strings.EqualFold(method, _route.method) && strings.EqualFold(pattern, _route.pattern) &&
			(checkParams(reflect.ValueOf(expectedParams).MapKeys(), _route.expectedParams)) {
			return _route, nil
		}
	}
	err := &Error{"No such route", 500}
	return nil, err
}
