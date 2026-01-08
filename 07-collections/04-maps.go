package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// var productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	productRanks := map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 3,
	}
	fmt.Println(productRanks)

	// Accessing a value
	fmt.Printf("Rank of pencil : %d\n", productRanks["pencil"])

	// adding new items
	productRanks["notebook"] = 4
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	const existingKey = "pen"
	const nonExistingKey = "mouse"

	fmt.Println("Removing an item")
	keyToRemove := existingKey
	// keyToRemove := nonExistingKey
	delete(productRanks, keyToRemove)
	fmt.Println(productRanks)

	// check for the existence of an item
	fmt.Println("check for the existence of an item")
	/*
		fmt.Printf("Rank of %q = %d\n", existingKey, productRanks[existingKey])
		fmt.Printf("Rank of %q = %d\n", nonExistingKey, productRanks[nonExistingKey])
	*/

	// keyToCheck := existingKey
	keyToCheck := nonExistingKey
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key [%q] does not exist!\n", keyToCheck)
	}
}
