package domain

type Cart struct {
	UUID     string
	Products []Product
}

type CartRepository interface {
	CreateCart() *Cart
	UpdateCart(cart *Cart) (*Cart, error)
	DeleteCart(cart string) error
	GetCart(cart string) (*Cart, error)
}
