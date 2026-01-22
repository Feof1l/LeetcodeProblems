package main

import "log"

// Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

// Example 1:

// Input: strs = ["act","pots","tops","cat","stop","hat"]

// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

func main() {

	log.Println(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
}

func groupAnagrams(strs []string) [][]string {
	dict := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s-'a']++
		}

		dict[count] = append(dict[count], str)
	}

	var result [][]string
	for _, v := range dict {
		result = append(result, v)
	}

	return result
}
