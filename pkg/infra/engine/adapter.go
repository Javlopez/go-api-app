package engine

import (
	"go-lana/pkg/infra/persistence/local"
	"go-lana/pkg/infra/storage"
)

type LocalAdapter struct {
	LocalStorage      *storage.LocalStorage
	cartRepository    *local.CartRepository
	productRepository *local.ProductRepository
}

func (la *LocalAdapter) CartRepository() *local.CartRepository {
	if la.cartRepository == nil {
		la.cartRepository = &local.CartRepository{
			LocalStorage: &storage.LocalStorage{},
		}
	}
	return la.cartRepository
}

//ProductRepository method
func (la *LocalAdapter) ProductRepository() *local.ProductRepository {
	if la.productRepository == nil {
		la.productRepository = &local.ProductRepository{}
	}
	return la.productRepository
}
