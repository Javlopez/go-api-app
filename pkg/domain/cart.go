package domain

type Cart struct {
	UUID       string
	Products   []Product
	Total      float32 `json:"-"`
	TotalPrice string  `json:"TotalPrice"`
}

type CartRepository interface {
	CreateCart() *Cart
	UpdateCart(cart *Cart) (*Cart, error)
	DeleteCart(cart string) error
	GetCart(cart string) (*Cart, error)
}
