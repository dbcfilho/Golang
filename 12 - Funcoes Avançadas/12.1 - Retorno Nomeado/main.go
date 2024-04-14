package main

import (
	"fmt"
)

// Não é preciso especificar as variaveis no retorno
// porque já nomeamos no tipo de retorno
func calculo(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("Retornos Nomeados")

	soma, sub := calculo(10, 20)
	fmt.Println(soma, sub)
}
