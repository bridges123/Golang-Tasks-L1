package main

import (
	"fmt"
	"math"
)

// HumanMoves - методы структуры Human
type HumanMoves interface {
	PrintName()
	AgeDiff(age int) int
}

type Human struct {
	name   string
	age    int
	height float32
}

type Action struct {
	// Встраивание методов Human
	Human
}

func (h *Human) PrintName() {
	fmt.Printf("My name is %s\n", h.name)
}

func (h *Human) AgeDiff(age int) int {
	return int(math.Abs(float64(h.age - age)))
}

func main() {
	action := &Action{Human{name: "Max", age: 19}}
	action.PrintName()
	fmt.Println(action.AgeDiff(13))
}
