package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// функция воркер, которая считывает данные с канала и выводит их
func worker(ch <-chan int, ctx context.Context) {
	for {
		// получение данных из канала
		select {
		case data := <-ch:
			fmt.Printf("Worker received: %d\n", data)
		// обработка завершения воркера
		case <-ctx.Done():
			fmt.Println("Worker has been finished")
			return
		}
	}
}

func main() {
	var N time.Duration
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	dataCh := make(chan int)
	defer close(dataCh)

	// контекст с закрытием через N секунд
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	// запуск горутины для считывания канала
	go worker(dataCh, ctx)

	// посылаем данные в канал (числа 0-9999)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			data := rand.Int() % 10000
			dataCh <- data
			time.Sleep(time.Millisecond) // для анти-перегрузки канала
		}
	}
}
