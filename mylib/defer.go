package mylib

import "fmt"

func Process(i int) {
	defer fmt.Println("Deferred i value: ", i)
	defer fmt.Println("First Deferred statement executed!")
	defer fmt.Println("Second Deferred statement executed!")
	defer fmt.Println("Third Deferred statement executed!")
	i++

	fmt.Println("Normal execution statement")
	fmt.Println("Value of i: ", i)
}
