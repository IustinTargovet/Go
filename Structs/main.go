package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Moran",
		contact: contactInfo{
			email: "alex.moran@yahoo.com",
			zip:   1092,
		},
	}
	fmt.Println(alex.lastName)
	fmt.Printf("%+v", alex)
}
