package main

import "fmt"

func moveZeroes(nums []int) []int {
	var nonZeroIndex int = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func main() {
	var nums []int = []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}
