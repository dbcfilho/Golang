package main

import "fmt"

// A INTERFACE VAZIA ACEITA QUALQUER TIPO DE DADO
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(10)
	generica(true)
}
