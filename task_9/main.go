package main

import "fmt"

func pushing(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		// в горутине в выходной канал выбрасываем числа из массива
		for _, num := range numbers {
			out <- num
		}
		// закрываем канал для добавления
		close(out)
	}()
	// возвращаем канал, пока элементы продолжают добавляться
	return out
}

func squaring(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// в горутине в выходной канал выбрасываем квадраты
		for num := range in {
			out <- num * num
		}
		// закрываем канал
		close(out)
	}()
	return out
}

func main() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	squared := squaring(pushing(numbers[:]))
	// считываем из финального канала все числа
	for num := range squared {
		fmt.Printf("%d ", num)
	}
}
