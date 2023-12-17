// Método é uma função associada a um tipo em particular (Aqui ele referencia o struct)
package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// A área vai receber um tipo retângulo
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r *retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{10, 5} //valor do retângulo a ser calculado
	fmt.Println("Área: ", r.area())
	fmt.Println("Perímetro: ", r.perimetro())
}
