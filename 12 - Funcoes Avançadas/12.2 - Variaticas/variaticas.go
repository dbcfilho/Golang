package main

import (
	"fmt"
)

// Função que recebe de 0 a N inteiros
func soma(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

// Função que recebe de 0 a N inteiros e uma string
// Podemos ter apenas um parâmetro variático por função
// E ele deve ser o último parâmetro
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("")
	resultado := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(resultado)

	resultado2 := soma()
	fmt.Println(resultado2)

	fmt.Println("--------------------")

	escrever("Olá Mundo", 10, 20, 100, 40)

}
