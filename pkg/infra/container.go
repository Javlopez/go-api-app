package infra

import (
	"go-lana/pkg/infra/engine"
	"go-lana/pkg/service"
)

//ContainerInfra interface contains all the services available to the handlers
type ContainerInfra interface {
	CartService() *service.CartService
	ProductService() *service.ProductService
}

//Container struct for DI
type Container struct {
	engine.LocalAdapter
	cartService    *service.CartService
	productService *service.ProductService
}

//ProductService method for create unique instance of service.ProductService
func (c *Container) ProductService() *service.ProductService {
	if c.productService == nil {
		c.productService = &service.ProductService{
			ProductRepo: c.ProductRepository(),
		}
	}
	return c.productService
}

//CartService method for create unique instance of service.CartService
func (c *Container) CartService() *service.CartService {
	if c.cartService == nil {
		c.cartService = &service.CartService{
			CartRepo:    c.CartRepository(),
			ProductRepo: c.ProductRepository(),
		}
	}
	return c.cartService
}
