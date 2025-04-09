package mylib

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math Error: square root of negative number")
	}

	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("error: Empty data")
	}

	return nil
}

func ErrorsFunc() {

	// result, err := sqrt(16)
	// if err != nil {
	// 	println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// fmt.Println(result1)

	// data := []byte{}
	// if err := process(data); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// fmt.Println("Data Processed Successfully")

	// if err1 := errorProcess(); err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// process(data)

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data read successfully")
}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func errorProcess() error {
	return &myError{"Custom error message"}
}

func readData() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("readData: %w", err)
	}

	return nil
}

func readConfig() error {
	return errors.New("config error")
}
