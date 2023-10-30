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

	// контекст c возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// горутина с выходом через контекст (закрытый извне)
	// пока контекст не закрыт - работает с задержкой
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done.")
				return
			default:
				fmt.Println("working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	// занимаем основной поток, затем закрывает контекст
	for i := 1; i < (1 << 30); i++ {
	}
	cancel()

	// ожидаем завершения
	time.Sleep(1 * time.Minute)
}
