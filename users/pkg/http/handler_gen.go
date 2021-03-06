// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	endpoint "ironchip.net/kit/users/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeListHandler(m, endpoints, options["List"])
	makeGetHandler(m, endpoints, options["Get"])
	makeCreateHandler(m, endpoints, options["Create"])
	makeDeleteHandler(m, endpoints, options["Delete"])
	return m
}
