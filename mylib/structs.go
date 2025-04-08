package mylib

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func StructProcess() {
	p1 := Person{
		firstName: "Dave",
		lastName:  "Olowo",
		age:       20,
		address: Address{
			city:    "London",
			country: "U.K",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "1234567833432",
			cell: "987654327656",
		},
	}

	p2 := Person{
		firstName: "John",
		age:       20,
	}

	p3 := Person{
		firstName: "John",
		age:       20,
	}

	// p2.address.city = "New York"
	// p2.address.country = "U.S.A"

	fmt.Println(p1.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p2.fullName())

	fmt.Println(p1.address)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)
	fmt.Println(p1.address.city)

	fmt.Println("Are p1 and p2 equal: ", p1 == p2)
	fmt.Println("Are p1 and p2 equal: ", p2 == p3)

	// Anonymous Struct
	user := struct {
		username string
		email    string
	}{
		username: "abc1",
		email:    "abc@gmail.com",
	}

	fmt.Println(user.email)

	println("Before Increment: ", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After Increment: ", p1.age)

}
