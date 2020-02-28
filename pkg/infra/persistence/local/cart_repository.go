package local

import (
	"encoding/json"
	"go-lana/pkg/domain"
	"go-lana/pkg/infra/storage"
	"go-lana/pkg/utils"
)

type CartRepository struct {
	LocalStorage *storage.LocalStorage
	cart         *domain.Cart
}

//CreateCart method: Create and save a new cart
func (cr *CartRepository) CreateCart() *domain.Cart {
	cr.cart = &domain.Cart{}
	cr.cart.UUID, _ = utils.GenerateUUID()
	cartJSON, _ := json.Marshal(cr.cart)
	cr.LocalStorage.Save(cr.cart.UUID, string(cartJSON))
	return cr.cart
}

//UpdateCart method: update cart already created
func (cr *CartRepository) UpdateCart(cart *domain.Cart) (*domain.Cart, error) {
	cartJSON, err := json.Marshal(cart)
	cr.LocalStorage.Save(cart.UUID, string(cartJSON))
	return cart, err
}

//GetCart method: Read content of the cart
func (cr *CartRepository) GetCart(name string) (*domain.Cart, error) {
	cart := &domain.Cart{}
	cartContent, err := cr.LocalStorage.Get(name)
	if err != nil {
		return cart, err
	}
	err = json.Unmarshal([]byte(cartContent), &cart)

	if err != nil {
		return cart, err
	}

	return cart, nil
}
