package mylib

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person // Embedded struct
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old. \n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f. \n", e.name, e.empId, e.salary)
}

func StructEmbedding() {
	emp := Employee{
		person: person{
			name: "John Doe",
			age:  30,
		},

		empId:  "E001",
		salary: 5000,
	}

	fmt.Println("Name: ", emp.name)
	fmt.Println("Age: ", emp.age)
	fmt.Println("Emp ID: ", emp.empId)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
}
