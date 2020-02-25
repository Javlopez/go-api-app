package engine

import "go-lana/pkg/infrastructure/persistence/local"

type LocalAdapter struct {
	productRepository *local.ProductRepository
}

//Connect method
func (la *LocalAdapter) Connect() LocalAdapter {
	return *la
}

//ProductRepository method
func (la *LocalAdapter) ProductRepository() *local.ProductRepository {
	if la.productRepository == nil {
		la.productRepository = &local.ProductRepository{}
	}

	return la.productRepository
}
