package domain

type Cart struct {
	UUID     string
	Products []Product
}

type CartRepository interface {
	CreateCart() *Cart
	UpdateCart(cart *Cart) (*Cart, error)
	//DeleteItem(product Product)
	GetCart(cart string) (*Cart, error)
}
