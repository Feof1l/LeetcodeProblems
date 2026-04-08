package main

import (
	"log"
)

func main() {
	log.Println(maxDistToClosest([]int{0, 1, 1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	currLen := 0
	maxLen := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			currLen = 0
		} else {
			currLen++
			if (currLen+1)/2 > maxLen {
				maxLen = (currLen + 1) / 2
			}
		}
	}

	flag := true
	for i := 0; i < len(seats) && flag; i++ {
		if seats[i] == 1 {
			if i > maxLen {
				maxLen = i
			}
			flag = false
		}
	}

	flag = true
	for i := len(seats) - 1; i >= 0 && flag; i-- {
		if seats[i] == 1 {
			if len(seats)-1-i > maxLen {
				maxLen = len(seats) - 1 - i
			}
			flag = false
		}
	}

	return maxLen
}
