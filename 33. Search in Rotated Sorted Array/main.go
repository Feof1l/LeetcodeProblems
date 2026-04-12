package main

import "log"

func main() {
	log.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 4))
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
