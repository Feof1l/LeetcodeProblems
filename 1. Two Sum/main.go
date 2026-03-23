package main

import "log"

// Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

// You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

// Return the answer with the smaller index first.

// Example 1:

// Input:
// nums = [3,4,5,6], target = 7

// Output: [0,1]

func main() {
	log.Println(twoSum([]int{4, 5, 6}, 10))
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)

	for i, elem := range nums {
		diff := target - elem
		if _, ok := dict[diff]; ok {
			return []int{i, dict[diff]}
		}

		dict[elem] = i
	}

	return []int{}
}
