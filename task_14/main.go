package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = "v1"
	var b interface{} = 123
	var c interface{} = 13.45
	// с помощью рефлексии можно в рантайме получить тип переменной
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))

	// также с помощью switch (type) можно получить тип в рантайме
	switch a.(type) {
	case string:
		fmt.Println("a is string")
	case int:
		fmt.Println("a is int")
	default:
		fmt.Println("a is unknown type")
	}
}
