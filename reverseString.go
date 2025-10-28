package main

import "fmt"

func reverseString(s []byte) []byte {
	left, rigth := 0, len(s)-1

	for left < rigth {
		s[left], s[rigth] = s[rigth], s[left]
		left++
		rigth--
	}
	return s
}

func main() {
	var s []byte = []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(reverseString(s))
}
