package main

import "fmt"

func main() {
	a := "v1"
	b := "v2"
	fmt.Printf("a = %v\tb = %v\n", a, b)
	// для переменных любых типов мы можем использовать стандартную конструкцию
	a, b = b, a
	fmt.Printf("a = %v\tb = %v\n", a, b)

	c := 13
	d := 21
	fmt.Printf("a = %v\tb = %v\n", c, d)
	// для чисел мы можем использовать операцию XOR и арифметику
	c = c ^ d
	d = d ^ c
	c = c ^ d
	fmt.Printf("a = %v\tb = %v\n", c, d)
	c = c + d
	d = c - d
	c = c - d
	fmt.Printf("a = %v\tb = %v\n", c, d)
}
