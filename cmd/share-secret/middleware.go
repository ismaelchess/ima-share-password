package main

import "net/http"

// Middleware represents a function which acts as man in the middle for a http request
type Middleware func(http.Handler) http.Handler

// Apply attaches the middlewares to a handler
func Apply(handler http.Handler, middlewares ...Middleware) http.Handler {
	chain := handler
	for _, m := range middlewares {
		chain = m(chain)
	}
	return chain
}
