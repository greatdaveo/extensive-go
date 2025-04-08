package mylib

import (
	"fmt"
	"os"
)

func ExitProcess() {
	defer fmt.Println("Deferred statement")
	
	fmt.Println("Starting of main func")

	os.Exit(1)

	fmt.Println("End of main func")
}
