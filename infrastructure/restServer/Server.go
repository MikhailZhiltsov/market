package restServer

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Server interface {
	RegisterPublicRoute(method, path string, handler http.HandlerFunc)
	Start(address string) error
}

type FiberServer interface {
	RegisterRoute(method, path string, handler fiber.Handler)
	Start(address string) error
	Shutdown() error
}
