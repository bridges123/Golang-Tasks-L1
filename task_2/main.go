package main

import (
	"fmt"
	"sync"
)

// Функция, вычисляющая квадрат, с использованием WG
func calcSquare(n int, wg *sync.WaitGroup, res chan<- int) {
	defer wg.Done()
	square := n * n
	res <- square
}

// Функция, оиждающая завершения всех WG, закрывающая канал
func waitSync(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
}

func main() {
	numArray := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}

	// Канал для передачи результатов в main
	results := make(chan int, len(numArray))

	// Итерируемся по массиву чисел и запускаем горутины для вычисления квадратов
	for _, num := range numArray {
		wg.Add(1)
		go calcSquare(num, wg, results)
	}
	go waitSync(wg, results)

	// Чтение результатов из канала и вывод их
	for square := range results {
		fmt.Printf("%d ", square)
	}
}
