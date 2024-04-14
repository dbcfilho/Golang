package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps!")

	usuario := map[string]string{
		"nome":      "Carol",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	estudantes := map[string]map[string]string{
		"estudante1": {
			"nome":  "João",
			"curso": "TI",
		},
	}
	fmt.Println(estudantes)
	delete(estudantes, "estudante1")
	fmt.Println(estudantes)
}
