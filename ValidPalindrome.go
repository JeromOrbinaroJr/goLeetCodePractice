package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		rLeft := rune(s[left])
		rRight := rune(s[right])

		if !unicode.IsLetter(rLeft) && !unicode.IsDigit(rLeft) {
			left++
			continue
		}

		if !unicode.IsLetter(rRight) && !unicode.IsDigit(rRight) {
			right--
			continue
		}

		if unicode.ToLower(rLeft) != unicode.ToLower(rRight) {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
}
