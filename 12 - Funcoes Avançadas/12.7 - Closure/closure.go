package main

import (
	"fmt"
)

// Função que retorna outra função
func teste() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	fmt.Println("Funções Closure!")

	texto := "Dentro da função Main!"
	fmt.Println(texto)

	funcaoNova := teste()
	funcaoNova()
}
