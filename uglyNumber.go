package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}

	var multiples []int
	d := 2
	for d*d <= num {
		if num%d == 0 {
			multiples = append(multiples, d)
			num /= d
		} else {
			d += 1
		}
	}
	if num > 1 {
		multiples = append(multiples, num)
	}

	for _, val := range multiples {
		if (val != 2) && (val != 3) && (val != 5) {
			return false
		}
	}
	return true
}

func main() {
	var num int = 22

	fmt.Println(isUgly(num))
}
