package main

import "log"

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// Example 1:

// Input: s = "anagram", t = "nagaram"

// Output: true

// Example 2:

// Input: s = "rat", t = "car"

// Output: false
func main() {
	log.Println(isAnagram3("racecar", "carrace"))

}

func isAnagram1(s string, t string) bool {
	dict := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, let := range s {
		dict[let]++
	}

	for _, let := range t {
		if _, ok := dict[let]; !ok {
			return false
		}

		dict[let]--
		if dict[let] < 0 {
			return false
		}
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dictS, dictT := make(map[rune]int), make(map[rune]int)

	for i, let := range s {
		dictS[let]++
		dictT[rune(t[i])]++
	}

	for k, v := range dictS {
		if dictT[k] != v {
			return false
		}
	}

	return true
}

func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
