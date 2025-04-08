package main

import (
	"extensive-go/mylib"
	"fmt"
)

func main() {
	fmt.Println(mylib.Factorial(5))
	fmt.Println(mylib.Factorial(10))

	fmt.Println(mylib.SumOfDigits(9))
	fmt.Println(mylib.SumOfDigits(12))
	fmt.Println(mylib.SumOfDigits(12345))
}
