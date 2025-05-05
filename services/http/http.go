package http

import "net/http"

// Server defines the interface for an HTTP server
type HTTPServer interface {
    GET(path string, handler func(http.ResponseWriter, *http.Request))
    POST(path string, handler func(http.ResponseWriter, *http.Request))
    Serve(port string) error
}