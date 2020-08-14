package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	str = strings.ReplaceAll(str, ", ", " ")
	str = strings.ReplaceAll(str, ".", "")
	wordList := strings.Fields(str)

	counts := make(map[string]int)

	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}
