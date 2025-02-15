package main

import (
	"github.com/gofiber/fiber/v2"
	"market/controller"
	"market/domain/usecase"
)

func InitializeServer() *fiber.App {
	app := fiber.New()
	productController := controller.NewProductController(usecase.NewDummyProductUseCase())
	SetupRoutes(app, productController)
	return app
}

func SetupRoutes(app *fiber.App, productController *controller.ProductController) {
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
