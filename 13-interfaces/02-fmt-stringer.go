/*
Create a type to represent a ShoppingCart
	- should be able to add products to the cart
	- each product added to the cart must have "quantity" as well
	- should be able to apply discounts to the products in the cart
	- should be able to list the products & their item values
	- should be able to print the overall cart value
*/

package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	// return a formatted string of the product
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.02f", p.Id, p.Name, p.Cost)
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return p.Format()
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* no return results */ {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

type LineItem struct {
	Item     *Product
	Count    int
	Discount float64
}

func NewLineItem(product *Product, count int) *LineItem {
	return &LineItem{
		Item:  product,
		Count: count,
	}
}

func (li *LineItem) ItemValue() float64 {
	x := li.Item.Cost * float64(li.Count)
	return x * ((100 - li.Discount) / 100)
}

func (li LineItem) Format() string {
	// return fmt.Sprintf("%s, Count = %d, Item Value = %0.02f", li.Item.Format(), li.Count, li.ItemValue())
	return fmt.Sprintf("%s, Count = %d, Item Value = %0.02f", li.Item, li.Count, li.ItemValue())
}

// fmt.Stringer interface implementation
func (li LineItem) String() string {
	return li.Format()
}

/* ShoppingCart */

type ShoppingCart struct {
	LineItems []*LineItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		LineItems: make([]*LineItem, 0, 10),
	}
}

func (sc *ShoppingCart) AddItem(p *Product, units int) {
	newLineItem := NewLineItem(p, units)
	sc.LineItems = append(sc.LineItems, newLineItem)
}

func (sc *ShoppingCart) ApplyDiscount(discountPercentage float64) {
	for _, li := range sc.LineItems {
		li.Discount = discountPercentage
	}
}

func (sc *ShoppingCart) CartValue() float64 {
	var cartValue float64
	for _, li := range sc.LineItems {
		cartValue += li.ItemValue()
	}
	return cartValue
}

func (sc ShoppingCart) Format() string {
	var sb strings.Builder
	sb.WriteString("===================================================\n")
	sb.WriteString("Shopping Cart\n")
	sb.WriteString("===================================================\n")
	for _, li := range sc.LineItems {
		// sb.WriteString(li.Format())
		sb.WriteString(fmt.Sprintf("%s", li))
		sb.WriteString("\n")
	}
	sb.WriteString("--------------------------------------------------\n")
	sb.WriteString(fmt.Sprintf("Cart Value : %0.02f\n", sc.CartValue()))
	sb.WriteString("--------------------------------------------------\n")
	return sb.String()
}

// fmt.Stringer interface implementation
func (sc ShoppingCart) String() string {
	return sc.Format()
}

var productMaster = map[int]*Product{
	101: NewProduct(101, "Pen", 10),
	102: NewProduct(102, "Pencil", 5),
	103: NewProduct(103, "Marker", 50),
	104: NewProduct(104, "Scribble Pad", 30),
}

func main() {
	sc := NewShoppingCart()
	sc.AddItem(productMaster[101], 10)
	sc.AddItem(productMaster[102], 5)
	sc.AddItem(productMaster[103], 6)
	sc.ApplyDiscount(10)
	// fmt.Println(sc.Format())
	fmt.Println(sc)
}
