package model

type Product struct {
	Id           int
	Product_name string `json:"name"`
	Price        float64
}
