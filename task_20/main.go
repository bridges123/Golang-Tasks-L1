package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	res := ""
	for _, sub := range strings.Split(str, " ") {
		res = sub + " " + res
	}
	return res
}

func main() {
	str := "snow dog sun"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseWords(str))

	str = "i walk in the park"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseWords(str))
}
