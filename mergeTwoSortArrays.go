package main

import (
	"fmt"
)

// func mergeTwoSortArrays(nums1 []int, nums2 []int) []int {
// 	var sortedArr []int

// 	for p1, p2 := 0, 0; p1 < len(nums1) && p2 < len(nums2); {
// 		if nums1[p1] > nums2[p2] {
// 			sortedArr = append(sortedArr, nums2[p2])
// 			p2++
// 		} else if nums1[p1] < nums2[p2] {
// 			sortedArr = append(sortedArr, nums1[p1])
// 			p1++
// 		} else {
// 			sortedArr = append(sortedArr, nums1[p1])
// 			p1++
// 			p2++
// 		}
// 	}

// 	return sortedArr
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: nil}

	current := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return dummyHead.Next
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	mergedList := mergeTwoLists(list1, list2)

	for mergedList != nil {
		fmt.Print(mergedList.Val, " ")
		mergedList = mergedList.Next
	}
}
