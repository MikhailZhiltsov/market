package main

import (
	"github.com/gofiber/fiber/v2"
	"market/adapters/controllers/rest"
	"market/domain/usecase"
)

func InitializeServer() *fiber.App {
	app := fiber.New()
	productController := rest.NewProductController(usecase.NewDummyProductUseCase())
	SetupRoutes(app, productController)
	return app
}

func SetupRoutes(app *fiber.App, productController *rest.ProductController) {
	app.Post("/products", productController.CreateProduct)
}

func StartServer(app *fiber.App, address string) error {
	return app.Listen(address)
}

func main() {
	app := InitializeServer()
	if err := StartServer(app, ":3000"); err != nil {
		panic(err)
	}
}
