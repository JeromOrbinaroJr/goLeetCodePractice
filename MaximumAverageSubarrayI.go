package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	var sum int

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for r := k; r < len(nums); r++ {
		sum += nums[r]
		sum -= nums[r-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}

	fmt.Println(findMaxAverage(nums, 4))
}
