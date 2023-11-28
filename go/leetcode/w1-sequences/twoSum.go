package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// func twoSum(nums []int, target int) []int {
// 	// make a map to store the index of the number (map of number:index)
// 	m := make(map[int]int)
// 	// loop through the nums
// 	for i := 0; i < len(nums); i++ {
// 		// find the difference between target and the current number
// 		diff := target - nums[i]
// 		// if the difference is in the map, return the index of the difference and the current index
// 		if _, ok := m[diff]; ok {
// 			return []int{m[diff], i}
// 		}
// 		// else, add the current number and index to the map
// 		m[nums[i]] = i
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := m[diff]; ok {
			return []int{m[diff], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
	fmt.Println(twoSum(nums, target))
}
