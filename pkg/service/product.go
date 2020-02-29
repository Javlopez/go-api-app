package service

import (
	"go-lana/pkg/domain"
	"go-lana/pkg/utils"
)

//ProductService struct
type ProductService struct {
	ProductRepo domain.ProductRepository
}

//GetAll method
func (p *ProductService) GetAll() ([]domain.Product, error) {

	var productsFormated []domain.Product

	products, err := p.ProductRepo.All()

	if err != nil {
		return products, err
	}

	for _, product := range products {
		product.PriceFormat = utils.FormatPrice(product.Price)
		productsFormated = append(productsFormated, product)
	}
	return productsFormated, nil
}

//GetProductByCode method
func (p *ProductService) GetProductByCode(code string) (domain.Product, error) {
	prod, err := p.ProductRepo.GetProductByCode(code)
	if err != nil {
		return prod, err
	}
	prod.PriceFormat = utils.FormatPrice(prod.Price)
	return prod, nil
}
