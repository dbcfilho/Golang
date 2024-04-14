package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64
	var numero int16 = 100
	fmt.Println(numero)

	// BITS é inferido de acordo com capacidade do PC
	numero2 := 100000
	fmt.Println(numero2)

	// UNSIGNED INT. NÃO PODE USAR "-"
	var numero3 uint = 200
	fmt.Println(numero3)

	// ALIAS
	// INT32 = RUNE
	var numero4 rune = 123
	fmt.Println(numero4)

	// BYTE = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)

	// STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// RETORNA A POSIÇÃO/NÚMERO NA TABELA ASCII
	char := 'A'
	fmt.Println(char)

	// VALOR ZERO
	var texto int
	fmt.Println(texto)

	// BOOLEANO
	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := false
	fmt.Println(booleano2)

	// ERRO É UM TIPO
	var erro1 error
	fmt.Println(erro1)

	// USANDO O PACOTE ERRORS
	var erro error = errors.New("Deu erro!")
	fmt.Println(erro)
}
