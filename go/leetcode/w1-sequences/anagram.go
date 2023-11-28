package main

import "fmt"

// Valig Anagram problem

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false

func isAnagram(s string, t string) bool {
	// condition 1 - if the length of s is not equal to the length of t, defo not an anagram
	if len(s) != len(t) {
		return false
	}

	// condition 2 -

	m := [26]int{} // 26 letters in the alphabet
	// loop through the string and increment the count of each letter
	for i := range s {
		// increment the count of each letter in s
		m[s[i]-'a']++
		// decrement the count of each letter in t
		m[t[i]-'a']--
	}
	// loop through the map and check if the count of each letter is 0
	for i := range m {
		// if the count of any letter is not 0, return false
		if m[i] != 0 {
			return false
		}
	}
	return true

}

func main() {
	s := "car"
	t := "arc"
	fmt.Println(isAnagram(s, t))
}
