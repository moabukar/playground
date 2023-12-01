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

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	left := 1
	right := 1

	for i := 0; i < len(nums); i++ {
		if i <= (len(nums)-1)/2 {
			answer[i] = 1
			answer[len(nums)-i-1] = 1
		}
		answer[len(nums)-i-1] *= right
		answer[i] *= left
		left *= nums[i]
		right *= nums[len(nums)-i-1]
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))

	nums = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
