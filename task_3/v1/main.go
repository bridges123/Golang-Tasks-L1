package main

import (
	"fmt"
	"sync"
)

// Функция, добавляющая квадрат к сумме, используя mutex, чтобы сохранить состояние переменной в потоке
func calcSquareSum(n int, wg *sync.WaitGroup, mx *sync.Mutex, squareSum *int) {
	defer wg.Done()
	mx.Lock()
	*squareSum += n * n
	mx.Unlock()
}

func main() {
	numArray := []int{2, 4, 6, 8, 10}
	squareSum := 0
	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// Итерируемся по массиву чисел и запускаем горутины для вычисления квадратов
	for _, num := range numArray {
		wg.Add(1)
		go calcSquareSum(num, wg, mx, &squareSum)
	}

	// Ожидаем завершения горутин
	wg.Wait()
	fmt.Printf("Sum of squares: %d\n", squareSum)
}
