package main

import "log"

// 217 https://leetcode.com/problems/contains-duplicate/description/
// Given an integer array nums, return true if any value appears
// at least twice in the array, and return false if every element is distinct.
// Input: nums = [1,2,3,1]

// Output: true
func main() {
	log.Println(hasDuplicate([]int{}))
}

func hasDuplicate(nums []int) bool {
	dict := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return true
		}

		dict[num] = true
	}

	return false
}
