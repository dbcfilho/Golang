package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Enviando dado para o canal
		time.Sleep(time.Second)
	}

	close(canal) //Fechando o canal depois de enviar todos os dados
}

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	// Isso será executado antes da mensagem
	fmt.Println("Depois da função escrever!")

	// for {
	// 	mensagem, aberto := <- canal // Canal recebe valor e armazena em uma variável
	// 	if !aberto {
	// 		// Se o canal já foi fechado, vamos encerrar o Loop
	// 		// Para não dar o erro de Deadlock
	// 		break
	// 	}
	// 	// Se ainda estiver aberto, continua
	// 	fmt.Println(mensagem)
	// }

	//Mesma coisa, numa sintaxe mais enxuta
	// Não precisa da variavel "aberto"
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	// Isso só será executado depois que o canal receber o dado, pois a execução trava
	fmt.Println("Depois do canal!")
}
