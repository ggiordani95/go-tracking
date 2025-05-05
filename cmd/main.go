package main

import (
	routes "github.com/ggiordani95/go-tracking/routes"
	customHttp "github.com/ggiordani95/go-tracking/services/http"
)

func main() {
	server := customHttp.NewGinServer(routes.Routes)
	server.RegisterRoutes()
	server.Serve(":8080")
}