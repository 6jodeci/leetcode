package main

import "strings"

func firstUniqChar(s string) int {
	for ind, letter := range s {
		if strings.Count(s, string(letter)) == 1 {
			return ind
		}
	}
	return -1
}
