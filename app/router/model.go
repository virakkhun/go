package router

import "net/http"

type RouterHandler struct {
	Get    func(pattern string, handler http.HandlerFunc)
	Post   func(handler http.HandlerFunc)
	Patch  func(pattern string, handler http.HandlerFunc)
	Put    func(pattern string, handler http.HandlerFunc)
	Delete func(pattern string, handler http.HandlerFunc)
}

type Router struct {
	Group func(prefix string) RouterHandler
}
