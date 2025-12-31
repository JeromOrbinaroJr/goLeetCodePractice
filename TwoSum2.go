package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{-1, 0}
	target := -1

	fmt.Println(twoSum(nums, target))
}
