package repository

import (
	"database/sql"
	"fmt"
	"go_ecommerce/model"
)

type ProductRepository struct {
	connection *sql.DB
}

// METHODS
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM products"
	row, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for row.Next() {
		err := row.Scan(
			&productObj.Id,
			&productObj.Product_name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	row.Close()

	return productList, nil
}

// FUNCTIONS
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}
