package main

import (
	"fmt"
)

func main() {
	// Declaração explicita de uma String
	var variavel1 string = "variavel 1"
	fmt.Println(variavel1)

	// Declaração implicita de uma String
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	// Explicita
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	// Multiplas variaveis
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//invertendo valores de variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
