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
	// &jim will give address to the value
	jimPointer := &jim
	jimPointer.updateName("james")
	jim.print()

	//without using address
	jim.updateName("jimmy")
	jim.print()

	// struct functions will use a copy of struct to do operation is operated without pointer
	// same is not true for slice, without pointers we can update slice.
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v",p)
}