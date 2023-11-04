package main

import (
	"math/rand"
	"strconv"
	"strings"
)

// глобальная переменная - плохая практика,
// т.к. значения постоянно хранится в памяти + переменная может вызвать путаницу с локальной
var justString string

func createHugeString(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += strconv.Itoa(rand.Intn(255) + 1)
	}
	return str
}

func someFunc() {
	// выделяем память под большую строку
	v := createHugeString(1 << 10)
	// храним в глобальной переменной ссылку на большую строку
	justString = v[:100]
}

func anotherFunc() string {
	// выделяем память под большую строку
	v := createHugeString(1 << 10)
	// возвращаем новую строку как результат, а не глобальной переменной
	// копируем эту подстроку в новый блок памяти, а не ссылаемся на hugeString
	return strings.Clone(v[0:100])
}

func main() {
	//someFunc()
	anotherFunc()
}
