package service

import (
	"go-lana/pkg/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionsForPENCode(t *testing.T) {
	for _, tc := range []struct {
		Title    string
		Products []domain.Product
		Want     float32
	}{
		{Title: "Should return 0.00 if dont have any product", Want: 0.00},
		{
			Title:    "Should return total = price*product",
			Products: []domain.Product{domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00}},
			Want:     5.00,
		},
		{
			Title: "Should return a discount when the products are 2",
			Products: []domain.Product{
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
			},
			Want: 5.00,
		},
		{
			Title: "Should return a discount when the products are mre than 2",
			Products: []domain.Product{
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
			},
			Want: 15.00,
		},
	} {
		t.Run(tc.Title, func(t *testing.T) {
			p := NewPromotionForPEN(tc.Products)
			got := p.ApplyPromotion()
			assert.Equal(t, tc.Want, got)

		})
	}
}

func TestPromotionsForTshirtCode(t *testing.T) {
	for _, tc := range []struct {
		Title    string
		Products []domain.Product
		Want     float32
	}{
		{Title: "Should return 0.00 if dont have any product", Want: 0.00},
		{
			Title:    "Should return total = price*product",
			Products: []domain.Product{domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00}},
			Want:     20.00,
		},
		{
			Title: "Should return right price if we have 2 items",
			Products: []domain.Product{
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
			},
			Want: 40.00,
		},
		{
			Title: "Should return a discount (25%) when the products are more than 2",
			Products: []domain.Product{
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
			},
			Want: 75.00,
		},
	} {
		t.Run(tc.Title, func(t *testing.T) {
			p := NewPromotionForTshirt(tc.Products)
			got := p.ApplyPromotion()
			assert.Equal(t, tc.Want, got)

		})
	}
}

func TestMultiPromotions(t *testing.T) {
	for _, tc := range []struct {
		Title          string
		PENProducts    []domain.Product
		TshirtProducts []domain.Product
		MugsProducts   []domain.Product
		Want           float32
	}{
		{Title: "Should return 0.00 if dont have any product", Want: 0.00},
		{
			Title:          "Should return total = price*product",
			PENProducts:    []domain.Product{domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00}},
			TshirtProducts: []domain.Product{domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00}},
			Want:           25.00,
		},
		{
			Title: "Should return right price if we have diferent items",
			TshirtProducts: []domain.Product{
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
			},
			PENProducts: []domain.Product{
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
			},
			Want: 45.00,
		},
		{
			Title: "Should return right price if we have diferent items with different promotions",
			TshirtProducts: []domain.Product{
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
				domain.Product{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.00},
			},
			PENProducts: []domain.Product{
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
				domain.Product{Code: "PEN", Name: "Lana Pen", Price: 5.00},
			},
			MugsProducts: []domain.Product{
				domain.Product{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.50},
				domain.Product{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.50},
			},
			Want: 105.00,
		},
	} {
		t.Run(tc.Title, func(t *testing.T) {
			total := NewPromotionForTshirt(tc.TshirtProducts).ApplyPromotion()
			total += NewPromotionForPEN(tc.PENProducts).ApplyPromotion()
			total += NewPromotionForMugs(tc.MugsProducts).ApplyPromotion()
			assert.Equal(t, tc.Want, total)

		})
	}
}
