package service

import (
	"errors"
	"go-lana/pkg/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cartExists = true
	fileExists = true
)

type CartRepositoryMock struct {
	mock.Mock
}

func (cm *CartRepositoryMock) CreateCart() *domain.Cart {
	cart := &domain.Cart{}
	cart.UUID = "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"
	return cart
}

func (cm *CartRepositoryMock) GetCart(name string) (*domain.Cart, error) {
	var err error
	cart := &domain.Cart{}
	if cartExists {
		cart.UUID = "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"
	} else {
		err = errors.New("The cart does not exists")
	}
	return cart, err
}

func (cm *CartRepositoryMock) UpdateCart(cart *domain.Cart) (*domain.Cart, error) {
	return cart, nil
}

func (cm *CartRepositoryMock) DeleteCart(cartName string) error {
	if fileExists {
		return nil
	}
	return errors.New("remove storage/11E5C5D2-B56B-B588-57F9-8F77A05FEEE8: no such file or directory")
}

func TestCart(t *testing.T) {

	t.Run("Cart service should be able to create new cart", func(t *testing.T) {
		cartRepo := CartRepositoryMock{}
		want := &domain.Cart{UUID: "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"}
		cartRepo.On("CreateCart").Return(want, nil)
		cartSvc := &CartService{CartRepo: &cartRepo}
		cart := cartSvc.CreateCart()
		assert.NotEmpty(t, cart.UUID)
		assert.Equal(t, want.UUID, cart.UUID)
		assert.Nil(t, cart.Products)
	})

	t.Run("Cart service should be able to get a cart", func(t *testing.T) {
		cartRepo := CartRepositoryMock{}
		want := &domain.Cart{UUID: "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"}
		cartRepo.On("CreateCart").Return(want, nil)

		wantCart := &domain.Cart{UUID: "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"}
		cartRepo.On("GetCart", want.UUID).Return(wantCart, nil)

		cartSvc := &CartService{CartRepo: &cartRepo}
		_ = cartSvc.CreateCart()

		cartExists, err := cartSvc.GetCart("11E5C5D2-B56B-B588-57F9-8F77A05FEEE8")

		if err != nil {
			t.Logf("Cannot write read the data due: %s", err.Error())
		}

		assert.NotEmpty(t, cartExists.UUID)
		assert.Equal(t, want.UUID, cartExists.UUID)
		assert.Equal(t, []domain.Product{}, cartExists.Products)
	})

	t.Run("Cart service should be able to add items in a cart", func(t *testing.T) {
		cartRepo := CartRepositoryMock{}
		want := &domain.Cart{
			UUID: "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8",
			Products: []domain.Product{
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00, PriceFormat: "20.00€"},
			},
		}

		cartRepo.On("UpdateCart", want.UUID, []string{"PEN"}).Return(want, nil)

		wantProduct := domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00, PriceFormat: "20.00€"}
		productRepo := ProductRepositoryMock{}
		productRepo.On("GetProductByCode", "PEN").Return(wantProduct, nil)

		cartSvc := &CartService{CartRepo: &cartRepo, ProductRepo: &productRepo}

		cartExists, err := cartSvc.AddItem(want.UUID, []string{"PEN"})

		if err != nil {
			t.Logf("Cannot write read the data due: %s", err.Error())
		}

		assert.Equal(t, want.UUID, cartExists.UUID)
		assert.Equal(t, want.Products, cartExists.Products)
	})

	t.Run("Cart service should be able to add items in a cart even if the cart does not exists", func(t *testing.T) {
		cartRepo := CartRepositoryMock{}
		cartExists = false
		want := &domain.Cart{
			UUID: "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8",
			Products: []domain.Product{
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00, PriceFormat: "20.00€"},
			},
		}

		wantCart := &domain.Cart{}
		cartRepo.On("GetCart", want.UUID).Return(wantCart, errors.New("The cart does not exists"))
		cartRepo.On("CreateCart").Return(domain.Cart{UUID: "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"})
		cartRepo.On("UpdateCart", want.UUID, []string{"PEN"}).Return(want)

		wantProduct := domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00, PriceFormat: "20.00€"}
		productRepo := ProductRepositoryMock{}
		productRepo.On("GetProductByCode", "PEN").Return(wantProduct)

		cartSvc := &CartService{CartRepo: &cartRepo, ProductRepo: &productRepo}

		cartExists, err := cartSvc.AddItem(want.UUID, []string{"PEN"})

		if err != nil {
			t.Logf("Cannot add data due: %s", err.Error())
		}

		assert.Equal(t, want.UUID, cartExists.UUID)
		assert.Equal(t, want.Products, cartExists.Products)
	})

	t.Run("Cart service should be able to delete any cart", func(t *testing.T) {
		cartRepo := CartRepositoryMock{}
		cartExists = false
		cartUUID := "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"
		cartRepo.On("DeleteCart", cartUUID).Return(nil)
		wantCart := &domain.Cart{}
		cartRepo.
			On("GetCart", cartUUID).
			Return(wantCart, errors.New("The cart does not exists"))

		cartSvc := &CartService{CartRepo: &cartRepo}
		err := cartSvc.DeleteCart(cartUUID)
		assert.Equal(t, nil, err)
	})

	t.Run("Cart service should show an error when we try to delete a not existent cart file", func(t *testing.T) {
		cartRepo := CartRepositoryMock{}
		fileExists = false
		cartExists = false
		cartUUID := "11E5C5D2-B56B-B588-57F9-8F77A05FEEE8"
		want := errors.New("remove storage/11E5C5D2-B56B-B588-57F9-8F77A05FEEE8: no such file or directory")
		cartRepo.On("DeleteCart", cartUUID).Return(want)
		wantCart := &domain.Cart{}
		cartRepo.
			On("GetCart", cartUUID).
			Return(wantCart, errors.New("The cart does not exists"))

		cartSvc := &CartService{CartRepo: &cartRepo}
		err := cartSvc.DeleteCart(cartUUID)
		assert.Equal(t, want, err)
	})
}
