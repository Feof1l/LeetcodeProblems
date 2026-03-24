package main

import "log"

func main() {
	log.Println(singleNumber([]int{4, 2, 1, 2, 1}))
}

func singleNumber(nums []int) int {
	res := 0

	for _, elem := range nums {
		res ^= elem
	}

	return res
}
