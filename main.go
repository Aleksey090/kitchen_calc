package main

import (
	"errors"
	"fmt"
)

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

func main() {

	products := []Product{
		{Name: "Salmon", Weight: 15.5, MinStock: 5.0, PricePerKg: 1800.0},
		{Name: "Avocado", Weight: 4.0, MinStock: 6.0, PricePerKg: 450.0},
		{Name: "Rice", Weight: 25.0, MinStock: 10.0, PricePerKg: 120.0},
		{Name: "Cheese", Weight: 6.5, MinStock: 7.0, PricePerKg: 950.0},
	}

	var totalWeight float64
	var totalCost float64

	err := products[0].Use(10.0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, p := range products {
		if p.NeedsRestock() {
			fmt.Println("Need to buy:", p.Name)
		}
		totalWeight += p.Weight
		totalCost += p.TotalPrice()
	}

	fmt.Println("Total product weight:", totalWeight)
	fmt.Println("Total product price:", totalCost)
}
