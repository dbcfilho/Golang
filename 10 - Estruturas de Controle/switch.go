package main

import (
	"fmt"
)

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	default:
		return "Numero invalido!"
	}
}

// Avaliando a conidção caso a caso
func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	default:
		diaDaSemana = "Numero invalido!"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(200)
	fmt.Println(dia)

	dia2 := diaDaSemana2(6)
	fmt.Println(dia2)
}
