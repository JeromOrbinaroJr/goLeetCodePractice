package main

import "fmt"

func hammingWeight(n int) int {
	var count int

	for n != 0 {
		if n%2 != 0 {
			count++
		}
		n = n / 2
	}

	return count
}

func main() {
	var num int = 2147483645

	fmt.Println(hammingWeight(num))
}
