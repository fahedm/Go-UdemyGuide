package main

import "fmt"

type contactInfo struct {
	email string
	zipCode string
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@test.com",
			zipCode: "10001",
		},
	}
	fmt.Printf("%+v", jim)
}