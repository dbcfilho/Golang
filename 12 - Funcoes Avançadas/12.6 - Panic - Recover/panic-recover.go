package main

import (
	"fmt"
)

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É 6")
}

func main() {
	fmt.Println("Panic - Recover!")

	media := alunoAprovado(6, 6)
	fmt.Println(media)

	fmt.Println("Final da Execução!")

}
