package main

import "log"

// Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

// Each product is guaranteed to fit in a 32-bit integer.

// Follow-up: Could you solve it in
// O
// (
// n
// )
// O(n) time without using the division operation?

// Example 1:

// Input: nums = [1,2,4,6]

// Output: [48,24,12,8]
// Example 2:

// Input: nums = [-1,0,1,2,3]

// Output: [0,-6,0,0,0]

func main() {
	log.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	pref := make([]int, len(nums))
	suf := make([]int, len(nums))
	res := make([]int, len(nums))

	pref[0] = 1
	suf[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		pref[i] = pref[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	for i := range nums {
		res[i] = pref[i] * suf[i]
	}

	return res
}
