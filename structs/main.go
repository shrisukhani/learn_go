package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) println() {
	fmt.Printf("%+v\n", p)
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Gatsby",
		contact: contactInfo{
			email:   "jimgatsby@gmail.com",
			zipCode: 94041,
		},
	}
	jim.println()
	jim.updateName("Jimmy")
	jim.println()
}
