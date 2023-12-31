/*
DESCRIÇÃO DO DESAFIO:

Vamos elevar o jogo, chegou a hora de você implementar
o algoritmo de uma calculadora, utilize seus conhecimentos
de Go até agora para implementar uma calculadora completa:
a função de soma, subtração, adição e multiplicação e
apresente os resultados corretos na tela.
*/
package main

import "fmt"

func main() {
	x := soma(1, 2, 3)               //6
	w := subtracao(13, 1, 2)         //10
	y := multiplicacao(10, 10, 2, 3) //600
	z := divisao(20, 2, 5)           //40,2,5=4 /20,2,5=2
	fmt.Printf("%d %d %d %.2f", x, w, y, z)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtracao(i ...int) int { // É i ... Porque n sabemos quantos argumentos serão adicionados.
	total := i[0]
	for _, v := range i {
		total -= v
	}
	return total + i[0]
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}
func divisao(i ...int) float64 {
	total := i[0]
	for _, v := range i {
		if v == i[0] {
			continue
			v = i[1]
		}
		total /= v
	}
	return float64(total)
}
