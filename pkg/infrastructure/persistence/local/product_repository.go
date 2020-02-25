package local

import "go-lana/pkg/domain"

var products = []domain.Product{
	domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
	domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
	domain.Product{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.50},
}

type ProductRepository struct {
}

//All method
func (pr *ProductRepository) All() ([]domain.Product, error) {
	return products, nil
}

//GetProductByCode method
func (pr *ProductRepository) GetProductByCode(code string) (domain.Product, error) {
	return products[0], nil
}
