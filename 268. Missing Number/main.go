package main

import "log"

func main() {
	log.Println(missingNumber([]int{3, 0, 1}))
}

// func missingNumber(nums []int) int {
// 	min := 0
// 	sum := 0

// 	for _, elem := range nums {
// 		if elem < min {
// 			min = elem
// 		}

// 		sum += elem
// 	}

// 	res := (2*min + len(nums)) * (len(nums) + 1) / 2

// 	return res - sum
// }

func missingNumber(nums []int) int {
	sum := 0
	for _, elem := range nums {
		sum += elem
	}

	an := len(nums)

	return (1+an)*an/2 - sum
}
