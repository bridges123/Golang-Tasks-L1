package main

import (
	"fmt"
)

func reverseString(str string) string {
	res := ""
	for _, let := range str {
		res = string(let) + res
	}
	return res
}

func main() {
	str := "главрыба"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseString(str))

	str = "amaрыба"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseString(str))
}
