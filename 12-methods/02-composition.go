package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	// return a formatted string of the product
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %0.02f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* no return results */ {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

// PerishableProduct IS a product
type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

// overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	// logic duplication
	// return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.02f, Expiry = %q", pp.Id, pp.Name, pp.Cost, pp.Expiry)

	// avoid logic duplication
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	milk := NewPerishableProduct(101, "Milk", 50, "2 Days")
	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())
}
