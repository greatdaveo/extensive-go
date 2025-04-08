package mylib

import "fmt"

func PanicProcess(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")

		// fmt.Println("Before Panic")
		// defer fmt.Println("Deferred 3")
	}

	fmt.Println("Processing input: ", input)

}
