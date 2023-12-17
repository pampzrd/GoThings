// Select funciona como um Switch para canais
// Select permite que vc aguarde operações em vários canais
// Combinar goroutines com canais é um recurso poderoso do go
// Para esse exemplo usaremos dois canais.
// Cada canal receberá um valor após algum tempo, para simular, por exemplo o bloqueio de operações RPC em execução em goroutines concorrentes.
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() { //recebemos os valores "um" e depois "dois" conforme o esperado.
		time.Sleep(1 * time.Second) // O tempo total será de 2 segundos,
		//pois o 1 e o 2 são executados simultaneamente.
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "2"
	}()

	for i := 0; i < 2; i++ {
		select { //Usaremos select para aguardar os dois valores simultaneamente.
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)

		}
	}
}

//Apesar de receber os valor ao mesmo tempo(Otimiza o processamento), ele imprime um por vez.
