package main

import (
	"fmt"
	"unicode"
)

// uniqueSymbols проходится по строке и добавляет символы в мап(сет),
// если символ уже есть - возвращает false. Если все символы уникальные - true
func uniqueSymbols(str string) bool {
	symbols := map[rune]bool{}
	for _, c := range str {
		c = unicode.ToLower(c)
		if _, exists := symbols[c]; exists {
			return false
		}
		symbols[c] = true
	}
	return true
}

func main() {
	str := "snow1dog2sun"
	fmt.Printf("%s - %t\n", str, uniqueSymbols(str))

	str = "my,Mine"
	fmt.Printf("%s - %t\n", str, uniqueSymbols(str))

	str = "I1get2of3hard"
	fmt.Printf("%s - %t\n", str, uniqueSymbols(str))
}
