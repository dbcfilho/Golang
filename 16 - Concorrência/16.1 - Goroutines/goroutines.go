package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!!")
	escrever("Programando em Go!!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
