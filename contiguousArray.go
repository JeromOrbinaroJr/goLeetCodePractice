package main

import "fmt"

func findMaxLength(nums []int) int {
	sumIndex := make(map[int]int)
	sumIndex[0] = -1
	maxLen := 0
	prefixSum := 0

	for i, num := range nums {
		if num == 0 {
			prefixSum -= 1
		} else {
			prefixSum += 1
		}

		if firstIndex, exists := sumIndex[prefixSum]; exists {
			length := i - firstIndex
			if length > maxLen {
				maxLen = length
			}
		} else {
			sumIndex[prefixSum] = i
		}
	}

	return maxLen
}

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(findMaxLength(arr))
}
