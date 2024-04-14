package main

import (
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
	cachorroEmJSON := `{"Nome":"Rex","Raca":"Dalmata","Idade":3}`

	var c cachorro
	fmt.Println(c)

	erro := json.Unmarshal([]byte(cachorroEmJSON), &c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)
}
