package main

import "fmt"

func main() {
	exec(f1) // f1 is invoked
	exec(f2) // f2 is invoked
	// exec(noExistentFnName)
	exec(func() {
		fmt.Println("anon fn invoked")
	})

}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
