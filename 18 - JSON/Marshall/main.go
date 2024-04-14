package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome`
	Raca  string `json:"raca`
	Idade uint   `json:"idade`
}

func main() {
	// Convertendo um Struct em JSON
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)                  // Slice de bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // Formato JSON

	// Convertendo um MAP em JSON
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Collier",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)                  // Slice de bytes
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON)) // Formato JSON
}
