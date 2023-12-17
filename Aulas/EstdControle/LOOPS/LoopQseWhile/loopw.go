// Simulando uma função WHILE

package main

import "fmt"

func main() {
	//While- enquanto for verdade, faça isso
	x := 0 //Inicializa o contador fora do FOR
	for x < 20 {
		fmt.Println(x, " É MENOR QUE 20.")
		x++
	}
}
