package routes

import (
	"net/http"

	"github.com/ggiordani95/go-tracking/controllers"
)

// Route struct to hold path and method
type Route struct {
	Path   string
	Method string
	Handler func(http.ResponseWriter, *http.Request)
}

// List of routes
var Routes = []Route{
	{Path: "/health-check", Method: "GET", Handler: controllers.HealthCheckController()},
	{Path: "/api/v1/track", Method: "POST", Handler: controllers.HealthCheckController()},
}
