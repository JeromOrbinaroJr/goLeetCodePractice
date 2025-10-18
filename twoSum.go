package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var ansNums []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ansNums = append(ansNums, i, j)
			}
		}
	}
	return ansNums
}

func main() {
	var nums []int = []int{3, 3}
	var target int = 6

	answer := twoSum(nums, target)
	fmt.Print("[", answer[0], ",", answer[1], "]")
}
