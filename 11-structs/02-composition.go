package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
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

func main() {
	// milk := PerishableProduct{Product{100, "Milk", 50}, "2 Days"}
	milk := PerishableProduct{
		Product: Product{100, "Milk", 50},
		Expiry:  "2 Days",
	}

	fmt.Println(milk.Expiry)
	/*
		fmt.Println(milk.Product.Cost)
		fmt.Println(milk.Product.Name)
		fmt.Println(milk.Product.Id)
	*/
	fmt.Println(milk.Id)
	fmt.Println(milk.Name)
	fmt.Println(milk.Cost)

	bread := NewPerishableProduct(101, "Bread", 40, "7 Days")

}
