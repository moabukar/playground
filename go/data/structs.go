package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	// jd := person{firstName: "John", lastName: "Doe"}
	// fmt.Println(jd)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	//////////////////////////

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}
	// fmt.Printf("%+v", jim)
	// jim.print()
	// jimPointer := &jim

	//// easy way to use pointer
	// the below knows that the type is a pointer to a person
	jim.updateName("jimmy")
	// jimPointer.updateLastName("party")
	jim.print()
	// jimPointer.updateName("jimmy")
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) updateLastName(newLastName string) {
	(*p).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
