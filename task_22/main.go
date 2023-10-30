package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 26)
	b := big.NewInt((1 << 25) - (1 << 10))
	fmt.Printf("a = %v\tb = %v\n", a, b)
	fmt.Printf("a + b = %v\n", big.NewInt(0).Add(a, b))
	fmt.Printf("a - b = %v\n", big.NewInt(0).Sub(a, b))
	fmt.Printf("a * b = %v\n", big.NewInt(0).Mul(a, b))
	fmt.Printf("a / b = %v\n", big.NewInt(0).Div(a, b))
}
