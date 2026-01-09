package main

import (
	"github.com/tkmagesh/wellfargo-go-jan-2026/13-interfaces/shapes"
	"github.com/tkmagesh/wellfargo-go-jan-2026/13-interfaces/utils"
)

func main() {
	c := shapes.Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	utils.PrintShape(c)

	r := shapes.Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	utils.PrintShape(r)

	/*
		s := shapes.Square{
			rectangle: shapes.Rectangle{
				Height: 12,
				Width:  12,
			},
		}
	*/
	s := shapes.NewSquare(12)
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/
	utils.PrintShape(s)
}
