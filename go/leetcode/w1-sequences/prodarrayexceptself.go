package main

import "fmt"

// Given an integer array nums, answerurn an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Method 1
// func productExceptSelf(nums []int) []int {
// 	// create a slice answer with the same length as nums
// 	answer := make([]int, len(nums))

// 	//
// 	left := 1
// 	right := 1

// 	for i := 1; i < len(nums); i++ {
// 		if i <= (len(nums)-1)/2 {
// 			answer[i] = 1
// 			answer[len(nums)-i-1] = 1
// 		}
// 		answer[len(nums)-i-1] *= right
// 		answer[i] *= left
// 		left *= nums[i]
// 		right *= nums[len(nums)-i-1]
// 	}
// 	return answer
// }

// Method 2

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	// first element of answer is set to 1.
	// As there are no elements to the left of the first element in nums so the product of elements to its left is 1
	answer[0] = 1
	// loop starts from 2nd element - and loops through array
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	rightProduct := 1
	// loop starts from last element - and loops through array
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))

	nums = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
