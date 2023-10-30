package main

import (
	"fmt"
)

func main() {
	slice := []string{"A", "B", "C", "D", "E"}
	i := 1

	// Копируем последний элемент в индекс i
	slice[i] = slice[len(slice)-1]

	// Стираем последний элемент
	slice[len(slice)-1] = ""

	// Обрезаем последний элемент слайса
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
