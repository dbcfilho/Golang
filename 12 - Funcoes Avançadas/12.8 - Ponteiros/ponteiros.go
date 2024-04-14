package main

import (
	"fmt"
)

func inverteSinal(numero int) int {
	return numero * -1
}

func invertSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Funções com ponteiro!")

	// SEM PONTEIRO
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println("Numero Invertido: ", numeroInvertido)
	fmt.Println("Numero Original: ", numero)

	fmt.Println("--------------------------")

	// COM PONTEIRO
	novoNumero := 40
	fmt.Println(novoNumero)
	invertSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
