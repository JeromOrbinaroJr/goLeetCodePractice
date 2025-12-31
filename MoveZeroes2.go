package main

import "fmt"

func moveZeroes(nums []int) []int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}

	fmt.Println(moveZeroes(nums))
}
