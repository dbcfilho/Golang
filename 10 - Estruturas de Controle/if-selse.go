package main

import (
	"fmt"
)

func main() {
	n := 10

	if n > 15 {
		fmt.Println("Sim")
	} else {
		fmt.Println("Não")
	}

	if x := 5; x > 10 {
		fmt.Println("Maior que dez!")
	} else if x > 0 {
		fmt.Println("Maior que zero!")
	} else {
		fmt.Println("Menor que zero!")
	}
}
