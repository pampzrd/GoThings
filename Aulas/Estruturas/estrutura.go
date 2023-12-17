package main

import "fmt"

// Registrando uma Estrutura
type pessoa struct {
	nome  string
	idade int
}

func main() {
	p1 := pessoa{"Ana", 54}
	p2 := pessoa{"Marmota", 35}
	fmt.Println(p1, p2)
}
