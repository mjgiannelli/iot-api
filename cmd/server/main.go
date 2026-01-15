package main

import (
	"log"
	"net/http"
	"os"
	"time"

	httpserver "github.com/mjgiannelli/iot-api/internal/http"
)

func main() {
	addr := ":3001"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	srv := http.Server{
		Addr: addr,
		Handler: httpserver.New(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Printf("listening on %s", addr)
	log.Fatal(srv.ListenAndServe())
}