package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	dict := make(map[rune]string)
	var patternDict []rune = []rune(pattern)
	var sDict []string = strings.Split(s, " ")

	if len(patternDict) != len(sDict) {
		return false
	}

	for i := 0; i < len(patternDict); i++ {
		dict[patternDict[i]] = sDict[i]
	}

	for i := 0; i < len(patternDict); i++ {
		if dict[patternDict[i]] != sDict[i] {
			return false
		}
	}

	seenWords := make(map[string]bool)
	for _, word := range dict {
		if seenWords[word] {
			return false
		}
		seenWords[word] = true
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "wf"))
}
