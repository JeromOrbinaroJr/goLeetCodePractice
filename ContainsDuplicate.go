package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, len(nums))

	for _, elem := range nums {
		if _, exists := set[elem]; exists {
			return true
		}
		set[elem] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
