package controller

import (
	"go_ecommerce/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

// METHODS
func (pc *ProductController) GetProducts(ctx echo.Context) error {
	products, err := pc.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusAccepted, products)
}

// FUNCTIONS
func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}
