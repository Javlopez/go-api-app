package service

import (
	"go-lana/pkg/domain"
)

type CartService struct {
	CartRepo    domain.CartRepository
	ProductRepo domain.ProductRepository
}

func (cs *CartService) CreateCart() *domain.Cart {
	return cs.CartRepo.CreateCart()
}

func (cs *CartService) AddItem(cartName string, items []string) (*domain.Cart, error) {
	var cart *domain.Cart

	cart, err := cs.CartRepo.GetCart(cartName)

	if err != nil {
		cart = cs.CreateCart()
	}

	for _, item := range items {
		product, _ := cs.ProductRepo.GetProductByCode(item)
		cart.Products = append(cart.Products, product)
	}

	return cs.CartRepo.UpdateCart(cart)
}

func (cs *CartService) GetCart(cart string) (*domain.Cart, error) {
	return cs.CartRepo.GetCart(cart)
}

func (cs *CartService) DeleteCart(cart string) error {
	return cs.CartRepo.DeleteCart(cart)
}
