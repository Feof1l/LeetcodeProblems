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
		buf := [26]int{}
		for _, elem := range str {
			buf[elem-'a']++
		}

		dict[buf] = append(dict[buf], str)
	}

	res := [][]string{}

	for _, value := range dict {
		res = append(res, value)
	}

	return res
}
