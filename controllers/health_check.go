package controllers

import (
	"net/http"
)

// HealthCheckHandler handles the health check request
func HealthCheckController() func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // Respond with a success message
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "OK"}`))
    }
}