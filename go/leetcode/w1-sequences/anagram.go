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
	for i := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for i := range m {
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
