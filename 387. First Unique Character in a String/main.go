package main

import "log"

func main() {
	log.Println(firstUniqChar("leetcode"))
}

func firstUniqChar(s string) int {
	dict := make(map[rune]bool)

	for _, value := range s {
		if _, ok := dict[value]; !ok {
			dict[value] = false
		} else {
			dict[value] = true
		}
	}

	for i := range s {
		if !dict[rune(s[i])] {
			return i
		}
	}

	return -1
}

func firstUniqChar2(s string) int {
	dict := make([]int, 26)

	for i := range s {
		dict[s[i]-'a']++
	}

	for i := range s {
		if dict[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
