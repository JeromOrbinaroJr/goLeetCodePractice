package main

import "fmt"

func RemoveElement(arr []int, val int) int {
	slow := 0
	for fast := 0; fast < len(arr); fast++ {
		if arr[fast] != val {
			arr[slow] = arr[fast]
			slow++
		}
	}

	return slow
}

func main() {
	arr := []int{3, 2, 2, 3}
	val := 3

	fmt.Println(RemoveElement(arr, val))
}
