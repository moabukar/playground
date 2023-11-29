package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false

func isValid(s string) bool {

	// create a map of runes

	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	// create a stack of runes
	stack := make([]rune, len(s))
	// create a top variable
	top := 0
	// loop through the string
	for _, char := range s {
		// case 1: if the char is an open bracket, push the corresponding closing bracket to the stack
		switch char {
		case '(', '{', '[':
			stack[top] = m[char]
			// increment the top
			top++
		// case 2: if the char is a closing bracket, check if the top of the stack is the same as the char
		case ')', '}', ']':
			// if it is, pop the stack
			if top > 0 && stack[top-1] == char {
				// decrement the top
				top--
			} else {
				// if it isn't, return false
				return false
			}
		}
	}
	// return true if the stack is empty
	return top == 0
}

func main() {
	fmt.Println(isValid("()"))     // should return true
	fmt.Println(isValid("()[]{}")) //should return true
	fmt.Println(isValid("(]"))     // should return false due to mismatched brackets
	fmt.Println(isValid(""))       // true since stack is empty
}
