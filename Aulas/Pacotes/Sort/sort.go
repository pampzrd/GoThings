package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

// Interface ParaNome
type ParaNome []Dados{}

func (ps ParaNome) Len() int { //retorna o tamanho
	return len(ps)
}

func (ps ParaNome) Less(i, j int) bool { //compara se i é menor q j
	return ps[i].Nome < ps[j].Nome
}

func (ps ParaNome) Swap(i, j int) { //troca os itens
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	cri := Dados{
		{"Júlia", 5},
		{"João", 10},
	}

	sort.Sort(cri)
	fmt.Println(cri)
// ERRO NA INTERFACE.
}
