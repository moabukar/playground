package main

import "fmt"

// Given an integer array nums,
// return true if any value appears at least twice in the array,
// and return false if every element is distinct.

func containsDuplicate(nums []int) bool {
	// make a map to store the number and a boolean
	m := make(map[int]bool)
	for i := range nums {
		// if the number is in the map, return true
		if _, ok := m[nums[i]]; ok {
			return true
		}
		// else, add the number to the map
		m[nums[i]] = true
	}
	// else, return false
	return false
}

func main() {
	nums := []int{7, 2, 3, 7}
	containsDuplicate(nums)
	fmt.Println(containsDuplicate(nums))
}
