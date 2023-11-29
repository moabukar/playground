package main

import "fmt"

// Given an integer array nums, find the
// subarray
//  with the largest sum, and return its sum.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Explanation: The subarray [1] has the largest sum 1.
// Example 3:

// Input: nums = [5,4,-1,7,8]
// Output: 23
// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

// Solving with Kadane's Algorithm [more efficient approach and uses O(N) time complexity and O(1) space complexity ]

func maxSubArray(nums []int) int {
	// set to the first element of the array & will hold the largest sum found
	max_so_far := nums[0]
	// also set to the first element - keeps track of the sum of the subarray ending at current position
	max_ending_here := 0

	for i := 1; i < len(nums); i++ {
		// in each iteration, we add current element to max_ending_here
		// if max_ending_here is negative, we set it to the current element
		if max_ending_here+nums[i] > nums[i] {
			max_ending_here += nums[i]
		} else {
			// otherwise
			max_ending_here = nums[i]
		}
		if max_ending_here > max_so_far {
			max_so_far = max_ending_here
		}
	}
	return max_so_far
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5}
	maxSubArray(nums)
	fmt.Println(maxSubArray(nums))
	/// test case 2
	// nums2 := []int{10, 10, 10, 10, 10, 10, 10}
	// maxSubArray(nums2)
	// fmt.Println(maxSubArray(nums2))
}
