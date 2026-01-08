package main

import (
	"errors"
	"fmt"
)

var ErrF1 = errors.New("unexpected error in f1")
var ErrF2 = errors.New("unexpected error in f1")

func main() {
	err := f1(false)
	if errors.Is(err, ErrF2) {
		fmt.Println("Something went wrong in f2")
	}
	if errors.Is(err, ErrF1) {
		fmt.Println("Something went wrong in f1")
	}
	if err == nil {
		fmt.Println("All good!")
	}
}

func f1(lucky bool) error {
	if f2_err := f2(lucky); f2_err != nil {
		return fmt.Errorf("%w, %w", ErrF1, f2_err)
	}
	return nil
}

func f2(lucky bool) error {
	if lucky {
		return nil
	}
	return ErrF2
}
