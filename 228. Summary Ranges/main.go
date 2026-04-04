package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	res := []string{}
	current := strconv.Itoa(nums[0])
	currLen := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if currLen > 0 {
				current += "->" + strconv.Itoa(nums[i-1])
			}
			res = append(res, current)
			current = strconv.Itoa(nums[i])
			currLen = 0
		} else {
			currLen++
		}

		if i == len(nums)-1 {
			if currLen > 0 {
				current += "->" + strconv.Itoa(nums[i])
			}
			res = append(res, current)
		}
	}

	return res
}
