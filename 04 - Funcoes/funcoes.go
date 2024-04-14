package main

import (
	"fmt"
)

func somar(x int8, y int8) int8 {
	return x + y
}

func calculosMatematicos(n1, n2 int16) (int16, int16) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	var soma1 int8 = somar(10, 20)
	soma2 := somar(20, 20)
	fmt.Println(soma1, soma2)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("Função F")

	resultado1, resultado2 := calculosMatematicos(20, 50)
	fmt.Println(resultado1)
	fmt.Println(resultado2)
}
