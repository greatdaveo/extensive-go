package mylib

import "fmt"

func Pointers() {

	var ptr *int

	var a int = 10

	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
