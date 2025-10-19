package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	for i := 2; i <= int(x/2); i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}

	return x / 2
}

func main() {
	var x int = 8

	fmt.Println(mySqrt(x))
}
