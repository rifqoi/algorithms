package main

import (
	"fmt"
)

func binary_search(haystack []int, needle int) (bool, int) {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		mid := (lo + hi) / 2

		v := haystack[mid]

		if v == needle {
			return true, mid
		} else if v > needle {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
		fmt.Println(lo, hi)

	}

	return false, -1
}

func binary_search2(haystack []int, needle int) (bool, int) {

	lo := 0
	hi := len(haystack) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		v := haystack[mid]

		if v == needle {
			return true, mid
		} else if v > needle {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
		fmt.Println(lo, hi)
	}

	return false, -1
}

// func main() {
// 	a := []int{1, 5, 19, 20, 21, 30, 90}
//
// 	fmt.Println("Binary Search 1")
// 	// fmt.Println(binary_search(a, 200))
// 	// fmt.Println(binary_search(a, 30))
// 	// fmt.Println(binary_search(a, 20))
// 	// fmt.Println(binary_search(a, 5))
// 	fmt.Println(binary_search(a, 1))
// 	// fmt.Println(binary_search(a, 90))
// 	fmt.Println("=================================")
// 	fmt.Println()
//
// 	fmt.Println("Binary Search 2")
// 	// fmt.Println(binary_search2(a, 200))
// 	// fmt.Println(binary_search2(a, 30))
// 	// fmt.Println(binary_search2(a, 20))
// 	// fmt.Println(binary_search2(a, 5))
// 	fmt.Println(binary_search2(a, 1))
// 	// fmt.Println(binary_search2(a, 90))
// }
