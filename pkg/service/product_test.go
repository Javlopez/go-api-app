package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-lana/pkg/domain"
	"testing"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) All() ([]domain.Product, error) {
	args := m.Called()
	r := []domain.Product{
		domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
		domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
		domain.Product{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.50},
	}
	return r, args.Error(1)
}

func (m *ProductRepositoryMock) GetProductByCode(code string) (domain.Product, error) {
	return domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00}, nil
}

func TestProducts(t *testing.T) {
	t.Run("Get all products", func(t *testing.T) {
		want := []domain.Product{
			domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
			domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
			domain.Product{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.50},
		}
		productRepo := ProductRepositoryMock{}
		productRepo.On("All").Return(want, nil)

		u := ProductService{ProductRepo: &productRepo}

		got, _ := u.GetAll()

		assert.Equal(t, len(got), len(want))
	})

	t.Run("Get product by code", func(t *testing.T) {
		want := domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00}
		productRepo := ProductRepositoryMock{}
		productRepo.On("GetProductByCode").Return(want, nil)

		u := ProductService{ProductRepo: &productRepo}

		got, _ := u.GetProductByCode("TSHIRT")

		assert.Equal(t, got, want)
	})
}
