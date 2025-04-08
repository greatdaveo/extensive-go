package mylib

import "fmt"

func RecoverProcess() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
}
