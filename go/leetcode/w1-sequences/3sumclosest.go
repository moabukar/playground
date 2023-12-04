package main

import "fmt"

// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

// Return the sum of the three integers.

// You may assume that each input would have exactly one solution.

 

// Example 1:

// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// Example 2:

// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).


func threeSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		if _, ok := seen[num]; ok {
			if i != seen[num] {
				return []int{nums[seen[num]], num, nums[i]}
			}
		}
		seen[target-num] = i
	}

	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 4, -5, 6, 7}
	target := 9
	sums := threeSum(nums, target)
	fmt.Println(sums)
}


// Method 2??
// func main() {
// 	// Define the array of integers and the target sum
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7}
// 	target := 12

// 	// Call the findTriplet function and store the result in a variable
// 	triplet := Triplet(numbers, target)

// 	// Print the result
// 	fmt.Println(triplet) // Output: [2 3 1]
// }

// func Triplet(numbers []int, target int) []int {
// 	seen := make(map[int]int)

// 	for i, num := range numbers {
// 		if _, ok := seen[num]; ok {
// 			if i != seen[num] {
// 				return []int{numbers[seen[num]], num, numbers[i]}
// 			}
// 		}
// 		seen[target-num] = i
// 	}

// 	return []int{}
// }
