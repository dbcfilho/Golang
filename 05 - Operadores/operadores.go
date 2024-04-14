package main

import (
	"fmt"
)

func main() {
	// ARITMETICOS
	fmt.Println("---------")
	soma := 1 + 2
	sub := 1 - 5
	divisao := 10 / 4
	resto := 10 % 4

	fmt.Println(soma, sub, divisao, resto)

	// OPERADORES LÓGICOS
	fmt.Println("---------")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false) // NOT FALSE
}
