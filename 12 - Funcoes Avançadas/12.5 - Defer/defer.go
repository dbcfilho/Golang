package main

import (
	"fmt"
)

func func1() {
	fmt.Println("Func 1")
}

func func2() {
	fmt.Println("Func 2")
}

func alunoAprovado(n1, n2 float32) string {
	defer fmt.Println("Média Calculada. Resultado do aluno: ")
	fmt.Println("Entrando na função!")

	media := (n1 + n2) / 2

	if media >= 7 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	defer func1()
	func2()

	statusAluno := alunoAprovado(10, 7)
	fmt.Println(statusAluno)
}
