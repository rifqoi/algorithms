package main

import "fmt"

func linear_search(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			fmt.Printf("index of %d\n", i)
			return true
		}
	}
	return false
}
