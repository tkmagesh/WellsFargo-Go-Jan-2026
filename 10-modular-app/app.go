package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/wellsfargo-go-jan-2026/10-modular-app/calculator"
	/* "github.com/tkmagesh/wellsfargo-go-jan-2026/10-modular-app/calculator/utils" */

	ut "github.com/tkmagesh/wellsfargo-go-jan-2026/10-modular-app/calculator/utils"
)

func run() {
	color.Yellow("application executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))

	fmt.Println(ut.IsEven(21))
}
