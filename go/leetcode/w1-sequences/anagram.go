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
	// Condition 1: if the length of both strings, then defo not an anagram
	if len(s) != len(t) {
		return false
	}
	// init an empty array of 26 elements
	m := [26]int{}

	// // loop through the string and increment the value of the index
	// for i := range s {
	// 	m[s[i]-'a']++
	// 	m[t[i]-'a']--
	// }

	// this is the EXACT same as the above code but reversed (the above might be better tbh)
	for i := range t {
		m[t[i]-'a']--
		m[s[i]-'a']++
	}
	// loop through the array and check if the value is 0
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

// Understanding

// In summary, this function works on the principle that if two strings are anagrams,
// they must contain the same characters in the same quantities,
// regardless of order.
// The array m tracks the frequency of each character in s and t,
// and the function verifies if these frequencies are equal,
// thereby determining if s and t are anagrams.
