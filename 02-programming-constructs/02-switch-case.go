package main

import "fmt"

func main() {
	/*
		no := 21
		switch no % 2 {
		case 0:
			fmt.Printf("%d is an even number\n", no)
		case 1:
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	/*
		switch no := 21; no % 2 {
		case 0:
			fmt.Printf("%d is an even number\n", no)
		case 1:
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	/*
		switch no := 21; {
		case no%2 == 0:
			fmt.Printf("%d is an even number\n", no)
		case no%2 == 1:
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	switch no := 16; {
	case no%3 == 0:
		fmt.Printf("%d is multiples of 3\n", no)
	case no%5 == 0:
		fmt.Printf("%d is multiples of 5\n", no)
	default:
		fmt.Printf("%d is neither multiples of 3 nor multiples of 5\n", no)
	}

	// fall-through (opposite of 'break')
	switch no := 4; no {
	case 0:
		fmt.Println("no == 0")
		fallthrough
	case 1:
		fmt.Println("no <= 1")
		fallthrough
	case 2:
		fmt.Println("no <= 2")
		fallthrough
	case 3:
		fmt.Println("no <= 3")
		fallthrough
	case 4:
		fmt.Println("no <= 4")
		fallthrough
	case 5:
		fmt.Println("no <= 5")
		fallthrough
	case 6:
		fmt.Println("no <= 6")
		// fallthrough
	case 7:
		fmt.Println("no <= 7")

	}

	// fallthrough applied
	fmt.Println("fallthrough applied")
	switch plan := "SUPER"; plan {
	case "SUPREME":
		fmt.Println("[SUPREME] : Offline download allowed!")
		fallthrough
	case "SUPER":
		fmt.Println("[SUPER] : For a family of 3")
		fallthrough
	case "PRO":
		fmt.Println("[PRO] : Create your own playlist")
		fallthrough
	case "FREE":
		fmt.Println("[FREE] : Listen to music!")
	}

	switch x := 3; x {
	case 1, 2, 3, 4:
		fmt.Println("x is <= 4 and >= 1")
	case 5, 6, 7, 8:
		fmt.Println("x is <= 8 and >= 5")
	}

}
