package main

import "fmt"

type Set struct {
	data map[string]bool
}

// InitSet - инициализация множества по args значений
func InitSet(values ...string) *Set {
	set := &Set{map[string]bool{}}
	for _, v := range values {
		set.data[v] = true
	}
	return set
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
	set := InitSet("cat", "cat", "dog", "cat", "tree")
	// собственное множество состоит из элементов родительского, но не копирует его, и не пустое
	// в нашем случае, мы создаем собственное множество, как неупорядоченное
	set.Print()
}
