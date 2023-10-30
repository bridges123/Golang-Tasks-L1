package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// создаем мап со значением - слайсом из флоатов
	m := map[int][]float32{}
	for _, temp := range temps {
		// задаем ключ - инт с отрезанной единичной частью
		key := int(temp/10) * 10
		// если не сущ слайса по этому ключу - инит
		_, exists := m[key]
		if !exists {
			m[key] = []float32{}
		}
		// добавляем темп по этому ключу
		m[key] = append(m[key], temp)
	}
	fmt.Println(m)
}
