package main

import "fmt"

func searchInsert(nums []int, target int) int {
	var currentIndex int
	for i, num := range nums {
		if num >= target {
			currentIndex = i
			break
		} else if i+1 == len(nums) {
			currentIndex = i + 1
			break
		}
	}
	return currentIndex
}

func main() {
	var nums []int = []int{1, 3, 5, 6}
	var target int = 7

	fmt.Println(searchInsert(nums, target))
}
