package main

import (
	"fmt"
	"sync"
)

type MapConcur struct {
	mu   *sync.Mutex
	wg   *sync.WaitGroup
	data map[int]int
}

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

	for i := 0; i < 1000; i++ {
		m.wg.Add(1)
		go m.Add(i, i*7%100)
	}
	m.wg.Wait()
}
