package local

import (
	"go-lana/pkg/domain"

	"errors"
)

var products = []domain.Product{
	domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
	domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
	domain.Product{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.50},
}

var ProductNotFound = errors.New("The product does not exists")

type ProductRepository struct {
}

//All method
func (pr *ProductRepository) All() ([]domain.Product, error) {
	return products, nil
}

//GetProductByCode method
func (pr *ProductRepository) GetProductByCode(code string) (domain.Product, error) {
	for _, p := range products {
		if p.Code == code {
			return p, nil
		}
	}
	return domain.Product{}, ProductNotFound
}
