package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// mark := person{firstName: "Mark", lastName: "Superfly"}

	var sue person
	sue.firstName = "Sue"
	sue.lastName = "Simpson"
	sue.contact.email = "sueSimpson@gmail.com"
	sue.contact.zipCode = 37112

	//fmt.Println(alex, mark, sue) // {Alex Anderson}

	//fmt.Printf("%+v \n", alex) // {firstName: Alex lastName: Anderson}

	jim := person{firstName: "Jim", lastName: "Party", contact: contactInfo{email: "jim@gmail.com", zipCode: 37013}}
	//fmt.Println(sue, jim)

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
	//sue.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
