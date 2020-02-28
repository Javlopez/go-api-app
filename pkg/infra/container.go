package infra

import (
	"go-lana/pkg/infra/engine"
	"go-lana/pkg/service"
)

//ContainerInfra interface
type ContainerInfra interface {
	CartService() *service.CartService
	ProductService() *service.ProductService
}

//Container struct
type Container struct {
	engine.LocalAdapter
	cartService    *service.CartService
	productService *service.ProductService
}

func (c *Container) ProductService() *service.ProductService {
	if c.productService == nil {
		c.productService = &service.ProductService{
			ProductRepo: c.ProductRepository(),
		}
	}
	return c.productService
}

func (c *Container) CartService() *service.CartService {
	if c.cartService == nil {
		c.cartService = &service.CartService{
			CartRepo:    c.CartRepository(),
			ProductRepo: c.ProductRepository(),
		}
	}
	return c.cartService
}
