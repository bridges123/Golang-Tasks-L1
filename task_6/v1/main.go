package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// канал данных
	ch := make(chan int)
	defer close(ch)

	// контекст с закрытием через N секунд
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// горутина с выходом через контекст (по таймауту извне)
	// пока контекст не закрыт - работает с задержкой
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done.")
				return
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	// ожидаем завершения
	time.Sleep(1 * time.Minute)
}
