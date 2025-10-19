package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	var lenSeach = len(haystack) - len(needle)
	if lenSeach == 0 {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}
	for i := 0; i < lenSeach+1; i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

func main() {
	var haystack string = "a"
	var needle string = "a"

	fmt.Println(strStr(haystack, needle))
}
