package main

import (
	"fmt"
)

// Executada antes da Função Main
// Mas ainda precisa da Função Main para rodar
// Pode ter uma por arquivo (e não por pacote)
// A main só pode ter uma por Pacote
func init() {
	fmt.Println("Executando a Função Init!")
}

func main() {
	fmt.Println("Executando  Função Main!")
}
