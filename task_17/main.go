package main

import "fmt"

func main() {
	numbers := []int{1, 4, 7, 15, 25, 99, 101, 1232}
	index := binarySearch(numbers, 101)
	fmt.Printf("number 101 is under index %d\n", index)
}

// binarySearch реализует бинарный поиск числа в массиве с помощью цикла
// если элемент был найден - возвращает его индекс, если не был найден, то -1
func binarySearch(arr []int, x int) int {
	r := -1
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == x {
			r = mid
			break
		} else if arr[mid] < x {
			start = mid + 1
		} else if arr[mid] > x {
			end = mid - 1
		}
	}
	return r
}
