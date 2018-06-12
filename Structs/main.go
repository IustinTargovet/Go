package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Moran",
		contactInfo: contactInfo{
			email: "alex.moran@yahoo.com",
			zip:   102,
		},
	}
	alex.print()
	alexPointer := &alex
	alexPointer.updateName("Alecu")
	alex.print()
}
