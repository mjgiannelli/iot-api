package http

import (
	"net/http"

	"github.com/mjgiannelli/iot-api/internal/health"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", health.Handler)
}