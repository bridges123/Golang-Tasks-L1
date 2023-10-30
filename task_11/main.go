package main

import "fmt"

type Set struct {
	data map[int]bool
}

// InitSet - инициализация множества по args значений
func InitSet(values ...int) *Set {
	set := &Set{map[int]bool{}}
	// создаем мап (реализацию сета) и записываем значения в ключи с true value
	for _, v := range values {
		set.data[v] = true
	}
	return set
}

// Intersect - функция пересечения множеств, возвращает новое
func (s *Set) Intersect(s2 *Set) *Set {
	res := InitSet()
	for v := range s.data {
		// если значение из 1 сущ. во 2, добавляем в res
		if _, exists := s2.data[v]; exists {
			res.data[v] = true
		}
	}
	return res
}

// Print - вывод значений множества
func (s *Set) Print() {
	fmt.Print("values: ")
	for v := range s.data {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func main() {
	s1 := InitSet(1, 2, 3, 4, 5, 6)
	s2 := InitSet(4, 5, 6, 7, 8, 9)
	res := s1.Intersect(s2)
	res.Print()
}
