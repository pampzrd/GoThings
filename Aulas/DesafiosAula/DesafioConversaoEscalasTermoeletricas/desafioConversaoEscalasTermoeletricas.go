/* PROPOSTA - Um professor de EM solicitou que seus alunos desenvolvessem um programa para converter o valor do ponto de ebulição da água de Kevin para graus Celsius. Dica: C = K -123
 */

package main

import "fmt"

// Valor fixo de temperatura Kelvin
const Kelvin = 273.15

func main() {
	// Referência Celsius de ebulição da água
	var CelsiusEbulicao float64 = 100
	ResultadoConversao := CelsiusEbulicao + Kelvin
	fmt.Printf("O ponto de ebulição da água de °K é %g", ResultadoConversao)
}
