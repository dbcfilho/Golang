package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
	// ROTAS - IDENTIFICA QUE O TIPO DE REQUISIÇÃO E O TIPO DE RESPOSTA
	// URI - IDENTIFICADOR DO RECURSO
	// MÉTODO - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
