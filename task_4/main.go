package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// функция воркер, которая считывает данные с канала и выводит их
func worker(id int, ch <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		// получение данных из канала
		select {
		case data := <-ch:
			fmt.Printf("Worker %d received: %d\n", id, data)
		// обработка завершения воркера
		case <-ctx.Done():
			fmt.Printf("Worker %d has been finished\n", id)
			return
		}
	}
}

func main() {
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	wg := &sync.WaitGroup{}
	// Канал для передачи данных в воркеры
	dataCh := make(chan int)

	// контекст с последующим закрытием для воркеров
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// запуск N воркеров-горутин
	for i := 1; i <= N; i++ {
		wg.Add(1)
		go worker(i, dataCh, wg, ctx)
	}

	// создаем канал для отлавливания Interrupt'a через Ctrl C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer close(dataCh)

	// горутина для ожидания сигнала завершения
	go func() {
		<-c
		cancel()
	}()

	// посылаем данные в канал (числа 0-9999)
	for {
		select {
		case <-ctx.Done():
			wg.Wait() // ожидание завершения всех воркеров
			fmt.Println("All workers are done")
			return
		default:
			data := rand.Int() % 10000
			dataCh <- data
		}
	}
}

/*
Выбран способ завершения воркеров через контекст, т.к. это более оптимальный и готовый вариант,
для передачи сигнала об остановке между потоками.
*/
