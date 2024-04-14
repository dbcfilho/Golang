package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays Internos")
	fmt.Println("--------")

	// SLICES
	fmt.Println("Slices!")
	slice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(slice)
	fmt.Println(len(slice)) // TAMANHO
	fmt.Println(cap(slice)) // CAPACIDADE
	slice = append(slice, 3)
	fmt.Println(slice)
	fmt.Println(len(slice)) // TAMANHO
	fmt.Println(cap(slice)) // CAPACIDADE

	// Tipo do dado, tamanho, capacidade máx
	slice2 := make([]float32, 10, 15)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // TAMANHO
	fmt.Println(cap(slice2)) // CAPACIDADE
}
