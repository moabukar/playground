package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

// func (englishBot) getGreeting() string {
// 	// Custom logic for generating an english greeting
// 	return "Hi there!"
// }

// func (spanishBot) getGreeting() string {
// 	// Custom logic for generating a spanish greeting
// 	return "Hola!"
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) string {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) string {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
