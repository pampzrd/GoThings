/* Utilitário para calcular quantas poções posso comprar no Ragnarok Online.*/
package main

import "fmt"

func main() {
	entrada := 0
	ehmercador := true
	tipopocao := 0
	valoresPocoes := [4]int{10, 50, 7, 38}
	valorBasePocao := 0
	quantidadeZenys := 0
	valorGasto := 0
	totalPocoes := 0

	fmt.Println("É mercador ? 1-SIM 2-NÃO")
	fmt.Scanln(&entrada)
	switch entrada {
	case 1:
		ehmercador = true
	case 2:
		ehmercador = false
	default:
		ehmercador = true
	}

	fmt.Println("1-Poção vermelha 2-Poção laranja")
	fmt.Scanln(&tipopocao)
	if ehmercador == true && tipopocao == 1 {
		valorBasePocao = valoresPocoes[2]
	} else if ehmercador == true && tipopocao == 2 {
		valorBasePocao = valoresPocoes[3]
	} else if ehmercador != true && tipopocao == 1 {
		valorBasePocao = valoresPocoes[0]
	} else if ehmercador != true && tipopocao == 2 {
		valorBasePocao = valoresPocoes[1]
	}

	// Recebe a entrada de quantidade de Zenys disponíveis

	fmt.Println("Inserir Zenys disponíveis:")
	fmt.Scanln(&quantidadeZenys)

	i := 0
	for {
		valorGasto += valorBasePocao

		if valorGasto > quantidadeZenys {
			break
		}
		totalPocoes++
		i++
	}

	// Calcula quantos Zenys vai sobrar na conta.
	troco := quantidadeZenys - (totalPocoes * valorBasePocao)
	fmt.Println("A quantidade de poções é ", totalPocoes, ".")
	fmt.Println("Sobrou ", troco, " Zenys de troco.")

}
