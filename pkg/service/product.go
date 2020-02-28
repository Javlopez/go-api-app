package service

import "go-lana/pkg/domain"

import "fmt"

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
		product.PriceFormat = fmt.Sprintf("%0.02f€", product.Price)
		productsFormated = append(productsFormated, product)
	}
	return productsFormated, nil
}

//GetProductByCode method
func (p *ProductService) GetProductByCode(code string) (domain.Product, error) {
	return p.ProductRepo.GetProductByCode(code)
}
