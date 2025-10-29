package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	letterCount := [26]int{}

	for _, char := range magazine {
		letterCount[char-'a']++
	}

	for _, char := range ransomNote {
		if letterCount[char-'a'] <= 0 {
			return false
		}
		letterCount[char-'a']--
	}

	return true
}
func main() {
	fmt.Println(canConstruct("aa", "baa"))
}
