package service

import (
	"go-lana/pkg/domain"
	"math"
)

const (
	DISCOUNT  = 0.25
	MIN_ITEMS = 3
)

type Promotional interface {
	ApplyPromotion() float32
}

type Promotion struct {
	Items int
	Price float32
}

type PromotionForPEN struct {
	Promotion
}

func (p *PromotionForPEN) ApplyPromotion() float32 {
	if p.Items == 0 {
		return 0.00
	}
	q := math.Ceil(float64(p.Items) / 2)
	return float32(q) * p.Price
}

type PromotionForTshirt struct {
	Promotion
}

func (p *PromotionForTshirt) ApplyPromotion() float32 {
	if p.Items == 0 {
		return 0.00
	}

	totalPrice := float32(p.Items) * p.Price

	if p.Items >= MIN_ITEMS {
		totalPrice = float32(p.Items) * (p.Price - (p.Price * DISCOUNT))
	}
	return float32(totalPrice)
}

type PromotionForMugs struct {
	Promotion
}

func (p *PromotionForMugs) ApplyPromotion() float32 {
	if p.Items == 0 {
		return 0.00
	}

	totalPrice := float32(p.Items) * p.Price

	return float32(totalPrice)
}

func newPromo(products []domain.Product) Promotion {
	p := Promotion{
		Items: len(products),
	}

	if len(products) > 0 {
		p.Price = products[0].Price
	}

	return p
}

func NewPromotionForPEN(products []domain.Product) Promotional {
	p := newPromo(products)
	return &PromotionForPEN{p}
}

func NewPromotionForTshirt(products []domain.Product) Promotional {
	p := newPromo(products)
	return &PromotionForTshirt{p}
}

func NewPromotionForMugs(products []domain.Product) Promotional {
	p := newPromo(products)
	return &PromotionForMugs{p}
}
