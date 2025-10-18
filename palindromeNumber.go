package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	var brokenNum []int

	if x < 0 {
		return false
	}

	var copyX int = x
	for copyX > 0 {
		brokenNum = append(brokenNum, copyX%10)
		copyX /= 10
	}

	var reverseX int
	var degree int = len(brokenNum) - 1
	for i := 0; i < len(brokenNum); i++ {
		reverseX += brokenNum[i] * int(math.Pow(10, float64(degree)))
		degree--
	}

	if x == reverseX {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(10))
}
