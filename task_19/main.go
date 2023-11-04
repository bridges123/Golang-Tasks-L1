package main

import (
	"fmt"
)

// reverseString переворачивает строку, посимвольно меняя начальный и конечные руны
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseString(str))

	str = "amaрыба"
	fmt.Printf("Before: %v\n", str)
	fmt.Printf("After: %v\n", reverseString(str))
}
