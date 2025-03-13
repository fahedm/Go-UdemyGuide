package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var fahed person
	fahed.firstName = "Fahed"
	fahed.lastName = "M"
	fmt.Println(fahed)
	fmt.Printf("%+v", fahed)
}