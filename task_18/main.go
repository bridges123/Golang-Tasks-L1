package main

import (
	"fmt"
	"sync"
)

// Counter описывает счетчик, который содержит mutex для конкурентного увеличения
type Counter struct {
	counter int64
	mu      *sync.Mutex
}

// Increment увеличивает счетчик с использованием Mutex для блокировки потока
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counter += 1
}

// NewCounter возвращает ссылку на новый объект счетчика
func NewCounter() *Counter {
	return &Counter{counter: 0, mu: &sync.Mutex{}}
}

func main() {
	counter := NewCounter()
	wg := &sync.WaitGroup{}
	// выполняем счет в 10 потоках
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c *Counter) {
			// в каждом потоке увеличиваем на 100
			defer wg.Done()
			for i := 0; i < 100; i++ {
				c.Increment()
			}
		}(counter)
	}

	// ждем выхода из всех горутин
	wg.Wait()
	fmt.Printf("Counter result: %d\n", counter.counter)
}
