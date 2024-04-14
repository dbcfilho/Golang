package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança - só que não!")

	var p1 pessoa = pessoa{"João", "Pedro", 20, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "TI", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.curso)
}
