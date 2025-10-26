package main

import "fmt"

func addDigits(num int) int {
	var sum int = 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	if sum >= 10 {
		sum = addDigits(sum)
	}
	return sum
}

func main() {
	var num int = 38
	fmt.Println(addDigits(num))
}
