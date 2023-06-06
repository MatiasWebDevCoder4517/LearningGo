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

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointer *person) updateFirstName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (pointer *person) updateLastName(newLastName string) {
	(*pointer).lastName = newLastName
}

func (pointer *contactInfo) updateContact(newEmail string, newZipCode int) {
	(*pointer).email = newEmail
	(*pointer).zipCode = newZipCode
}

func main() {
	//alex := person{firstName: "John", lastName: "Rambo"}
	//fmt.Println(alex)

	//var elliot person
	//elliot.firstName = "Elliot"
	//elliot.lastName = "Alderson"
	//fmt.Println(elliot)
	//fmt.Printf("%+v", elliot)

	//thomas := person{
	//	firstName: " Thomas |",
	//	lastName:  " Anderson |",
	//	contact:   contactInfo{email: " imtheone@matrix.com |", zipCode: 2004},
	//}
	//thomas.updateFirstName("Neo")
	//thomas.print()

	thomas := person{
		firstName: " Thomas |",
		lastName:  " Anderson |",
		contact:   contactInfo{email: " mr.anderson@matrix.com |", zipCode: 2004},
	}

	/* change firstname on memory using pointers */
	personPointer := &thomas
	personPointer.updateFirstName(" Neo |")
	thomas.print()

	/* change lastname on memory using pointers */
	personPointer.updateLastName(" I'm the One |")
	thomas.print()

	/* change contact on memory using pointers */
	contactPointer := &thomas
	contactPointer.contact.updateContact("imtheOne@matrix.cl", 1)
	thomas.print()

	

}
