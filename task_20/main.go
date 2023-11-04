package main

import (
	"fmt"
	"strings"
)

// reverseWords переворачивает слова в строке, проходясь по слайсу слов и добавляя их в билдер
func reverseWords(str string) string {
	split := strings.Split(str, " ")
	var res strings.Builder
	res.Grow(len(str))
	for i := len(split) - 1; i >= 0; i-- {
		res.WriteString(split[i])
		res.WriteString(" ")
	}
	return res.String()
}

func main() {
	str := "snow dog sun"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseWords(str))

	str = "i walk in the park"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseWords(str))
}
