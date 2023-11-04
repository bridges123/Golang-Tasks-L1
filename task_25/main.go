package main

import (
	"fmt"
	"time"
)

func sleep1(d time.Duration) {
	// считаем время к которому закончится sleep
	dstTime := time.Now().Add(d).Unix()
	now := time.Now().Unix()
	// запускаем цикл пока Now не равно конечному (занимаем поток)
	for now != dstTime {
		now = time.Now().Unix()
	}
}

// с помощью Tick ждем пока в канал вернется новый Тайм
func sleep2(d time.Duration) {
	<-time.Tick(d)
}

// с помощью After ждем пока в канал вернется новый Тайм
func sleep3(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Before sleep 1")
	sleep1(2 * time.Second)
	fmt.Println("After sleep 1")

	fmt.Println("Before sleep 2")
	sleep2(2 * time.Second)
	fmt.Println("After sleep 2")

	fmt.Println("Before sleep 3")
	sleep3(2 * time.Second)
	fmt.Println("After sleep 3")
}
