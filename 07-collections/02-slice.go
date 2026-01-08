package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println("nos =", nos)

	fmt.Println("Iteration using index")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iteration using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos // a copy of the pointer to the underlying array is created
	nos[0] = 99
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	// reset the nos slice
	nos = []int{3, 1, 4, 2, 5}
	fmt.Println("Before sorting, nos =", nos)
	sortSlice(nos)
	fmt.Println("After sorting, nos =", nos)

	// adding new items
	// reset the nos slice
	nos = []int{3, 1, 4, 2, 5}
	nos = append(nos, 7)
	nos = append(nos, 9, 6, 8)

	hundreds := []int{300, 100, 200}
	nos = append(nos, hundreds...)
	fmt.Println("nos =", nos)

	// slicing
	fmt.Println("nos[2:5] =", nos[2:5])
	fmt.Println("nos[2:] =", nos[2:])
	fmt.Println("nos[:5] =", nos[:5])

}

func sortSlice(list []int) /* return return values */ {
	// bubble sort
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j <= len(list)-1; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
