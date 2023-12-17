package main

import "fmt"

func main() {

	bandas := [6]string{"Oficina G3", "Skillet", "Jesus Culture", "Kari Jobe", "Hillsong", "Flyleaf"}

	fatia1 := []int{5, 6, 7, 8}

	// Faz uma fatia com zeros porque não usa uma array como referência
	fatia := make([]string, 5)
	//Cria uma fatia da array bandas
	fatia2 := bandas[3:5]

	//Adiciona no final da array
	fatia3acc := append(fatia2, "Resgate")

	//Copia para outra array
	fatiacopia := make([]int, 2)
	copy(fatiacopia, fatia1)

	fmt.Println(fatia)
	fmt.Println(fatia2)
	fmt.Println(fatia2, fatia3acc)
	fmt.Println(fatia1, fatiacopia)
}
