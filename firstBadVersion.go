package main

import "fmt"

func isBadVersion(n int) bool {
	if n >= 1 {
		return true
	} else {
		return false
	}
}

func firstBadVersion(n int) int {
	if !isBadVersion(n / 2) {
		for i := n / 2; i <= n; i++ {
			if isBadVersion(i) {
				return i
			}
		}
	} else if isBadVersion(n / 2) {
		for i := (n / 2); i >= 0; i-- {
			if !isBadVersion(i) {
				return i + 1
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(firstBadVersion(1))
}
