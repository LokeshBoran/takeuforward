package main

import "fmt"

func minimumLength(s string) int {
	strMap := make(map[string]int, 0)
	minLength := 0

	for _, c := range s {
		strMap[string(c)]++
	}

	for _, v := range strMap {
		if v%2 != 0 {
			minLength++
		} else {
			minLength += 2
		}
	}
	return minLength
}

func main() {
	fmt.Println(minimumLength("abccde"))
	fmt.Println(minimumLength("abcabcbb"))
	fmt.Println(minimumLength("bbbbb"))
	fmt.Println(minimumLength("pwwkew"))
	fmt.Println(minimumLength("abaacbcbb"))
	fmt.Println(minimumLength("aa"))
}
