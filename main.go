package main

import (
	"go_ecommerce/controller"
	"go_ecommerce/db"
	"go_ecommerce/repository"
	"go_ecommerce/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbConnnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//CAMADA DE REPOSITORY
	ProductRepository := repository.NewProductRepository(dbConnnection)
	//CAMADA DE USE CASE
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// CAMADA DE CONTROLLER
	ProductController := controller.NewProductController(ProductUsecase)

	e.GET("/products", ProductController.GetProducts)
	e.Logger.Fatal(e.Start(":9090"))
}
