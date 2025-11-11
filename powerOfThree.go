package main

import "fmt"

func isPowerOfThree(n int) bool {
	MAX_POWER_OF_3 := 1162261467
	return n > 0 && MAX_POWER_OF_3%n == 0
}

func main() {
	fmt.Println(isPowerOfThree(-1))
}
