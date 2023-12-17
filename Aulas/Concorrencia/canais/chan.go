//canal: Modo de duas goroutines se comunicarem e sincronizarem sua execução.

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { //palavra reservada chan //declara o canal
	for i := 0; ; i++ {
		c <- "ping" //usado para enviar e receber mensagem pelo canal
	}
}

func imprimir(c chan string) { //declara o canal
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string) //cria o canal

	go pingar(c)
	go imprimir(c)

	//sem o Scan ele vai fechar o programa. Com o Scan ele exibe a cada segundo a palavra "ping". pq? O tempo é por causa do
	//Modo de deixar o programa aberto até 'apertar o enter'
	var entrada string
	fmt.Scanln(&entrada)
}
