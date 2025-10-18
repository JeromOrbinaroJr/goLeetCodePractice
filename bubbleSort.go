package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-i-1; i++ {
		if arr[i] > arr[i+1] {
			tempVal := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = tempVal
		}
	}
	return arr
}

func main() {
	var arr []int = []int{2, 1, 5}
	fmt.Println(bubbleSort(arr))
}
