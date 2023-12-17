package main

import "fmt"

// Declarando uma constante
const ebulicaoF float64 = 212.0

func main() {
	//var F = ebulicaoF
	//var C = (F - 32) * 5 / 9
	// Usando Gopher (:=) para declarar uma variável, e elimina a palavra var
	F := ebulicaoF
	C := (F - 32) * 5 / 9
	//fmt.Println("O ponto de ebulição da água em °F é: ", F)
	//fmt.Println("O ponto de ebulição da água em °C é: ", C)
	// Usando o PRINT FORMAT  referencia para float -> %g
	fmt.Printf("O ponto de ebulição da água em °F é: %g  \nO ponto de ebulição da água em °C é: %g", F, C)
}
