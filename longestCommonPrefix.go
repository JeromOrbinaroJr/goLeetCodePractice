package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	var minLen int = math.MaxInt

	for _, word := range strs {
		if len(word) < minLen {
			minLen = len(word)
		}
	}

	var maxCommonPrefix string

	for i := 0; i < minLen; i++ { // для индексов в слове
		var tempWord int
		for j := 1; j < len(strs); j++ { // сравнение каждого слова
			tempWord = j
			if (strs[j-1])[i] != (strs[j])[i] {
				return maxCommonPrefix
			}
		}
		maxCommonPrefix += string((strs[tempWord])[i])
	}
	return maxCommonPrefix
}

func main() {
	var strArr []string = []string{"flower", "flower", "flower", "flower"}
	fmt.Println(longestCommonPrefix(strArr))
}
