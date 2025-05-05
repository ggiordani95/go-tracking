package http

import (
	"errors"
	"net/http"

	"github.com/ggiordani95/go-tracking/routes"
	"github.com/gin-gonic/gin"
)

// GinServer is an implementation of the Server interface using Gin
type GinServer struct {
    router         *gin.Engine
    routes   [] routes.Route // Lista de caminhos permitidos
}

// NewGinServer creates a new GinServer instance
func NewGinServer(allRoutes []routes.Route) *GinServer {
    return &GinServer{
        router:       gin.Default(),
        routes: allRoutes,
    }
}

// registerRoutes loops through the routes and registers them based on the method
func (g *GinServer) RegisterRoutes() error {
	for _, route := range g.routes {
		switch route.Method {
		case http.MethodGet:
			if err := g.GET(route.Path, route.Handler); err != nil {
				return err
			}
		case http.MethodPost:
			if err := g.POST(route.Path, route.Handler); err != nil {
				return err
			}
		default:
			return errors.New("unsupported method: " + route.Method)
		}
	}
	return nil
}

// GET registers a GET route
func (g *GinServer) GET(path string, handler func(http.ResponseWriter, *http.Request)) error {
    g.router.GET(path, func(c *gin.Context) {
        handler(c.Writer, c.Request)
    })
    return nil
}

// POST registers a POST route
func (g *GinServer) POST(path string, handler func(http.ResponseWriter, *http.Request)) error {
    g.router.POST(path, func(c *gin.Context) {
        handler(c.Writer, c.Request)
    })
    return nil
}

// Serve starts the HTTP server
func (g *GinServer) Serve(port string) error {
    return g.router.Run(port)
}