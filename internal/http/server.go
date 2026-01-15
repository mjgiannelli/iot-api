package http

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}