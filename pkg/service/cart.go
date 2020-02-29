package service

import (
	"go-lana/pkg/domain"
	"go-lana/pkg/utils"
)

type CartService struct {
	Cart        *domain.Cart
	CartRepo    domain.CartRepository
	ProductRepo domain.ProductRepository
}

func (cs *CartService) CreateCart() *domain.Cart {
	cs.Cart = cs.CartRepo.CreateCart()
	return cs.CalculateTotalPrice().Cart
}

func (cs *CartService) AddItem(cartName string, items []string) (*domain.Cart, error) {
	var cart *domain.Cart
	cart, err := cs.GetCart(cartName)

	if err != nil {
		cart = cs.CreateCart()
	}

	for _, item := range items {
		product, err := cs.ProductRepo.GetProductByCode(item)
		product.PriceFormat = utils.FormatPrice(product.Price)
		if err != nil {
			return cart, err
		}
		cart.Products = append(cart.Products, product)
	}

	cartUpdated, err := cs.CartRepo.UpdateCart(cart)
	if err != nil {
		return cartUpdated, err
	}

	cs.Cart = cartUpdated
	cs.Cart = cs.CalculateTotalPrice().Cart

	return cs.Cart, nil
}

func (cs *CartService) CalculateTotalPrice() *CartService {
	var (
		penItems    []domain.Product
		tshirtItems []domain.Product
		mugItems    []domain.Product
		total       float32
	)

	for _, product := range cs.Cart.Products {
		switch product.Code {
		case "PEN":
			penItems = append(penItems, product)
		case "TSHIRT":
			tshirtItems = append(tshirtItems, product)
		case "MUG":
			mugItems = append(mugItems, product)
		}
	}

	total += NewPromotionForPEN(penItems).ApplyPromotion()
	total += NewPromotionForTshirt(tshirtItems).ApplyPromotion()
	total += NewPromotionForMugs(mugItems).ApplyPromotion()
	cs.Cart.Total = total
	cs.Cart.TotalPrice = utils.FormatPrice(cs.Cart.Total)
	return cs
}

func (cs *CartService) GetCart(cartName string) (*domain.Cart, error) {

	items := []domain.Product{}

	cart, err := cs.CartRepo.GetCart(cartName)
	if err != nil {
		return cart, err
	}

	for _, item := range cart.Products {
		product, err := cs.ProductRepo.GetProductByCode(item.Code)
		product.PriceFormat = utils.FormatPrice(product.Price)
		if err != nil {
			return cart, err
		}
		items = append(items, product)
	}

	cart.Products = items
	cs.Cart = cart
	cs.CalculateTotalPrice()
	return cs.Cart, nil
}

func (cs *CartService) DeleteCart(cart string) error {
	return cs.CartRepo.DeleteCart(cart)
}
