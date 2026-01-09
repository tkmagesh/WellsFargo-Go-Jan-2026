package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	// return a formatted string of the product
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.02f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* no return results */ {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {

	/* Format() and ApplyDiscount() as plain simple "functions" */
	/*
		fmt.Println("Before applying discount :", Format(p1))
		ApplyDiscount(&p1, 10)
		fmt.Println("After applying discount :", Format(p1))
	*/

	/* Using Format() and ApplyDiscount() as methods */
	/*
		var p1 Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	var p1 *Product = &Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	fmt.Println("Before applying discount :", p1.Format())
	p1.ApplyDiscount(10)
	fmt.Println("After applying discount :", p1.Format())
}
