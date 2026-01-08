package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iteration using index")
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iteration using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos // a copy of nos will be created
	nos[0] = 99
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	// creating a reference
	var nosPtr *[5]int
	nosPtr = &nos

	// deferencing
	fmt.Println("(*nosPtr)[0] =", (*nosPtr)[0])

	// members of a pointer to a container can be accessed directly without explicit dereferencing
	fmt.Println("nosPtr[0] =", nosPtr[0])

	// reset the nos array
	nos = [5]int{3, 1, 4, 2, 5}
	fmt.Println("Before sorting, nos =", nos)
	fmt.Printf("[main] &nos = %p\n", &nos)
	sortArray(&nos)
	fmt.Println("After sorting, nos =", nos)
}

func sortArray(list *[5]int) /* return return values */ {
	// bubble sort
	fmt.Printf("[sort] &list = %p\n", list)
	for i := 0; i < 4; i++ {
		for j := i + 1; j <= 4; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
