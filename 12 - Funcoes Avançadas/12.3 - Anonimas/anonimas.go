package main

import (
	"fmt"
)

func main() {
	// SEM parametro
	func() {
		fmt.Println("Ola")
	}()

	// COM parametro
	func(texto string) {
		fmt.Println(texto)
	}("Parametro passado!")

	// COM RETORNO
	ret := func(n int) int {
		return n + 1
	}(10)
	fmt.Println(ret)
}
