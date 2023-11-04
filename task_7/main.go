package main

import (
	"fmt"
	"sync"
)

// MapConcur реализует map, способный выполнять конкурентную запись
type MapConcur struct {
	mu   *sync.Mutex
	wg   *sync.WaitGroup
	data map[int]int
}

// Add конкурентно добавляет в мап ключ-значение
func (m *MapConcur) Add(key int, value int) {
	defer m.wg.Done()
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()

	fmt.Printf("Add new k-v: %d-%d\n", key, value)
}

func main() {
	m := &MapConcur{
		&sync.Mutex{},
		&sync.WaitGroup{},
		map[int]int{},
	}

	// Конкурентная запись в мап с использованием WaitGroup
	for i := 0; i < 20; i++ {
		go func(m *MapConcur, i int) {
			for j := 0; j < 50; j++ {
				m.wg.Add(1)
				go m.Add(i*50+j, (i*50+j)*7%100)
			}
		}(m, i)
	}
	m.wg.Wait()
}
