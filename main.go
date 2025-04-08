package main

import (
	"extensive-go/mylib"
	"fmt"
)

func init() {
	fmt.Println("1 Initializing package...")
}

func init() {
	fmt.Println("2 Initializing package...")
}

func init() {
	fmt.Println("3 Initializing package...")
}

func main() {
	// mylib.Process(10)
	// mylib.PanicProcess(10)

	// mylib.RecoverProcess()
	// fmt.Println("Returned from process!")

	// mylib.ExitProcess()

	// fmt.Println("Inside the main func")

	// FOR CLOSURE
	sequence := mylib.Adder()
	sequence2 := mylib.Adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	fmt.Println(sequence2())

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))

}
