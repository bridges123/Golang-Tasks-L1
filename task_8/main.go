package main

import "fmt"

// Функция устанавливает в index бит числа значение value
func setBit(num *int64, index int64, value int64) {
	mask := int64(1) << index
	// обнуляем i бит
	*num &= ^mask
	// применяем лог. или на i бит
	*num |= value << index
}

func main() {
	var num, i, value int64
	fmt.Printf("Enter number: ")
	fmt.Scan(&num)
	fmt.Printf("Enter i: ")
	fmt.Scan(&i)
	fmt.Printf("Enter value: ")
	fmt.Scan(&value)
	setBit(&num, i, value)
	fmt.Printf("Changed number: %d\tbinary: %b\n", num, num)
}
