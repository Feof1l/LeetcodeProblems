package main

import "log"

// Given an integer array nums and an integer k, return the k most frequent elements within the array.

// The test cases are generated such that the answer is always unique.

// You may return the output in any order.

// Example 1:

// Input: nums = [1,2,2,3,3,3], k = 2

// Output: [2,3]
// Example 2:

// Input: nums = [7,7], k = 1

// Output: [7]

func main() {
	log.Println(topKFrequent([]int{1, 1, 1, 1, 2, 3, 3, 3, 3}, 1))
}

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, value := range nums {
		dict[value]++
	}

	for key, value := range dict {
		freq[value] = append(freq[value], key)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, value := range freq[i] {
			res = append(res, value)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
