package inventory

import "errors"

type Product struct {
	Name       string
	Weight     float64
	MinStock   float64
	PricePerKg float64
}

func (p Product) TotalPrice() float64 {
	return p.Weight * p.PricePerKg
}
func (p Product) NeedsRestock() bool {
	return p.Weight < p.MinStock
}
func (p *Product) Restock(kg float64) {
	p.Weight += kg
}

func (p *Product) Use(kg float64) error {
	if kg > p.Weight {
		return errors.New("Not enough stock")
	}
	p.Weight -= kg
	return nil
}
