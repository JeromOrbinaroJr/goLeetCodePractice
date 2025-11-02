package main

import "fmt"

func removeDuplicates(nums []int) int {
	count, unique, temp := 1, 0, 1

	for unique < len(nums) && temp < len(nums) {
		if nums[unique] == nums[temp] {
			temp++
		} else {
			nums[unique+1] = nums[temp]
			unique++
			temp++
			count++
		}
	}

	return count
}

func main() {
	var numbers []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(numbers))
}
