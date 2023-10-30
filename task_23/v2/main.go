package main

import (
	"fmt"
)

func main() {
	slice := []string{"A", "B", "C", "D", "E"}
	i := 1

	// Выполняем сдвиг влево на 1
	copy(slice[i:], slice[i+1:])

	// Стираем последний элемент
	slice[len(slice)-1] = ""

	// Обрезаем последний элемент слайса
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
