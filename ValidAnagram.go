package main

import "fmt"

func countChars(str string) map[string]int {
	m := make(map[string]int)
	for _, ch := range str {
		m[string(ch)]++
	}
	return m
}

func areMapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for key, val1 := range m1 {
		if val2, ok := m2[key]; !ok || val1 != val2 {
			return false
		}
	}
	return true
}

func isValid(s string, t string) bool {
	return areMapsEqual(countChars(s), countChars(t))
}

func main() {
	fmt.Println(isValid("rat", "car"))
}
