package main

import "fmt"

type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	s := MyString("Officia voluptate occaecat deserunt esse ea do tempor amet sit eiusmod veniam irure consectetur. Sint et dolore mollit elit ea. Quis sint deserunt culpa cupidatat incididunt sunt in Lorem do duis aute. Id deserunt in ullamco duis dolore officia mollit eu aliquip sit quis officia in in. In dolor irure proident deserunt.")
	fmt.Println(s.Length())
}
