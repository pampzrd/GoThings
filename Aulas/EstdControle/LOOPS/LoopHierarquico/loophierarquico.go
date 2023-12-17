package main

import "fmt"

func main() {
	//Relógio é um Loop Hierárquico/ Loop aninhado. O loop interno precisa executar para o externo executar.
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("\nHora ", horas)
		for minutos := 0; minutos <= 59; minutos++ {
			fmt.Print(" ", minutos, " ")
		}
	}
}
