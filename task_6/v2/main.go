package main

import (
	"fmt"
	"time"
)

func main() {
	// канал данных
	ch := make(chan int)
	defer close(ch)

	// горутина, выход из которой обеспечивает закрытие канала извне
	go func(ch <-chan int) {
		for {
			val, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(val)
		}
		fmt.Println("done.")
	}(ch)

	// пускаем в канал значения и закрываем
	for i := 1; i < 1000; i += 150 {
		ch <- i
	}
	close(ch)

	// ожидаем завершения
	time.Sleep(1 * time.Minute)
}
