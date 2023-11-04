package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := map[int]int{}
	m[0] = 1
	m[1] = 124
	m[2] = 281
	for k, v := range m {
		fmt.Printf("k: %v\tv: %v\n", k, v)
	}

	fmt.Println(m)

	numbers := [9]int{15, 21, 8, 1, 7232, 77, -12, 23, 0}
	sortedNumbers := quickSort(numbers[:])
	fmt.Printf("Sorted: %v\n", sortedNumbers)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// выбираем случайный индекс опорного элемента, чтобы получить среднее время
	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	// ставим текущий элемент слева от стены, если он меньше опорного
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	// вызываем для элементов до левого и после
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}
