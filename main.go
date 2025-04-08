package main

import (
	"extensive-go/mylib"
	"fmt"
)

func main() {
	// mylib.Process(10)
	// mylib.PanicProcess(10)

	mylib.RecoverProcess()
	fmt.Println("Returned from process!")
}
