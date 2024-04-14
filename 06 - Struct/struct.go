package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint64
}

func main() {
	fmt.Println("Arquivo structs!")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	endereco1 := endereco{"Rua Iperoig", 690}
	u2 := usuario{"Joao", 29, endereco1}
	fmt.Println(u2)

	u3 := usuario{nome: "Jose"}
	fmt.Println(u3)
}
