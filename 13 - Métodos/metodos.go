package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no BD! \n", u.nome)
}

func (u usuario) maiorIdade() string {
	if u.idade > 18 {
		return "maior de idade!"
	} else {
		return "menor de idade!"
	}
}

func (u *usuario) fazAniversario() {
	u.idade++
}

func main() {
	var usu1 usuario = usuario{"Joao", 29}
	fmt.Println(usu1)
	usu1.salvar()
	maiorIdade := usu1.maiorIdade()
	fmt.Printf("O usuario 1 é %s \n", maiorIdade)

	usu2 := usuario{"Carol", 28}
	fmt.Println(usu2)
	usu2.salvar()
	maiorIdade2 := usu2.maiorIdade()
	fmt.Printf("O usuario 2 é %s \n", maiorIdade2)

	usu2.fazAniversario()
	fmt.Println(usu2)
}
