package service

import "go-lana/pkg/domain"

//ProductService struct
type ProductService struct {
	ProductRepo domain.ProductRepository
}

//GetAll method
func (p *ProductService) GetAll() ([]domain.Product, error) {
	return p.ProductRepo.All()
}

//GetProductByCode method
func (p *ProductService) GetProductByCode(code string) (domain.Product, error) {
	return p.ProductRepo.GetProductByCode(code)
}
