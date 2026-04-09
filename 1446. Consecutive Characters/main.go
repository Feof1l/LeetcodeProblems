package main

import "log"

func main() {
	log.Println(maxPower("leetcode"))
}

func maxPower(s string) int {
	maxPower := 1
	currentPower := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			currentPower++
		} else {
			currentPower = 1
		}
		if currentPower > maxPower {
			maxPower = currentPower
		}
	}

	return maxPower
}
