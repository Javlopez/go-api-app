package domain

//Product struct
type Product struct {
	Code  string
	Name  string
	Price float32
}

//ProductRepository repository
type ProductRepository interface {
	All() ([]Product, error)
	GetProductByCode(code string) (Product, error)
}
