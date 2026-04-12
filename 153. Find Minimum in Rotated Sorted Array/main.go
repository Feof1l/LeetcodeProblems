package main

import "log"

func main() {
	log.Println(findMin([]int{1, 2, 3, 4, 5, 6, 7}))
}

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
