package infrastructure

import (
	"go-lana/pkg/infrastructure/engine"
	"go-lana/pkg/service"
)

// ContainerInfrastructure interface
type ContainerInfrastructure interface {
	ProductService() *service.ProductService
}

// Container struct
type Container struct {
	localAdapter   engine.LocalAdapter
	productService *service.ProductService
}

func (c *Container) ProductService() *service.ProductService {
	if c.productService == nil {
		c.productService = &service.ProductService{
			ProductRepo: c.localAdapter.ProductRepository(),
		}
	}
	return c.productService
}
