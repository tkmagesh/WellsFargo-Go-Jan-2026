package main

import "fmt"

/* product => id, name, cost */

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	// var p1 Product
	// var p1 Product = Product{100, "Pen", 10}
	// var p1 Product = Product{Id: 100, Name: "Pen", Cost: 10}

	// var p1 = Product{Id: 100, Name: "Pen", Cost: 10}

	p1 := Product{Id: 100, Name: "Pen", Cost: 10}
	// fmt.Println(p1)
	// fmt.Printf("%#v\n", p1)
	fmt.Printf("%+v\n", p1)

	fmt.Println("Accessing the attributes")
	fmt.Printf("Id = %d, Name = %q, Cost = %0.02f\n", p1.Id, p1.Name, p1.Cost)

	var p2 Product = p1 // creating a copy (coz structs are values)
	p1.Cost = 20
	fmt.Printf("p1.Cost = %v, p2.Cost = %v\n", p1.Cost, p2.Cost)

	// Create a reference
	var p1Ptr *Product
	p1Ptr = &p1
	fmt.Println("Accessing the attributes through the pointer")
	fmt.Printf("p1Ptr.Id = %d, p1Ptr.Name = %q, p1Ptr.Cost = %0.02f\n", p1Ptr.Id, p1Ptr.Name, p1Ptr.Cost)

	fmt.Println("Before applying discount :", Format(p1))
	ApplyDiscount(&p1, 10)
	fmt.Println("After applying discount :", Format(p1))
}

func Format(p Product) string {
	// return a formatted string of the product
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.02f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) /* no return results */ {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
