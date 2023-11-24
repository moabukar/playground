package main

import "fmt"

// create a custom type of 'book'
type book []string

// and any instance of the class book can call the printTitle method
func (b book) printTitle() {
	fmt.Println(b)
}

// main entry point
// create new books using the receiver method
func main() {
	b := book{"The Lord of the Rings", "The Hobbit", "The Silmarillion", "Harry Potter"}
	b.printTitle()
}
