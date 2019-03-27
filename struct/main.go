package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func (p *person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) setName(firstName string) {
	p.firstName = firstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contact{
			email:   "jim.halpert@gmail.com",
			zipCode: 94000,
		},
	}

	jim.print()
	jim.setName("Hello")
	jim.print()
}
