package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays & Slices!")

	// ARRAYS
	fmt.Println("Arrays!")
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]string{"Posição 1", "Posição 2", "Posição 3"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	fmt.Println("---------")
	// SLICES
	fmt.Println("Slices!")
	slice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	// Provando que são tipos de dados diferentes
	fmt.Println(reflect.TypeOf(slice)) // tipo do slice
	fmt.Println(reflect.TypeOf(arr3))  // tipo do array

	slice2 := arr2[1:3]
	slice2[1] = "Posição Alterada"
	fmt.Println(slice2)
	fmt.Println(arr2)
	//fmt.Println(reflect.TypeOf(slice2))
	//fmt.Println(reflect.TypeOf(arr2))
}
