package main

import (
	"fmt"
	"math"
)

// ESTRUTURA DO RETANGULO
type retangulo struct {
	altura  float64
	largura float64
}

// MÉTODO DO RETANGULO QUE CÁLCULA A ÁREA
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// ESTRUTURA DO CIRCULO
type circulo struct {
	raio float64
}

// MÉTODO DO RETANGULO QUE CÁLCULA A ÁREA
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// INTERFACE FORMA
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	fmt.Println("Interfaces!!")
	fmt.Println()

	r := retangulo{10, 20}
	escreverArea(r)

	c := circulo{4}
	escreverArea(c)
}
