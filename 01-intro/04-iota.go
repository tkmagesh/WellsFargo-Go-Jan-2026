package main

import "fmt"

func main() {
	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/

	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/

	/*
		const (
			red int = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota
			green
			_ // skipping 2
			blue
		)
	*/

	const (
		red   = (iota * 2) + 1 // (0 * 2) + 1
		green                  // (1 * 2) + 1
		blue                   // (2 * 2) + 1
	)
	fmt.Printf("red = %d, green = %d and blue = %d\n", red, green, blue)

	// Mimicking unix file permissions
	const (
		X = 1 << iota
		W
		R

		XW  = X | W
		WR  = W | R
		XWR = X | W | R
	)
	fmt.Printf("%b %b %b %b %b %b\n", X, W, XW, R, WR, XWR)
}
