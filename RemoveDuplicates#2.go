package main

import "fmt"

func removeDuplicates(nums []int) int {
	countDupl := 0
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if fast == 0 || nums[fast] != nums[fast-1] {
			countDupl = 1
		} else {
			countDupl++
		}

		if countDupl <= 2 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(removeDuplicates(nums))
}
