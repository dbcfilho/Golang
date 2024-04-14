package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 3 {
		i++
		fmt.Println("Valor de i:", i)
		time.Sleep(time.Second)
	}

	// INIT DA VARIAVEL NO FOR LOOP
	for j := 0; j < 3; j++ {
		fmt.Println("J: ", j)
		time.Sleep(time.Second)
	}

	// FOR LOOP COM RANGE
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// FOR LOOP EM UM MAP
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Caio",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
