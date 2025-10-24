package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n != 0 {
		if n == 1 {
			return true
		}
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}

func main() {
	var N int = 3

	fmt.Println("N =", N, "is Power of Two:", isPowerOfTwo(N))
}
