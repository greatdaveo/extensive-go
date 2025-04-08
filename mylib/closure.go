package mylib

import "fmt"

func Adder() func() int {
	i := 0

	fmt.Println("Previous value of i: ", i)

	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}
