package rest

import (
	"github.com/gofiber/fiber/v2"
	"market/adapters/controllers/rest/requests"
	"market/boundary/mapper"
	"market/domain/usecase"
)

type ProductController struct {
	useCase usecase.ProductUseCase
}

func NewProductController(useCase usecase.ProductUseCase) *ProductController {
	return &ProductController{useCase: useCase}
}

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	var req requests.ProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	productDTO := mapper.MapProductRequestToDTO(&req)

	if err := pc.useCase.CreateProduct(productDTO); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully"})
}
