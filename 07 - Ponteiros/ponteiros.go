package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros!!")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	fmt.Println("----------")

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3

	fmt.Println(var3)
	fmt.Println(ponteiro)  //endereço de memória
	fmt.Println(*ponteiro) //valor da variável - DESREFERÊNCIAÇÃO

	fmt.Println("-----------TESTE----------")
	var var4 int = 800
	var p *int = &var4
	fmt.Println(var4, p)
	fmt.Println(*p)
}
